package service

import (
	"context"
	"errors"
	"net/url"
	"path"
	"strconv"
	"strings"

	appproviderv1beta1 "github.com/cs3org/go-cs3apis/cs3/app/provider/v1beta1"
	gatewayv1beta1 "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	rpcv1beta1 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/v2/pkg/utils"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/config"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/helpers"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/middleware"
)

// NewHandler creates a new grpc service implementing the OpenInApp interface
func NewHandler(opts ...Option) (*Service, func(), error) {
	teardown := func() {
		/* this is required as a argument for the return value to satisfy the interface */
		/* in case you are wondering about the necessity of this comment, sonarcloud is asking for it */
	}
	options := newOptions(opts...)

	gwc := options.Gwc
	var err error
	if gwc == nil {
		gwc, err = pool.GetGatewayServiceClient(options.Config.CS3Api.Gateway.Name)
		if err != nil {
			return nil, teardown, err
		}
	}

	return &Service{
		id:      options.Config.GRPC.Namespace + "." + options.Config.Service.Name + "." + options.Config.App.Name,
		appURLs: options.AppURLs,
		logger:  options.Logger,
		config:  options.Config,
		gwc:     gwc,
	}, teardown, nil
}

// Service implements the OpenInApp interface
type Service struct {
	id      string
	appURLs map[string]map[string]string
	logger  log.Logger
	config  *config.Config
	gwc     gatewayv1beta1.GatewayAPIClient
}

// OpenInApp will implement the OpenInApp interface of the app provider
func (s *Service) OpenInApp(
	ctx context.Context,
	req *appproviderv1beta1.OpenInAppRequest,
) (*appproviderv1beta1.OpenInAppResponse, error) {

	// get the current user
	var user *userv1beta1.User = nil
	meReq := &gatewayv1beta1.WhoAmIRequest{
		Token: req.GetAccessToken(),
	}
	meResp, err := s.gwc.WhoAmI(ctx, meReq)
	if err == nil {
		if meResp.GetStatus().GetCode() == rpcv1beta1.Code_CODE_OK {
			user = meResp.GetUser()
		}
	}

	// required for the response, it will be used also for logs
	providerFileRef := providerv1beta1.Reference{
		ResourceId: req.GetResourceInfo().GetId(),
		Path:       ".",
	}

	logger := s.logger.With().
		Str("FileReference", providerFileRef.String()).
		Str("ViewMode", req.GetViewMode().String()).
		Str("Requester", user.GetId().String()).
		Logger()

	// get the file extension to use the right wopi app url
	fileExt := path.Ext(req.GetResourceInfo().GetPath())

	// get the appURL we need to use
	appURL := s.getAppUrl(fileExt, req.GetViewMode())
	if appURL == "" {
		logger.Error().Msg("OpenInApp: neither edit nor view app URL found")
		return nil, errors.New("neither edit nor view app URL found")
	}

	// append the parameters we need
	appURL, err = s.addQueryToURL(appURL, req)
	if err != nil {
		logger.Error().Err(err).Msg("OpenInApp: error parsing appUrl")
		return nil, err
	}

	// create the wopiContext and generate the token
	wopiContext := middleware.WopiContext{
		AccessToken:   req.GetAccessToken(), // it will be encrypted
		ViewOnlyToken: utils.ReadPlainFromOpaque(req.GetOpaque(), "viewOnlyToken"),
		FileReference: &providerFileRef,
		User:          user,
		ViewMode:      req.GetViewMode(),
	}

	accessToken, accessExpiration, err := middleware.GenerateWopiToken(wopiContext, s.config)
	if err != nil {
		logger.Error().Err(err).Msg("OpenInApp: error generating the token")
		return &appproviderv1beta1.OpenInAppResponse{
			Status: &rpcv1beta1.Status{Code: rpcv1beta1.Code_CODE_INTERNAL},
		}, err
	}

	logger.Debug().Msg("OpenInApp: success")

	return &appproviderv1beta1.OpenInAppResponse{
		Status: &rpcv1beta1.Status{Code: rpcv1beta1.Code_CODE_OK},
		AppUrl: &appproviderv1beta1.OpenInAppURL{
			AppUrl: appURL,
			Method: "POST",
			FormParameters: map[string]string{
				// these parameters will be passed to the web server by the app provider application
				"access_token": accessToken,
				// milliseconds since Jan 1, 1970 UTC as required in https://docs.microsoft.com/en-us/microsoft-365/cloud-storage-partner-program/rest/concepts#access_token_ttl
				//"access_token_ttl": strconv.FormatInt(claims.ExpiresAt.UnixMilli(), 10),
				"access_token_ttl": strconv.FormatInt(accessExpiration, 10),
			},
		},
	}, nil
}

// getAppUrlFor gets the appURL from the list of appURLs based on the
// action and file extension provided. If there is no match, an empty
// string will be returned.
func (s *Service) getAppUrlFor(action, fileExt string) string {
	if actionURL, ok := s.appURLs[action]; ok {
		if actionExtensionURL, ok := actionURL[fileExt]; ok {
			return actionExtensionURL
		}
	}
	return ""
}

// getAppUrl will get the appURL that should be used based on the extension
// and the provided view mode.
// "view" urls will be chosen first, then if the view mode is "read/write",
// "edit" urls will be prioritized. Note that "view" url might be returned for
// "read/write" view mode if no "edit" url is found.
func (s *Service) getAppUrl(fileExt string, viewMode appproviderv1beta1.ViewMode) string {
	// check view_comment action first (for collabora)
	appURL := s.getAppUrlFor("view_comment", fileExt)

	// prioritize view action if possible
	if viewAppURL := s.getAppUrlFor("view", fileExt); viewAppURL != "" {
		appURL = viewAppURL
	}

	// If read/write mode has been requested, prioritize edit action.
	// Special case for collabora because it only provides one action per
	// extension,
	if viewMode == appproviderv1beta1.ViewMode_VIEW_MODE_READ_WRITE || strings.ToLower(s.config.App.Name) == "collabora" {
		if editAppURL := s.getAppUrlFor("edit", fileExt); editAppURL != "" {
			appURL = editAppURL
		}
	}

	return appURL
}

// addQueryToURL will add specific query parameters to the baseURL. These
// parameters are:
// * "WOPISrc" pointing to the requested resource in the OpenInAppRequest
// * "dchat" to disable the chat, based on configuration
// * "lang" (WOPI app dependent) with the language in the request. "lang"
// for collabora, "ui" for onlyoffice and "UI_LLCC" for the rest
func (s *Service) addQueryToURL(baseURL string, req *appproviderv1beta1.OpenInAppRequest) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// build a urlsafe and stable file reference that can be used for proxy routing,
	// so that all sessions on one file end on the same office server
	fileRef := helpers.HashResourceId(req.GetResourceInfo().GetId())

	wopiSrcURL, err := url.Parse(s.config.Wopi.WopiSrc)
	if err != nil {
		return "", err
	}
	wopiSrcURL.Path = path.Join("wopi", "files", fileRef)

	q := u.Query()
	q.Add("WOPISrc", wopiSrcURL.String())

	if s.config.Wopi.DisableChat {
		q.Add("dchat", "1")
	}

	lang := utils.ReadPlainFromOpaque(req.GetOpaque(), "lang")

	if lang != "" {
		switch strings.ToLower(s.config.App.Name) {
		case "collabora":
			q.Add("lang", lang)
		case "onlyoffice":
			q.Add("ui", lang)
		default:
			q.Add("UI_LLCC", lang)
		}
	}

	qs := q.Encode()
	u.RawQuery = qs

	return u.String(), nil
}

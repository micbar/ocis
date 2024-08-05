package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"path"
	"strconv"

	appproviderv1beta1 "github.com/cs3org/go-cs3apis/cs3/app/provider/v1beta1"
	gatewayv1beta1 "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	rpcv1beta1 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/cs3org/reva/v2/pkg/utils"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/wopisrc"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/config"
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

	// build a urlsafe and stable file reference that can be used for proxy routing,
	// so that all sessions on one file end on the same office server

	c := sha256.New()
	c.Write([]byte(req.GetResourceInfo().GetId().GetStorageId() + "$" + req.GetResourceInfo().GetId().GetSpaceId() + "!" + req.GetResourceInfo().GetId().GetOpaqueId()))
	fileRef := hex.EncodeToString(c.Sum(nil))

	// get the file extension to use the right wopi app url
	fileExt := path.Ext(req.GetResourceInfo().GetPath())

	var viewCommentAppURL string
	var viewAppURL string
	var editAppURL string
	if viewCommentAppURLs, ok := s.appURLs["view_comment"]; ok {
		if u, ok := viewCommentAppURLs[fileExt]; ok {
			viewCommentAppURL = u
		}
	}
	if viewAppURLs, ok := s.appURLs["view"]; ok {
		if u, ok := viewAppURLs[fileExt]; ok {
			viewAppURL = u
		}
	}
	if editAppURLs, ok := s.appURLs["edit"]; ok {
		if u, ok := editAppURLs[fileExt]; ok {
			editAppURL = u
		}
	}
	if editAppURL == "" && viewAppURL == "" && viewCommentAppURL == "" {
		err := fmt.Errorf("OpenInApp: neither edit nor view app url found")
		s.logger.Error().
			Err(err).
			Str("FileReference", providerFileRef.String()).
			Str("ViewMode", req.GetViewMode().String()).
			Str("Requester", user.GetId().String()).Send()
		return nil, err
	}

	if editAppURL == "" {
		// assuming that an view action is always available in the /hosting/discovery manifest
		// eg. Collabora does support viewing jpgs but no editing
		// eg. OnlyOffice does support viewing pdfs but no editing
		// there is no known case of supporting edit only without view
		editAppURL = viewAppURL
	}
	if viewAppURL == "" {
		// the URL of the end-user application in view mode when different (defaults to edit mod URL)
		viewAppURL = editAppURL
	}
	// TODO: check if collabora will support an "edit" url in the future
	if viewAppURL == "" && editAppURL == "" && viewCommentAppURL != "" {
		// there are rare cases where neither view nor edit is supported but view_comment is
		viewAppURL = viewCommentAppURL
		// that can be the case for editable and viewable files
		if req.GetViewMode() == appproviderv1beta1.ViewMode_VIEW_MODE_READ_WRITE {
			editAppURL = viewCommentAppURL
		}
	}
	wopiSrcURL, err := wopisrc.GenerateWopiSrc(fileRef, s.config)
	if err != nil {
		return nil, err
	}

	addWopiSrcQueryParam := func(baseURL string) (string, error) {
		u, err := url.Parse(baseURL)
		if err != nil {
			return "", err
		}

		q := u.Query()
		q.Add("WOPISrc", wopiSrcURL.String())

		if s.config.Wopi.DisableChat {
			q.Add("dchat", "1")
		}

		lang := utils.ReadPlainFromOpaque(req.GetOpaque(), "lang")

		// @TODO: this is a temporary solution until we figure out how to send these from oc web
		switch lang {
		case "bg":
			lang = "bg-BG"
		case "cs":
			lang = "cs-CZ"
		case "de":
			lang = "de-DE"
		case "en":
			lang = "en-US"
		case "es":
			lang = "es-ES"
		case "fr":
			lang = "fr-FR"
		case "gl":
			lang = "gl-ES"
		case "it":
			lang = "it-IT"
		case "nl":
			lang = "nl-NL"
		case "ko":
			lang = "ko-KR"
		case "sq":
			lang = "sq-AL"
		case "sv":
			lang = "sv-SE"
		case "tr":
			lang = "tr-TR"
		case "zh":
			lang = "zh-CN"
		}

		if lang != "" {
			q.Add("ui", lang)      // OnlyOffice
			q.Add("lang", lang)    // Collabora, Impact on the default document language of OnlyOffice
			q.Add("UI_LLCC", lang) // Office365
		}
		qs := q.Encode()
		u.RawQuery = qs

		return u.String(), nil
	}

	viewAppURL, err = addWopiSrcQueryParam(viewAppURL)
	if err != nil {
		s.logger.Error().
			Err(err).
			Str("FileReference", providerFileRef.String()).
			Str("ViewMode", req.GetViewMode().String()).
			Str("Requester", user.GetId().String()).
			Msg("OpenInApp: error parsing viewAppUrl")
		return nil, err
	}
	editAppURL, err = addWopiSrcQueryParam(editAppURL)
	if err != nil {
		s.logger.Error().
			Err(err).
			Str("FileReference", providerFileRef.String()).
			Str("ViewMode", req.GetViewMode().String()).
			Str("Requester", user.GetId().String()).
			Msg("OpenInApp: error parsing editAppUrl")
		return nil, err
	}

	appURL := viewAppURL
	if req.GetViewMode() == appproviderv1beta1.ViewMode_VIEW_MODE_READ_WRITE {
		appURL = editAppURL
	}

	wopiContext := middleware.WopiContext{
		AccessToken:   req.GetAccessToken(), // it will be encrypted
		ViewOnlyToken: utils.ReadPlainFromOpaque(req.GetOpaque(), "viewOnlyToken"),
		FileReference: &providerFileRef,
		ViewMode:      req.GetViewMode(),
	}

	accessToken, accessExpiration, err := middleware.GenerateWopiToken(wopiContext, s.config)
	if err != nil {
		s.logger.Error().
			Err(err).
			Str("FileReference", providerFileRef.String()).
			Str("ViewMode", req.GetViewMode().String()).
			Str("Requester", user.GetId().String()).
			Msg("OpenInApp: error generating the token")
		return &appproviderv1beta1.OpenInAppResponse{
			Status: &rpcv1beta1.Status{Code: rpcv1beta1.Code_CODE_INTERNAL},
		}, err
	}

	s.logger.Debug().
		Str("FileReference", providerFileRef.String()).
		Str("ViewMode", req.GetViewMode().String()).
		Str("Requester", user.GetId().String()).
		Str("AppUrl", appURL).
		Msg("OpenInApp: success")

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

package service_test

import (
	"context"
	"strconv"
	"time"

	"github.com/cs3org/reva/v2/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	appproviderv1beta1 "github.com/cs3org/go-cs3apis/cs3/app/provider/v1beta1"
	authpb "github.com/cs3org/go-cs3apis/cs3/auth/provider/v1beta1"
	gatewayv1beta1 "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	rpcv1beta1 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	providerv1beta1 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	types "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"

	"github.com/cs3org/reva/v2/pkg/rgrpc/status"
	cs3mocks "github.com/cs3org/reva/v2/tests/cs3mocks/mocks"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/config"
	service "github.com/owncloud/ocis/v2/services/collaboration/pkg/service/grpc/v0"
)

// Based on https://github.com/cs3org/reva/blob/b99ad4865401144a981d4cfd1ae28b5a018ea51d/pkg/token/manager/jwt/jwt.go#L82
func MintToken(u *userv1beta1.User, secret string, nowTime time.Time) string {
	scopes := make(map[string]*authpb.Scope)
	scopes["user"] = &authpb.Scope{
		Resource: &types.OpaqueEntry{
			Decoder: "json",
			Value:   []byte("{\"Path\":\"/\"}"),
		},
		Role: authpb.Role_ROLE_OWNER,
	}

	claims := jwt.MapClaims{
		"exp":   nowTime.Add(5 * time.Hour).Unix(),
		"iss":   "myself",
		"aud":   "reva",
		"iat":   nowTime.Unix(),
		"user":  u,
		"scope": scopes,
	}
	/*
		claims := claims{
			StandardClaims: jwt.RegisteredClaims{
				ExpiresAt: time.Now().Add(5 * time.Hour),
				Issuer:    "myself",
				Audience:  "reva",
				IssuedAt:  time.Now(),
			},
			User:  u,
			Scope: scopes,
		}
	*/

	t := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	tkn, _ := t.SignedString([]byte(secret))

	return tkn
}

var _ = Describe("Discovery", func() {
	var (
		cfg           *config.Config
		gatewayClient *cs3mocks.GatewayAPIClient
		srv           *service.Service
		srvTear       func()
	)

	BeforeEach(func() {
		cfg = &config.Config{}
		gatewayClient = &cs3mocks.GatewayAPIClient{}

		srv, srvTear, _ = service.NewHandler(
			service.Logger(log.NopLogger()),
			service.Config(cfg),
			service.AppURLs(map[string]map[string]string{
				"view": map[string]string{
					".pdf":  "https://test.server.prv/hosting/wopi/word/view",
					".djvu": "https://test.server.prv/hosting/wopi/word/view",
					".docx": "https://test.server.prv/hosting/wopi/word/view",
					".xls":  "https://test.server.prv/hosting/wopi/cell/view",
					".xlsb": "https://test.server.prv/hosting/wopi/cell/view",
				},
				"edit": map[string]string{
					".docx": "https://test.server.prv/hosting/wopi/word/edit",
				},
			}),
			service.GatewayAPIClient(gatewayClient),
		)
	})

	AfterEach(func() {
		srvTear()
	})

	Describe("OpenInApp", func() {
		It("Invalid access token", func() {
			ctx := context.Background()

			cfg.Wopi.WopiSrc = "https://wopiserver.test.prv"

			req := &appproviderv1beta1.OpenInAppRequest{
				ResourceInfo: &providerv1beta1.ResourceInfo{
					Id: &providerv1beta1.ResourceId{
						StorageId: "myStorage",
						OpaqueId:  "storageOpaque001",
						SpaceId:   "SpaceA",
					},
					Path: "/path/to/file",
				},
				ViewMode:    appproviderv1beta1.ViewMode_VIEW_MODE_READ_WRITE,
				AccessToken: "goodAccessToken",
			}

			gatewayClient.On("WhoAmI", mock.Anything, mock.Anything).Times(1).Return(&gatewayv1beta1.WhoAmIResponse{
				Status: status.NewOK(ctx),
				User: &userv1beta1.User{
					Id: &userv1beta1.UserId{
						Idp:      "myIdp",
						OpaqueId: "opaque001",
						Type:     userv1beta1.UserType_USER_TYPE_PRIMARY,
					},
					Username: "username",
				},
			}, nil)

			resp, err := srv.OpenInApp(ctx, req)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})

		DescribeTable(
			"Success",
			func(appName, lang string, disableChat bool, expectedAppUrl string) {
				ctx := context.Background()
				nowTime := time.Now()

				cfg.Wopi.WopiSrc = "https://wopiserver.test.prv"
				cfg.Wopi.Secret = "my_supa_secret"
				cfg.Wopi.DisableChat = disableChat
				cfg.App.Name = appName

				myself := &userv1beta1.User{
					Id: &userv1beta1.UserId{
						Idp:      "myIdp",
						OpaqueId: "opaque001",
						Type:     userv1beta1.UserType_USER_TYPE_PRIMARY,
					},
					Username: "username",
				}

				req := &appproviderv1beta1.OpenInAppRequest{
					ResourceInfo: &providerv1beta1.ResourceInfo{
						Id: &providerv1beta1.ResourceId{
							StorageId: "myStorage",
							OpaqueId:  "storageOpaque001",
							SpaceId:   "SpaceA",
						},
						Path: "/path/to/file.docx",
					},
					ViewMode:    appproviderv1beta1.ViewMode_VIEW_MODE_READ_WRITE,
					AccessToken: MintToken(myself, cfg.Wopi.Secret, nowTime),
				}
				if lang != "" {
					req.Opaque = utils.AppendPlainToOpaque(req.Opaque, "lang", lang)
				}

				gatewayClient.On("WhoAmI", mock.Anything, mock.Anything).Times(1).Return(&gatewayv1beta1.WhoAmIResponse{
					Status: status.NewOK(ctx),
					User:   myself,
				}, nil)

				resp, err := srv.OpenInApp(ctx, req)
				Expect(err).To(Succeed())
				Expect(resp.GetStatus().GetCode()).To(Equal(rpcv1beta1.Code_CODE_OK))
				Expect(resp.GetAppUrl().GetMethod()).To(Equal("POST"))
				Expect(resp.GetAppUrl().GetAppUrl()).To(Equal(expectedAppUrl))
				Expect(resp.GetAppUrl().GetFormParameters()["access_token_ttl"]).To(Equal(strconv.FormatInt(nowTime.Add(5*time.Hour).Unix()*1000, 10)))
			},
			Entry("Microsoft chat no lang", "Microsoft", "", false, "https://test.server.prv/hosting/wopi/word/edit?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e"),
			Entry("Collabora chat no lang", "Collabora", "", false, "https://test.server.prv/hosting/wopi/word/view?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e"),
			Entry("OnlyOffice chat no lang", "OnlyOffice", "", false, "https://test.server.prv/hosting/wopi/word/edit?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e"),
			Entry("Microsoft chat lang", "Microsoft", "de", false, "https://test.server.prv/hosting/wopi/word/edit?UI_LLCC=de-DE&WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e"),
			Entry("Collabora chat lang", "Collabora", "de", false, "https://test.server.prv/hosting/wopi/word/view?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&lang=de-DE"),
			Entry("OnlyOffice chat lang", "OnlyOffice", "de", false, "https://test.server.prv/hosting/wopi/word/edit?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&ui=de-DE"),
			Entry("Microsoft no chat no lang", "Microsoft", "", true, "https://test.server.prv/hosting/wopi/word/edit?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&dchat=1"),
			Entry("Collabora no chat no lang", "Collabora", "", true, "https://test.server.prv/hosting/wopi/word/view?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&dchat=1"),
			Entry("OnlyOffice no chat no lang", "OnlyOffice", "", true, "https://test.server.prv/hosting/wopi/word/edit?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&dchat=1"),
			Entry("Microsoft no chat lang", "Microsoft", "de", true, "https://test.server.prv/hosting/wopi/word/edit?UI_LLCC=de-DE&WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&dchat=1"),
			Entry("Collabora no chat lang", "Collabora", "de", true, "https://test.server.prv/hosting/wopi/word/view?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&dchat=1&lang=de-DE"),
			Entry("OnlyOffice no chat lang", "OnlyOffice", "de", true, "https://test.server.prv/hosting/wopi/word/edit?WOPISrc=https%3A%2F%2Fwopiserver.test.prv%2Fwopi%2Ffiles%2F2f6ec18696dd1008106749bd94106e5cfad5c09e15de7b77088d03843e71b43e&dchat=1&ui=de-DE"),
		)
	})
})

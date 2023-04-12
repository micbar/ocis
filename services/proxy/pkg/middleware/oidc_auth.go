package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/ocis-pkg/oidc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/shamaton/msgpack/v2"
	store "go-micro.dev/v4/store"
	"golang.org/x/crypto/sha3"
	"golang.org/x/oauth2"
)

const (
	_headerAuthorization = "Authorization"
	_bearerPrefix        = "Bearer "
)

// NewOIDCAuthenticator returns a ready to use authenticator which can handle OIDC authentication.
func NewOIDCAuthenticator(opts ...Option) *OIDCAuthenticator {
	options := newOptions(opts...)

	return &OIDCAuthenticator{
		Logger:                  options.Logger,
		userInfoCache:           options.UserInfoCache,
		sessionLookupCache:      options.SessionLookupCache,
		DefaultTokenCacheTTL:    options.DefaultAccessTokenTTL,
		HTTPClient:              options.HTTPClient,
		OIDCIss:                 options.OIDCIss,
		oidcClient:              options.OIDCClient,
		AccessTokenVerifyMethod: options.AccessTokenVerifyMethod,
	}
}

// OIDCAuthenticator is an authenticator responsible for OIDC authentication.
type OIDCAuthenticator struct {
	Logger                  log.Logger
	HTTPClient              *http.Client
	OIDCIss                 string
	userInfoCache           store.Store
	sessionLookupCache      store.Store
	DefaultTokenCacheTTL    time.Duration
	oidcClient              oidc.OIDCProvider
	AccessTokenVerifyMethod string
}

func (m *OIDCAuthenticator) getClaims(token string, req *http.Request) (map[string]interface{}, error) {
	var claims map[string]interface{}

	// use a 64 bytes long hash to have 256-bit collision resistance.
	hash := make([]byte, 64)
	sha3.ShakeSum256(hash, []byte(token))
	encodedHash := base64.URLEncoding.EncodeToString(hash)

	record, err := m.userInfoCache.Read(encodedHash)
	if err != nil && err != store.ErrNotFound {
		m.Logger.Error().Err(err).Msg("could not read from userinfo cache")
	}
	if len(record) > 0 {
		if err = msgpack.Unmarshal(record[0].Value, &claims); err == nil {
			m.Logger.Debug().Interface("claims", claims).Msg("cache hit for userinfo")
			return claims, nil
		}
		m.Logger.Error().Err(err).Msg("could not unmarshal userinfo")
	}

	// TODO: use mClaims
	aClaims, _, err := m.oidcClient.VerifyAccessToken(req.Context(), token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to verify access token")
	}

	oauth2Token := &oauth2.Token{
		AccessToken: token,
	}

	userInfo, err := m.oidcClient.UserInfo(
		context.WithValue(req.Context(), oauth2.HTTPClient, m.HTTPClient),
		oauth2.StaticTokenSource(oauth2Token),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get userinfo")
	}
	if err := userInfo.Claims(&claims); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal userinfo claims")
	}

	expiration := m.extractExpiration(aClaims)
	go func() {
		if d, err := msgpack.Marshal(claims); err != nil {
			m.Logger.Error().Err(err).Msg("failed to marshal claims for userinfo cache")
		} else {
			err = m.userInfoCache.Write(&store.Record{
				Key:    encodedHash,
				Value:  d,
				Expiry: time.Until(expiration),
			})
			if err != nil {
				m.Logger.Error().Err(err).Msg("failed to write to userinfo cache")
			}

			if sid, ok := claims["sid"]; ok {
				err = m.sessionLookupCache.Write(&store.Record{
					Key:    fmt.Sprintf("%s", sid),
					Value:  []byte(encodedHash),
					Expiry: time.Until(expiration),
				})
			}
			if err != nil {
				m.Logger.Error().Err(err).Msg("failed to write session lookup cache")
			}
		}
	}()

	m.Logger.Debug().Interface("claims", claims).Msg("extracted claims")
	return claims, nil
}

// extractExpiration tries to extract the expriration time from the access token
// If the access token does not have an exp claim it will fallback to the configured
// default expiration
func (m OIDCAuthenticator) extractExpiration(aClaims jwt.RegisteredClaims) time.Time {
	defaultExpiration := time.Now().Add(m.DefaultTokenCacheTTL)
	if aClaims.ExpiresAt != nil {
		m.Logger.Debug().Str("exp", aClaims.ExpiresAt.String()).Msg("Expiration Time from access_token")
		return aClaims.ExpiresAt.Time
	}
	return defaultExpiration
}

func (m OIDCAuthenticator) shouldServe(req *http.Request) bool {
	if m.OIDCIss == "" {
		return false
	}

	header := req.Header.Get(_headerAuthorization)
	return strings.HasPrefix(header, _bearerPrefix)
}

// Authenticate implements the authenticator interface to authenticate requests via oidc auth.
func (m *OIDCAuthenticator) Authenticate(r *http.Request) (*http.Request, bool) {
	// there is no bearer token on the request,
	if !m.shouldServe(r) || isPublicPath(r.URL.Path) {
		// The authentication of public path requests is handled by another authenticator.
		// Since we can't guarantee the order of execution of the authenticators, we better
		// implement an early return here for paths we can't authenticate in this authenticator.
		return nil, false
	}
	token := strings.TrimPrefix(r.Header.Get(_headerAuthorization), _bearerPrefix)

	claims, err := m.getClaims(token, r)
	if err != nil {
		m.Logger.Error().
			Err(err).
			Str("authenticator", "oidc").
			Str("path", r.URL.Path).
			Msg("failed to authenticate the request")
		return nil, false
	}
	m.Logger.Debug().
		Str("authenticator", "oidc").
		Str("path", r.URL.Path).
		Msg("successfully authenticated request")
	return r.WithContext(oidc.NewContext(r.Context(), claims)), true
}

package wopisrc

import (
	"net/url"
	"path"

	"github.com/golang-jwt/jwt/v4"
	"github.com/owncloud/ocis/v2/services/collaboration/pkg/config"
)

// GenerateWopiSrc generates a WOPI src URL for the given file reference.
// If a proxy URL and proxy secret are configured, the URL will be generated
// as a jwt token that is signed with the proxy secret and contains the file reference
// and the WOPI src URL.
// Example:
// https://localhost:9300/wopi/files/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiZm9vIiwiZiI6ImJhciJ9.123456?access_token=123456&access_token_ttl=1234
//
// If no proxy URL and proxy secret are configured, the URL will be generated
// as a direct URL that contains the file reference.
// Example:
// https://localhost:9300/wopi/files/12312678470610632091729803710923&access_token=123456&access_token_ttl=1234
func GenerateWopiSrc(fileRef string, cfg *config.Config) (*url.URL, error) {
	wopiSrcURL, err := url.Parse(cfg.Wopi.WopiSrc)
	if err != nil {
		return nil, err
	}

	if cfg.Wopi.ProxyURL != "" && cfg.Wopi.ProxySecret != "" {
		return generateProxySrc(fileRef, cfg.Wopi.ProxyURL, cfg.Wopi.ProxySecret, wopiSrcURL)
	}

	return generateDirectSrc(fileRef, wopiSrcURL)
}

func generateDirectSrc(fileRef string, wopiSrcURL *url.URL) (*url.URL, error) {
	wopiSrcURL.Path = path.Join("wopi", "files", fileRef)
	return wopiSrcURL, nil
}

func generateProxySrc(fileRef string, proxyUrl string, proxySecret string, wopiSrcURL *url.URL) (*url.URL, error) {
	proxyURL, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, err
	}

	wopiSrcURL.Path = path.Join("wopi", "files")

	type tokenClaims struct {
		URL    string `json:"u"`
		FileID string `json:"f"`
		jwt.RegisteredClaims
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		FileID: fileRef,
		// the string value from the URL package always ends with a slash
		// the office365 proxy assumes that we have a trailing slash
		URL: wopiSrcURL.String() + "/",
	})
	tokenString, err := token.SignedString([]byte(proxySecret))
	if err != nil {
		return nil, err
	}
	proxyURL.Path = path.Join("wopi", "files", tokenString)
	return proxyURL, nil
}

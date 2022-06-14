package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
)

// Config combines all available configuration parts.
type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing"`
	Log     *Log     `yaml:"log"`
	Debug   Debug    `yaml:"debug"`

	HTTP HTTP `yaml:"http"`

	TokenManager *TokenManager `yaml:"token_manager"`
	Reva         *Reva         `yaml:"reva"`

	IdentityManagement IdentityManagement `yaml:"identity_management"`

	AccountBackend    string `yaml:"account_backend" env:"OCS_ACCOUNT_BACKEND_TYPE"`
	MachineAuthAPIKey string `yaml:"machine_auth_api_key" env:"OCIS_MACHINE_AUTH_API_KEY;OCS_MACHINE_AUTH_API_KEY" desc: "Machine auth API key used for accessing the 'auth-machine' service to impersonate users."`

	Context context.Context `yaml:"-"`
}

// IdentityManagement keeps track of the OIDC address. This is because Reva requisite of uniqueness for users
// is based in the combination of IDP hostname + UserID. For more information see:
// https://github.com/cs3org/reva/blob/4fd0229f13fae5bc9684556a82dbbd0eced65ef9/pkg/storage/utils/decomposedfs/node/node.go#L856-L865
type IdentityManagement struct {
	Address string `yaml:"address" env:"OCIS_URL;OCIS_OIDC_ISSUER;OCS_IDM_ADDRESS" desc:"URL of the OIDC issuer. It defaults to URL of the builtin IDP."`
}

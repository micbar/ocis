// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	jwt "github.com/golang-jwt/jwt/v5"
	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"

	oidc "github.com/owncloud/ocis/v2/ocis-pkg/oidc"
)

// OIDCClient is an autogenerated mock type for the OIDCClient type
type OIDCClient struct {
	mock.Mock
}

type OIDCClient_Expecter struct {
	mock *mock.Mock
}

func (_m *OIDCClient) EXPECT() *OIDCClient_Expecter {
	return &OIDCClient_Expecter{mock: &_m.Mock}
}

// UserInfo provides a mock function with given fields: ctx, ts
func (_m *OIDCClient) UserInfo(ctx context.Context, ts oauth2.TokenSource) (*oidc.UserInfo, error) {
	ret := _m.Called(ctx, ts)

	if len(ret) == 0 {
		panic("no return value specified for UserInfo")
	}

	var r0 *oidc.UserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.TokenSource) (*oidc.UserInfo, error)); ok {
		return rf(ctx, ts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oauth2.TokenSource) *oidc.UserInfo); ok {
		r0 = rf(ctx, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.UserInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, oauth2.TokenSource) error); ok {
		r1 = rf(ctx, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OIDCClient_UserInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfo'
type OIDCClient_UserInfo_Call struct {
	*mock.Call
}

// UserInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - ts oauth2.TokenSource
func (_e *OIDCClient_Expecter) UserInfo(ctx interface{}, ts interface{}) *OIDCClient_UserInfo_Call {
	return &OIDCClient_UserInfo_Call{Call: _e.mock.On("UserInfo", ctx, ts)}
}

func (_c *OIDCClient_UserInfo_Call) Run(run func(ctx context.Context, ts oauth2.TokenSource)) *OIDCClient_UserInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oauth2.TokenSource))
	})
	return _c
}

func (_c *OIDCClient_UserInfo_Call) Return(_a0 *oidc.UserInfo, _a1 error) *OIDCClient_UserInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OIDCClient_UserInfo_Call) RunAndReturn(run func(context.Context, oauth2.TokenSource) (*oidc.UserInfo, error)) *OIDCClient_UserInfo_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyAccessToken provides a mock function with given fields: ctx, token
func (_m *OIDCClient) VerifyAccessToken(ctx context.Context, token string) (oidc.RegClaimsWithSID, jwt.MapClaims, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyAccessToken")
	}

	var r0 oidc.RegClaimsWithSID
	var r1 jwt.MapClaims
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (oidc.RegClaimsWithSID, jwt.MapClaims, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) oidc.RegClaimsWithSID); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(oidc.RegClaimsWithSID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) jwt.MapClaims); ok {
		r1 = rf(ctx, token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(jwt.MapClaims)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OIDCClient_VerifyAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyAccessToken'
type OIDCClient_VerifyAccessToken_Call struct {
	*mock.Call
}

// VerifyAccessToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *OIDCClient_Expecter) VerifyAccessToken(ctx interface{}, token interface{}) *OIDCClient_VerifyAccessToken_Call {
	return &OIDCClient_VerifyAccessToken_Call{Call: _e.mock.On("VerifyAccessToken", ctx, token)}
}

func (_c *OIDCClient_VerifyAccessToken_Call) Run(run func(ctx context.Context, token string)) *OIDCClient_VerifyAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OIDCClient_VerifyAccessToken_Call) Return(_a0 oidc.RegClaimsWithSID, _a1 jwt.MapClaims, _a2 error) *OIDCClient_VerifyAccessToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *OIDCClient_VerifyAccessToken_Call) RunAndReturn(run func(context.Context, string) (oidc.RegClaimsWithSID, jwt.MapClaims, error)) *OIDCClient_VerifyAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyLogoutToken provides a mock function with given fields: ctx, token
func (_m *OIDCClient) VerifyLogoutToken(ctx context.Context, token string) (*oidc.LogoutToken, error) {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyLogoutToken")
	}

	var r0 *oidc.LogoutToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*oidc.LogoutToken, error)); ok {
		return rf(ctx, token)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *oidc.LogoutToken); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oidc.LogoutToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OIDCClient_VerifyLogoutToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyLogoutToken'
type OIDCClient_VerifyLogoutToken_Call struct {
	*mock.Call
}

// VerifyLogoutToken is a helper method to define mock.On call
//   - ctx context.Context
//   - token string
func (_e *OIDCClient_Expecter) VerifyLogoutToken(ctx interface{}, token interface{}) *OIDCClient_VerifyLogoutToken_Call {
	return &OIDCClient_VerifyLogoutToken_Call{Call: _e.mock.On("VerifyLogoutToken", ctx, token)}
}

func (_c *OIDCClient_VerifyLogoutToken_Call) Run(run func(ctx context.Context, token string)) *OIDCClient_VerifyLogoutToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OIDCClient_VerifyLogoutToken_Call) Return(_a0 *oidc.LogoutToken, _a1 error) *OIDCClient_VerifyLogoutToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OIDCClient_VerifyLogoutToken_Call) RunAndReturn(run func(context.Context, string) (*oidc.LogoutToken, error)) *OIDCClient_VerifyLogoutToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewOIDCClient creates a new instance of OIDCClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOIDCClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *OIDCClient {
	mock := &OIDCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

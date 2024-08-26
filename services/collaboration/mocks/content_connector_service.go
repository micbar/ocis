// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	connector "github.com/owncloud/ocis/v2/services/collaboration/pkg/connector"

	http "net/http"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// ContentConnectorService is an autogenerated mock type for the ContentConnectorService type
type ContentConnectorService struct {
	mock.Mock
}

type ContentConnectorService_Expecter struct {
	mock *mock.Mock
}

func (_m *ContentConnectorService) EXPECT() *ContentConnectorService_Expecter {
	return &ContentConnectorService_Expecter{mock: &_m.Mock}
}

// GetFile provides a mock function with given fields: ctx, w
func (_m *ContentConnectorService) GetFile(ctx context.Context, w http.ResponseWriter) error {
	ret := _m.Called(ctx, w)

	if len(ret) == 0 {
		panic("no return value specified for GetFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, http.ResponseWriter) error); ok {
		r0 = rf(ctx, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContentConnectorService_GetFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFile'
type ContentConnectorService_GetFile_Call struct {
	*mock.Call
}

// GetFile is a helper method to define mock.On call
//   - ctx context.Context
//   - w http.ResponseWriter
func (_e *ContentConnectorService_Expecter) GetFile(ctx interface{}, w interface{}) *ContentConnectorService_GetFile_Call {
	return &ContentConnectorService_GetFile_Call{Call: _e.mock.On("GetFile", ctx, w)}
}

func (_c *ContentConnectorService_GetFile_Call) Run(run func(ctx context.Context, w http.ResponseWriter)) *ContentConnectorService_GetFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(http.ResponseWriter))
	})
	return _c
}

func (_c *ContentConnectorService_GetFile_Call) Return(_a0 error) *ContentConnectorService_GetFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContentConnectorService_GetFile_Call) RunAndReturn(run func(context.Context, http.ResponseWriter) error) *ContentConnectorService_GetFile_Call {
	_c.Call.Return(run)
	return _c
}

// PutFile provides a mock function with given fields: ctx, stream, streamLength, lockID
func (_m *ContentConnectorService) PutFile(ctx context.Context, stream io.Reader, streamLength int64, lockID string) (*connector.ConnectorResponse, error) {
	ret := _m.Called(ctx, stream, streamLength, lockID)

	if len(ret) == 0 {
		panic("no return value specified for PutFile")
	}

	var r0 *connector.ConnectorResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, int64, string) (*connector.ConnectorResponse, error)); ok {
		return rf(ctx, stream, streamLength, lockID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, int64, string) *connector.ConnectorResponse); ok {
		r0 = rf(ctx, stream, streamLength, lockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connector.ConnectorResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, int64, string) error); ok {
		r1 = rf(ctx, stream, streamLength, lockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContentConnectorService_PutFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutFile'
type ContentConnectorService_PutFile_Call struct {
	*mock.Call
}

// PutFile is a helper method to define mock.On call
//   - ctx context.Context
//   - stream io.Reader
//   - streamLength int64
//   - lockID string
func (_e *ContentConnectorService_Expecter) PutFile(ctx interface{}, stream interface{}, streamLength interface{}, lockID interface{}) *ContentConnectorService_PutFile_Call {
	return &ContentConnectorService_PutFile_Call{Call: _e.mock.On("PutFile", ctx, stream, streamLength, lockID)}
}

func (_c *ContentConnectorService_PutFile_Call) Run(run func(ctx context.Context, stream io.Reader, streamLength int64, lockID string)) *ContentConnectorService_PutFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(io.Reader), args[2].(int64), args[3].(string))
	})
	return _c
}

func (_c *ContentConnectorService_PutFile_Call) Return(_a0 *connector.ConnectorResponse, _a1 error) *ContentConnectorService_PutFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContentConnectorService_PutFile_Call) RunAndReturn(run func(context.Context, io.Reader, int64, string) (*connector.ConnectorResponse, error)) *ContentConnectorService_PutFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewContentConnectorService creates a new instance of ContentConnectorService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContentConnectorService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContentConnectorService {
	mock := &ContentConnectorService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	collaborationv1beta1 "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"

	libregraph "github.com/owncloud/libre-graph-api-go"

	mock "github.com/stretchr/testify/mock"

	ocmv1beta1 "github.com/cs3org/go-cs3apis/cs3/sharing/ocm/v1beta1"
)

// BaseGraphProvider is an autogenerated mock type for the BaseGraphProvider type
type BaseGraphProvider struct {
	mock.Mock
}

type BaseGraphProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *BaseGraphProvider) EXPECT() *BaseGraphProvider_Expecter {
	return &BaseGraphProvider_Expecter{mock: &_m.Mock}
}

// CS3ReceivedOCMSharesToDriveItems provides a mock function with given fields: ctx, receivedOCMShares
func (_m *BaseGraphProvider) CS3ReceivedOCMSharesToDriveItems(ctx context.Context, receivedOCMShares []*ocmv1beta1.ReceivedShare) ([]libregraph.DriveItem, error) {
	ret := _m.Called(ctx, receivedOCMShares)

	if len(ret) == 0 {
		panic("no return value specified for CS3ReceivedOCMSharesToDriveItems")
	}

	var r0 []libregraph.DriveItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*ocmv1beta1.ReceivedShare) ([]libregraph.DriveItem, error)); ok {
		return rf(ctx, receivedOCMShares)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*ocmv1beta1.ReceivedShare) []libregraph.DriveItem); ok {
		r0 = rf(ctx, receivedOCMShares)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]libregraph.DriveItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*ocmv1beta1.ReceivedShare) error); ok {
		r1 = rf(ctx, receivedOCMShares)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CS3ReceivedOCMSharesToDriveItems'
type BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call struct {
	*mock.Call
}

// CS3ReceivedOCMSharesToDriveItems is a helper method to define mock.On call
//   - ctx context.Context
//   - receivedOCMShares []*ocmv1beta1.ReceivedShare
func (_e *BaseGraphProvider_Expecter) CS3ReceivedOCMSharesToDriveItems(ctx interface{}, receivedOCMShares interface{}) *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call {
	return &BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call{Call: _e.mock.On("CS3ReceivedOCMSharesToDriveItems", ctx, receivedOCMShares)}
}

func (_c *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call) Run(run func(ctx context.Context, receivedOCMShares []*ocmv1beta1.ReceivedShare)) *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*ocmv1beta1.ReceivedShare))
	})
	return _c
}

func (_c *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call) Return(_a0 []libregraph.DriveItem, _a1 error) *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call) RunAndReturn(run func(context.Context, []*ocmv1beta1.ReceivedShare) ([]libregraph.DriveItem, error)) *BaseGraphProvider_CS3ReceivedOCMSharesToDriveItems_Call {
	_c.Call.Return(run)
	return _c
}

// CS3ReceivedSharesToDriveItems provides a mock function with given fields: ctx, receivedShares
func (_m *BaseGraphProvider) CS3ReceivedSharesToDriveItems(ctx context.Context, receivedShares []*collaborationv1beta1.ReceivedShare) ([]libregraph.DriveItem, error) {
	ret := _m.Called(ctx, receivedShares)

	if len(ret) == 0 {
		panic("no return value specified for CS3ReceivedSharesToDriveItems")
	}

	var r0 []libregraph.DriveItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*collaborationv1beta1.ReceivedShare) ([]libregraph.DriveItem, error)); ok {
		return rf(ctx, receivedShares)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*collaborationv1beta1.ReceivedShare) []libregraph.DriveItem); ok {
		r0 = rf(ctx, receivedShares)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]libregraph.DriveItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*collaborationv1beta1.ReceivedShare) error); ok {
		r1 = rf(ctx, receivedShares)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CS3ReceivedSharesToDriveItems'
type BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call struct {
	*mock.Call
}

// CS3ReceivedSharesToDriveItems is a helper method to define mock.On call
//   - ctx context.Context
//   - receivedShares []*collaborationv1beta1.ReceivedShare
func (_e *BaseGraphProvider_Expecter) CS3ReceivedSharesToDriveItems(ctx interface{}, receivedShares interface{}) *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call {
	return &BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call{Call: _e.mock.On("CS3ReceivedSharesToDriveItems", ctx, receivedShares)}
}

func (_c *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call) Run(run func(ctx context.Context, receivedShares []*collaborationv1beta1.ReceivedShare)) *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]*collaborationv1beta1.ReceivedShare))
	})
	return _c
}

func (_c *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call) Return(_a0 []libregraph.DriveItem, _a1 error) *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call) RunAndReturn(run func(context.Context, []*collaborationv1beta1.ReceivedShare) ([]libregraph.DriveItem, error)) *BaseGraphProvider_CS3ReceivedSharesToDriveItems_Call {
	_c.Call.Return(run)
	return _c
}

// NewBaseGraphProvider creates a new instance of BaseGraphProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBaseGraphProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *BaseGraphProvider {
	mock := &BaseGraphProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

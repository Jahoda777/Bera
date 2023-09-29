// Code generated by mockery v2.34.1. DO NOT EDIT.

package mocks

import (
	proto "github.com/cosmos/gogoproto/proto"
	mock "github.com/stretchr/testify/mock"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// SdkTx is an autogenerated mock type for the SdkTx type
type SdkTx struct {
	mock.Mock
}

type SdkTx_Expecter struct {
	mock *mock.Mock
}

func (_m *SdkTx) EXPECT() *SdkTx_Expecter {
	return &SdkTx_Expecter{mock: &_m.Mock}
}

// GetMsgs provides a mock function with given fields:
func (_m *SdkTx) GetMsgs() []proto.Message {
	ret := _m.Called()

	var r0 []proto.Message
	if rf, ok := ret.Get(0).(func() []proto.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]proto.Message)
		}
	}

	return r0
}

// SdkTx_GetMsgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMsgs'
type SdkTx_GetMsgs_Call struct {
	*mock.Call
}

// GetMsgs is a helper method to define mock.On call
func (_e *SdkTx_Expecter) GetMsgs() *SdkTx_GetMsgs_Call {
	return &SdkTx_GetMsgs_Call{Call: _e.mock.On("GetMsgs")}
}

func (_c *SdkTx_GetMsgs_Call) Run(run func()) *SdkTx_GetMsgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SdkTx_GetMsgs_Call) Return(_a0 []proto.Message) *SdkTx_GetMsgs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SdkTx_GetMsgs_Call) RunAndReturn(run func() []proto.Message) *SdkTx_GetMsgs_Call {
	_c.Call.Return(run)
	return _c
}

// GetMsgsV2 provides a mock function with given fields:
func (_m *SdkTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) {
	ret := _m.Called()

	var r0 []protoreflect.ProtoMessage
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]protoreflect.ProtoMessage, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []protoreflect.ProtoMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]protoreflect.ProtoMessage)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SdkTx_GetMsgsV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMsgsV2'
type SdkTx_GetMsgsV2_Call struct {
	*mock.Call
}

// GetMsgsV2 is a helper method to define mock.On call
func (_e *SdkTx_Expecter) GetMsgsV2() *SdkTx_GetMsgsV2_Call {
	return &SdkTx_GetMsgsV2_Call{Call: _e.mock.On("GetMsgsV2")}
}

func (_c *SdkTx_GetMsgsV2_Call) Run(run func()) *SdkTx_GetMsgsV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SdkTx_GetMsgsV2_Call) Return(_a0 []protoreflect.ProtoMessage, _a1 error) *SdkTx_GetMsgsV2_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SdkTx_GetMsgsV2_Call) RunAndReturn(run func() ([]protoreflect.ProtoMessage, error)) *SdkTx_GetMsgsV2_Call {
	_c.Call.Return(run)
	return _c
}

// NewSdkTx creates a new instance of SdkTx. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSdkTx(t interface {
	mock.TestingT
	Cleanup(func())
}) *SdkTx {
	mock := &SdkTx{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

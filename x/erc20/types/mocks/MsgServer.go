// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/sacasnetwork/sacas/v1/x/erc20/types"
	mock "github.com/stretchr/testify/mock"
)

// MsgServer is an autogenerated mock type for the MsgServer type
type MsgServer struct {
	mock.Mock
}

// ConvertERC20 provides a mock function with given fields: _a0, _a1
func (_m *MsgServer) ConvertERC20(_a0 context.Context, _a1 *types.MsgConvertERC20) (*types.MsgConvertERC20Response, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ConvertERC20")
	}

	var r0 *types.MsgConvertERC20Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.MsgConvertERC20) (*types.MsgConvertERC20Response, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.MsgConvertERC20) *types.MsgConvertERC20Response); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MsgConvertERC20Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.MsgConvertERC20) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateParams provides a mock function with given fields: _a0, _a1
func (_m *MsgServer) UpdateParams(_a0 context.Context, _a1 *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateParams")
	}

	var r0 *types.MsgUpdateParamsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.MsgUpdateParams) *types.MsgUpdateParamsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MsgUpdateParamsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.MsgUpdateParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMsgServer creates a new instance of MsgServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMsgServer(t interface {
	mock.TestingT
	Cleanup(func())
},
) *MsgServer {
	mock := &MsgServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

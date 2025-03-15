// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/sacasnetwork/sacas/v15/x/evm/types"
)

// EVMQueryClient is an autogenerated mock type for the EVMQueryClient type
type EVMQueryClient struct {
	mock.Mock
}

// Account provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) Account(ctx context.Context, in *types.QueryAccountRequest, opts ...grpc.CallOption) (*types.QueryAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAccountRequest, ...grpc.CallOption) *types.QueryAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Balance provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) Balance(ctx context.Context, in *types.QueryBalanceRequest, opts ...grpc.CallOption) (*types.QueryBalanceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryBalanceResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryBalanceRequest, ...grpc.CallOption) *types.QueryBalanceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryBalanceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryBalanceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BaseFee provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) BaseFee(ctx context.Context, in *types.QueryBaseFeeRequest, opts ...grpc.CallOption) (*types.QueryBaseFeeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryBaseFeeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryBaseFeeRequest, ...grpc.CallOption) *types.QueryBaseFeeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryBaseFeeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryBaseFeeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) Code(ctx context.Context, in *types.QueryCodeRequest, opts ...grpc.CallOption) (*types.QueryCodeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryCodeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryCodeRequest, ...grpc.CallOption) *types.QueryCodeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryCodeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryCodeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CosmosAccount provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) CosmosAccount(ctx context.Context, in *types.QueryCosmosAccountRequest, opts ...grpc.CallOption) (*types.QueryCosmosAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryCosmosAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryCosmosAccountRequest, ...grpc.CallOption) *types.QueryCosmosAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryCosmosAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryCosmosAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) EstimateGas(ctx context.Context, in *types.EthCallRequest, opts ...grpc.CallOption) (*types.EstimateGasResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.EstimateGasResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.EthCallRequest, ...grpc.CallOption) *types.EstimateGasResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.EstimateGasResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.EthCallRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EthCall provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) EthCall(ctx context.Context, in *types.EthCallRequest, opts ...grpc.CallOption) (*types.MsgEthereumTxResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.MsgEthereumTxResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.EthCallRequest, ...grpc.CallOption) *types.MsgEthereumTxResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.MsgEthereumTxResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.EthCallRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Params provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryParamsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) *types.QueryParamsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryParamsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryParamsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) Storage(ctx context.Context, in *types.QueryStorageRequest, opts ...grpc.CallOption) (*types.QueryStorageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryStorageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryStorageRequest, ...grpc.CallOption) *types.QueryStorageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryStorageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryStorageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TraceBlock provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) TraceBlock(ctx context.Context, in *types.QueryTraceBlockRequest, opts ...grpc.CallOption) (*types.QueryTraceBlockResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryTraceBlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTraceBlockRequest, ...grpc.CallOption) *types.QueryTraceBlockResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryTraceBlockResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryTraceBlockRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TraceTx provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) TraceTx(ctx context.Context, in *types.QueryTraceTxRequest, opts ...grpc.CallOption) (*types.QueryTraceTxResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryTraceTxResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryTraceTxRequest, ...grpc.CallOption) *types.QueryTraceTxResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryTraceTxResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryTraceTxRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatorAccount provides a mock function with given fields: ctx, in, opts
func (_m *EVMQueryClient) ValidatorAccount(ctx context.Context, in *types.QueryValidatorAccountRequest, opts ...grpc.CallOption) (*types.QueryValidatorAccountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryValidatorAccountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryValidatorAccountRequest, ...grpc.CallOption) *types.QueryValidatorAccountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryValidatorAccountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryValidatorAccountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEVMQueryClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewEVMQueryClient creates a new instance of EVMQueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEVMQueryClient(t mockConstructorTestingTNewEVMQueryClient) *EVMQueryClient {
	mock := &EVMQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

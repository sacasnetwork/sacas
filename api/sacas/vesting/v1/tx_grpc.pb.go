// SPDX-License-Identifier:ENCL-1.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sacas/vesting/v1/tx.proto

package vestingv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_CreateClawbackVestingAccount_FullMethodName = "/sacas.vesting.v1.Msg/CreateClawbackVestingAccount"
	Msg_Clawback_FullMethodName                     = "/sacas.vesting.v1.Msg/Clawback"
	Msg_UpdateVestingFunder_FullMethodName          = "/sacas.vesting.v1.Msg/UpdateVestingFunder"
	Msg_ConvertVestingAccount_FullMethodName        = "/sacas.vesting.v1.Msg/ConvertVestingAccount"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateClawbackVestingAccount creats a vesting account that is subject to
	// clawback and the configuration of vesting and lockup schedules.
	CreateClawbackVestingAccount(ctx context.Context, in *MsgCreateClawbackVestingAccount, opts ...grpc.CallOption) (*MsgCreateClawbackVestingAccountResponse, error)
	// Clawback removes the unvested tokens from a ClawbackVestingAccount.
	Clawback(ctx context.Context, in *MsgClawback, opts ...grpc.CallOption) (*MsgClawbackResponse, error)
	// UpdateVestingFunder updates the funder address of an existing
	// ClawbackVestingAccount.
	UpdateVestingFunder(ctx context.Context, in *MsgUpdateVestingFunder, opts ...grpc.CallOption) (*MsgUpdateVestingFunderResponse, error)
	// ConvertVestingAccount converts a ClawbackVestingAccount to a Eth account
	ConvertVestingAccount(ctx context.Context, in *MsgConvertVestingAccount, opts ...grpc.CallOption) (*MsgConvertVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateClawbackVestingAccount(ctx context.Context, in *MsgCreateClawbackVestingAccount, opts ...grpc.CallOption) (*MsgCreateClawbackVestingAccountResponse, error) {
	out := new(MsgCreateClawbackVestingAccountResponse)
	err := c.cc.Invoke(ctx, Msg_CreateClawbackVestingAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Clawback(ctx context.Context, in *MsgClawback, opts ...grpc.CallOption) (*MsgClawbackResponse, error) {
	out := new(MsgClawbackResponse)
	err := c.cc.Invoke(ctx, Msg_Clawback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateVestingFunder(ctx context.Context, in *MsgUpdateVestingFunder, opts ...grpc.CallOption) (*MsgUpdateVestingFunderResponse, error) {
	out := new(MsgUpdateVestingFunderResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateVestingFunder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConvertVestingAccount(ctx context.Context, in *MsgConvertVestingAccount, opts ...grpc.CallOption) (*MsgConvertVestingAccountResponse, error) {
	out := new(MsgConvertVestingAccountResponse)
	err := c.cc.Invoke(ctx, Msg_ConvertVestingAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateClawbackVestingAccount creats a vesting account that is subject to
	// clawback and the configuration of vesting and lockup schedules.
	CreateClawbackVestingAccount(context.Context, *MsgCreateClawbackVestingAccount) (*MsgCreateClawbackVestingAccountResponse, error)
	// Clawback removes the unvested tokens from a ClawbackVestingAccount.
	Clawback(context.Context, *MsgClawback) (*MsgClawbackResponse, error)
	// UpdateVestingFunder updates the funder address of an existing
	// ClawbackVestingAccount.
	UpdateVestingFunder(context.Context, *MsgUpdateVestingFunder) (*MsgUpdateVestingFunderResponse, error)
	// ConvertVestingAccount converts a ClawbackVestingAccount to a Eth account
	ConvertVestingAccount(context.Context, *MsgConvertVestingAccount) (*MsgConvertVestingAccountResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateClawbackVestingAccount(context.Context, *MsgCreateClawbackVestingAccount) (*MsgCreateClawbackVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClawbackVestingAccount not implemented")
}
func (UnimplementedMsgServer) Clawback(context.Context, *MsgClawback) (*MsgClawbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clawback not implemented")
}
func (UnimplementedMsgServer) UpdateVestingFunder(context.Context, *MsgUpdateVestingFunder) (*MsgUpdateVestingFunderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVestingFunder not implemented")
}
func (UnimplementedMsgServer) ConvertVestingAccount(context.Context, *MsgConvertVestingAccount) (*MsgConvertVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertVestingAccount not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateClawbackVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateClawbackVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateClawbackVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateClawbackVestingAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateClawbackVestingAccount(ctx, req.(*MsgCreateClawbackVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Clawback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClawback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Clawback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Clawback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Clawback(ctx, req.(*MsgClawback))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateVestingFunder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateVestingFunder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateVestingFunder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateVestingFunder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateVestingFunder(ctx, req.(*MsgUpdateVestingFunder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConvertVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ConvertVestingAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertVestingAccount(ctx, req.(*MsgConvertVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sacas.vesting.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClawbackVestingAccount",
			Handler:    _Msg_CreateClawbackVestingAccount_Handler,
		},
		{
			MethodName: "Clawback",
			Handler:    _Msg_Clawback_Handler,
		},
		{
			MethodName: "UpdateVestingFunder",
			Handler:    _Msg_UpdateVestingFunder_Handler,
		},
		{
			MethodName: "ConvertVestingAccount",
			Handler:    _Msg_ConvertVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sacas/vesting/v1/tx.proto",
}

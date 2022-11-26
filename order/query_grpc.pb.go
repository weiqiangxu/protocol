// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: query.proto

package order

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

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	GetOrderList(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*CommonReply, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, "/order.Login/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) GetOrderList(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*CommonReply, error) {
	out := new(CommonReply)
	err := c.cc.Invoke(ctx, "/order.Login/GetOrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	GetOrderList(context.Context, *GetOrderRequest) (*CommonReply, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedLoginServer) GetOrderList(context.Context, *GetOrderRequest) (*CommonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Login/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Login/GetOrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).GetOrderList(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _Login_GetOrder_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _Login_GetOrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}

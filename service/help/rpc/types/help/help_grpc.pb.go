// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: help.proto

package help

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

// HelpClient is the client API for Help service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelpClient interface {
	RegisterUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Reply, error)
	UpDateStatus(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Reply, error)
}

type helpClient struct {
	cc grpc.ClientConnInterface
}

func NewHelpClient(cc grpc.ClientConnInterface) HelpClient {
	return &helpClient{cc}
}

func (c *helpClient) RegisterUser(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/help.help/registerUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helpClient) UpDateStatus(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/help.help/upDateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelpServer is the server API for Help service.
// All implementations must embed UnimplementedHelpServer
// for forward compatibility
type HelpServer interface {
	RegisterUser(context.Context, *IdReq) (*Reply, error)
	UpDateStatus(context.Context, *IdReq) (*Reply, error)
	mustEmbedUnimplementedHelpServer()
}

// UnimplementedHelpServer must be embedded to have forward compatible implementations.
type UnimplementedHelpServer struct {
}

func (UnimplementedHelpServer) RegisterUser(context.Context, *IdReq) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedHelpServer) UpDateStatus(context.Context, *IdReq) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpDateStatus not implemented")
}
func (UnimplementedHelpServer) mustEmbedUnimplementedHelpServer() {}

// UnsafeHelpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelpServer will
// result in compilation errors.
type UnsafeHelpServer interface {
	mustEmbedUnimplementedHelpServer()
}

func RegisterHelpServer(s grpc.ServiceRegistrar, srv HelpServer) {
	s.RegisterService(&Help_ServiceDesc, srv)
}

func _Help_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/help.help/registerUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpServer).RegisterUser(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Help_UpDateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelpServer).UpDateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/help.help/upDateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelpServer).UpDateStatus(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Help_ServiceDesc is the grpc.ServiceDesc for Help service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Help_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "help.help",
	HandlerType: (*HelpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerUser",
			Handler:    _Help_RegisterUser_Handler,
		},
		{
			MethodName: "upDateStatus",
			Handler:    _Help_UpDateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "help.proto",
}

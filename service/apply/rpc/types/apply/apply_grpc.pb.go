// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: apply.proto

package apply

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

// ApplyClient is the client API for Apply service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApplyClient interface {
	CreateIdentify(ctx context.Context, in *CreateIdentifyReq, opts ...grpc.CallOption) (*CreateIdentifyReply, error)
	CheckIdentify(ctx context.Context, in *CheckIdentifyReq, opts ...grpc.CallOption) (*CheckIdentifyReply, error)
	CheckUser(ctx context.Context, in *CheckUserReq, opts ...grpc.CallOption) (*CheckUserReply, error)
}

type applyClient struct {
	cc grpc.ClientConnInterface
}

func NewApplyClient(cc grpc.ClientConnInterface) ApplyClient {
	return &applyClient{cc}
}

func (c *applyClient) CreateIdentify(ctx context.Context, in *CreateIdentifyReq, opts ...grpc.CallOption) (*CreateIdentifyReply, error) {
	out := new(CreateIdentifyReply)
	err := c.cc.Invoke(ctx, "/apply.Apply/CreateIdentify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applyClient) CheckIdentify(ctx context.Context, in *CheckIdentifyReq, opts ...grpc.CallOption) (*CheckIdentifyReply, error) {
	out := new(CheckIdentifyReply)
	err := c.cc.Invoke(ctx, "/apply.Apply/CheckIdentify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applyClient) CheckUser(ctx context.Context, in *CheckUserReq, opts ...grpc.CallOption) (*CheckUserReply, error) {
	out := new(CheckUserReply)
	err := c.cc.Invoke(ctx, "/apply.Apply/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplyServer is the server API for Apply service.
// All implementations must embed UnimplementedApplyServer
// for forward compatibility
type ApplyServer interface {
	CreateIdentify(context.Context, *CreateIdentifyReq) (*CreateIdentifyReply, error)
	CheckIdentify(context.Context, *CheckIdentifyReq) (*CheckIdentifyReply, error)
	CheckUser(context.Context, *CheckUserReq) (*CheckUserReply, error)
	mustEmbedUnimplementedApplyServer()
}

// UnimplementedApplyServer must be embedded to have forward compatible implementations.
type UnimplementedApplyServer struct {
}

func (UnimplementedApplyServer) CreateIdentify(context.Context, *CreateIdentifyReq) (*CreateIdentifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdentify not implemented")
}
func (UnimplementedApplyServer) CheckIdentify(context.Context, *CheckIdentifyReq) (*CheckIdentifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIdentify not implemented")
}
func (UnimplementedApplyServer) CheckUser(context.Context, *CheckUserReq) (*CheckUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (UnimplementedApplyServer) mustEmbedUnimplementedApplyServer() {}

// UnsafeApplyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApplyServer will
// result in compilation errors.
type UnsafeApplyServer interface {
	mustEmbedUnimplementedApplyServer()
}

func RegisterApplyServer(s grpc.ServiceRegistrar, srv ApplyServer) {
	s.RegisterService(&Apply_ServiceDesc, srv)
}

func _Apply_CreateIdentify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIdentifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplyServer).CreateIdentify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apply.Apply/CreateIdentify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplyServer).CreateIdentify(ctx, req.(*CreateIdentifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apply_CheckIdentify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIdentifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplyServer).CheckIdentify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apply.Apply/CheckIdentify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplyServer).CheckIdentify(ctx, req.(*CheckIdentifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Apply_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplyServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apply.Apply/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplyServer).CheckUser(ctx, req.(*CheckUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Apply_ServiceDesc is the grpc.ServiceDesc for Apply service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Apply_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apply.Apply",
	HandlerType: (*ApplyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdentify",
			Handler:    _Apply_CreateIdentify_Handler,
		},
		{
			MethodName: "CheckIdentify",
			Handler:    _Apply_CheckIdentify_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _Apply_CheckUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apply.proto",
}

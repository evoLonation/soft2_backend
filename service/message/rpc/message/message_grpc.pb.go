// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: message.proto

package message

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
	CreateMessage(ctx context.Context, in *CreateMessageReq, opts ...grpc.CallOption) (*CreateMessageReply, error)
}

type applyClient struct {
	cc grpc.ClientConnInterface
}

func NewApplyClient(cc grpc.ClientConnInterface) ApplyClient {
	return &applyClient{cc}
}

func (c *applyClient) CreateMessage(ctx context.Context, in *CreateMessageReq, opts ...grpc.CallOption) (*CreateMessageReply, error) {
	out := new(CreateMessageReply)
	err := c.cc.Invoke(ctx, "/message.Apply/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplyServer is the server API for Apply service.
// All implementations must embed UnimplementedApplyServer
// for forward compatibility
type ApplyServer interface {
	CreateMessage(context.Context, *CreateMessageReq) (*CreateMessageReply, error)
	mustEmbedUnimplementedApplyServer()
}

// UnimplementedApplyServer must be embedded to have forward compatible implementations.
type UnimplementedApplyServer struct {
}

func (UnimplementedApplyServer) CreateMessage(context.Context, *CreateMessageReq) (*CreateMessageReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
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

func _Apply_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplyServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Apply/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplyServer).CreateMessage(ctx, req.(*CreateMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Apply_ServiceDesc is the grpc.ServiceDesc for Apply service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Apply_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.Apply",
	HandlerType: (*ApplyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _Apply_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

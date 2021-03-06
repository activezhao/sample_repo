// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stubs

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

// Zhaoh0311Client is the client API for Zhaoh0311 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Zhaoh0311Client interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type zhaoh0311Client struct {
	cc grpc.ClientConnInterface
}

func NewZhaoh0311Client(cc grpc.ClientConnInterface) Zhaoh0311Client {
	return &zhaoh0311Client{cc}
}

func (c *zhaoh0311Client) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/sample_module_0311.sample_package_0311.Zhaoh0311/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Zhaoh0311Server is the server API for Zhaoh0311 service.
// All implementations must embed UnimplementedZhaoh0311Server
// for forward compatibility
type Zhaoh0311Server interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedZhaoh0311Server()
}

// UnimplementedZhaoh0311Server must be embedded to have forward compatible implementations.
type UnimplementedZhaoh0311Server struct {
}

func (UnimplementedZhaoh0311Server) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedZhaoh0311Server) mustEmbedUnimplementedZhaoh0311Server() {}

// UnsafeZhaoh0311Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Zhaoh0311Server will
// result in compilation errors.
type UnsafeZhaoh0311Server interface {
	mustEmbedUnimplementedZhaoh0311Server()
}

func RegisterZhaoh0311Server(s grpc.ServiceRegistrar, srv Zhaoh0311Server) {
	s.RegisterService(&Zhaoh0311_ServiceDesc, srv)
}

func _Zhaoh0311_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Zhaoh0311Server).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample_module_0311.sample_package_0311.Zhaoh0311/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Zhaoh0311Server).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Zhaoh0311_ServiceDesc is the grpc.ServiceDesc for Zhaoh0311 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Zhaoh0311_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sample_module_0311.sample_package_0311.Zhaoh0311",
	HandlerType: (*Zhaoh0311Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Zhaoh0311_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Zhaoh0311.proto",
}

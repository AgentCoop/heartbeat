// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package heartbeat

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

// HeartBeatClient is the client API for HeartBeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartBeatClient interface {
	Foo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type heartBeatClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartBeatClient(cc grpc.ClientConnInterface) HeartBeatClient {
	return &heartBeatClient{cc}
}

func (c *heartBeatClient) Foo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/heartbeat.utils.HeartBeat/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeatServer is the server API for HeartBeat service.
// All implementations must embed UnimplementedHeartBeatServer
// for forward compatibility
type HeartBeatServer interface {
	Foo(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedHeartBeatServer()
}

// UnimplementedHeartBeatServer must be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServer struct {
}

func (UnimplementedHeartBeatServer) Foo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (UnimplementedHeartBeatServer) mustEmbedUnimplementedHeartBeatServer() {}

// UnsafeHeartBeatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartBeatServer will
// result in compilation errors.
type UnsafeHeartBeatServer interface {
	mustEmbedUnimplementedHeartBeatServer()
}

func RegisterHeartBeatServer(s grpc.ServiceRegistrar, srv HeartBeatServer) {
	s.RegisterService(&HeartBeat_ServiceDesc, srv)
}

func _HeartBeat_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heartbeat.utils.HeartBeat/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServer).Foo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartBeat_ServiceDesc is the grpc.ServiceDesc for HeartBeat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartBeat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heartbeat.utils.HeartBeat",
	HandlerType: (*HeartBeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _HeartBeat_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}
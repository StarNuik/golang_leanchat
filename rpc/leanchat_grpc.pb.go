// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: rpc/leanchat.proto

package rpc

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
	Leanchat_PeekChannel_FullMethodName    = "/Leanchat/PeekChannel"
	Leanchat_MessageChannel_FullMethodName = "/Leanchat/MessageChannel"
	Leanchat_ListChannels_FullMethodName   = "/Leanchat/ListChannels"
	Leanchat_CreateChannel_FullMethodName  = "/Leanchat/CreateChannel"
)

// LeanchatClient is the client API for Leanchat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeanchatClient interface {
	PeekChannel(ctx context.Context, in *PeekChannelRequest, opts ...grpc.CallOption) (*PeekChannelResponse, error)
	MessageChannel(ctx context.Context, in *MessageChannelRequest, opts ...grpc.CallOption) (*MessageChannelResponse, error)
	ListChannels(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error)
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
}

type leanchatClient struct {
	cc grpc.ClientConnInterface
}

func NewLeanchatClient(cc grpc.ClientConnInterface) LeanchatClient {
	return &leanchatClient{cc}
}

func (c *leanchatClient) PeekChannel(ctx context.Context, in *PeekChannelRequest, opts ...grpc.CallOption) (*PeekChannelResponse, error) {
	out := new(PeekChannelResponse)
	err := c.cc.Invoke(ctx, Leanchat_PeekChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leanchatClient) MessageChannel(ctx context.Context, in *MessageChannelRequest, opts ...grpc.CallOption) (*MessageChannelResponse, error) {
	out := new(MessageChannelResponse)
	err := c.cc.Invoke(ctx, Leanchat_MessageChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leanchatClient) ListChannels(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error) {
	out := new(ListChannelsResponse)
	err := c.cc.Invoke(ctx, Leanchat_ListChannels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leanchatClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, Leanchat_CreateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeanchatServer is the server API for Leanchat service.
// All implementations must embed UnimplementedLeanchatServer
// for forward compatibility
type LeanchatServer interface {
	PeekChannel(context.Context, *PeekChannelRequest) (*PeekChannelResponse, error)
	MessageChannel(context.Context, *MessageChannelRequest) (*MessageChannelResponse, error)
	ListChannels(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error)
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	mustEmbedUnimplementedLeanchatServer()
}

// UnimplementedLeanchatServer must be embedded to have forward compatible implementations.
type UnimplementedLeanchatServer struct {
}

func (UnimplementedLeanchatServer) PeekChannel(context.Context, *PeekChannelRequest) (*PeekChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeekChannel not implemented")
}
func (UnimplementedLeanchatServer) MessageChannel(context.Context, *MessageChannelRequest) (*MessageChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageChannel not implemented")
}
func (UnimplementedLeanchatServer) ListChannels(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}
func (UnimplementedLeanchatServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedLeanchatServer) mustEmbedUnimplementedLeanchatServer() {}

// UnsafeLeanchatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeanchatServer will
// result in compilation errors.
type UnsafeLeanchatServer interface {
	mustEmbedUnimplementedLeanchatServer()
}

func RegisterLeanchatServer(s grpc.ServiceRegistrar, srv LeanchatServer) {
	s.RegisterService(&Leanchat_ServiceDesc, srv)
}

func _Leanchat_PeekChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeekChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeanchatServer).PeekChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leanchat_PeekChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeanchatServer).PeekChannel(ctx, req.(*PeekChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leanchat_MessageChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeanchatServer).MessageChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leanchat_MessageChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeanchatServer).MessageChannel(ctx, req.(*MessageChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leanchat_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeanchatServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leanchat_ListChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeanchatServer).ListChannels(ctx, req.(*ListChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Leanchat_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeanchatServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leanchat_CreateChannel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeanchatServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Leanchat_ServiceDesc is the grpc.ServiceDesc for Leanchat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leanchat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Leanchat",
	HandlerType: (*LeanchatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PeekChannel",
			Handler:    _Leanchat_PeekChannel_Handler,
		},
		{
			MethodName: "MessageChannel",
			Handler:    _Leanchat_MessageChannel_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _Leanchat_ListChannels_Handler,
		},
		{
			MethodName: "CreateChannel",
			Handler:    _Leanchat_CreateChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/leanchat.proto",
}

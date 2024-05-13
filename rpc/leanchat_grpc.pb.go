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
	Leanchat_CreateChannel_FullMethodName  = "/Leanchat/CreateChannel"
	Leanchat_SearchChannels_FullMethodName = "/Leanchat/SearchChannels"
	Leanchat_PeekChannel_FullMethodName    = "/Leanchat/PeekChannel"
	Leanchat_MessageChannel_FullMethodName = "/Leanchat/MessageChannel"
)

// LeanchatClient is the client API for Leanchat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeanchatClient interface {
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error)
	SearchChannels(ctx context.Context, in *SearchChannelsRequest, opts ...grpc.CallOption) (*SearchChannelsResponse, error)
	PeekChannel(ctx context.Context, in *PeekChannelRequest, opts ...grpc.CallOption) (*PeekChannelResponse, error)
	MessageChannel(ctx context.Context, in *MessageChannelRequest, opts ...grpc.CallOption) (*MessageChannelResponse, error)
}

type leanchatClient struct {
	cc grpc.ClientConnInterface
}

func NewLeanchatClient(cc grpc.ClientConnInterface) LeanchatClient {
	return &leanchatClient{cc}
}

func (c *leanchatClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*CreateChannelResponse, error) {
	out := new(CreateChannelResponse)
	err := c.cc.Invoke(ctx, Leanchat_CreateChannel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leanchatClient) SearchChannels(ctx context.Context, in *SearchChannelsRequest, opts ...grpc.CallOption) (*SearchChannelsResponse, error) {
	out := new(SearchChannelsResponse)
	err := c.cc.Invoke(ctx, Leanchat_SearchChannels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

// LeanchatServer is the server API for Leanchat service.
// All implementations must embed UnimplementedLeanchatServer
// for forward compatibility
type LeanchatServer interface {
	CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error)
	SearchChannels(context.Context, *SearchChannelsRequest) (*SearchChannelsResponse, error)
	PeekChannel(context.Context, *PeekChannelRequest) (*PeekChannelResponse, error)
	MessageChannel(context.Context, *MessageChannelRequest) (*MessageChannelResponse, error)
	mustEmbedUnimplementedLeanchatServer()
}

// UnimplementedLeanchatServer must be embedded to have forward compatible implementations.
type UnimplementedLeanchatServer struct {
}

func (UnimplementedLeanchatServer) CreateChannel(context.Context, *CreateChannelRequest) (*CreateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedLeanchatServer) SearchChannels(context.Context, *SearchChannelsRequest) (*SearchChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChannels not implemented")
}
func (UnimplementedLeanchatServer) PeekChannel(context.Context, *PeekChannelRequest) (*PeekChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeekChannel not implemented")
}
func (UnimplementedLeanchatServer) MessageChannel(context.Context, *MessageChannelRequest) (*MessageChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageChannel not implemented")
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

func _Leanchat_SearchChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeanchatServer).SearchChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Leanchat_SearchChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeanchatServer).SearchChannels(ctx, req.(*SearchChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
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

// Leanchat_ServiceDesc is the grpc.ServiceDesc for Leanchat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Leanchat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Leanchat",
	HandlerType: (*LeanchatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _Leanchat_CreateChannel_Handler,
		},
		{
			MethodName: "SearchChannels",
			Handler:    _Leanchat_SearchChannels_Handler,
		},
		{
			MethodName: "PeekChannel",
			Handler:    _Leanchat_PeekChannel_Handler,
		},
		{
			MethodName: "MessageChannel",
			Handler:    _Leanchat_MessageChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/leanchat.proto",
}

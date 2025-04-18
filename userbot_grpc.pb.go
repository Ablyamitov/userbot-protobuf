// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc3
// source: userbot.proto

package userbotpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserbotService_GetMessages_FullMethodName = "/userbot.UserbotService/GetMessages"
	UserbotService_SendMessage_FullMethodName = "/userbot.UserbotService/SendMessage"
)

// UserbotServiceClient is the client API for UserbotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserbotServiceClient interface {
	GetMessages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (*MessagesResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type userbotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserbotServiceClient(cc grpc.ClientConnInterface) UserbotServiceClient {
	return &userbotServiceClient{cc}
}

func (c *userbotServiceClient) GetMessages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (*MessagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MessagesResponse)
	err := c.cc.Invoke(ctx, UserbotService_GetMessages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userbotServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, UserbotService_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserbotServiceServer is the server API for UserbotService service.
// All implementations must embed UnimplementedUserbotServiceServer
// for forward compatibility.
type UserbotServiceServer interface {
	GetMessages(context.Context, *MessagesRequest) (*MessagesResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	mustEmbedUnimplementedUserbotServiceServer()
}

// UnimplementedUserbotServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserbotServiceServer struct{}

func (UnimplementedUserbotServiceServer) GetMessages(context.Context, *MessagesRequest) (*MessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedUserbotServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedUserbotServiceServer) mustEmbedUnimplementedUserbotServiceServer() {}
func (UnimplementedUserbotServiceServer) testEmbeddedByValue()                        {}

// UnsafeUserbotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserbotServiceServer will
// result in compilation errors.
type UnsafeUserbotServiceServer interface {
	mustEmbedUnimplementedUserbotServiceServer()
}

func RegisterUserbotServiceServer(s grpc.ServiceRegistrar, srv UserbotServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserbotServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserbotService_ServiceDesc, srv)
}

func _UserbotService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserbotServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserbotService_GetMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserbotServiceServer).GetMessages(ctx, req.(*MessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserbotService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserbotServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserbotService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserbotServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserbotService_ServiceDesc is the grpc.ServiceDesc for UserbotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserbotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userbot.UserbotService",
	HandlerType: (*UserbotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessages",
			Handler:    _UserbotService_GetMessages_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _UserbotService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userbot.proto",
}

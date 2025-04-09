// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: chat.proto

package chat

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
	Chat_CreateConnect_FullMethodName = "/Chat/CreateConnect"
	Chat_SendMsg_FullMethodName       = "/Chat/SendMsg"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	CreateConnect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MessageResponse], error)
	SendMsg(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*None, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) CreateConnect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MessageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], Chat_CreateConnect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ConnectRequest, MessageResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Chat_CreateConnectClient = grpc.ServerStreamingClient[MessageResponse]

func (c *chatClient) SendMsg(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*None, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(None)
	err := c.cc.Invoke(ctx, Chat_SendMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility.
type ChatServer interface {
	CreateConnect(*ConnectRequest, grpc.ServerStreamingServer[MessageResponse]) error
	SendMsg(context.Context, *MessageRequest) (*None, error)
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServer struct{}

func (UnimplementedChatServer) CreateConnect(*ConnectRequest, grpc.ServerStreamingServer[MessageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method CreateConnect not implemented")
}
func (UnimplementedChatServer) SendMsg(context.Context, *MessageRequest) (*None, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}
func (UnimplementedChatServer) testEmbeddedByValue()              {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	// If the following call pancis, it indicates UnimplementedChatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_CreateConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).CreateConnect(m, &grpc.GenericServerStream[ConnectRequest, MessageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Chat_CreateConnectServer = grpc.ServerStreamingServer[MessageResponse]

func _Chat_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chat_SendMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMsg(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _Chat_SendMsg_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateConnect",
			Handler:       _Chat_CreateConnect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}

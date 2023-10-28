// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: event/v1/archive.proto

package event

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArchiveEventServiceClient is the client API for ArchiveEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchiveEventServiceClient interface {
	PublishEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PublishEventResponse, error)
	ReceiveEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ArchiveEventService_ReceiveEventClient, error)
}

type archiveEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveEventServiceClient(cc grpc.ClientConnInterface) ArchiveEventServiceClient {
	return &archiveEventServiceClient{cc}
}

func (c *archiveEventServiceClient) PublishEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PublishEventResponse, error) {
	out := new(PublishEventResponse)
	err := c.cc.Invoke(ctx, "/satellite.event.v1.ArchiveEventService/PublishEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *archiveEventServiceClient) ReceiveEvent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ArchiveEventService_ReceiveEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArchiveEventService_ServiceDesc.Streams[0], "/satellite.event.v1.ArchiveEventService/ReceiveEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &archiveEventServiceReceiveEventClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArchiveEventService_ReceiveEventClient interface {
	Recv() (*ArchiveEvent, error)
	grpc.ClientStream
}

type archiveEventServiceReceiveEventClient struct {
	grpc.ClientStream
}

func (x *archiveEventServiceReceiveEventClient) Recv() (*ArchiveEvent, error) {
	m := new(ArchiveEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArchiveEventServiceServer is the server API for ArchiveEventService service.
// All implementations must embed UnimplementedArchiveEventServiceServer
// for forward compatibility
type ArchiveEventServiceServer interface {
	PublishEvent(context.Context, *emptypb.Empty) (*PublishEventResponse, error)
	ReceiveEvent(*emptypb.Empty, ArchiveEventService_ReceiveEventServer) error
	mustEmbedUnimplementedArchiveEventServiceServer()
}

// UnimplementedArchiveEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArchiveEventServiceServer struct {
}

func (UnimplementedArchiveEventServiceServer) PublishEvent(context.Context, *emptypb.Empty) (*PublishEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishEvent not implemented")
}
func (UnimplementedArchiveEventServiceServer) ReceiveEvent(*emptypb.Empty, ArchiveEventService_ReceiveEventServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveEvent not implemented")
}
func (UnimplementedArchiveEventServiceServer) mustEmbedUnimplementedArchiveEventServiceServer() {}

// UnsafeArchiveEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchiveEventServiceServer will
// result in compilation errors.
type UnsafeArchiveEventServiceServer interface {
	mustEmbedUnimplementedArchiveEventServiceServer()
}

func RegisterArchiveEventServiceServer(s grpc.ServiceRegistrar, srv ArchiveEventServiceServer) {
	s.RegisterService(&ArchiveEventService_ServiceDesc, srv)
}

func _ArchiveEventService_PublishEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveEventServiceServer).PublishEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/satellite.event.v1.ArchiveEventService/PublishEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveEventServiceServer).PublishEvent(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArchiveEventService_ReceiveEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArchiveEventServiceServer).ReceiveEvent(m, &archiveEventServiceReceiveEventServer{stream})
}

type ArchiveEventService_ReceiveEventServer interface {
	Send(*ArchiveEvent) error
	grpc.ServerStream
}

type archiveEventServiceReceiveEventServer struct {
	grpc.ServerStream
}

func (x *archiveEventServiceReceiveEventServer) Send(m *ArchiveEvent) error {
	return x.ServerStream.SendMsg(m)
}

// ArchiveEventService_ServiceDesc is the grpc.ServiceDesc for ArchiveEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArchiveEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "satellite.event.v1.ArchiveEventService",
	HandlerType: (*ArchiveEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishEvent",
			Handler:    _ArchiveEventService_PublishEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveEvent",
			Handler:       _ArchiveEventService_ReceiveEvent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event/v1/archive.proto",
}

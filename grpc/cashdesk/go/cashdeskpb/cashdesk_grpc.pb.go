// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: grpc/cashdesk/cashdesk.proto

package cashdeskpb

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

const (
	Cashdesk_StreamEvents_FullMethodName        = "/net.vseinstrumenti.git.fd.proto.cashdesk.Cashdesk/StreamEvents"
	Cashdesk_SendWorkstationLogs_FullMethodName = "/net.vseinstrumenti.git.fd.proto.cashdesk.Cashdesk/SendWorkstationLogs"
	Cashdesk_SendEventResult_FullMethodName     = "/net.vseinstrumenti.git.fd.proto.cashdesk.Cashdesk/SendEventResult"
)

// CashdeskClient is the client API for Cashdesk service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CashdeskClient interface {
	StreamEvents(ctx context.Context, opts ...grpc.CallOption) (Cashdesk_StreamEventsClient, error)
	SendWorkstationLogs(ctx context.Context, in *SendWorkstationLogsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendEventResult(ctx context.Context, in *SendEventResultRequest, opts ...grpc.CallOption) (*SendEventResultResponse, error)
}

type cashdeskClient struct {
	cc grpc.ClientConnInterface
}

func NewCashdeskClient(cc grpc.ClientConnInterface) CashdeskClient {
	return &cashdeskClient{cc}
}

func (c *cashdeskClient) StreamEvents(ctx context.Context, opts ...grpc.CallOption) (Cashdesk_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cashdesk_ServiceDesc.Streams[0], Cashdesk_StreamEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &cashdeskStreamEventsClient{stream}
	return x, nil
}

type Cashdesk_StreamEventsClient interface {
	Send(*StreamEventsAck) error
	Recv() (*StreamWorkstationEventsResponse, error)
	grpc.ClientStream
}

type cashdeskStreamEventsClient struct {
	grpc.ClientStream
}

func (x *cashdeskStreamEventsClient) Send(m *StreamEventsAck) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cashdeskStreamEventsClient) Recv() (*StreamWorkstationEventsResponse, error) {
	m := new(StreamWorkstationEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cashdeskClient) SendWorkstationLogs(ctx context.Context, in *SendWorkstationLogsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Cashdesk_SendWorkstationLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashdeskClient) SendEventResult(ctx context.Context, in *SendEventResultRequest, opts ...grpc.CallOption) (*SendEventResultResponse, error) {
	out := new(SendEventResultResponse)
	err := c.cc.Invoke(ctx, Cashdesk_SendEventResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CashdeskServer is the server API for Cashdesk service.
// All implementations must embed UnimplementedCashdeskServer
// for forward compatibility
type CashdeskServer interface {
	StreamEvents(Cashdesk_StreamEventsServer) error
	SendWorkstationLogs(context.Context, *SendWorkstationLogsRequest) (*emptypb.Empty, error)
	SendEventResult(context.Context, *SendEventResultRequest) (*SendEventResultResponse, error)
	mustEmbedUnimplementedCashdeskServer()
}

// UnimplementedCashdeskServer must be embedded to have forward compatible implementations.
type UnimplementedCashdeskServer struct {
}

func (UnimplementedCashdeskServer) StreamEvents(Cashdesk_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedCashdeskServer) SendWorkstationLogs(context.Context, *SendWorkstationLogsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWorkstationLogs not implemented")
}
func (UnimplementedCashdeskServer) SendEventResult(context.Context, *SendEventResultRequest) (*SendEventResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEventResult not implemented")
}
func (UnimplementedCashdeskServer) mustEmbedUnimplementedCashdeskServer() {}

// UnsafeCashdeskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CashdeskServer will
// result in compilation errors.
type UnsafeCashdeskServer interface {
	mustEmbedUnimplementedCashdeskServer()
}

func RegisterCashdeskServer(s grpc.ServiceRegistrar, srv CashdeskServer) {
	s.RegisterService(&Cashdesk_ServiceDesc, srv)
}

func _Cashdesk_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CashdeskServer).StreamEvents(&cashdeskStreamEventsServer{stream})
}

type Cashdesk_StreamEventsServer interface {
	Send(*StreamWorkstationEventsResponse) error
	Recv() (*StreamEventsAck, error)
	grpc.ServerStream
}

type cashdeskStreamEventsServer struct {
	grpc.ServerStream
}

func (x *cashdeskStreamEventsServer) Send(m *StreamWorkstationEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cashdeskStreamEventsServer) Recv() (*StreamEventsAck, error) {
	m := new(StreamEventsAck)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Cashdesk_SendWorkstationLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendWorkstationLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashdeskServer).SendWorkstationLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cashdesk_SendWorkstationLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashdeskServer).SendWorkstationLogs(ctx, req.(*SendWorkstationLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cashdesk_SendEventResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEventResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CashdeskServer).SendEventResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cashdesk_SendEventResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CashdeskServer).SendEventResult(ctx, req.(*SendEventResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cashdesk_ServiceDesc is the grpc.ServiceDesc for Cashdesk service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cashdesk_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "net.vseinstrumenti.git.fd.proto.cashdesk.Cashdesk",
	HandlerType: (*CashdeskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendWorkstationLogs",
			Handler:    _Cashdesk_SendWorkstationLogs_Handler,
		},
		{
			MethodName: "SendEventResult",
			Handler:    _Cashdesk_SendEventResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Cashdesk_StreamEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/cashdesk/cashdesk.proto",
}

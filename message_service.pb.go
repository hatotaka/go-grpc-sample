// Code generated by protoc-gen-go.
// source: message_service.proto
// DO NOT EDIT!

/*
Package gogrpcsample is a generated protocol buffer package.

It is generated from these files:
	message_service.proto

It has these top-level messages:
	Message
	Result
	WatchType
*/
package gogrpcsample

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Message struct {
	Body string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type WatchType struct {
}

func (m *WatchType) Reset()                    { *m = WatchType{} }
func (m *WatchType) String() string            { return proto.CompactTextString(m) }
func (*WatchType) ProtoMessage()               {}
func (*WatchType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Message)(nil), "gogrpcsample.Message")
	proto.RegisterType((*Result)(nil), "gogrpcsample.Result")
	proto.RegisterType((*WatchType)(nil), "gogrpcsample.WatchType")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for MessageService service

type MessageServiceClient interface {
	SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Result, error)
	WatchMessage(ctx context.Context, in *WatchType, opts ...grpc.CallOption) (MessageService_WatchMessageClient, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/gogrpcsample.MessageService/SendMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) WatchMessage(ctx context.Context, in *WatchType, opts ...grpc.CallOption) (MessageService_WatchMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageService_serviceDesc.Streams[0], c.cc, "/gogrpcsample.MessageService/WatchMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceWatchMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageService_WatchMessageClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messageServiceWatchMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceWatchMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	SendMessage(context.Context, *Message) (*Result, error)
	WatchMessage(*WatchType, MessageService_WatchMessageServer) error
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MessageServiceServer).SendMessage(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _MessageService_WatchMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchType)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageServiceServer).WatchMessage(m, &messageServiceWatchMessageServer{stream})
}

type MessageService_WatchMessageServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messageServiceWatchMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceWatchMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gogrpcsample.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMessage",
			Handler:       _MessageService_WatchMessage_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x49, 0xcf, 0x4f, 0x2f, 0x2a, 0x48, 0x2e, 0x4e, 0xcc, 0x2d, 0xc8, 0x49, 0x55,
	0x92, 0xe5, 0x62, 0xf7, 0x85, 0x28, 0x13, 0x12, 0xe2, 0x62, 0x49, 0xca, 0x4f, 0xa9, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x38, 0xb8, 0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73,
	0x4a, 0x94, 0xb8, 0xb9, 0x38, 0xc3, 0x13, 0x4b, 0x92, 0x33, 0x42, 0x2a, 0x0b, 0x52, 0x8d, 0x26,
	0x31, 0x72, 0xf1, 0x41, 0xb5, 0x05, 0x43, 0x0c, 0x17, 0xb2, 0xe1, 0xe2, 0x0e, 0x4e, 0xcd, 0x4b,
	0x81, 0x19, 0x26, 0xaa, 0x87, 0x6c, 0x8d, 0x1e, 0x54, 0x58, 0x4a, 0x04, 0x55, 0x18, 0x6a, 0x36,
	0x83, 0x90, 0x13, 0x17, 0x0f, 0xd8, 0x74, 0x98, 0x76, 0x71, 0x54, 0x75, 0x70, 0x9b, 0xa5, 0xb0,
	0x9b, 0xab, 0xc4, 0x60, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x9f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x90, 0xb3, 0x8a, 0xde, 0xf8, 0x00, 0x00, 0x00,
}

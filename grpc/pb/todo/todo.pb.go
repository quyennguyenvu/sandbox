// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Reminder             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=reminder,proto3" json:"reminder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_385a15c54c9b7faa, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Request) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Request) GetReminder() *timestamp.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

type Response struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_385a15c54c9b7faa, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListFilter struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFilter) Reset()         { *m = ListFilter{} }
func (m *ListFilter) String() string { return proto.CompactTextString(m) }
func (*ListFilter) ProtoMessage()    {}
func (*ListFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_385a15c54c9b7faa, []int{2}
}
func (m *ListFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFilter.Unmarshal(m, b)
}
func (m *ListFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFilter.Marshal(b, m, deterministic)
}
func (dst *ListFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFilter.Merge(dst, src)
}
func (m *ListFilter) XXX_Size() int {
	return xxx_messageInfo_ListFilter.Size(m)
}
func (m *ListFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListFilter proto.InternalMessageInfo

func (m *ListFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "todo.Request")
	proto.RegisterType((*Response)(nil), "todo.Response")
	proto.RegisterType((*ListFilter)(nil), "todo.ListFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	List(ctx context.Context, in *ListFilter, opts ...grpc.CallOption) (Todo_ListClient, error)
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) List(ctx context.Context, in *ListFilter, opts ...grpc.CallOption) (Todo_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Todo_serviceDesc.Streams[0], "/todo.Todo/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todo_ListClient interface {
	Recv() (*Request, error)
	grpc.ClientStream
}

type todoListClient struct {
	grpc.ClientStream
}

func (x *todoListClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/todo.Todo/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	List(*ListFilter, Todo_ListServer) error
	Create(context.Context, *Request) (*Response, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServer).List(m, &todoListServer{stream})
}

type Todo_ListServer interface {
	Send(*Request) error
	grpc.ServerStream
}

type todoListServer struct {
	grpc.ServerStream
}

func (x *todoListServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func _Todo_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Todo_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _Todo_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_todo_385a15c54c9b7faa) }

var fileDescriptor_todo_385a15c54c9b7faa = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4e, 0xf3, 0x40,
	0x10, 0x85, 0xb5, 0xfe, 0x9d, 0xc4, 0x99, 0x28, 0x51, 0xb4, 0x3f, 0x85, 0x65, 0x21, 0x61, 0xb9,
	0x40, 0x11, 0x85, 0x0d, 0x01, 0x51, 0x00, 0x1d, 0x12, 0x15, 0xd5, 0x2a, 0x17, 0x70, 0xe2, 0x21,
	0x5a, 0x91, 0x78, 0xcc, 0xee, 0x04, 0x44, 0x09, 0x1d, 0x35, 0x47, 0xe3, 0x0a, 0x1c, 0x04, 0xd9,
	0x1b, 0x07, 0x10, 0xdd, 0xbe, 0x79, 0xef, 0xad, 0xbe, 0x19, 0x00, 0xa6, 0x82, 0xd2, 0xca, 0x10,
	0x93, 0xf4, 0xeb, 0x77, 0x74, 0xb0, 0x24, 0x5a, 0xae, 0x30, 0x6b, 0x66, 0xf3, 0xcd, 0x5d, 0xc6,
	0x7a, 0x8d, 0x96, 0xf3, 0x75, 0xe5, 0x62, 0xd1, 0xfe, 0x36, 0x90, 0x57, 0x3a, 0xcb, 0xcb, 0x92,
	0x38, 0x67, 0x4d, 0xa5, 0x75, 0x6e, 0xf2, 0x26, 0xa0, 0xa7, 0xf0, 0x61, 0x83, 0x96, 0xe5, 0x08,
	0x3c, 0x5d, 0x84, 0x22, 0x16, 0x93, 0x8e, 0xf2, 0x74, 0x21, 0xf7, 0xa0, 0xc3, 0x9a, 0x57, 0x18,
	0x7a, 0xb1, 0x98, 0xf4, 0x95, 0x13, 0x32, 0x86, 0x41, 0x81, 0x76, 0x61, 0x74, 0x55, 0xff, 0x13,
	0xfe, 0x6b, 0xbc, 0x9f, 0x23, 0x79, 0x0e, 0x81, 0xc1, 0xb5, 0x2e, 0x0b, 0x34, 0xa1, 0x1f, 0x8b,
	0xc9, 0x60, 0x1a, 0xa5, 0x0e, 0x22, 0x6d, 0x29, 0xd3, 0x59, 0x4b, 0xa9, 0x76, 0xd9, 0xe4, 0x0c,
	0x02, 0x85, 0xb6, 0xa2, 0xd2, 0xe2, 0x1f, 0x96, 0x10, 0x7a, 0x76, 0xb3, 0x58, 0xa0, 0xb5, 0x0d,
	0x4d, 0xa0, 0x5a, 0x99, 0x1c, 0x02, 0xdc, 0x6a, 0xcb, 0x37, 0x7a, 0xc5, 0x68, 0xea, 0xdc, 0x3d,
	0x3e, 0x3f, 0x91, 0x71, 0xe5, 0xbe, 0x6a, 0xe5, 0xf4, 0x45, 0x80, 0x3f, 0xa3, 0x82, 0xe4, 0x25,
	0xf8, 0x75, 0x41, 0x8e, 0xd3, 0xe6, 0x98, 0xdf, 0xe5, 0x68, 0xe8, 0x26, 0xdb, 0x7b, 0x24, 0xe3,
	0xd7, 0x8f, 0xcf, 0x77, 0x0f, 0x64, 0x90, 0x3d, 0x9e, 0x64, 0xb5, 0x73, 0x2c, 0xe4, 0x15, 0x74,
	0xaf, 0x0d, 0xe6, 0x8c, 0xf2, 0x77, 0x38, 0x1a, 0xb5, 0xd2, 0x2d, 0x90, 0xfc, 0x6f, 0xca, 0xc3,
	0x64, 0x57, 0xbe, 0x10, 0x47, 0xf3, 0x6e, 0xb3, 0xff, 0xe9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3e, 0x25, 0x42, 0xef, 0xc7, 0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package customer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Id                   int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string             `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string             `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Addresses            []*Request_Address `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Request) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Request) GetAddresses() []*Request_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Request_Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string   `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	IsShippingAddress    bool     `protobuf:"varint,5,opt,name=isShippingAddress,proto3" json:"isShippingAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Address) Reset()         { *m = Request_Address{} }
func (m *Request_Address) String() string { return proto.CompactTextString(m) }
func (*Request_Address) ProtoMessage()    {}
func (*Request_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0, 0}
}

func (m *Request_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Address.Unmarshal(m, b)
}
func (m *Request_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Address.Marshal(b, m, deterministic)
}
func (m *Request_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Address.Merge(m, src)
}
func (m *Request_Address) XXX_Size() int {
	return xxx_messageInfo_Request_Address.Size(m)
}
func (m *Request_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Address proto.InternalMessageInfo

func (m *Request_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *Request_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Request_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Request_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Request_Address) GetIsShippingAddress() bool {
	if m != nil {
		return m.IsShippingAddress
	}
	return false
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
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
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

type Filter struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "customer.Request")
	proto.RegisterType((*Request_Address)(nil), "customer.Request.Address")
	proto.RegisterType((*Response)(nil), "customer.Response")
	proto.RegisterType((*Filter)(nil), "customer.Filter")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x9b, 0xa4, 0x49, 0xd3, 0xfb, 0x41, 0x69, 0x2f, 0x1f, 0x32, 0x76, 0x15, 0xb2, 0xca,
	0x42, 0x82, 0x56, 0xc1, 0xb5, 0x14, 0x74, 0x3f, 0x3e, 0x41, 0x4c, 0x2e, 0x76, 0xb0, 0xf9, 0xe3,
	0xcc, 0x04, 0xa9, 0xaf, 0xe0, 0x4b, 0xf8, 0xa8, 0x32, 0x93, 0x89, 0x15, 0xb2, 0xbb, 0xe7, 0xe4,
	0x5c, 0xe6, 0x77, 0x4f, 0x60, 0x55, 0xf6, 0x4a, 0xb7, 0x35, 0xc9, 0xbc, 0x93, 0xad, 0x6e, 0x31,
	0x1e, 0x75, 0xfa, 0xed, 0xc3, 0x82, 0xd3, 0x7b, 0x4f, 0x4a, 0xe3, 0x0a, 0x7c, 0x51, 0x31, 0x2f,
	0xf1, 0xb2, 0x90, 0xfb, 0xa2, 0x42, 0x84, 0x79, 0x53, 0xd4, 0xc4, 0xfc, 0xc4, 0xcb, 0x96, 0xdc,
	0xce, 0xf8, 0x1f, 0x42, 0xaa, 0x0b, 0x71, 0x64, 0x81, 0x35, 0x07, 0x61, 0xdc, 0xee, 0xd0, 0x36,
	0xc4, 0xe6, 0x83, 0x6b, 0x05, 0xde, 0xc3, 0xb2, 0xa8, 0x2a, 0x49, 0x4a, 0x91, 0x62, 0x61, 0x12,
	0x64, 0xff, 0x76, 0x97, 0xf9, 0x2f, 0x89, 0x7b, 0x35, 0x7f, 0x18, 0x22, 0xfc, 0x9c, 0xdd, 0x7e,
	0x79, 0xb0, 0x70, 0x36, 0x5e, 0x40, 0xa4, 0xb4, 0x24, 0xd2, 0x16, 0x6c, 0xc9, 0x9d, 0x32, 0x70,
	0xa5, 0xd0, 0xa7, 0x11, 0xce, 0xcc, 0x06, 0x43, 0xe9, 0x42, 0xd3, 0x08, 0x67, 0x05, 0xae, 0x21,
	0xf8, 0x14, 0x9d, 0x43, 0x33, 0x23, 0x5e, 0xc1, 0x46, 0xa8, 0xe7, 0x83, 0xe8, 0x3a, 0xd1, 0xbc,
	0xba, 0x87, 0x58, 0x98, 0x78, 0x59, 0xcc, 0xa7, 0x1f, 0xd2, 0x3b, 0x88, 0x39, 0xa9, 0xae, 0x6d,
	0x14, 0x4d, 0x2a, 0x62, 0xb0, 0x50, 0x7d, 0x59, 0x9a, 0x7d, 0xdf, 0xee, 0x8f, 0x32, 0x4d, 0x21,
	0x7a, 0x14, 0x47, 0x4d, 0xd2, 0x64, 0xde, 0xe8, 0xf4, 0xd1, 0xca, 0xca, 0x9d, 0x30, 0xca, 0x5d,
	0x0d, 0xf1, 0xde, 0xd5, 0x81, 0x39, 0x04, 0x4f, 0xa4, 0x71, 0x7d, 0x2e, 0x68, 0x58, 0xdf, 0x6e,
	0x26, 0x95, 0xa5, 0xb3, 0x6b, 0x0f, 0x6f, 0x20, 0xda, 0x4b, 0x32, 0xf7, 0x4d, 0x03, 0x5b, 0xfc,
	0x6b, 0x0d, 0xe8, 0xe9, 0xec, 0x25, 0xb2, 0x3f, 0xff, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x98,
	0x35, 0x5d, 0xcf, 0x0e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (Customer_GetClient, error)
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (Customer_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Customer_serviceDesc.Streams[0], "/customer.Customer/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &customerGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Customer_GetClient interface {
	Recv() (*Request, error)
	grpc.ClientStream
}

type customerGetClient struct {
	grpc.ClientStream
}

func (x *customerGetClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *customerClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/customer.Customer/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	Get(*Filter, Customer_GetServer) error
	Create(context.Context, *Request) (*Response, error)
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Filter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CustomerServer).Get(m, &customerGetServer{stream})
}

type Customer_GetServer interface {
	Send(*Request) error
	grpc.ServerStream
}

type customerGetServer struct {
	grpc.ServerStream
}

func (x *customerGetServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func _Customer_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer.Customer/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Customer_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _Customer_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "customer.proto",
}

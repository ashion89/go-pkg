// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Respose struct {
	List                 []*Respose_List `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Respose) Reset()         { *m = Respose{} }
func (m *Respose) String() string { return proto.CompactTextString(m) }
func (*Respose) ProtoMessage()    {}
func (*Respose) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Respose) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Respose.Unmarshal(m, b)
}
func (m *Respose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Respose.Marshal(b, m, deterministic)
}
func (m *Respose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Respose.Merge(m, src)
}
func (m *Respose) XXX_Size() int {
	return xxx_messageInfo_Respose.Size(m)
}
func (m *Respose) XXX_DiscardUnknown() {
	xxx_messageInfo_Respose.DiscardUnknown(m)
}

var xxx_messageInfo_Respose proto.InternalMessageInfo

func (m *Respose) GetList() []*Respose_List {
	if m != nil {
		return m.List
	}
	return nil
}

type Respose_List struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Respose_List) Reset()         { *m = Respose_List{} }
func (m *Respose_List) String() string { return proto.CompactTextString(m) }
func (*Respose_List) ProtoMessage()    {}
func (*Respose_List) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1, 0}
}

func (m *Respose_List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Respose_List.Unmarshal(m, b)
}
func (m *Respose_List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Respose_List.Marshal(b, m, deterministic)
}
func (m *Respose_List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Respose_List.Merge(m, src)
}
func (m *Respose_List) XXX_Size() int {
	return xxx_messageInfo_Respose_List.Size(m)
}
func (m *Respose_List) XXX_DiscardUnknown() {
	xxx_messageInfo_Respose_List.DiscardUnknown(m)
}

var xxx_messageInfo_Respose_List proto.InternalMessageInfo

func (m *Respose_List) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Respose_List) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterType((*Respose)(nil), "protos.Respose")
	proto.RegisterType((*Respose_List)(nil), "protos.Respose.List")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	GetList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Respose, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) GetList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Respose, error) {
	out := new(Respose)
	err := c.cc.Invoke(ctx, "/protos.Test/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	GetList(context.Context, *Request) (*Respose, error)
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Test/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetList(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _Test_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xb2, 0x5c, 0xec, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x52, 0x3c, 0x48, 0xba, 0xb8, 0x20, 0xbf, 0x38, 0x55,
	0x48, 0x83, 0x8b, 0x25, 0x27, 0xb3, 0xb8, 0x44, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x04,
	0x62, 0x4e, 0xb1, 0x1e, 0x54, 0x5a, 0xcf, 0x27, 0xb3, 0xb8, 0x24, 0x08, 0xac, 0x42, 0x4a, 0x8b,
	0x8b, 0x05, 0xc4, 0x13, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x81, 0x1a, 0xc7, 0x94, 0x99, 0x02, 0xb7,
	0x80, 0x09, 0x61, 0x81, 0x91, 0x29, 0x17, 0x4b, 0x08, 0xc8, 0x72, 0x5d, 0x2e, 0x76, 0xf7, 0xd4,
	0x12, 0xb0, 0x36, 0x7e, 0x84, 0xd1, 0x60, 0x87, 0x49, 0xf1, 0xa3, 0xd9, 0xa5, 0xc4, 0x90, 0x04,
	0x71, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xb0, 0x99, 0xd5, 0xd3, 0x00, 0x00, 0x00,
}
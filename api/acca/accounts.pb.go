// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/acca/accounts.proto

package acca

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateCurrencyRequest struct {
	Key                  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateCurrencyRequest) Reset()         { *m = CreateCurrencyRequest{} }
func (m *CreateCurrencyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCurrencyRequest) ProtoMessage()    {}
func (*CreateCurrencyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_14b03f126f2e888c, []int{0}
}
func (m *CreateCurrencyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCurrencyRequest.Unmarshal(m, b)
}
func (m *CreateCurrencyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCurrencyRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCurrencyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCurrencyRequest.Merge(dst, src)
}
func (m *CreateCurrencyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCurrencyRequest.Size(m)
}
func (m *CreateCurrencyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCurrencyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCurrencyRequest proto.InternalMessageInfo

func (m *CreateCurrencyRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateCurrencyRequest) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type CreateCurrencyResponse struct {
	CurrencyId           int64    `protobuf:"varint,1,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCurrencyResponse) Reset()         { *m = CreateCurrencyResponse{} }
func (m *CreateCurrencyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCurrencyResponse) ProtoMessage()    {}
func (*CreateCurrencyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_14b03f126f2e888c, []int{1}
}
func (m *CreateCurrencyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCurrencyResponse.Unmarshal(m, b)
}
func (m *CreateCurrencyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCurrencyResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCurrencyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCurrencyResponse.Merge(dst, src)
}
func (m *CreateCurrencyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCurrencyResponse.Size(m)
}
func (m *CreateCurrencyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCurrencyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCurrencyResponse proto.InternalMessageInfo

func (m *CreateCurrencyResponse) GetCurrencyId() int64 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

type CreateAccountRequest struct {
	Key                  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	CurrencyId           int64             `protobuf:"varint,2,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,3,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_14b03f126f2e888c, []int{2}
}
func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (dst *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(dst, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateAccountRequest) GetCurrencyId() int64 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

func (m *CreateAccountRequest) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type CreateAccountResponse struct {
	AccId                int64    `protobuf:"varint,1,opt,name=acc_id,json=accId,proto3" json:"acc_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_accounts_14b03f126f2e888c, []int{3}
}
func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (dst *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(dst, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetAccId() int64 {
	if m != nil {
		return m.AccId
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateCurrencyRequest)(nil), "acca.CreateCurrencyRequest")
	proto.RegisterMapType((map[string]string)(nil), "acca.CreateCurrencyRequest.MetaEntry")
	proto.RegisterType((*CreateCurrencyResponse)(nil), "acca.CreateCurrencyResponse")
	proto.RegisterType((*CreateAccountRequest)(nil), "acca.CreateAccountRequest")
	proto.RegisterMapType((map[string]string)(nil), "acca.CreateAccountRequest.MetaEntry")
	proto.RegisterType((*CreateAccountResponse)(nil), "acca.CreateAccountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error)
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) CreateCurrency(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CreateCurrencyResponse, error) {
	out := new(CreateCurrencyResponse)
	err := c.cc.Invoke(ctx, "/acca.Accounts/CreateCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/acca.Accounts/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	CreateCurrency(context.Context, *CreateCurrencyRequest) (*CreateCurrencyResponse, error)
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_CreateCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).CreateCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acca.Accounts/CreateCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).CreateCurrency(ctx, req.(*CreateCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/acca.Accounts/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "acca.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCurrency",
			Handler:    _Accounts_CreateCurrency_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Accounts_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/acca/accounts.proto",
}

func init() { proto.RegisterFile("api/acca/accounts.proto", fileDescriptor_accounts_14b03f126f2e888c) }

var fileDescriptor_accounts_14b03f126f2e888c = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0x4c, 0x4e, 0x4e, 0x04, 0x11, 0xf9, 0xa5, 0x79, 0x25, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x2c, 0x20, 0x41, 0xa5, 0xb9, 0x8c, 0x5c, 0xa2, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9,
	0xce, 0xa5, 0x45, 0x45, 0xa9, 0x79, 0xc9, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42,
	0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x90,
	0x25, 0x17, 0x4b, 0x6e, 0x6a, 0x49, 0xa2, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xaa, 0x1e,
	0xc8, 0x00, 0x3d, 0xac, 0x9a, 0xf5, 0x7c, 0x53, 0x4b, 0x12, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83,
	0xc0, 0x5a, 0xa4, 0xcc, 0xb9, 0x38, 0xe1, 0x42, 0x58, 0x4c, 0x16, 0xe1, 0x62, 0x2d, 0x4b, 0xcc,
	0x29, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4a, 0x96, 0x5c, 0x62,
	0xe8, 0x36, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x73, 0x71, 0x27, 0x43, 0xc5, 0xe2,
	0x33, 0x53, 0xc0, 0xa6, 0x31, 0x07, 0x71, 0xc1, 0x84, 0x3c, 0x53, 0x94, 0xf6, 0x30, 0x72, 0x89,
	0x40, 0xf4, 0x3a, 0x42, 0x7c, 0x8e, 0xdb, 0x67, 0x68, 0x66, 0x31, 0xa1, 0x9b, 0x25, 0x64, 0x01,
	0xf5, 0x3a, 0x33, 0xd8, 0xeb, 0x2a, 0xc8, 0x5e, 0x47, 0x35, 0x9c, 0x7a, 0x3e, 0xd7, 0x83, 0x45,
	0x0c, 0xdc, 0x02, 0xa8, 0xc7, 0x45, 0xb9, 0xd8, 0x12, 0x93, 0x93, 0x11, 0x7e, 0x66, 0x4d, 0x4c,
	0x4e, 0xf6, 0x4c, 0x31, 0x5a, 0xca, 0xc8, 0xc5, 0x01, 0x55, 0x5a, 0x2c, 0xe4, 0xcb, 0xc5, 0x87,
	0x1a, 0x6c, 0x42, 0xd2, 0x78, 0xa2, 0x4b, 0x4a, 0x06, 0xbb, 0x24, 0xc4, 0x42, 0x25, 0x06, 0x21,
	0x2f, 0x2e, 0x5e, 0x14, 0xb7, 0x08, 0x49, 0xe1, 0x0e, 0x01, 0x29, 0x69, 0xac, 0x72, 0x30, 0xb3,
	0x92, 0xd8, 0xc0, 0xc9, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xef, 0x48, 0x9a, 0x99,
	0x02, 0x00, 0x00,
}
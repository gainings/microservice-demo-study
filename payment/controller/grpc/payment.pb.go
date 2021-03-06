// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

package payment_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//Request はI3セッションを更新しつつ取得するためのリクエストです。
type PayRequest struct {
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayRequest) Reset()         { *m = PayRequest{} }
func (m *PayRequest) String() string { return proto.CompactTextString(m) }
func (*PayRequest) ProtoMessage()    {}
func (*PayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{0}
}

func (m *PayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayRequest.Unmarshal(m, b)
}
func (m *PayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayRequest.Marshal(b, m, deterministic)
}
func (m *PayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayRequest.Merge(m, src)
}
func (m *PayRequest) XXX_Size() int {
	return xxx_messageInfo_PayRequest.Size(m)
}
func (m *PayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PayRequest proto.InternalMessageInfo

func (m *PayRequest) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

//PaymentStatus は支払いの状況を返します。
type PaymentStatus struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Authorized           bool     `protobuf:"varint,2,opt,name=authorized,proto3" json:"authorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentStatus) Reset()         { *m = PaymentStatus{} }
func (m *PaymentStatus) String() string { return proto.CompactTextString(m) }
func (*PaymentStatus) ProtoMessage()    {}
func (*PaymentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6362648dfa63d410, []int{1}
}

func (m *PaymentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentStatus.Unmarshal(m, b)
}
func (m *PaymentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentStatus.Marshal(b, m, deterministic)
}
func (m *PaymentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentStatus.Merge(m, src)
}
func (m *PaymentStatus) XXX_Size() int {
	return xxx_messageInfo_PaymentStatus.Size(m)
}
func (m *PaymentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentStatus proto.InternalMessageInfo

func (m *PaymentStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PaymentStatus) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func init() {
	proto.RegisterType((*PayRequest)(nil), "payment.v1.PayRequest")
	proto.RegisterType((*PaymentStatus)(nil), "payment.v1.PaymentStatus")
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor_6362648dfa63d410) }

var fileDescriptor_6362648dfa63d410 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xac, 0xcc,
	0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x71, 0xcb, 0x0c, 0x95,
	0x54, 0xb8, 0xb8, 0x02, 0x12, 0x2b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8,
	0xd8, 0x12, 0x73, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xa0, 0x3c,
	0x25, 0x4f, 0x2e, 0xde, 0x00, 0x88, 0x9e, 0xe0, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x21, 0x09, 0x2e,
	0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0xb0, 0x4a, 0xce, 0x20, 0x18, 0x57, 0x48, 0x8e,
	0x8b, 0x2b, 0xb1, 0xb4, 0x24, 0x23, 0xbf, 0x28, 0xb3, 0x2a, 0x35, 0x45, 0x82, 0x49, 0x81, 0x51,
	0x83, 0x23, 0x08, 0x49, 0xc4, 0xc8, 0x87, 0x8b, 0x0f, 0x66, 0x54, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x90, 0x15, 0x17, 0x73, 0x40, 0x62, 0xa5, 0x90, 0x98, 0x1e, 0xc2, 0x59, 0x7a, 0x08, 0x37,
	0x49, 0x49, 0xa2, 0x89, 0x23, 0x5c, 0xa1, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0x91, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x93, 0x09, 0xfe, 0xe2, 0x00, 0x00, 0x00,
}

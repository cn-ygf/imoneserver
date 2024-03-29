// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hbp.proto

package proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 心跳包请求
type HBPREQ struct {
	Time                 *int64   `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Msg                  *string  `protobuf:"bytes,2,req,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HBPREQ) Reset()         { *m = HBPREQ{} }
func (m *HBPREQ) String() string { return proto.CompactTextString(m) }
func (*HBPREQ) ProtoMessage()    {}
func (*HBPREQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_c75c376d9aa1b70d, []int{0}
}

func (m *HBPREQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HBPREQ.Unmarshal(m, b)
}
func (m *HBPREQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HBPREQ.Marshal(b, m, deterministic)
}
func (m *HBPREQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HBPREQ.Merge(m, src)
}
func (m *HBPREQ) XXX_Size() int {
	return xxx_messageInfo_HBPREQ.Size(m)
}
func (m *HBPREQ) XXX_DiscardUnknown() {
	xxx_messageInfo_HBPREQ.DiscardUnknown(m)
}

var xxx_messageInfo_HBPREQ proto.InternalMessageInfo

func (m *HBPREQ) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *HBPREQ) GetMsg() string {
	if m != nil && m.Msg != nil {
		return *m.Msg
	}
	return ""
}

// 心跳包应答
type HBPACK struct {
	Code                 *int32   `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HBPACK) Reset()         { *m = HBPACK{} }
func (m *HBPACK) String() string { return proto.CompactTextString(m) }
func (*HBPACK) ProtoMessage()    {}
func (*HBPACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_c75c376d9aa1b70d, []int{1}
}

func (m *HBPACK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HBPACK.Unmarshal(m, b)
}
func (m *HBPACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HBPACK.Marshal(b, m, deterministic)
}
func (m *HBPACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HBPACK.Merge(m, src)
}
func (m *HBPACK) XXX_Size() int {
	return xxx_messageInfo_HBPACK.Size(m)
}
func (m *HBPACK) XXX_DiscardUnknown() {
	xxx_messageInfo_HBPACK.DiscardUnknown(m)
}

var xxx_messageInfo_HBPACK proto.InternalMessageInfo

func (m *HBPACK) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*HBPREQ)(nil), "proto.HBPREQ")
	proto.RegisterType((*HBPACK)(nil), "proto.HBPACK")
}

func init() { proto.RegisterFile("proto/hbp.proto", fileDescriptor_c75c376d9aa1b70d) }

var fileDescriptor_c75c376d9aa1b70d = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0x2a, 0xd0, 0x03, 0xb3, 0x84, 0x58, 0xc1, 0x94, 0x92, 0x1e, 0x17, 0x9b,
	0x87, 0x53, 0x40, 0x90, 0x6b, 0xa0, 0x90, 0x10, 0x17, 0x4b, 0x49, 0x66, 0x6e, 0xaa, 0x04, 0xa3,
	0x02, 0x93, 0x06, 0x73, 0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x2e, 0xc1, 0xa4,
	0xc0, 0xa4, 0xc1, 0x19, 0x04, 0x62, 0x2a, 0xc9, 0x80, 0xd5, 0x3b, 0x3a, 0x7b, 0x83, 0xd4, 0x27,
	0xe7, 0xa7, 0x40, 0xd4, 0xb3, 0x06, 0x81, 0xd9, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xca,
	0xbf, 0x48, 0x66, 0x00, 0x00, 0x00,
}

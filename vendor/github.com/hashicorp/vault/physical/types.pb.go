// Code generated by protoc-gen-go. DO NOT EDIT.
// source: physical/types.proto

package physical // import "github.com/hashicorp/vault/physical"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SealWrapEntry struct {
	Ciphertext           []byte   `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	IV                   []byte   `protobuf:"bytes,2,opt,name=iv,proto3" json:"iv,omitempty"`
	HMAC                 []byte   `protobuf:"bytes,3,opt,name=hmac,proto3" json:"hmac,omitempty"`
	Wrapped              bool     `protobuf:"varint,4,opt,name=wrapped,proto3" json:"wrapped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SealWrapEntry) Reset()         { *m = SealWrapEntry{} }
func (m *SealWrapEntry) String() string { return proto.CompactTextString(m) }
func (*SealWrapEntry) ProtoMessage()    {}
func (*SealWrapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_0f1d299b2a49592e, []int{0}
}
func (m *SealWrapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SealWrapEntry.Unmarshal(m, b)
}
func (m *SealWrapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SealWrapEntry.Marshal(b, m, deterministic)
}
func (dst *SealWrapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SealWrapEntry.Merge(dst, src)
}
func (m *SealWrapEntry) XXX_Size() int {
	return xxx_messageInfo_SealWrapEntry.Size(m)
}
func (m *SealWrapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SealWrapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SealWrapEntry proto.InternalMessageInfo

func (m *SealWrapEntry) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

func (m *SealWrapEntry) GetIV() []byte {
	if m != nil {
		return m.IV
	}
	return nil
}

func (m *SealWrapEntry) GetHMAC() []byte {
	if m != nil {
		return m.HMAC
	}
	return nil
}

func (m *SealWrapEntry) GetWrapped() bool {
	if m != nil {
		return m.Wrapped
	}
	return false
}

func init() {
	proto.RegisterType((*SealWrapEntry)(nil), "physical.SealWrapEntry")
}

func init() { proto.RegisterFile("physical/types.proto", fileDescriptor_types_0f1d299b2a49592e) }

var fileDescriptor_types_0f1d299b2a49592e = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0xc6, 0x71, 0x5a, 0x8b, 0x96, 0x43, 0x1d, 0x82, 0x43, 0x26, 0x29, 0x8a, 0xd0, 0xa9, 0x19,
	0x7c, 0x03, 0xc1, 0x17, 0xa8, 0x83, 0xe0, 0x76, 0x8d, 0xc1, 0x04, 0x5a, 0x73, 0xa4, 0xd7, 0x6a,
	0xdf, 0x5e, 0x08, 0x14, 0xdc, 0xbe, 0xef, 0x37, 0xfd, 0x61, 0x47, 0x76, 0xea, 0x9d, 0xc6, 0x56,
	0xf1, 0x44, 0xa6, 0xaf, 0x28, 0x78, 0xf6, 0x22, 0x9f, 0xf5, 0xd0, 0xc1, 0xe6, 0x66, 0xb0, 0xbd,
	0x07, 0xa4, 0xeb, 0x9b, 0xc3, 0x24, 0xf6, 0x00, 0xda, 0x91, 0x35, 0x81, 0xcd, 0x97, 0x65, 0x52,
	0x24, 0xe5, 0xba, 0xfe, 0x13, 0xb1, 0x85, 0xd4, 0x8d, 0x32, 0x8d, 0x9e, 0xba, 0x51, 0x08, 0xc8,
	0x6c, 0x87, 0x5a, 0x2e, 0xa2, 0xc4, 0x2d, 0x24, 0xac, 0x3e, 0x01, 0x89, 0xcc, 0x53, 0x66, 0x45,
	0x52, 0xe6, 0xf5, 0x7c, 0x2f, 0xa7, 0xc7, 0xf1, 0xe5, 0xd8, 0x0e, 0x4d, 0xa5, 0x7d, 0xa7, 0x2c,
	0xf6, 0xd6, 0x69, 0x1f, 0x48, 0x8d, 0x38, 0xb4, 0xac, 0xe6, 0xaa, 0x66, 0x19, 0x33, 0xcf, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xe9, 0xa0, 0xf8, 0xbe, 0x00, 0x00, 0x00,
}

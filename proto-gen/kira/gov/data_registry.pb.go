// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/data_registry.proto

package gov

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

type DataRegistryEntry struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Reference            string   `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
	Encoding             string   `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Size                 uint64   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataRegistryEntry) Reset()         { *m = DataRegistryEntry{} }
func (m *DataRegistryEntry) String() string { return proto.CompactTextString(m) }
func (*DataRegistryEntry) ProtoMessage()    {}
func (*DataRegistryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_740f8fb5adb6d79b, []int{0}
}

func (m *DataRegistryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataRegistryEntry.Unmarshal(m, b)
}
func (m *DataRegistryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataRegistryEntry.Marshal(b, m, deterministic)
}
func (m *DataRegistryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataRegistryEntry.Merge(m, src)
}
func (m *DataRegistryEntry) XXX_Size() int {
	return xxx_messageInfo_DataRegistryEntry.Size(m)
}
func (m *DataRegistryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DataRegistryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DataRegistryEntry proto.InternalMessageInfo

func (m *DataRegistryEntry) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *DataRegistryEntry) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *DataRegistryEntry) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *DataRegistryEntry) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*DataRegistryEntry)(nil), "kira.gov.DataRegistryEntry")
}

func init() {
	proto.RegisterFile("kira/gov/data_registry.proto", fileDescriptor_740f8fb5adb6d79b)
}

var fileDescriptor_740f8fb5adb6d79b = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0x3d, 0x8f, 0x83, 0x30,
	0x0c, 0x06, 0x60, 0x71, 0x87, 0x4e, 0x90, 0xed, 0x32, 0x45, 0x27, 0x06, 0x74, 0x13, 0x0b, 0x64,
	0xe8, 0x3f, 0xe8, 0xc7, 0xd4, 0x8d, 0xb1, 0x4b, 0x15, 0xc0, 0x0d, 0x51, 0xd5, 0xb8, 0x32, 0x06,
	0x95, 0xfe, 0xfa, 0x8a, 0xa8, 0x1f, 0xdb, 0xeb, 0xf7, 0x91, 0x2c, 0x5b, 0x64, 0x67, 0x47, 0x46,
	0x5b, 0x9c, 0x74, 0x67, 0xd8, 0x1c, 0x09, 0xac, 0x1b, 0x98, 0xe6, 0xea, 0x4a, 0xc8, 0x28, 0x93,
	0x45, 0x2b, 0x8b, 0xd3, 0xff, 0x28, 0x7e, 0xb7, 0x86, 0x4d, 0xfd, 0xf4, 0x9d, 0x67, 0x9a, 0xa5,
	0x14, 0x71, 0x6f, 0x86, 0x5e, 0x45, 0x79, 0x54, 0xa4, 0x75, 0xc8, 0x32, 0x13, 0x29, 0xc1, 0x09,
	0x08, 0x7c, 0x0b, 0xea, 0x2b, 0xc0, 0xa7, 0x90, 0x7f, 0x22, 0x01, 0xdf, 0x62, 0xe7, 0xbc, 0x55,
	0xdf, 0x01, 0xdf, 0xf3, 0xb2, 0x6d, 0x70, 0x77, 0x50, 0x71, 0x1e, 0x15, 0x71, 0x1d, 0xf2, 0x5a,
	0x1f, 0x4a, 0xeb, 0xb8, 0x1f, 0x9b, 0xaa, 0xc5, 0x8b, 0xde, 0x3b, 0x32, 0x1b, 0x24, 0xd0, 0xce,
	0x33, 0xd0, 0x4d, 0x87, 0x23, 0x4b, 0x0b, 0x5e, 0xbf, 0xbe, 0x68, 0x7e, 0x42, 0xb7, 0x7a, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x30, 0x94, 0xe6, 0x2a, 0xd8, 0x00, 0x00, 0x00,
}

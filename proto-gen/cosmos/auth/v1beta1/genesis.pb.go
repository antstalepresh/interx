// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cosmos/auth/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// GenesisState defines the auth module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	// accounts are the accounts present at genesis.
	Accounts             []*anypb.Any `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d897ccbce9822332, []int{0}
}

func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenesisState.Unmarshal(m, b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return xxx_messageInfo_GenesisState.Size(m)
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetAccounts() []*anypb.Any {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.auth.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("cosmos/auth/v1beta1/genesis.proto", fileDescriptor_d897ccbce9822332) }

var fileDescriptor_d897ccbce9822332 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4,
	0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86,
	0x28, 0xd1, 0x03, 0x29, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x07, 0x2b, 0x49, 0x2a, 0x4d, 0xd3, 0x4f, 0xcc, 0xab, 0x84, 0xa8, 0x97, 0x12, 0x49, 0xcf, 0x4f,
	0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0xa8, 0xa8, 0x1c, 0x36, 0x8b, 0xc0, 0x46, 0x82, 0xe5, 0x95,
	0xaa, 0xb9, 0x78, 0xdc, 0x21, 0xd6, 0x06, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x59, 0x72, 0xb1, 0x15,
	0x24, 0x16, 0x25, 0xe6, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xeb, 0x61, 0x71,
	0x86, 0x5e, 0x00, 0x58, 0x89, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x50, 0x0d, 0x42, 0x06,
	0x5c, 0x1c, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc,
	0x46, 0x22, 0x7a, 0x10, 0xe7, 0xea, 0xc1, 0x9c, 0xab, 0xe7, 0x98, 0x57, 0x19, 0x04, 0x57, 0xe5,
	0xa4, 0x1d, 0xa5, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x75,
	0x29, 0x84, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0xaf, 0x80, 0x38, 0xbb, 0xa4, 0xb2, 0x20, 0xb5, 0x38,
	0x89, 0x0d, 0x6c, 0x88, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xcc, 0x8d, 0x09, 0x3b, 0x01,
	0x00, 0x00,
}

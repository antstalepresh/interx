// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cosmos/upgrade/v1beta1/upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Plan specifies information about a planned upgrade and when it should occur.
type Plan struct {
	// Sets the name for the upgrade. This name will be used by the upgraded
	// version of the software to apply any special "on-upgrade" commands during
	// the first BeginBlock method after the upgrade is applied. It is also used
	// to detect whether a software version can handle a given upgrade. If no
	// upgrade handler with this name has been set in the software, it will be
	// assumed that the software is out-of-date when the upgrade Time or Height is
	// reached and the software will exit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Time based upgrades have been deprecated. Time based upgrade logic
	// has been removed from the SDK.
	// If this field is not empty, an error will be thrown.
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"` // Deprecated: Do not use.
	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height int64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	// Deprecated: UpgradedClientState field has been deprecated. IBC upgrade logic has been
	// moved to the IBC module in the sub module 02-client.
	// If this field is not empty, an error will be thrown.
	UpgradedClientState  *anypb.Any `protobuf:"bytes,5,opt,name=upgraded_client_state,json=upgradedClientState,proto3" json:"upgraded_client_state,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{0}
}

func (m *Plan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plan.Unmarshal(m, b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return xxx_messageInfo_Plan.Size(m)
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Deprecated: Do not use.
func (m *Plan) GetTime() *timestamppb.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Plan) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Plan) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

// Deprecated: Do not use.
func (m *Plan) GetUpgradedClientState() *anypb.Any {
	if m != nil {
		return m.UpgradedClientState
	}
	return nil
}

// SoftwareUpgradeProposal is a gov Content type for initiating a software
// upgrade.
type SoftwareUpgradeProposal struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Plan                 *Plan    `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoftwareUpgradeProposal) Reset()         { *m = SoftwareUpgradeProposal{} }
func (m *SoftwareUpgradeProposal) String() string { return proto.CompactTextString(m) }
func (*SoftwareUpgradeProposal) ProtoMessage()    {}
func (*SoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{1}
}

func (m *SoftwareUpgradeProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoftwareUpgradeProposal.Unmarshal(m, b)
}
func (m *SoftwareUpgradeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoftwareUpgradeProposal.Marshal(b, m, deterministic)
}
func (m *SoftwareUpgradeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoftwareUpgradeProposal.Merge(m, src)
}
func (m *SoftwareUpgradeProposal) XXX_Size() int {
	return xxx_messageInfo_SoftwareUpgradeProposal.Size(m)
}
func (m *SoftwareUpgradeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SoftwareUpgradeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SoftwareUpgradeProposal proto.InternalMessageInfo

func (m *SoftwareUpgradeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SoftwareUpgradeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SoftwareUpgradeProposal) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

// CancelSoftwareUpgradeProposal is a gov Content type for cancelling a software
// upgrade.
type CancelSoftwareUpgradeProposal struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelSoftwareUpgradeProposal) Reset()         { *m = CancelSoftwareUpgradeProposal{} }
func (m *CancelSoftwareUpgradeProposal) String() string { return proto.CompactTextString(m) }
func (*CancelSoftwareUpgradeProposal) ProtoMessage()    {}
func (*CancelSoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{2}
}

func (m *CancelSoftwareUpgradeProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelSoftwareUpgradeProposal.Unmarshal(m, b)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelSoftwareUpgradeProposal.Marshal(b, m, deterministic)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelSoftwareUpgradeProposal.Merge(m, src)
}
func (m *CancelSoftwareUpgradeProposal) XXX_Size() int {
	return xxx_messageInfo_CancelSoftwareUpgradeProposal.Size(m)
}
func (m *CancelSoftwareUpgradeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelSoftwareUpgradeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CancelSoftwareUpgradeProposal proto.InternalMessageInfo

func (m *CancelSoftwareUpgradeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CancelSoftwareUpgradeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// ModuleVersion specifies a module and its consensus version.
//
// Since: cosmos-sdk 0.43
type ModuleVersion struct {
	// name of the app module
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// consensus version of the app module
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModuleVersion) Reset()         { *m = ModuleVersion{} }
func (m *ModuleVersion) String() string { return proto.CompactTextString(m) }
func (*ModuleVersion) ProtoMessage()    {}
func (*ModuleVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccf2a7d4d7b48dca, []int{3}
}

func (m *ModuleVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModuleVersion.Unmarshal(m, b)
}
func (m *ModuleVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModuleVersion.Marshal(b, m, deterministic)
}
func (m *ModuleVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleVersion.Merge(m, src)
}
func (m *ModuleVersion) XXX_Size() int {
	return xxx_messageInfo_ModuleVersion.Size(m)
}
func (m *ModuleVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleVersion proto.InternalMessageInfo

func (m *ModuleVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModuleVersion) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Plan)(nil), "cosmos.upgrade.v1beta1.Plan")
	proto.RegisterType((*SoftwareUpgradeProposal)(nil), "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal")
	proto.RegisterType((*CancelSoftwareUpgradeProposal)(nil), "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal")
	proto.RegisterType((*ModuleVersion)(nil), "cosmos.upgrade.v1beta1.ModuleVersion")
}

func init() {
	proto.RegisterFile("cosmos/upgrade/v1beta1/upgrade.proto", fileDescriptor_ccf2a7d4d7b48dca)
}

var fileDescriptor_ccf2a7d4d7b48dca = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3f, 0x6f, 0xd4, 0x30,
	0x14, 0x3f, 0xd3, 0xb4, 0x50, 0x9f, 0x58, 0xcc, 0x51, 0xcc, 0xa9, 0x90, 0x28, 0x62, 0xb8, 0x01,
	0x1c, 0xf5, 0x90, 0x18, 0x6e, 0x23, 0x1d, 0x98, 0x90, 0xaa, 0x14, 0x18, 0x58, 0x2a, 0x27, 0xf1,
	0xe5, 0x2c, 0x1c, 0x3b, 0x8a, 0x9d, 0x42, 0xbe, 0x05, 0x12, 0x0b, 0x63, 0x3f, 0x4e, 0x3f, 0x45,
	0x11, 0x1b, 0x33, 0x23, 0x13, 0xb2, 0x9d, 0xa0, 0x13, 0xdc, 0xc8, 0xe4, 0xf7, 0xe7, 0xf7, 0x7e,
	0xbf, 0xe7, 0xf7, 0x1e, 0x7c, 0x52, 0x28, 0x5d, 0x2b, 0x9d, 0x74, 0x4d, 0xd5, 0xd2, 0x92, 0x25,
	0x97, 0x27, 0x39, 0x33, 0xf4, 0x64, 0xf4, 0x49, 0xd3, 0x2a, 0xa3, 0xd0, 0x91, 0x47, 0x91, 0x31,
	0x3a, 0xa0, 0xe6, 0x0f, 0x2b, 0xa5, 0x2a, 0xc1, 0x12, 0x87, 0xca, 0xbb, 0x75, 0x42, 0x65, 0xef,
	0x4b, 0xe6, 0xb3, 0x4a, 0x55, 0xca, 0x99, 0x89, 0xb5, 0x86, 0x68, 0xf8, 0x77, 0x81, 0xe1, 0x35,
	0xd3, 0x86, 0xd6, 0x8d, 0x07, 0xc4, 0xbf, 0x00, 0x0c, 0xce, 0x04, 0x95, 0x08, 0xc1, 0x40, 0xd2,
	0x9a, 0x61, 0x10, 0x81, 0xc5, 0x61, 0xe6, 0x6c, 0xb4, 0x82, 0x81, 0xc5, 0xe3, 0x5b, 0x11, 0x58,
	0x4c, 0x97, 0x73, 0xe2, 0xc9, 0xc8, 0x48, 0x46, 0xde, 0x8c, 0x64, 0x29, 0xbc, 0xbe, 0x09, 0x27,
	0x9f, 0xbf, 0x85, 0x00, 0x83, 0xcc, 0xd5, 0xa0, 0x23, 0x78, 0xb0, 0x61, 0xbc, 0xda, 0x18, 0xbc,
	0x17, 0x81, 0xc5, 0x5e, 0x36, 0x78, 0x56, 0x87, 0xcb, 0xb5, 0xc2, 0x81, 0xd7, 0xb1, 0x36, 0x12,
	0xf0, 0xfe, 0xf0, 0xd3, 0xf2, 0xa2, 0x10, 0x9c, 0x49, 0x73, 0xa1, 0x0d, 0x35, 0x0c, 0xef, 0x3b,
	0xe1, 0xd9, 0x3f, 0xc2, 0x2f, 0x65, 0x9f, 0xc6, 0x3f, 0x6f, 0xc2, 0xe3, 0x9e, 0xd6, 0x62, 0x15,
	0xef, 0x2c, 0x8e, 0x31, 0xc8, 0xee, 0x8d, 0x99, 0x53, 0x97, 0x38, 0xb7, 0xf1, 0xd5, 0x9d, 0xaf,
	0x57, 0xe1, 0xe4, 0xc7, 0x55, 0x08, 0xe2, 0x2f, 0x00, 0x3e, 0x38, 0x57, 0x6b, 0xf3, 0x91, 0xb6,
	0xec, 0xad, 0x47, 0x9e, 0xb5, 0xaa, 0x51, 0x9a, 0x0a, 0x34, 0x83, 0xfb, 0x86, 0x1b, 0x31, 0x0e,
	0xc4, 0x3b, 0x28, 0x82, 0xd3, 0x92, 0xe9, 0xa2, 0xe5, 0x8d, 0xe1, 0x4a, 0xba, 0xc1, 0x1c, 0x66,
	0xdb, 0x21, 0xf4, 0x02, 0x06, 0x8d, 0xa0, 0xd2, 0xfd, 0x7a, 0xba, 0x3c, 0x26, 0xbb, 0x37, 0x49,
	0xec, 0xcc, 0xd3, 0xc0, 0x4e, 0x2d, 0x73, 0xf8, 0xad, 0xae, 0x28, 0x7c, 0x74, 0x4a, 0x65, 0xc1,
	0xc4, 0x7f, 0x6e, 0x6d, 0x4b, 0xe2, 0x15, 0xbc, 0xfb, 0x5a, 0x95, 0x9d, 0x60, 0xef, 0x58, 0xab,
	0x6d, 0xd7, 0xbb, 0xb6, 0x8f, 0xe1, 0xed, 0x4b, 0x9f, 0x76, 0x64, 0x41, 0x36, 0xba, 0x8e, 0x08,
	0x58, 0xa2, 0x74, 0x79, 0xfd, 0xfd, 0xf1, 0xe4, 0xfd, 0xd3, 0x8a, 0x9b, 0x4d, 0x97, 0x93, 0x42,
	0xd5, 0xc9, 0x70, 0xdf, 0xfe, 0x79, 0xa6, 0xcb, 0x0f, 0xc9, 0xa7, 0x3f, 0xc7, 0x6e, 0xfa, 0x86,
	0xe9, 0xfc, 0xc0, 0xad, 0xf1, 0xf9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0xe3, 0xfc, 0x14,
	0x0b, 0x03, 0x00, 0x00,
}

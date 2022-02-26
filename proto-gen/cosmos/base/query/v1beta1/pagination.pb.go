// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cosmos/base/query/v1beta1/pagination.proto

package query

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

// PageRequest is to be embedded in gRPC request messages for efficient
// pagination. Ex:
//
//  message SomeRequest {
//          Foo some_parameter = 1;
//          PageRequest pagination = 2;
//  }
type PageRequest struct {
	// key is a value returned in PageResponse.next_key to begin
	// querying the next page most efficiently. Only one of offset or key
	// should be set.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// offset is a numeric offset that can be used when key is unavailable.
	// It is less efficient than using key. Only one of offset or key should
	// be set.
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// limit is the total number of results to be returned in the result page.
	// If left empty it will default to a value to be set by each app.
	Limit uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// count_total is set to true  to indicate that the result set should include
	// a count of the total number of items available for pagination in UIs.
	// count_total is only respected when offset is used. It is ignored when key
	// is set.
	CountTotal bool `protobuf:"varint,4,opt,name=count_total,json=countTotal,proto3" json:"count_total,omitempty"`
	// reverse is set to true if results are to be returned in the descending order.
	//
	// Since: cosmos-sdk 0.43
	Reverse              bool     `protobuf:"varint,5,opt,name=reverse,proto3" json:"reverse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageRequest) Reset()         { *m = PageRequest{} }
func (m *PageRequest) String() string { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()    {}
func (*PageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53d6d609fe6828af, []int{0}
}

func (m *PageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageRequest.Unmarshal(m, b)
}
func (m *PageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageRequest.Marshal(b, m, deterministic)
}
func (m *PageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageRequest.Merge(m, src)
}
func (m *PageRequest) XXX_Size() int {
	return xxx_messageInfo_PageRequest.Size(m)
}
func (m *PageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PageRequest proto.InternalMessageInfo

func (m *PageRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PageRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *PageRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PageRequest) GetCountTotal() bool {
	if m != nil {
		return m.CountTotal
	}
	return false
}

func (m *PageRequest) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

// PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
type PageResponse struct {
	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `protobuf:"bytes,1,opt,name=next_key,json=nextKey,proto3" json:"next_key,omitempty"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total                uint64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageResponse) Reset()         { *m = PageResponse{} }
func (m *PageResponse) String() string { return proto.CompactTextString(m) }
func (*PageResponse) ProtoMessage()    {}
func (*PageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53d6d609fe6828af, []int{1}
}

func (m *PageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageResponse.Unmarshal(m, b)
}
func (m *PageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageResponse.Marshal(b, m, deterministic)
}
func (m *PageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageResponse.Merge(m, src)
}
func (m *PageResponse) XXX_Size() int {
	return xxx_messageInfo_PageResponse.Size(m)
}
func (m *PageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PageResponse proto.InternalMessageInfo

func (m *PageResponse) GetNextKey() []byte {
	if m != nil {
		return m.NextKey
	}
	return nil
}

func (m *PageResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*PageRequest)(nil), "cosmos.base.query.v1beta1.PageRequest")
	proto.RegisterType((*PageResponse)(nil), "cosmos.base.query.v1beta1.PageResponse")
}

func init() {
	proto.RegisterFile("cosmos/base/query/v1beta1/pagination.proto", fileDescriptor_53d6d609fe6828af)
}

var fileDescriptor_53d6d609fe6828af = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0xfd, 0xcb, 0xb4, 0x07, 0x59, 0x44, 0xb6, 0x27, 0x43, 0x4f, 0xa1, 0x60, 0x96,
	0xe2, 0x07, 0x10, 0xbc, 0x7a, 0x91, 0xe0, 0xc9, 0x4b, 0xd9, 0xc4, 0x69, 0x0c, 0x6d, 0x32, 0x69,
	0x66, 0x52, 0xcc, 0x37, 0xf0, 0x63, 0xcb, 0xee, 0x46, 0x3c, 0xed, 0xfe, 0xde, 0x3c, 0x66, 0x1e,
	0x0f, 0x76, 0x05, 0x71, 0x4d, 0x6c, 0x72, 0xcb, 0x68, 0x2e, 0x3d, 0x76, 0x83, 0xb9, 0xee, 0x73,
	0x14, 0xbb, 0x37, 0xad, 0x2d, 0xab, 0xc6, 0x4a, 0x45, 0x4d, 0xda, 0x76, 0x24, 0xa4, 0x36, 0xc1,
	0x9b, 0x3a, 0x6f, 0xea, 0xbd, 0xe9, 0xe8, 0xdd, 0xfe, 0x44, 0xb0, 0x7a, 0xb3, 0x25, 0x66, 0x78,
	0xe9, 0x91, 0x45, 0xdd, 0xc2, 0xe4, 0x84, 0x83, 0x8e, 0xe2, 0x28, 0x59, 0x67, 0xee, 0xab, 0xee,
	0x61, 0x4e, 0xc7, 0x23, 0xa3, 0xe8, 0x9b, 0x38, 0x4a, 0xa6, 0xd9, 0x48, 0xea, 0x0e, 0x66, 0xe7,
	0xaa, 0xae, 0x44, 0x4f, 0xbc, 0x1c, 0x40, 0x3d, 0xc0, 0xaa, 0xa0, 0xbe, 0x91, 0x83, 0x90, 0xd8,
	0xb3, 0x9e, 0xc6, 0x51, 0xb2, 0xcc, 0xc0, 0x4b, 0xef, 0x4e, 0x51, 0x1a, 0x16, 0x1d, 0x5e, 0xb1,
	0x63, 0xd4, 0x33, 0x3f, 0xfc, 0xc3, 0xed, 0x33, 0xac, 0x43, 0x12, 0x6e, 0xa9, 0x61, 0x54, 0x1b,
	0x58, 0x36, 0xf8, 0x2d, 0x87, 0xff, 0x3c, 0x0b, 0xc7, 0xaf, 0x38, 0xb8, 0xdb, 0x61, 0x7f, 0x88,
	0x14, 0xe0, 0x65, 0xf7, 0x91, 0x94, 0x95, 0x7c, 0xf5, 0x79, 0x5a, 0x50, 0x6d, 0xc6, 0x7e, 0xc2,
	0xf3, 0xc8, 0x9f, 0x27, 0x23, 0x43, 0x8b, 0x1c, 0xba, 0xca, 0xe7, 0xbe, 0x99, 0xa7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xba, 0x07, 0xed, 0x5e, 0x47, 0x01, 0x00, 0x00,
}

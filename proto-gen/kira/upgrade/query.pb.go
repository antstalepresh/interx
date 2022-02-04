// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/upgrade/query.proto

package upgrade

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryCurrentPlanRequest is the request type for the Query/CurrentPlan RPC
// method.
type QueryCurrentPlanRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryCurrentPlanRequest) Reset()         { *m = QueryCurrentPlanRequest{} }
func (m *QueryCurrentPlanRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPlanRequest) ProtoMessage()    {}
func (*QueryCurrentPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6e2e2517e583f0, []int{0}
}

func (m *QueryCurrentPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCurrentPlanRequest.Unmarshal(m, b)
}
func (m *QueryCurrentPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCurrentPlanRequest.Marshal(b, m, deterministic)
}
func (m *QueryCurrentPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPlanRequest.Merge(m, src)
}
func (m *QueryCurrentPlanRequest) XXX_Size() int {
	return xxx_messageInfo_QueryCurrentPlanRequest.Size(m)
}
func (m *QueryCurrentPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPlanRequest proto.InternalMessageInfo

// QueryCurrentPlanResponse is the response type for the Query/CurrentPlan RPC
// method.
type QueryCurrentPlanResponse struct {
	// plan is the current upgrade plan.
	Plan                 *Plan    `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryCurrentPlanResponse) Reset()         { *m = QueryCurrentPlanResponse{} }
func (m *QueryCurrentPlanResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentPlanResponse) ProtoMessage()    {}
func (*QueryCurrentPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6e2e2517e583f0, []int{1}
}

func (m *QueryCurrentPlanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCurrentPlanResponse.Unmarshal(m, b)
}
func (m *QueryCurrentPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCurrentPlanResponse.Marshal(b, m, deterministic)
}
func (m *QueryCurrentPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentPlanResponse.Merge(m, src)
}
func (m *QueryCurrentPlanResponse) XXX_Size() int {
	return xxx_messageInfo_QueryCurrentPlanResponse.Size(m)
}
func (m *QueryCurrentPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentPlanResponse proto.InternalMessageInfo

func (m *QueryCurrentPlanResponse) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

// QueryNextPlanRequest is the request type for the Query/CurrentPlan RPC
// method.
type QueryNextPlanRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryNextPlanRequest) Reset()         { *m = QueryNextPlanRequest{} }
func (m *QueryNextPlanRequest) String() string { return proto.CompactTextString(m) }
func (*QueryNextPlanRequest) ProtoMessage()    {}
func (*QueryNextPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6e2e2517e583f0, []int{2}
}

func (m *QueryNextPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryNextPlanRequest.Unmarshal(m, b)
}
func (m *QueryNextPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryNextPlanRequest.Marshal(b, m, deterministic)
}
func (m *QueryNextPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNextPlanRequest.Merge(m, src)
}
func (m *QueryNextPlanRequest) XXX_Size() int {
	return xxx_messageInfo_QueryNextPlanRequest.Size(m)
}
func (m *QueryNextPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNextPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNextPlanRequest proto.InternalMessageInfo

// QueryNextPlanResponse is the response type for the Query/CurrentPlan RPC
// method.
type QueryNextPlanResponse struct {
	// plan is the next upgrade plan.
	Plan                 *Plan    `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryNextPlanResponse) Reset()         { *m = QueryNextPlanResponse{} }
func (m *QueryNextPlanResponse) String() string { return proto.CompactTextString(m) }
func (*QueryNextPlanResponse) ProtoMessage()    {}
func (*QueryNextPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6e2e2517e583f0, []int{3}
}

func (m *QueryNextPlanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryNextPlanResponse.Unmarshal(m, b)
}
func (m *QueryNextPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryNextPlanResponse.Marshal(b, m, deterministic)
}
func (m *QueryNextPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNextPlanResponse.Merge(m, src)
}
func (m *QueryNextPlanResponse) XXX_Size() int {
	return xxx_messageInfo_QueryNextPlanResponse.Size(m)
}
func (m *QueryNextPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNextPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNextPlanResponse proto.InternalMessageInfo

func (m *QueryNextPlanResponse) GetPlan() *Plan {
	if m != nil {
		return m.Plan
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryCurrentPlanRequest)(nil), "kira.upgrade.QueryCurrentPlanRequest")
	proto.RegisterType((*QueryCurrentPlanResponse)(nil), "kira.upgrade.QueryCurrentPlanResponse")
	proto.RegisterType((*QueryNextPlanRequest)(nil), "kira.upgrade.QueryNextPlanRequest")
	proto.RegisterType((*QueryNextPlanResponse)(nil), "kira.upgrade.QueryNextPlanResponse")
}

func init() {
	proto.RegisterFile("kira/upgrade/query.proto", fileDescriptor_4d6e2e2517e583f0)
}

var fileDescriptor_4d6e2e2517e583f0 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0x86, 0x33, 0x1f, 0x1f, 0xc6, 0x14, 0x56, 0x8d, 0x0a, 0x8c, 0x3f, 0x99, 0x8c, 0x91, 0x18,
	0x13, 0xa6, 0x32, 0xae, 0x74, 0x21, 0x01, 0x96, 0x26, 0x06, 0x49, 0xdc, 0xb8, 0x21, 0x05, 0xeb,
	0x30, 0x11, 0xdb, 0xd2, 0xe9, 0x28, 0x6c, 0xdd, 0xb8, 0xc7, 0x5b, 0xf1, 0x1a, 0xbc, 0x01, 0x6f,
	0xc1, 0x0b, 0x31, 0xd3, 0x96, 0x84, 0xe1, 0xcf, 0xb8, 0x6a, 0xd2, 0xe7, 0xed, 0xe9, 0x73, 0x7a,
	0x52, 0x50, 0x7c, 0x0c, 0x05, 0x46, 0x31, 0x0f, 0x04, 0xbe, 0x27, 0x68, 0x18, 0x13, 0x31, 0xf6,
	0xb8, 0x60, 0x92, 0xc1, 0x7c, 0x42, 0x3c, 0x43, 0xec, 0x52, 0xc0, 0x58, 0x30, 0x20, 0x48, 0xb1,
	0x6e, 0xfc, 0x80, 0x30, 0x35, 0x41, 0x7b, 0xcf, 0x20, 0xcc, 0x43, 0x84, 0x29, 0x65, 0x12, 0xcb,
	0x90, 0xd1, 0xc8, 0x50, 0xbd, 0xf4, 0x2a, 0x01, 0xa1, 0x15, 0xc6, 0x09, 0xc5, 0x3c, 0x7c, 0xf6,
	0x11, 0xe3, 0x2a, 0xb3, 0x24, 0x6f, 0xa7, 0x84, 0xcc, 0x6a, 0x58, 0x21, 0xc5, 0xf8, 0x00, 0x53,
	0x0d, 0xdc, 0x12, 0x28, 0xdc, 0x24, 0xea, 0xcd, 0x58, 0x08, 0x42, 0x65, 0x6b, 0x80, 0x69, 0x9b,
	0x0c, 0x63, 0x12, 0x49, 0xb7, 0x01, 0x8a, 0x8b, 0x28, 0xe2, 0x8c, 0x46, 0x04, 0x96, 0xc1, 0xff,
	0xa4, 0x48, 0xd1, 0x72, 0xac, 0xe3, 0x9c, 0x0f, 0xbd, 0xd9, 0x8e, 0x3d, 0x95, 0x54, 0xdc, 0xdd,
	0x01, 0x5b, 0xaa, 0xc6, 0x35, 0x19, 0xa5, 0x6a, 0xd7, 0xc0, 0xf6, 0xdc, 0xfe, 0xdf, 0x0a, 0xfb,
	0x6f, 0x19, 0x90, 0x55, 0x15, 0xe0, 0xa7, 0x05, 0x72, 0x33, 0x8a, 0xf0, 0x28, 0x7d, 0x66, 0x45,
	0x77, 0x76, 0xf9, 0xb7, 0x98, 0x16, 0x72, 0xfb, 0x93, 0x7a, 0x0d, 0x64, 0xd5, 0x78, 0xa1, 0xad,
	0x52, 0x8e, 0x89, 0x39, 0xb7, 0xfa, 0xac, 0x93, 0xe4, 0xed, 0x35, 0xec, 0xf5, 0xeb, 0xfb, 0xfd,
	0x9f, 0x03, 0x0f, 0xd4, 0x9c, 0x53, 0x83, 0xe8, 0xe9, 0x74, 0x27, 0x69, 0x09, 0x7e, 0x58, 0x60,
	0x73, 0xfa, 0x1e, 0xd0, 0x5d, 0xa2, 0x37, 0xf7, 0x88, 0xf6, 0xe1, 0xda, 0x8c, 0xf1, 0xef, 0x4c,
	0xea, 0xe7, 0x53, 0x7f, 0x3d, 0x6a, 0x27, 0xc9, 0xa4, 0xe5, 0x57, 0x01, 0x65, 0xbe, 0x0f, 0x77,
	0x17, 0xcd, 0x29, 0x19, 0x69, 0xed, 0x86, 0x9c, 0xd4, 0x2f, 0x61, 0xd6, 0xcf, 0x54, 0xbd, 0xd3,
	0x13, 0xcb, 0x12, 0x3e, 0xc8, 0x07, 0xed, 0x56, 0xb3, 0x12, 0x60, 0x49, 0x5e, 0xf0, 0x18, 0xba,
	0x7d, 0x29, 0x79, 0x74, 0x81, 0x50, 0x10, 0xca, 0x7e, 0xdc, 0xf5, 0x7a, 0xec, 0x09, 0x5d, 0x85,
	0x02, 0x37, 0x99, 0x20, 0x28, 0xa4, 0x92, 0x88, 0xd1, 0x5d, 0x75, 0x35, 0xd3, 0xdf, 0x26, 0xf9,
	0x03, 0xa9, 0xfb, 0xbb, 0x1b, 0x6a, 0xff, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff, 0x36, 0x17, 0x12,
	0x21, 0x86, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// CurrentPlan queries the current upgrade plan.
	CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error)
	// NextPlan queries the next upgrade plan.
	NextPlan(ctx context.Context, in *QueryNextPlanRequest, opts ...grpc.CallOption) (*QueryNextPlanResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CurrentPlan(ctx context.Context, in *QueryCurrentPlanRequest, opts ...grpc.CallOption) (*QueryCurrentPlanResponse, error) {
	out := new(QueryCurrentPlanResponse)
	err := c.cc.Invoke(ctx, "/kira.upgrade.Query/CurrentPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NextPlan(ctx context.Context, in *QueryNextPlanRequest, opts ...grpc.CallOption) (*QueryNextPlanResponse, error) {
	out := new(QueryNextPlanResponse)
	err := c.cc.Invoke(ctx, "/kira.upgrade.Query/NextPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CurrentPlan queries the current upgrade plan.
	CurrentPlan(context.Context, *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error)
	// NextPlan queries the next upgrade plan.
	NextPlan(context.Context, *QueryNextPlanRequest) (*QueryNextPlanResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CurrentPlan(ctx context.Context, req *QueryCurrentPlanRequest) (*QueryCurrentPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentPlan not implemented")
}
func (*UnimplementedQueryServer) NextPlan(ctx context.Context, req *QueryNextPlanRequest) (*QueryNextPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextPlan not implemented")
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CurrentPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.upgrade.Query/CurrentPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentPlan(ctx, req.(*QueryCurrentPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NextPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNextPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NextPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.upgrade.Query/NextPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NextPlan(ctx, req.(*QueryNextPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.upgrade.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentPlan",
			Handler:    _Query_CurrentPlan_Handler,
		},
		{
			MethodName: "NextPlan",
			Handler:    _Query_NextPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kira/upgrade/query.proto",
}

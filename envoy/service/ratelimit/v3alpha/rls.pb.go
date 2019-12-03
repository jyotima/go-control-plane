// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/ratelimit/v3alpha/rls.proto

package envoy_service_ratelimit_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/ratelimit"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}

var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}

func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	RateLimitResponse_RateLimit_SECOND  RateLimitResponse_RateLimit_Unit = 1
	RateLimitResponse_RateLimit_MINUTE  RateLimitResponse_RateLimit_Unit = 2
	RateLimitResponse_RateLimit_HOUR    RateLimitResponse_RateLimit_Unit = 3
	RateLimitResponse_RateLimit_DAY     RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}

func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 0, 0}
}

type RateLimitRequest struct {
	Domain               string                           `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	HitsAddend           uint32                           `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{0}
}

func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (m *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(m, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

type RateLimitResponse struct {
	OverallCode          RateLimitResponse_Code                `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v3alpha.RateLimitResponse_Code" json:"overall_code,omitempty"`
	Statuses             []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	Headers              []*core.HeaderValue                   `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	RequestHeadersToAdd  []*core.HeaderValue                   `protobuf:"bytes,4,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1}
}

func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(m, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RateLimitResponse) GetRequestHeadersToAdd() []*core.HeaderValue {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

type RateLimitResponse_RateLimit struct {
	RequestsPerUnit      uint32                           `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v3alpha.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Unmarshal(m, b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(m, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Size(m)
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	Code                 RateLimitResponse_Code       `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v3alpha.RateLimitResponse_Code" json:"code,omitempty"`
	CurrentLimit         *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	LimitRemaining       uint32                       `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_5929a49ed6c3ed07, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(m, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.service.ratelimit.v3alpha.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v3alpha.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v3alpha.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v3alpha.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v3alpha.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v3alpha.RateLimitResponse.DescriptorStatus")
}

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v3alpha/rls.proto", fileDescriptor_5929a49ed6c3ed07)
}

var fileDescriptor_5929a49ed6c3ed07 = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x4e, 0xdb, 0x48,
	0x14, 0xc7, 0x71, 0x12, 0x85, 0x70, 0xc2, 0x87, 0x99, 0x95, 0xd8, 0x28, 0x17, 0x0b, 0x9b, 0xd5,
	0xaa, 0xb4, 0x14, 0xa7, 0x35, 0x52, 0x5b, 0x45, 0x85, 0x2a, 0x10, 0x24, 0x10, 0x90, 0xa0, 0x09,
	0xa1, 0x1f, 0x37, 0xd6, 0x10, 0x8f, 0xc8, 0x48, 0xc6, 0xe3, 0xce, 0x8c, 0xa3, 0xd2, 0x27, 0xe0,
	0xa6, 0xed, 0x7d, 0xdf, 0xa2, 0x0f, 0xd1, 0x27, 0xe8, 0x0b, 0x55, 0x1e, 0x3b, 0x4e, 0x1a, 0x5a,
	0x41, 0x68, 0xef, 0x3c, 0x67, 0xfe, 0xe7, 0x37, 0x3e, 0xff, 0x73, 0x66, 0xe0, 0x3e, 0xf5, 0xfb,
	0xfc, 0xb2, 0x2a, 0xa9, 0xe8, 0xb3, 0x2e, 0xad, 0x0a, 0xa2, 0xa8, 0xc7, 0x2e, 0x98, 0xaa, 0xf6,
	0x37, 0x88, 0x17, 0xf4, 0x48, 0x55, 0x78, 0xd2, 0x0a, 0x04, 0x57, 0x1c, 0x2d, 0x6b, 0xa9, 0x95,
	0x48, 0xad, 0x54, 0x6a, 0x25, 0xd2, 0xf2, 0xbf, 0x31, 0x8b, 0x04, 0x2c, 0xcd, 0xee, 0x72, 0x41,
	0xab, 0x67, 0x44, 0xd2, 0x98, 0x51, 0x5e, 0xbb, 0x2e, 0x19, 0x1e, 0x39, 0x24, 0xc6, 0xe2, 0xff,
	0x43, 0x37, 0x20, 0x5a, 0x4b, 0x7c, 0x9f, 0x2b, 0xa2, 0x18, 0xf7, 0x65, 0xb5, 0x4f, 0x85, 0x64,
	0xdc, 0x67, 0xfe, 0x79, 0x22, 0xfb, 0xbb, 0x4f, 0x3c, 0xe6, 0x12, 0x45, 0xab, 0x83, 0x8f, 0x78,
	0xa3, 0xf2, 0xcd, 0x00, 0x13, 0x13, 0x45, 0x0f, 0x23, 0x26, 0xa6, 0x6f, 0x43, 0x2a, 0x15, 0x5a,
	0x82, 0xbc, 0xcb, 0x2f, 0x08, 0xf3, 0x4b, 0xc6, 0x8a, 0xb1, 0x3a, 0x83, 0x93, 0x15, 0xc2, 0x50,
	0x74, 0xa9, 0xec, 0x0a, 0x16, 0x28, 0x2e, 0x64, 0x29, 0xb3, 0x92, 0x5d, 0x2d, 0xda, 0x8f, 0xac,
	0xb8, 0x66, 0x12, 0xb0, 0x41, 0x95, 0x23, 0x75, 0xa7, 0xec, 0x46, 0x9a, 0x88, 0x47, 0x21, 0x68,
	0x19, 0x8a, 0x3d, 0xa6, 0xa4, 0x43, 0x5c, 0x97, 0xfa, 0x6e, 0x29, 0xbb, 0x62, 0xac, 0xce, 0x61,
	0x88, 0x42, 0x75, 0x1d, 0xa9, 0xd9, 0x9f, 0xbf, 0x5e, 0xfd, 0xb3, 0x0e, 0x6b, 0xbf, 0x74, 0xd6,
	0xb6, 0xc6, 0x0b, 0xa8, 0x7c, 0x2a, 0xc0, 0xe2, 0x48, 0x50, 0x06, 0xdc, 0x97, 0x14, 0xbd, 0x81,
	0x59, 0xde, 0xa7, 0x82, 0x78, 0x9e, 0xd3, 0xe5, 0x2e, 0xd5, 0xc5, 0xcd, 0xdb, 0x4f, 0xad, 0x1b,
	0x7a, 0x66, 0x5d, 0x23, 0x59, 0x3b, 0xdc, 0xa5, 0xb8, 0x98, 0xc0, 0xa2, 0x05, 0x72, 0xa0, 0x20,
	0x15, 0x51, 0xa1, 0xa4, 0x03, 0x5f, 0x76, 0xee, 0xc0, 0x1d, 0x9a, 0xd4, 0xd6, 0x30, 0x9c, 0x42,
	0xd1, 0x26, 0x4c, 0xf7, 0x28, 0x71, 0xa9, 0x90, 0xa5, 0xac, 0xe6, 0xff, 0xf7, 0x13, 0xdf, 0xa3,
	0x51, 0xb2, 0xf6, 0xb4, 0xec, 0x94, 0x78, 0x21, 0xc5, 0x83, 0x1c, 0xf4, 0x0a, 0x96, 0x44, 0x6c,
	0x8e, 0x93, 0x84, 0x1c, 0xc5, 0x23, 0xd3, 0x4b, 0xb9, 0xdb, 0xd3, 0xfe, 0x4a, 0x10, 0x71, 0x4c,
	0x9e, 0xf0, 0xba, 0xeb, 0x96, 0x3f, 0x64, 0x60, 0x26, 0xad, 0x04, 0x3d, 0x80, 0xc5, 0x44, 0x24,
	0x9d, 0x80, 0x0a, 0x27, 0xf4, 0x99, 0xd2, 0x46, 0xcf, 0xe1, 0x85, 0xc1, 0xc6, 0x31, 0x15, 0x1d,
	0x9f, 0x29, 0xd4, 0x81, 0x9c, 0xde, 0xce, 0xe8, 0x3e, 0xd4, 0xef, 0xe0, 0x57, 0x1a, 0xb1, 0x22,
	0x20, 0xd6, 0xb8, 0xca, 0x16, 0xe4, 0x34, 0xbe, 0x08, 0xd3, 0x9d, 0xe6, 0x41, 0xb3, 0xf5, 0xb2,
	0x69, 0x4e, 0x21, 0x80, 0x7c, 0x7b, 0x77, 0xa7, 0xd5, 0x6c, 0x98, 0x46, 0xf4, 0x7d, 0xb4, 0xdf,
	0xec, 0x9c, 0xec, 0x9a, 0x19, 0x54, 0x80, 0xdc, 0x5e, 0xab, 0x83, 0xcd, 0x2c, 0x9a, 0x86, 0x6c,
	0xa3, 0xfe, 0xda, 0xcc, 0xd5, 0x36, 0xa3, 0x81, 0x7b, 0x06, 0x4f, 0x6e, 0x37, 0x70, 0xe3, 0x7f,
	0x52, 0xfe, 0x92, 0x01, 0x73, 0xbc, 0x8f, 0xe8, 0x00, 0x72, 0x7f, 0x62, 0xe4, 0x34, 0x04, 0x11,
	0x98, 0xeb, 0x86, 0x42, 0x50, 0x5f, 0x39, 0x3a, 0x4b, 0x1b, 0x58, 0xb4, 0x9f, 0xff, 0x8e, 0x81,
	0x78, 0x36, 0x41, 0xc6, 0x6d, 0xbc, 0x07, 0x0b, 0x3a, 0xd5, 0x11, 0x34, 0xba, 0xf9, 0xcc, 0x3f,
	0x4f, 0x6e, 0xe6, 0xbc, 0x17, 0xe7, 0x27, 0xd1, 0x5a, 0x23, 0x32, 0xeb, 0x05, 0x6c, 0x4e, 0x64,
	0xd6, 0xb8, 0x3d, 0x95, 0x35, 0xc8, 0xe9, 0x5b, 0xf4, 0x43, 0xcb, 0xf2, 0x90, 0x69, 0x1d, 0x98,
	0x06, 0x9a, 0x07, 0x68, 0x9d, 0xee, 0x62, 0xe7, 0x70, 0xff, 0x68, 0xff, 0xc4, 0xcc, 0xd4, 0x36,
	0xa2, 0x23, 0x2d, 0x78, 0x38, 0xc9, 0x91, 0xf6, 0xc7, 0xd1, 0x77, 0xae, 0x1d, 0xe7, 0xa0, 0xf7,
	0xb0, 0xd0, 0xee, 0xf1, 0xd0, 0x73, 0x87, 0xf3, 0xfb, 0x78, 0x12, 0x13, 0xf5, 0x38, 0x97, 0xed,
	0xc9, 0x7d, 0xaf, 0x4c, 0x6d, 0x6f, 0xc1, 0x3a, 0xe3, 0x71, 0x66, 0x20, 0xf8, 0xbb, 0xcb, 0x9b,
	0x20, 0xdb, 0x05, 0xec, 0xc9, 0xe3, 0xe8, 0xcd, 0x3e, 0x36, 0xae, 0x0c, 0xe3, 0x2c, 0xaf, 0xdf,
	0xef, 0x8d, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0xe2, 0x2a, 0x63, 0x9d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v3alpha.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct {
}

func (*UnimplementedRateLimitServiceServer) ShouldRateLimit(ctx context.Context, req *RateLimitRequest) (*RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v3alpha.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v3alpha.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v3alpha/rls.proto",
}

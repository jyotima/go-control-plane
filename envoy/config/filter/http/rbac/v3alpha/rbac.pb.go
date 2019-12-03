// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rbac/v3alpha/rbac.proto

package envoy_config_filter_http_rbac_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RBAC struct {
	Rules                *v3alpha.RBAC `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v3alpha.RBAC `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f8a7c98dce62dc2, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

type RBACPerRoute struct {
	Rbac                 *RBAC    `protobuf:"bytes,2,opt,name=rbac,proto3" json:"rbac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RBACPerRoute) Reset()         { *m = RBACPerRoute{} }
func (m *RBACPerRoute) String() string { return proto.CompactTextString(m) }
func (*RBACPerRoute) ProtoMessage()    {}
func (*RBACPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f8a7c98dce62dc2, []int{1}
}

func (m *RBACPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBACPerRoute.Unmarshal(m, b)
}
func (m *RBACPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBACPerRoute.Marshal(b, m, deterministic)
}
func (m *RBACPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBACPerRoute.Merge(m, src)
}
func (m *RBACPerRoute) XXX_Size() int {
	return xxx_messageInfo_RBACPerRoute.Size(m)
}
func (m *RBACPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_RBACPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_RBACPerRoute proto.InternalMessageInfo

func (m *RBACPerRoute) GetRbac() *RBAC {
	if m != nil {
		return m.Rbac
	}
	return nil
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.http.rbac.v3alpha.RBAC")
	proto.RegisterType((*RBACPerRoute)(nil), "envoy.config.filter.http.rbac.v3alpha.RBACPerRoute")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rbac/v3alpha/rbac.proto", fileDescriptor_0f8a7c98dce62dc2)
}

var fileDescriptor_0f8a7c98dce62dc2 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc7, 0x49, 0xe9, 0xf7, 0xa1, 0x69, 0x0f, 0xb2, 0x17, 0xa5, 0x07, 0x95, 0x62, 0x51, 0x50,
	0x13, 0x69, 0xf5, 0xd2, 0x8b, 0xb8, 0xde, 0x3c, 0x2d, 0xfb, 0x02, 0x32, 0xbb, 0x9b, 0x76, 0x03,
	0x4b, 0x26, 0x64, 0xb3, 0xab, 0x7d, 0x03, 0xc1, 0x37, 0xf0, 0x25, 0x7c, 0x0a, 0xdf, 0x4b, 0x92,
	0x6c, 0xc1, 0x1e, 0xc4, 0xbd, 0x4d, 0xc8, 0xff, 0x37, 0xbf, 0x61, 0x86, 0xde, 0x08, 0xd5, 0xe2,
	0x86, 0xe7, 0xa8, 0x56, 0x72, 0xcd, 0x57, 0xb2, 0xb2, 0xc2, 0xf0, 0xd2, 0x5a, 0xcd, 0x4d, 0x06,
	0x39, 0x6f, 0x17, 0x50, 0xe9, 0x12, 0xfc, 0x83, 0x69, 0x83, 0x16, 0xa3, 0x99, 0x27, 0x58, 0x20,
	0x58, 0x20, 0x98, 0x23, 0x98, 0x0f, 0x75, 0xc4, 0xe4, 0x6c, 0xa7, 0xf1, 0x2f, 0xcd, 0x26, 0xb3,
	0xa6, 0xd0, 0xc0, 0x41, 0x4b, 0x0e, 0x4a, 0xa1, 0x05, 0x2b, 0x51, 0xd5, 0xbc, 0x15, 0xa6, 0x96,
	0xa8, 0xa4, 0x5a, 0x77, 0xb1, 0xc3, 0x16, 0x2a, 0x59, 0x80, 0x15, 0x7c, 0x5b, 0x84, 0x8f, 0xe9,
	0x27, 0xa1, 0xc3, 0x34, 0x7e, 0x78, 0x8c, 0xee, 0xe8, 0x3f, 0xd3, 0x54, 0xa2, 0x3e, 0x22, 0xa7,
	0xe4, 0x62, 0x34, 0x3f, 0x61, 0x3b, 0x53, 0xfe, 0x9c, 0x8c, 0xb9, 0x7c, 0x1a, 0xd2, 0x51, 0x4c,
	0xc7, 0x75, 0x09, 0x05, 0xbe, 0x3c, 0x07, 0x7a, 0xd0, 0x8f, 0x1e, 0x05, 0x28, 0x75, 0xcc, 0xf2,
	0xea, 0xe3, 0xeb, 0xed, 0xf8, 0x9c, 0xfe, 0xb5, 0x97, 0xb9, 0x47, 0xa7, 0xef, 0x84, 0x8e, 0x5d,
	0x91, 0x08, 0x93, 0x62, 0x63, 0x45, 0x74, 0x4f, 0x87, 0x2e, 0xd0, 0xa9, 0x2f, 0x59, 0xaf, 0xf5,
	0x86, 0x31, 0x3c, 0xb8, 0xbc, 0x75, 0x7e, 0x4e, 0xaf, 0x7b, 0xf9, 0xb7, 0xda, 0xa7, 0xe1, 0x1e,
	0x39, 0x18, 0xc4, 0x31, 0x5d, 0x48, 0x0c, 0x4a, 0x6d, 0xf0, 0x75, 0xd3, 0xcf, 0x1e, 0xef, 0xa7,
	0x19, 0xe4, 0x89, 0xbb, 0x40, 0x42, 0xb2, 0xff, 0xfe, 0x14, 0x8b, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0x51, 0x0a, 0xb0, 0x4b, 0x02, 0x00, 0x00,
}

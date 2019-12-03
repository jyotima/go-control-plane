// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v3alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/common/dynamic_forward_proxy/v3alpha"
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

type FilterConfig struct {
	DnsCacheConfig       *v3alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0bf0b4ef6d47a8b, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v3alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

type PerRouteConfig struct {
	// Types that are valid to be assigned to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewrite
	//	*PerRouteConfig_AutoHostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0bf0b4ef6d47a8b, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
}

type PerRouteConfig_HostRewrite struct {
	HostRewrite string `protobuf:"bytes,1,opt,name=host_rewrite,json=hostRewrite,proto3,oneof"`
}

type PerRouteConfig_AutoHostRewriteHeader struct {
	AutoHostRewriteHeader string `protobuf:"bytes,2,opt,name=auto_host_rewrite_header,json=autoHostRewriteHeader,proto3,oneof"`
}

func (*PerRouteConfig_HostRewrite) isPerRouteConfig_HostRewriteSpecifier() {}

func (*PerRouteConfig_AutoHostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (m *PerRouteConfig) GetHostRewrite() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewrite); ok {
		return x.HostRewrite
	}
	return ""
}

func (m *PerRouteConfig) GetAutoHostRewriteHeader() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_AutoHostRewriteHeader); ok {
		return x.AutoHostRewriteHeader
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerRouteConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerRouteConfig_HostRewrite)(nil),
		(*PerRouteConfig_AutoHostRewriteHeader)(nil),
	}
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v3alpha.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v3alpha.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v3alpha/dynamic_forward_proxy.proto", fileDescriptor_b0bf0b4ef6d47a8b)
}

var fileDescriptor_b0bf0b4ef6d47a8b = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x87, 0x6f, 0x0a, 0xf7, 0x72, 0xef, 0xb4, 0x94, 0x12, 0xb8, 0x1a, 0xba, 0x10, 0xa9, 0x08,
	0xae, 0x32, 0xd0, 0x82, 0x60, 0x77, 0x4d, 0xff, 0x58, 0xc4, 0x45, 0xc8, 0x0b, 0x84, 0x69, 0x32,
	0x69, 0x06, 0xda, 0x39, 0xc3, 0x64, 0x92, 0xda, 0x37, 0x10, 0x1f, 0xc1, 0xf7, 0xf1, 0x01, 0xdc,
	0xf9, 0x2c, 0xae, 0x64, 0x32, 0x51, 0x1b, 0x28, 0x8a, 0xdd, 0x95, 0xf9, 0x9d, 0xf3, 0x9d, 0xef,
	0x34, 0x07, 0x05, 0x94, 0x17, 0xb0, 0xc5, 0x11, 0xf0, 0x84, 0x2d, 0x71, 0xc2, 0x56, 0x8a, 0x4a,
	0x9c, 0x2a, 0x25, 0x70, 0xbc, 0xe5, 0x64, 0xcd, 0xa2, 0x30, 0x01, 0xb9, 0x21, 0x32, 0x0e, 0x85,
	0x84, 0xbb, 0x2d, 0x2e, 0x06, 0x64, 0x25, 0x52, 0xb2, 0x3f, 0x75, 0x85, 0x04, 0x05, 0xf6, 0x65,
	0xc9, 0x74, 0x0d, 0xd3, 0x35, 0x4c, 0x57, 0x33, 0xdd, 0xfd, 0x5d, 0x15, 0xb3, 0x3b, 0xaa, 0xb9,
	0x44, 0xb0, 0x5e, 0x03, 0xff, 0x4e, 0x83, 0x67, 0x61, 0x44, 0xa2, 0x94, 0x9a, 0xd1, 0xdd, 0xf3,
	0x3c, 0x16, 0x04, 0x13, 0xc1, 0x30, 0xe1, 0x1c, 0x14, 0x51, 0x0c, 0x78, 0x86, 0x0b, 0x2a, 0x33,
	0x06, 0x9c, 0xf1, 0x65, 0x55, 0x76, 0x5c, 0x90, 0x15, 0x8b, 0x89, 0xa2, 0xf8, 0xfd, 0x87, 0x09,
	0x7a, 0xcf, 0x16, 0x6a, 0xcd, 0x4a, 0xe1, 0x71, 0x69, 0x61, 0xe7, 0xa8, 0xf3, 0x31, 0x23, 0x34,
	0x66, 0x8e, 0x75, 0x6a, 0x5d, 0x34, 0xfb, 0x23, 0xb7, 0xb6, 0xa6, 0xd1, 0xfd, 0x7a, 0x43, 0x77,
	0xc2, 0xb3, 0xb1, 0x26, 0x19, 0xb8, 0xf7, 0xf7, 0xd5, 0xfb, 0xfd, 0x60, 0x35, 0x3a, 0x56, 0xd0,
	0x8e, 0x6b, 0xc9, 0xf0, 0xe6, 0xf1, 0xe9, 0xfe, 0x64, 0x8a, 0xc6, 0x3f, 0xfd, 0x27, 0xfb, 0x66,
	0xce, 0xee, 0x0a, 0xbd, 0x17, 0x0b, 0xb5, 0x7d, 0x2a, 0x03, 0xc8, 0x55, 0x85, 0xb7, 0xcf, 0x50,
	0x2b, 0x85, 0x4c, 0x85, 0x92, 0x6e, 0x24, 0x53, 0xb4, 0xdc, 0xe8, 0xdf, 0xfc, 0x57, 0xd0, 0xd4,
	0xaf, 0x81, 0x79, 0xb4, 0xaf, 0x90, 0x43, 0x72, 0x05, 0xe1, 0x6e, 0x65, 0x98, 0x52, 0x12, 0x53,
	0xe9, 0x34, 0xaa, 0x86, 0xff, 0xba, 0x62, 0xfe, 0xd9, 0x34, 0x2f, 0xe3, 0xe1, 0xad, 0xd6, 0xbf,
	0x46, 0xd3, 0x03, 0xf5, 0xeb, 0xb6, 0x9e, 0x83, 0x8e, 0x6a, 0x0e, 0x99, 0xa0, 0x11, 0x4b, 0x18,
	0x95, 0xde, 0x02, 0x4d, 0x18, 0x98, 0xef, 0x60, 0x08, 0x87, 0x5d, 0x9e, 0xe7, 0x4c, 0x4c, 0x3c,
	0x33, 0xa9, 0xaf, 0x43, 0x5f, 0x1f, 0x84, 0x6f, 0x2d, 0xfe, 0x94, 0x97, 0x31, 0x78, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x26, 0x62, 0x8e, 0x70, 0x2a, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3alpha/metadata.proto

package envoy_type_matcher_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
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

type MetadataMatcher struct {
	Filter               string                         `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Path                 []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Value                *ValueMatcher                  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_d306cd29ee202178, []int{0}
}

func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher.Unmarshal(m, b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(m, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher.Size(m)
}
func (m *MetadataMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher proto.InternalMessageInfo

func (m *MetadataMatcher) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *MetadataMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

type MetadataMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment              isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MetadataMatcher_PathSegment) Reset()         { *m = MetadataMatcher_PathSegment{} }
func (m *MetadataMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher_PathSegment) ProtoMessage()    {}
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d306cd29ee202178, []int{0, 0}
}

func (m *MetadataMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Unmarshal(m, b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(m, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Size(m)
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.v3alpha.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.v3alpha.MetadataMatcher.PathSegment")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3alpha/metadata.proto", fileDescriptor_d306cd29ee202178)
}

var fileDescriptor_d306cd29ee202178 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x9d, 0xa4, 0x37, 0xa7, 0x78, 0x21, 0x1b, 0x4b, 0x04, 0x0d, 0x45, 0x25, 0xdd, 0x64,
	0xa4, 0x45, 0x84, 0xba, 0x9b, 0x55, 0x37, 0x85, 0x12, 0xd1, 0xfd, 0xd1, 0x8e, 0xcd, 0x60, 0x3a,
	0x33, 0xa6, 0xd3, 0x60, 0xde, 0x40, 0x5c, 0xba, 0xf4, 0x7d, 0x7c, 0x20, 0x77, 0xd2, 0x95, 0x24,
	0x99, 0x82, 0x14, 0x8d, 0xbb, 0x70, 0xf2, 0xcd, 0xf7, 0x9f, 0x9f, 0x83, 0x7b, 0x4c, 0xa4, 0x32,
	0x23, 0x3a, 0x53, 0x8c, 0xcc, 0x41, 0xdf, 0x47, 0x2c, 0x21, 0xe9, 0x00, 0x62, 0x15, 0x01, 0x99,
	0x33, 0x0d, 0x53, 0xd0, 0x10, 0xa8, 0x44, 0x6a, 0xe9, 0xb8, 0x05, 0x1a, 0xe4, 0x68, 0x60, 0xd0,
	0xc0, 0xa0, 0xee, 0x59, 0x85, 0x26, 0x85, 0x78, 0xc9, 0x4a, 0x87, 0x7b, 0xba, 0x9c, 0x2a, 0x20,
	0xa0, 0x38, 0x01, 0x21, 0xa4, 0x06, 0xcd, 0xa5, 0x58, 0x90, 0x94, 0x25, 0x0b, 0x2e, 0x05, 0x17,
	0x33, 0x83, 0x1d, 0xa4, 0x10, 0xf3, 0x29, 0x68, 0x46, 0xd6, 0x1f, 0xe5, 0x8f, 0xee, 0xa7, 0x85,
	0xf7, 0xc6, 0x66, 0xad, 0x71, 0x99, 0xe3, 0x1c, 0xe3, 0xc6, 0x03, 0x8f, 0x35, 0x4b, 0x3a, 0xc8,
	0x43, 0xfe, 0x36, 0x6d, 0xae, 0x68, 0x2d, 0xb1, 0x3c, 0x14, 0x9a, 0xb1, 0x73, 0x83, 0x6b, 0x0a,
	0x74, 0xd4, 0xb1, 0x3c, 0xdb, 0x6f, 0xf7, 0x2f, 0x83, 0xbf, 0x7b, 0x04, 0x1b, 0xee, 0x60, 0x02,
	0x3a, 0xba, 0x66, 0xb3, 0x39, 0x13, 0x9a, 0xb6, 0x56, 0xb4, 0xfe, 0x86, 0xac, 0x16, 0x0a, 0x0b,
	0x9d, 0x33, 0xc2, 0xf5, 0xa2, 0x5a, 0xc7, 0xf6, 0x90, 0xdf, 0xee, 0xfb, 0x55, 0xde, 0xdb, 0x1c,
	0x34, 0xd2, 0x42, 0xf4, 0x8a, 0xac, 0x7d, 0x14, 0x96, 0x02, 0xf7, 0x09, 0xb7, 0x7f, 0x04, 0x39,
	0x87, 0xd8, 0x7e, 0x64, 0xd9, 0x46, 0x9b, 0xd1, 0x56, 0x98, 0x4f, 0x87, 0x17, 0xef, 0x1f, 0x2f,
	0x47, 0xe7, 0xf8, 0xb7, 0xb0, 0xaa, 0xe5, 0x77, 0x71, 0x73, 0x61, 0xf4, 0xf6, 0x17, 0x45, 0xc3,
	0x5e, 0xae, 0x39, 0xc1, 0xdd, 0xff, 0x35, 0xf4, 0x0a, 0xfb, 0x5c, 0x96, 0x79, 0x2a, 0x91, 0xcf,
	0x59, 0x45, 0x4f, 0xba, 0xb3, 0x7e, 0x3c, 0xc9, 0xcf, 0x35, 0x41, 0x77, 0x8d, 0xe2, 0x6e, 0x83,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xb4, 0xbc, 0xe7, 0x68, 0x02, 0x00, 0x00,
}

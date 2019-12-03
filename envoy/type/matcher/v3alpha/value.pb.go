// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3alpha/value.proto

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

type ValueMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*ValueMatcher_NullMatch_
	//	*ValueMatcher_DoubleMatch
	//	*ValueMatcher_StringMatch
	//	*ValueMatcher_BoolMatch
	//	*ValueMatcher_PresentMatch
	//	*ValueMatcher_ListMatch
	MatchPattern         isValueMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ValueMatcher) Reset()         { *m = ValueMatcher{} }
func (m *ValueMatcher) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher) ProtoMessage()    {}
func (*ValueMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e35f51f298eeff, []int{0}
}

func (m *ValueMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher.Unmarshal(m, b)
}
func (m *ValueMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher.Marshal(b, m, deterministic)
}
func (m *ValueMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher.Merge(m, src)
}
func (m *ValueMatcher) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher.Size(m)
}
func (m *ValueMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher proto.InternalMessageInfo

type isValueMatcher_MatchPattern interface {
	isValueMatcher_MatchPattern()
}

type ValueMatcher_NullMatch_ struct {
	NullMatch *ValueMatcher_NullMatch `protobuf:"bytes,1,opt,name=null_match,json=nullMatch,proto3,oneof"`
}

type ValueMatcher_DoubleMatch struct {
	DoubleMatch *DoubleMatcher `protobuf:"bytes,2,opt,name=double_match,json=doubleMatch,proto3,oneof"`
}

type ValueMatcher_StringMatch struct {
	StringMatch *StringMatcher `protobuf:"bytes,3,opt,name=string_match,json=stringMatch,proto3,oneof"`
}

type ValueMatcher_BoolMatch struct {
	BoolMatch bool `protobuf:"varint,4,opt,name=bool_match,json=boolMatch,proto3,oneof"`
}

type ValueMatcher_PresentMatch struct {
	PresentMatch bool `protobuf:"varint,5,opt,name=present_match,json=presentMatch,proto3,oneof"`
}

type ValueMatcher_ListMatch struct {
	ListMatch *ListMatcher `protobuf:"bytes,6,opt,name=list_match,json=listMatch,proto3,oneof"`
}

func (*ValueMatcher_NullMatch_) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_DoubleMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_StringMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_BoolMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_PresentMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_ListMatch) isValueMatcher_MatchPattern() {}

func (m *ValueMatcher) GetMatchPattern() isValueMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ValueMatcher) GetNullMatch() *ValueMatcher_NullMatch {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_NullMatch_); ok {
		return x.NullMatch
	}
	return nil
}

func (m *ValueMatcher) GetDoubleMatch() *DoubleMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_DoubleMatch); ok {
		return x.DoubleMatch
	}
	return nil
}

func (m *ValueMatcher) GetStringMatch() *StringMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_StringMatch); ok {
		return x.StringMatch
	}
	return nil
}

func (m *ValueMatcher) GetBoolMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_BoolMatch); ok {
		return x.BoolMatch
	}
	return false
}

func (m *ValueMatcher) GetPresentMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_PresentMatch); ok {
		return x.PresentMatch
	}
	return false
}

func (m *ValueMatcher) GetListMatch() *ListMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_ListMatch); ok {
		return x.ListMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValueMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValueMatcher_NullMatch_)(nil),
		(*ValueMatcher_DoubleMatch)(nil),
		(*ValueMatcher_StringMatch)(nil),
		(*ValueMatcher_BoolMatch)(nil),
		(*ValueMatcher_PresentMatch)(nil),
		(*ValueMatcher_ListMatch)(nil),
	}
}

type ValueMatcher_NullMatch struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueMatcher_NullMatch) Reset()         { *m = ValueMatcher_NullMatch{} }
func (m *ValueMatcher_NullMatch) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher_NullMatch) ProtoMessage()    {}
func (*ValueMatcher_NullMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e35f51f298eeff, []int{0, 0}
}

func (m *ValueMatcher_NullMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher_NullMatch.Unmarshal(m, b)
}
func (m *ValueMatcher_NullMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher_NullMatch.Marshal(b, m, deterministic)
}
func (m *ValueMatcher_NullMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher_NullMatch.Merge(m, src)
}
func (m *ValueMatcher_NullMatch) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher_NullMatch.Size(m)
}
func (m *ValueMatcher_NullMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher_NullMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher_NullMatch proto.InternalMessageInfo

type ListMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*ListMatcher_OneOf
	MatchPattern         isListMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListMatcher) Reset()         { *m = ListMatcher{} }
func (m *ListMatcher) String() string { return proto.CompactTextString(m) }
func (*ListMatcher) ProtoMessage()    {}
func (*ListMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e35f51f298eeff, []int{1}
}

func (m *ListMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatcher.Unmarshal(m, b)
}
func (m *ListMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatcher.Marshal(b, m, deterministic)
}
func (m *ListMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatcher.Merge(m, src)
}
func (m *ListMatcher) XXX_Size() int {
	return xxx_messageInfo_ListMatcher.Size(m)
}
func (m *ListMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatcher proto.InternalMessageInfo

type isListMatcher_MatchPattern interface {
	isListMatcher_MatchPattern()
}

type ListMatcher_OneOf struct {
	OneOf *ValueMatcher `protobuf:"bytes,1,opt,name=one_of,json=oneOf,proto3,oneof"`
}

func (*ListMatcher_OneOf) isListMatcher_MatchPattern() {}

func (m *ListMatcher) GetMatchPattern() isListMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ListMatcher) GetOneOf() *ValueMatcher {
	if x, ok := m.GetMatchPattern().(*ListMatcher_OneOf); ok {
		return x.OneOf
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListMatcher_OneOf)(nil),
	}
}

func init() {
	proto.RegisterType((*ValueMatcher)(nil), "envoy.type.matcher.v3alpha.ValueMatcher")
	proto.RegisterType((*ValueMatcher_NullMatch)(nil), "envoy.type.matcher.v3alpha.ValueMatcher.NullMatch")
	proto.RegisterType((*ListMatcher)(nil), "envoy.type.matcher.v3alpha.ListMatcher")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3alpha/value.proto", fileDescriptor_65e35f51f298eeff)
}

var fileDescriptor_65e35f51f298eeff = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0xee, 0xae, 0xa2, 0xaf, 0xbd, 0x25, 0x42, 0x02, 0x75, 0xb8, 0x2b, 0x27, 0xf5,
	0x2e, 0x15, 0x52, 0x8c, 0xda, 0x05, 0x55, 0x62, 0x20, 0x62, 0xc8, 0x00, 0xa5, 0x6a, 0x25, 0xd6,
	0xca, 0x21, 0x2e, 0xb5, 0xe4, 0xda, 0x96, 0xe3, 0x44, 0xf4, 0x1b, 0xb0, 0xb0, 0x30, 0xf2, 0x7d,
	0xf8, 0x4e, 0x8c, 0xc8, 0x8e, 0x43, 0x32, 0x84, 0xf4, 0xb6, 0xf8, 0xe5, 0xf7, 0xfe, 0xff, 0xff,
	0x7b, 0x7a, 0x70, 0x4f, 0x78, 0x21, 0x4e, 0x48, 0x9f, 0x24, 0x41, 0x47, 0xac, 0xbf, 0x1c, 0x88,
	0x42, 0xc5, 0x02, 0x33, 0x79, 0xc0, 0xa8, 0xc0, 0x2c, 0x27, 0xa1, 0x54, 0x42, 0x0b, 0x7f, 0x6c,
	0xb9, 0xd0, 0x70, 0xa1, 0xe3, 0x42, 0xc7, 0x8d, 0x1f, 0x3a, 0x34, 0x78, 0x7e, 0x4c, 0x88, 0x2a,
	0x45, 0x3a, 0xc1, 0x4c, 0x2b, 0xca, 0xbf, 0x3a, 0x70, 0x9a, 0xa7, 0x12, 0x23, 0x2c, 0x29, 0xc2,
	0x9c, 0x0b, 0x8d, 0x35, 0x15, 0x3c, 0x43, 0x05, 0x51, 0x19, 0x15, 0xbc, 0xc6, 0x9e, 0x17, 0x98,
	0xd1, 0x14, 0x6b, 0x82, 0xaa, 0x8f, 0xf2, 0xc7, 0xdd, 0xcf, 0x4b, 0x18, 0x7d, 0x36, 0xe9, 0x3f,
	0x96, 0x2e, 0xfe, 0x16, 0x80, 0xe7, 0x8c, 0xed, 0xac, 0xeb, 0x0b, 0x6f, 0xe2, 0x05, 0xc3, 0xf9,
	0x3c, 0xfc, 0xff, 0x4c, 0x61, 0xb3, 0x3b, 0x5c, 0xe5, 0x8c, 0xd9, 0xef, 0xb8, 0xb7, 0x19, 0xf0,
	0xea, 0xe1, 0xaf, 0x60, 0x94, 0x8a, 0x3c, 0x61, 0xc4, 0xc9, 0x3e, 0xb1, 0xb2, 0xb3, 0x2e, 0xd9,
	0xf7, 0x96, 0x77, 0xba, 0x71, 0x6f, 0x33, 0x4c, 0xeb, 0x82, 0xd1, 0x2b, 0xb7, 0xe0, 0xf4, 0x2e,
	0xce, 0xeb, 0x6d, 0x2d, 0xdf, 0xd0, 0xcb, 0xea, 0x82, 0x7f, 0x0b, 0x90, 0x08, 0x51, 0x0d, 0x7d,
	0x39, 0xf1, 0x82, 0xa7, 0x66, 0x00, 0x53, 0x2b, 0x81, 0x29, 0x5c, 0x4b, 0x45, 0x32, 0xc2, 0xb5,
	0x63, 0xae, 0x1c, 0x33, 0x72, 0xe5, 0x12, 0x8b, 0x01, 0x18, 0xcd, 0x2a, 0xa6, 0x6f, 0x53, 0x3d,
	0x74, 0xa5, 0xfa, 0x40, 0x33, 0x5d, 0x67, 0x1a, 0xb0, 0xea, 0x39, 0x7e, 0x0b, 0x83, 0x7f, 0xbb,
	0x5c, 0xbe, 0xfe, 0xf5, 0xfb, 0xfb, 0xcd, 0x2b, 0x98, 0xb5, 0x08, 0xb5, 0x6f, 0x7f, 0x79, 0x6f,
	0x3a, 0x5e, 0xc2, 0xed, 0x99, 0x8e, 0xe8, 0x19, 0x5c, 0xdb, 0xfa, 0x4e, 0x62, 0xad, 0x89, 0xe2,
	0xfe, 0xc5, 0x9f, 0xc8, 0xbb, 0xfb, 0xe1, 0xc1, 0xb0, 0x91, 0xcc, 0x7f, 0x07, 0x7d, 0xc1, 0xc9,
	0x4e, 0xec, 0xdd, 0x3d, 0x04, 0x8f, 0xbd, 0x87, 0xb8, 0xb7, 0xb9, 0x12, 0x9c, 0x7c, 0xda, 0x2f,
	0xa7, 0x26, 0xd0, 0x04, 0x6e, 0x5a, 0x1a, 0x1b, 0x4e, 0xed, 0x79, 0xa2, 0x37, 0x10, 0x50, 0x51,
	0x7a, 0x4a, 0x25, 0xbe, 0x9d, 0x3a, 0xec, 0x23, 0xb0, 0xfe, 0x6b, 0x73, 0xdc, 0x6b, 0x2f, 0xe9,
	0xdb, 0x2b, 0x5f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x69, 0x33, 0x4b, 0xbd, 0x03, 0x00,
	0x00,
}

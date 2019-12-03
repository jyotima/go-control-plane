// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/overload/v3alpha/overload.proto

package envoy_config_overload_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type ResourceMonitor struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ResourceMonitor_TypedConfig
	ConfigType           isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResourceMonitor) Reset()         { *m = ResourceMonitor{} }
func (m *ResourceMonitor) String() string { return proto.CompactTextString(m) }
func (*ResourceMonitor) ProtoMessage()    {}
func (*ResourceMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{0}
}

func (m *ResourceMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMonitor.Unmarshal(m, b)
}
func (m *ResourceMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMonitor.Marshal(b, m, deterministic)
}
func (m *ResourceMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMonitor.Merge(m, src)
}
func (m *ResourceMonitor) XXX_Size() int {
	return xxx_messageInfo_ResourceMonitor.Size(m)
}
func (m *ResourceMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMonitor proto.InternalMessageInfo

func (m *ResourceMonitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
}

type ResourceMonitor_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ResourceMonitor) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResourceMonitor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResourceMonitor_TypedConfig)(nil),
	}
}

type ThresholdTrigger struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThresholdTrigger) Reset()         { *m = ThresholdTrigger{} }
func (m *ThresholdTrigger) String() string { return proto.CompactTextString(m) }
func (*ThresholdTrigger) ProtoMessage()    {}
func (*ThresholdTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{1}
}

func (m *ThresholdTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThresholdTrigger.Unmarshal(m, b)
}
func (m *ThresholdTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThresholdTrigger.Marshal(b, m, deterministic)
}
func (m *ThresholdTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThresholdTrigger.Merge(m, src)
}
func (m *ThresholdTrigger) XXX_Size() int {
	return xxx_messageInfo_ThresholdTrigger.Size(m)
}
func (m *ThresholdTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_ThresholdTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_ThresholdTrigger proto.InternalMessageInfo

func (m *ThresholdTrigger) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Trigger struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TriggerOneof:
	//	*Trigger_Threshold
	TriggerOneof         isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{2}
}

func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (m *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(m, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,proto3,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (m *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := m.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Trigger) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Trigger_Threshold)(nil),
	}
}

type OverloadAction struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Triggers             []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OverloadAction) Reset()         { *m = OverloadAction{} }
func (m *OverloadAction) String() string { return proto.CompactTextString(m) }
func (*OverloadAction) ProtoMessage()    {}
func (*OverloadAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{3}
}

func (m *OverloadAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadAction.Unmarshal(m, b)
}
func (m *OverloadAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadAction.Marshal(b, m, deterministic)
}
func (m *OverloadAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadAction.Merge(m, src)
}
func (m *OverloadAction) XXX_Size() int {
	return xxx_messageInfo_OverloadAction.Size(m)
}
func (m *OverloadAction) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadAction.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadAction proto.InternalMessageInfo

func (m *OverloadAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OverloadAction) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type OverloadManager struct {
	RefreshInterval      *duration.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	ResourceMonitors     []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors,proto3" json:"resource_monitors,omitempty"`
	Actions              []*OverloadAction  `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverloadManager) Reset()         { *m = OverloadManager{} }
func (m *OverloadManager) String() string { return proto.CompactTextString(m) }
func (*OverloadManager) ProtoMessage()    {}
func (*OverloadManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{4}
}

func (m *OverloadManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadManager.Unmarshal(m, b)
}
func (m *OverloadManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadManager.Marshal(b, m, deterministic)
}
func (m *OverloadManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadManager.Merge(m, src)
}
func (m *OverloadManager) XXX_Size() int {
	return xxx_messageInfo_OverloadManager.Size(m)
}
func (m *OverloadManager) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadManager.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadManager proto.InternalMessageInfo

func (m *OverloadManager) GetRefreshInterval() *duration.Duration {
	if m != nil {
		return m.RefreshInterval
	}
	return nil
}

func (m *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if m != nil {
		return m.ResourceMonitors
	}
	return nil
}

func (m *OverloadManager) GetActions() []*OverloadAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceMonitor)(nil), "envoy.config.overload.v3alpha.ResourceMonitor")
	proto.RegisterType((*ThresholdTrigger)(nil), "envoy.config.overload.v3alpha.ThresholdTrigger")
	proto.RegisterType((*Trigger)(nil), "envoy.config.overload.v3alpha.Trigger")
	proto.RegisterType((*OverloadAction)(nil), "envoy.config.overload.v3alpha.OverloadAction")
	proto.RegisterType((*OverloadManager)(nil), "envoy.config.overload.v3alpha.OverloadManager")
}

func init() {
	proto.RegisterFile("envoy/config/overload/v3alpha/overload.proto", fileDescriptor_542f910eb6259f77)
}

var fileDescriptor_542f910eb6259f77 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xd3, 0x6e, 0xed, 0x5c, 0x4a, 0x4b, 0x54, 0x69, 0xed, 0x80, 0xa9, 0xaa, 0x34, 0x28,
	0x62, 0x8d, 0x51, 0x0b, 0x07, 0x7a, 0x00, 0x2d, 0x4c, 0x62, 0x20, 0xa6, 0x4d, 0xd1, 0xee, 0x95,
	0xd7, 0xb8, 0x69, 0xa4, 0xcc, 0x8e, 0x1c, 0x27, 0x5a, 0xff, 0x01, 0x67, 0x8e, 0xfc, 0x0a, 0xee,
	0xc0, 0x91, 0xff, 0xc3, 0x11, 0xf5, 0x84, 0x62, 0x3b, 0x9d, 0x9a, 0x69, 0xed, 0x72, 0x8a, 0xde,
	0xfb, 0xbe, 0xf7, 0xbe, 0xef, 0xd3, 0x33, 0x3c, 0x24, 0x34, 0x61, 0x73, 0x34, 0x61, 0x74, 0xea,
	0x7b, 0x88, 0x25, 0x84, 0x07, 0x0c, 0xbb, 0x28, 0x19, 0xe2, 0x20, 0x9c, 0xe1, 0x65, 0xc1, 0x0a,
	0x39, 0x13, 0xcc, 0x7c, 0x2a, 0xd1, 0x96, 0x42, 0x5b, 0xcb, 0xa6, 0x46, 0xef, 0xb5, 0x3d, 0xc6,
	0xbc, 0x80, 0x20, 0x09, 0xbe, 0x8c, 0xa7, 0x08, 0xd3, 0xb9, 0x62, 0xee, 0xed, 0xe7, 0x5b, 0x6e,
	0xcc, 0xb1, 0xf0, 0x19, 0xd5, 0xfd, 0x27, 0xf9, 0x7e, 0x24, 0x78, 0x3c, 0x11, 0xba, 0x7b, 0x10,
	0xbb, 0x21, 0x46, 0x38, 0xf4, 0x11, 0xa6, 0x94, 0x09, 0x49, 0x8c, 0x50, 0x42, 0x78, 0xe4, 0x33,
	0xea, 0x53, 0x4f, 0xc3, 0x76, 0x13, 0x1c, 0xf8, 0x2e, 0x16, 0x04, 0x65, 0x3f, 0xaa, 0xd1, 0xfd,
	0x05, 0x60, 0xdd, 0x21, 0x11, 0x8b, 0xf9, 0x84, 0x9c, 0x32, 0xea, 0x0b, 0xc6, 0xcd, 0xc7, 0xb0,
	0x44, 0xf1, 0x15, 0x69, 0x81, 0x0e, 0xe8, 0xed, 0xd8, 0xe5, 0x85, 0x5d, 0xe2, 0x46, 0x07, 0x38,
	0xb2, 0x68, 0xbe, 0x85, 0x0f, 0xc4, 0x3c, 0x24, 0xee, 0x58, 0x59, 0x6d, 0x15, 0x3b, 0xa0, 0x57,
	0x1d, 0x34, 0x2d, 0xa5, 0xd2, 0xca, 0x54, 0x5a, 0x47, 0x74, 0x7e, 0x52, 0x70, 0xaa, 0x12, 0xfb,
	0x41, 0x42, 0x47, 0xaf, 0xbf, 0xff, 0xf9, 0xba, 0x8f, 0x60, 0xff, 0x8e, 0xa8, 0x06, 0x32, 0x2a,
	0x2b, 0xa7, 0xc6, 0xae, 0xc1, 0xaa, 0x82, 0x8e, 0xd3, 0x59, 0x9f, 0x4b, 0x15, 0xa3, 0x51, 0x74,
	0xb6, 0x55, 0xa9, 0x7b, 0x0d, 0x1b, 0x17, 0x33, 0x4e, 0xa2, 0x19, 0x0b, 0xdc, 0x0b, 0xee, 0x7b,
	0x1e, 0xe1, 0x66, 0x1f, 0x6e, 0x25, 0x38, 0x88, 0x95, 0x7e, 0x60, 0xef, 0x2e, 0xec, 0xa6, 0x69,
	0xb6, 0x0b, 0xf2, 0xfb, 0xfb, 0xfe, 0x45, 0x41, 0x7f, 0x8e, 0x42, 0x8d, 0xde, 0xa4, 0xaa, 0x5e,
	0x41, 0x6b, 0xbd, 0xaa, 0xfc, 0x96, 0xee, 0x4f, 0x00, 0xcb, 0xd9, 0xc6, 0xb5, 0x81, 0x9d, 0xc1,
	0x1d, 0x91, 0x91, 0x5b, 0x86, 0x4c, 0x0b, 0x59, 0x6b, 0xaf, 0xe5, 0xd6, 0xb2, 0x93, 0x82, 0x73,
	0x33, 0x63, 0x74, 0x98, 0x0a, 0x7e, 0x0e, 0x0f, 0x36, 0x08, 0x56, 0x54, 0xbb, 0x09, 0x6b, 0x42,
	0xfd, 0x8e, 0x19, 0x25, 0x6c, 0x6a, 0x16, 0xff, 0xd9, 0xa0, 0xfb, 0x03, 0xc0, 0x87, 0x67, 0x9a,
	0x72, 0x34, 0x49, 0x8f, 0x66, 0xbd, 0x89, 0x2f, 0xb0, 0xa2, 0xa7, 0x44, 0x2d, 0xa3, 0x53, 0xec,
	0x55, 0x07, 0xcf, 0x36, 0x79, 0xd0, 0xfb, 0x2b, 0x0b, 0x7b, 0xeb, 0x1b, 0x30, 0x2a, 0xc0, 0x59,
	0x4e, 0x18, 0x0d, 0x53, 0x07, 0x96, 0x7e, 0x61, 0x77, 0x3a, 0x58, 0xd5, 0xd7, 0xfd, 0x6d, 0xc0,
	0x7a, 0x56, 0x3a, 0xc5, 0x14, 0xa7, 0xc1, 0x1f, 0xc3, 0x06, 0x27, 0xd3, 0x34, 0x98, 0xb1, 0x4f,
	0x05, 0xe1, 0x09, 0x0e, 0xa4, 0xfe, 0xea, 0xa0, 0x7d, 0xeb, 0x20, 0x8f, 0xf5, 0xb3, 0x72, 0xea,
	0x9a, 0xf2, 0x49, 0x33, 0x4c, 0x02, 0x1f, 0x71, 0x7d, 0x74, 0xe3, 0x2b, 0x75, 0x75, 0x99, 0x4b,
	0x6b, 0x83, 0xcb, 0xfc, 0xb1, 0xde, 0xb8, 0x6d, 0xf0, 0xd5, 0x56, 0x64, 0x7e, 0x84, 0x65, 0x2c,
	0xad, 0x44, 0xad, 0xa2, 0x1c, 0xde, 0xdf, 0x30, 0x7c, 0x35, 0x00, 0x27, 0x63, 0xdf, 0xf3, 0x1d,
	0xe5, 0xb2, 0xb2, 0xdf, 0xc1, 0x97, 0x3e, 0x53, 0x1b, 0x43, 0xce, 0xae, 0xe7, 0xeb, 0x97, 0xdb,
	0xb5, 0x8c, 0x7f, 0x9e, 0x06, 0x78, 0x0e, 0x2e, 0xb7, 0x65, 0x92, 0xc3, 0xff, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc4, 0x66, 0xa2, 0x9e, 0x18, 0x05, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/discovery.proto

package envoy_api_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type DiscoveryRequest struct {
	VersionInfo          string         `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Node                 *core.Node     `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	ResourceNames        []string       `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	TypeUrl              string         `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResponseNonce        string         `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail          *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ac02770f363096, []int{0}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	VersionInfo          string             `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Resources            []*any.Any         `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	Canary               bool               `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	TypeUrl              string             `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Nonce                string             `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ControlPlane         *core.ControlPlane `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ac02770f363096, []int{1}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []*any.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *DiscoveryResponse) GetControlPlane() *core.ControlPlane {
	if m != nil {
		return m.ControlPlane
	}
	return nil
}

type DeltaDiscoveryRequest struct {
	Node                     *core.Node        `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	TypeUrl                  string            `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResourceNamesSubscribe   []string          `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	ResourceNamesUnsubscribe []string          `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	InitialResourceVersions  map[string]string `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseNonce            string            `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail              *status.Status    `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *DeltaDiscoveryRequest) Reset()         { *m = DeltaDiscoveryRequest{} }
func (m *DeltaDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryRequest) ProtoMessage()    {}
func (*DeltaDiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ac02770f363096, []int{2}
}

func (m *DeltaDiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryRequest.Unmarshal(m, b)
}
func (m *DeltaDiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryRequest.Merge(m, src)
}
func (m *DeltaDiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryRequest.Size(m)
}
func (m *DeltaDiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryRequest proto.InternalMessageInfo

func (m *DeltaDiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if m != nil {
		return m.ResourceNamesSubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if m != nil {
		return m.ResourceNamesUnsubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if m != nil {
		return m.InitialResourceVersions
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DeltaDiscoveryResponse struct {
	SystemVersionInfo    string      `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	Resources            []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	TypeUrl              string      `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	RemovedResources     []string    `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	Nonce                string      `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeltaDiscoveryResponse) Reset()         { *m = DeltaDiscoveryResponse{} }
func (m *DeltaDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryResponse) ProtoMessage()    {}
func (*DeltaDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ac02770f363096, []int{3}
}

func (m *DeltaDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryResponse.Unmarshal(m, b)
}
func (m *DeltaDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryResponse.Merge(m, src)
}
func (m *DeltaDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryResponse.Size(m)
}
func (m *DeltaDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryResponse proto.InternalMessageInfo

func (m *DeltaDiscoveryResponse) GetSystemVersionInfo() string {
	if m != nil {
		return m.SystemVersionInfo
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetRemovedResources() []string {
	if m != nil {
		return m.RemovedResources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type Resource struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Aliases              []string `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Resource             *any.Any `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_94ac02770f363096, []int{4}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetResource() *any.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v3alpha.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v3alpha.DiscoveryResponse")
	proto.RegisterType((*DeltaDiscoveryRequest)(nil), "envoy.api.v3alpha.DeltaDiscoveryRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.api.v3alpha.DeltaDiscoveryRequest.InitialResourceVersionsEntry")
	proto.RegisterType((*DeltaDiscoveryResponse)(nil), "envoy.api.v3alpha.DeltaDiscoveryResponse")
	proto.RegisterType((*Resource)(nil), "envoy.api.v3alpha.Resource")
}

func init() { proto.RegisterFile("envoy/api/v3alpha/discovery.proto", fileDescriptor_94ac02770f363096) }

var fileDescriptor_94ac02770f363096 = []byte{
	// 748 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0x96, 0x93, 0x10, 0x92, 0x09, 0x20, 0xb2, 0x2f, 0x1f, 0x26, 0x2f, 0x85, 0x90, 0x12, 0x29,
	0x6d, 0x25, 0x1b, 0x05, 0x55, 0xa2, 0x51, 0x2f, 0xa5, 0x70, 0xa0, 0x07, 0x84, 0x8c, 0xe0, 0x6a,
	0x6d, 0x9c, 0x85, 0xae, 0x6a, 0x76, 0xdd, 0x5d, 0x3b, 0xaa, 0x8f, 0xbd, 0xf5, 0xd2, 0x3f, 0xd0,
	0x63, 0xa5, 0x1e, 0x7b, 0xec, 0x4f, 0xe8, 0xff, 0xaa, 0xbc, 0x5e, 0xe7, 0x83, 0x98, 0x28, 0x37,
	0xcf, 0xcc, 0x33, 0xb3, 0x33, 0xcf, 0x3c, 0x63, 0x38, 0x20, 0x6c, 0xc8, 0x63, 0x1b, 0x07, 0xd4,
	0x1e, 0x1e, 0x63, 0x3f, 0xf8, 0x88, 0xed, 0x01, 0x95, 0x1e, 0x1f, 0x12, 0x11, 0x5b, 0x81, 0xe0,
	0x21, 0x47, 0x75, 0x05, 0xb1, 0x70, 0x40, 0x2d, 0x0d, 0x69, 0xe4, 0x64, 0x79, 0x5c, 0x10, 0xbb,
	0x8f, 0x25, 0x49, 0xb3, 0x1a, 0x3b, 0xf7, 0x9c, 0xdf, 0xfb, 0xc4, 0x56, 0x56, 0x3f, 0xba, 0xb3,
	0x31, 0xd3, 0x05, 0x1b, 0xdb, 0x3a, 0x24, 0x02, 0xcf, 0x96, 0x21, 0x0e, 0x23, 0xa9, 0x03, 0xed,
	0x68, 0x10, 0x60, 0x55, 0x15, 0x33, 0xc6, 0x43, 0x1c, 0x52, 0xce, 0xa4, 0x3d, 0x24, 0x42, 0x52,
	0xce, 0x28, 0xbb, 0x4f, 0x61, 0xad, 0xdf, 0x05, 0x58, 0x3f, 0xcb, 0x9a, 0x74, 0xc8, 0xe7, 0x88,
	0xc8, 0x10, 0x1d, 0xc0, 0x8a, 0x06, 0xba, 0x94, 0xdd, 0x71, 0xd3, 0x68, 0x1a, 0x9d, 0xaa, 0x53,
	0xd3, 0xbe, 0x0b, 0x76, 0xc7, 0xd1, 0x11, 0x94, 0x18, 0x1f, 0x10, 0xb3, 0xd0, 0x34, 0x3a, 0xb5,
	0xee, 0xae, 0x35, 0x33, 0x97, 0x95, 0x0c, 0x61, 0x5d, 0xf2, 0x01, 0x71, 0x14, 0x12, 0xb5, 0x61,
	0x4d, 0x10, 0xc9, 0x23, 0xe1, 0x11, 0x97, 0xe1, 0x07, 0x22, 0xcd, 0x62, 0xb3, 0xd8, 0xa9, 0x3a,
	0xab, 0x99, 0xf7, 0x32, 0x71, 0xa2, 0x1d, 0xa8, 0x84, 0x71, 0x40, 0xdc, 0x48, 0xf8, 0x66, 0x49,
	0xbd, 0xbb, 0x9c, 0xd8, 0x37, 0xc2, 0xd7, 0x15, 0x02, 0xce, 0x24, 0x71, 0x19, 0x67, 0x1e, 0x31,
	0x97, 0x14, 0x60, 0x35, 0xf3, 0x5e, 0x26, 0x4e, 0xf4, 0x1a, 0x56, 0x88, 0x10, 0x5c, 0xb8, 0x03,
	0x12, 0x62, 0xea, 0x9b, 0x65, 0xd5, 0x22, 0xb2, 0x52, 0xa6, 0x2c, 0x11, 0x78, 0xd6, 0xb5, 0x62,
	0xca, 0xa9, 0x29, 0xdc, 0x99, 0x82, 0xf5, 0x0e, 0x7f, 0xfc, 0xfd, 0xb6, 0xb7, 0x0f, 0xcf, 0x26,
	0x26, 0xe9, 0x5a, 0x8f, 0xa9, 0x69, 0xfd, 0x2a, 0x40, 0x7d, 0xc2, 0x99, 0xbe, 0xbb, 0x08, 0x61,
	0x5d, 0xa8, 0x66, 0x83, 0x4a, 0xb3, 0xd0, 0x2c, 0x76, 0x6a, 0xdd, 0x8d, 0xac, 0xa5, 0x6c, 0xaf,
	0xd6, 0x3b, 0x16, 0x3b, 0x63, 0x18, 0xda, 0x82, 0xb2, 0x87, 0x19, 0x16, 0xb1, 0x59, 0x6c, 0x1a,
	0x9d, 0x8a, 0xa3, 0xad, 0x79, 0x1c, 0x6d, 0xc0, 0xd2, 0x24, 0x35, 0xa9, 0x81, 0x2e, 0x60, 0xd5,
	0xe3, 0x2c, 0x14, 0xdc, 0x77, 0x03, 0x1f, 0x33, 0xa2, 0x39, 0x39, 0x7c, 0x6a, 0x6d, 0xef, 0x53,
	0xf0, 0x55, 0x82, 0x75, 0x56, 0xbc, 0x09, 0xab, 0xd7, 0x4e, 0x68, 0x6a, 0xc2, 0xde, 0x53, 0x34,
	0xa5, 0x8c, 0xb4, 0xfe, 0x94, 0x60, 0xf3, 0x8c, 0xf8, 0x21, 0x9e, 0x11, 0x57, 0xa6, 0x1c, 0x63,
	0x61, 0xe5, 0x4c, 0x8e, 0x5b, 0x98, 0x1e, 0xf7, 0x04, 0xcc, 0x69, 0x51, 0xb9, 0x32, 0xea, 0x4b,
	0x4f, 0xd0, 0x3e, 0xd1, 0xf2, 0xda, 0x9a, 0x92, 0xd7, 0x75, 0x16, 0x45, 0x6f, 0xa1, 0xf1, 0x28,
	0x33, 0x62, 0xe3, 0xdc, 0x92, 0xca, 0x35, 0xa7, 0x72, 0x6f, 0xc6, 0x71, 0xf4, 0xd5, 0x80, 0x1d,
	0xca, 0x68, 0x48, 0xb1, 0xef, 0x8e, 0xca, 0xe8, 0x75, 0x4b, 0x73, 0x49, 0xad, 0xf7, 0x3c, 0x67,
	0xb4, 0x5c, 0x4a, 0xac, 0x8b, 0xb4, 0x92, 0xa3, 0x0b, 0xdd, 0xea, 0x3a, 0xe7, 0x2c, 0x14, 0xb1,
	0xb3, 0x4d, 0xf3, 0xa3, 0x39, 0xe7, 0x50, 0x5e, 0xe4, 0x1c, 0x96, 0x17, 0x3a, 0x87, 0xc6, 0x07,
	0xd8, 0x9d, 0xd7, 0x16, 0x5a, 0x87, 0xe2, 0x27, 0x12, 0x6b, 0xa5, 0x27, 0x9f, 0x89, 0xf4, 0x86,
	0xd8, 0x8f, 0x88, 0xde, 0x51, 0x6a, 0xf4, 0x0a, 0x27, 0x46, 0xef, 0x45, 0xa2, 0x99, 0x43, 0x68,
	0x4d, 0x6b, 0x26, 0x8f, 0x8a, 0xd6, 0xf7, 0x02, 0x6c, 0x3d, 0x8e, 0xe8, 0x23, 0xb3, 0xe0, 0x3f,
	0x19, 0xcb, 0x90, 0x3c, 0xb8, 0x39, 0xb7, 0x56, 0x4f, 0x43, 0xb7, 0x13, 0x17, 0xf7, 0x66, 0xf6,
	0xe2, 0xfe, 0xcf, 0x59, 0x49, 0x36, 0xde, 0xe4, 0xe1, 0xcd, 0x39, 0xb0, 0x57, 0x50, 0x17, 0xe4,
	0x81, 0x0f, 0xc9, 0xc0, 0x1d, 0x57, 0x2f, 0x2b, 0xb9, 0xac, 0xeb, 0x80, 0x33, 0xaa, 0x93, 0x7b,
	0x8d, 0xbd, 0x97, 0x09, 0x1d, 0x6d, 0x78, 0x3e, 0x97, 0x0e, 0x7d, 0x47, 0x3f, 0x0d, 0xa8, 0x64,
	0xf5, 0x10, 0x82, 0x52, 0x22, 0x55, 0xf5, 0x37, 0xa8, 0x3a, 0xea, 0x1b, 0x99, 0xb0, 0x8c, 0x7d,
	0x8a, 0x25, 0x91, 0x5a, 0xb4, 0x99, 0x99, 0x44, 0x34, 0x51, 0x9a, 0xa3, 0xcc, 0x44, 0x47, 0x50,
	0xc9, 0x7a, 0xd7, 0x3f, 0xf0, 0xfc, 0x5f, 0xd1, 0x08, 0xd5, 0xdb, 0x4d, 0x5a, 0xde, 0x86, 0xcd,
	0xa9, 0x96, 0xb3, 0xbe, 0x4e, 0x8f, 0x61, 0x9f, 0xf2, 0x94, 0xda, 0x40, 0xf0, 0x2f, 0xf1, 0x2c,
	0xcb, 0xa7, 0x6b, 0xa3, 0xd1, 0xae, 0x92, 0x17, 0xae, 0x8c, 0x7e, 0x59, 0x3d, 0x75, 0xfc, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x87, 0xf9, 0x12, 0xc0, 0x36, 0x07, 0x00, 0x00,
}

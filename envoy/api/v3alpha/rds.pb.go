// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/rds.proto

package envoy_api_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type RouteConfiguration struct {
	Name                            string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VirtualHosts                    []*route.VirtualHost      `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	Vhds                            *Vhds                     `protobuf:"bytes,9,opt,name=vhds,proto3" json:"vhds,omitempty"`
	InternalOnlyHeaders             []string                  `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	ResponseHeadersToAdd            []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	ResponseHeadersToRemove         []string                  `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	RequestHeadersToAdd             []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove          []string                  `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	MostSpecificHeaderMutationsWins bool                      `protobuf:"varint,10,opt,name=most_specific_header_mutations_wins,json=mostSpecificHeaderMutationsWins,proto3" json:"most_specific_header_mutations_wins,omitempty"`
	ValidateClusters                *wrappers.BoolValue       `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                  `json:"-"`
	XXX_unrecognized                []byte                    `json:"-"`
	XXX_sizecache                   int32                     `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaba07addf704458, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetVhds() *Vhds {
	if m != nil {
		return m.Vhds
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetMostSpecificHeaderMutationsWins() bool {
	if m != nil {
		return m.MostSpecificHeaderMutationsWins
	}
	return false
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

type Vhds struct {
	ConfigSource         *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Vhds) Reset()         { *m = Vhds{} }
func (m *Vhds) String() string { return proto.CompactTextString(m) }
func (*Vhds) ProtoMessage()    {}
func (*Vhds) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaba07addf704458, []int{1}
}

func (m *Vhds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vhds.Unmarshal(m, b)
}
func (m *Vhds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vhds.Marshal(b, m, deterministic)
}
func (m *Vhds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vhds.Merge(m, src)
}
func (m *Vhds) XXX_Size() int {
	return xxx_messageInfo_Vhds.Size(m)
}
func (m *Vhds) XXX_DiscardUnknown() {
	xxx_messageInfo_Vhds.DiscardUnknown(m)
}

var xxx_messageInfo_Vhds proto.InternalMessageInfo

func (m *Vhds) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v3alpha.RouteConfiguration")
	proto.RegisterType((*Vhds)(nil), "envoy.api.v3alpha.Vhds")
}

func init() { proto.RegisterFile("envoy/api/v3alpha/rds.proto", fileDescriptor_eaba07addf704458) }

var fileDescriptor_eaba07addf704458 = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x41, 0x4f, 0xdb, 0x48,
	0x18, 0x5d, 0x27, 0x01, 0x92, 0x09, 0x48, 0x64, 0x58, 0x88, 0x09, 0x2b, 0x08, 0x81, 0x5d, 0x05,
	0x56, 0x72, 0x50, 0x38, 0x6d, 0xf6, 0xb4, 0x01, 0xed, 0xb2, 0xd2, 0xae, 0x40, 0x4e, 0x95, 0x1e,
	0xad, 0xc1, 0x1e, 0x92, 0x51, 0x9d, 0x19, 0x77, 0x66, 0x6c, 0x9a, 0x4b, 0x0f, 0x3d, 0xa1, 0x4a,
	0xbd, 0xb4, 0xb7, 0xfe, 0x9f, 0xfe, 0x82, 0xfe, 0x83, 0xb6, 0x87, 0xfe, 0x06, 0x4e, 0x95, 0xc7,
	0xe3, 0x36, 0xc1, 0x6e, 0x55, 0xa9, 0xed, 0x25, 0xb2, 0xf3, 0xbd, 0xf7, 0xbe, 0x37, 0xdf, 0x3c,
	0x7f, 0x60, 0x0b, 0xd3, 0x88, 0x4d, 0x3b, 0x28, 0x20, 0x9d, 0xe8, 0x18, 0xf9, 0xc1, 0x18, 0x75,
	0xb8, 0x27, 0xac, 0x80, 0x33, 0xc9, 0x60, 0x4d, 0x15, 0x2d, 0x14, 0x10, 0x4b, 0x17, 0x1b, 0xbb,
	0x59, 0xbc, 0xcb, 0x38, 0xee, 0x5c, 0x22, 0x81, 0x13, 0x56, 0xe3, 0xf0, 0x33, 0x10, 0x97, 0xd1,
	0x2b, 0x32, 0x72, 0x04, 0x0b, 0xb9, 0x9b, 0x62, 0x73, 0xe4, 0x3c, 0x22, 0x5c, 0x16, 0x61, 0x3e,
	0xd5, 0x90, 0xbd, 0x1c, 0x87, 0x2c, 0x94, 0x38, 0xf9, 0xd5, 0xa0, 0x5f, 0x46, 0x8c, 0x8d, 0x7c,
	0xac, 0x50, 0x88, 0x52, 0x26, 0x91, 0x24, 0x8c, 0xea, 0x73, 0x34, 0xb6, 0x75, 0x55, 0xbd, 0x5d,
	0x86, 0x57, 0x9d, 0x6b, 0x8e, 0x82, 0x00, 0xf3, 0xb4, 0xfe, 0x6b, 0xe8, 0x05, 0xe8, 0x2e, 0xb7,
	0x13, 0x61, 0x2e, 0x08, 0xa3, 0x84, 0x8e, 0x34, 0xac, 0x1e, 0x21, 0x9f, 0x78, 0x48, 0xe2, 0x4e,
	0xfa, 0x90, 0x14, 0x5a, 0x6f, 0x17, 0x00, 0xb4, 0x63, 0x37, 0x27, 0xea, 0x88, 0x21, 0x57, 0x0a,
	0x10, 0x82, 0x12, 0x45, 0x13, 0x6c, 0x1a, 0x4d, 0xa3, 0x5d, 0xb1, 0xd5, 0x33, 0xfc, 0x17, 0xac,
	0x44, 0x84, 0xcb, 0x10, 0xf9, 0xce, 0x98, 0x09, 0x29, 0xcc, 0x42, 0xb3, 0xd8, 0xae, 0x76, 0xf7,
	0xad, 0xcc, 0xa8, 0xad, 0xe4, 0x7c, 0xc3, 0x04, 0x7d, 0xc6, 0x84, 0xb4, 0x97, 0xa3, 0x4f, 0x2f,
	0x02, 0xfe, 0x0e, 0x4a, 0xd1, 0xd8, 0x13, 0x66, 0xa5, 0x69, 0xb4, 0xab, 0xdd, 0x7a, 0x8e, 0xc2,
	0x70, 0xec, 0x09, 0x5b, 0x81, 0x60, 0x17, 0xac, 0x13, 0x2a, 0x31, 0xa7, 0xc8, 0x77, 0x18, 0xf5,
	0xa7, 0xce, 0x18, 0x23, 0x0f, 0x73, 0x61, 0x16, 0x9b, 0xc5, 0x76, 0xc5, 0x5e, 0x4b, 0x8b, 0xe7,
	0xd4, 0x9f, 0x9e, 0x25, 0x25, 0xf8, 0x00, 0xd4, 0x39, 0x16, 0x01, 0xa3, 0x02, 0xa7, 0x70, 0x47,
	0x32, 0x07, 0x79, 0x9e, 0x59, 0x52, 0xae, 0x0f, 0x72, 0x7a, 0xc6, 0x57, 0x6d, 0x25, 0x0a, 0x43,
	0xe4, 0x87, 0xf8, 0x3c, 0x88, 0x67, 0xd1, 0xaf, 0xdc, 0xf6, 0x17, 0x9f, 0x1b, 0xc5, 0xd5, 0xf7,
	0x4b, 0xf6, 0xcf, 0xa9, 0xa8, 0xee, 0x73, 0x8f, 0xfd, 0xe5, 0x79, 0xf0, 0x4f, 0xd0, 0xc8, 0x6b,
	0xc6, 0xf1, 0x84, 0x45, 0xd8, 0x5c, 0x50, 0x2e, 0xeb, 0x19, 0xa6, 0xad, 0xca, 0x90, 0x80, 0x0d,
	0x8e, 0x1f, 0x86, 0x58, 0xc8, 0xbb, 0x46, 0x17, 0xbf, 0xc1, 0xe8, 0x9a, 0xd6, 0x9c, 0xf3, 0xf9,
	0x07, 0xd8, 0xcc, 0x69, 0xa5, 0x6d, 0x96, 0x95, 0xcd, 0x8d, 0xbb, 0x3c, 0xed, 0xf2, 0x3f, 0xb0,
	0x37, 0x61, 0x42, 0x3a, 0x22, 0xc0, 0x2e, 0xb9, 0x22, 0xae, 0x16, 0x70, 0x26, 0xa1, 0x8e, 0x9c,
	0x73, 0x4d, 0xa8, 0x30, 0x41, 0xd3, 0x68, 0x97, 0xed, 0x9d, 0x18, 0x3a, 0xd0, 0xc8, 0x44, 0xe9,
	0xff, 0x14, 0x77, 0x9f, 0x50, 0x01, 0xff, 0x01, 0xb5, 0x34, 0x86, 0x8e, 0xeb, 0x87, 0x42, 0xc6,
	0xb7, 0xb9, 0xa4, 0xb2, 0xd0, 0xb0, 0x92, 0xc0, 0x5b, 0x69, 0xe0, 0xad, 0x3e, 0x63, 0xbe, 0x3a,
	0xa5, 0xbd, 0x9a, 0x92, 0x4e, 0x34, 0xa7, 0xf7, 0xdb, 0xcb, 0x57, 0x37, 0xdb, 0xbb, 0x60, 0x67,
	0x66, 0x44, 0x5d, 0x2b, 0x1b, 0xe7, 0x56, 0x08, 0x4a, 0x71, 0xa0, 0xe0, 0x00, 0xac, 0xcc, 0x7d,
	0xca, 0x2a, 0xdf, 0xf9, 0x11, 0x56, 0x33, 0x4e, 0x54, 0x06, 0x0a, 0xdb, 0x2f, 0xdf, 0xf6, 0x17,
	0x9e, 0x1a, 0x85, 0x55, 0xc3, 0x5e, 0x76, 0x67, 0xfe, 0xef, 0x99, 0xb1, 0x89, 0x35, 0x50, 0x9b,
	0x33, 0x11, 0xb7, 0xeb, 0xbe, 0x29, 0x80, 0x75, 0xe5, 0xe6, 0x34, 0x5d, 0x0c, 0x03, 0xcc, 0x23,
	0xe2, 0x62, 0xe8, 0x80, 0xe5, 0x81, 0xe4, 0x18, 0x4d, 0x54, 0x59, 0xc0, 0xbd, 0x1c, 0x07, 0x1f,
	0x49, 0x76, 0x72, 0x27, 0x8d, 0xfd, 0x2f, 0x83, 0x92, 0x7c, 0xb5, 0x7e, 0x6a, 0x1b, 0x47, 0x06,
	0x1c, 0x83, 0xea, 0x29, 0xf6, 0x25, 0xd2, 0xfa, 0xed, 0x3c, 0x6a, 0x5c, 0xcf, 0x34, 0x39, 0xf8,
	0x0a, 0xe4, 0x5c, 0xa7, 0xc7, 0xa0, 0xfa, 0x37, 0x96, 0xee, 0xf8, 0xfb, 0x9f, 0x64, 0xff, 0xc9,
	0xeb, 0x77, 0x2f, 0x0a, 0xdb, 0xad, 0xcd, 0xec, 0x82, 0xed, 0xa9, 0x0d, 0x23, 0x7a, 0xc6, 0x61,
	0xf7, 0x99, 0x01, 0xb6, 0x66, 0x36, 0x4d, 0x66, 0xd4, 0x14, 0xd4, 0x94, 0xff, 0xe1, 0xec, 0x02,
	0xfa, 0x71, 0xf3, 0xe8, 0x1f, 0x81, 0x1d, 0xc2, 0x12, 0x52, 0xc0, 0xd9, 0xa3, 0x69, 0x96, 0xdf,
	0x2f, 0xdb, 0x9e, 0xb8, 0x88, 0xf3, 0x7d, 0x61, 0xdc, 0x18, 0xc6, 0xe5, 0xa2, 0xca, 0xfa, 0xf1,
	0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x4e, 0xe7, 0x75, 0xd1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error)
	DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error)
	FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) StreamRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_StreamRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.RouteDiscoveryService/StreamRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceStreamRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_StreamRoutesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceStreamRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceStreamRoutesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) DeltaRoutes(ctx context.Context, opts ...grpc.CallOption) (RouteDiscoveryService_DeltaRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v3alpha.RouteDiscoveryService/DeltaRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeDiscoveryServiceDeltaRoutesClient{stream}
	return x, nil
}

type RouteDiscoveryService_DeltaRoutesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type routeDiscoveryServiceDeltaRoutesClient struct {
	grpc.ClientStream
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeDiscoveryServiceClient) FetchRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v3alpha.RouteDiscoveryService/FetchRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteDiscoveryServiceServer is the server API for RouteDiscoveryService service.
type RouteDiscoveryServiceServer interface {
	StreamRoutes(RouteDiscoveryService_StreamRoutesServer) error
	DeltaRoutes(RouteDiscoveryService_DeltaRoutesServer) error
	FetchRoutes(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedRouteDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRouteDiscoveryServiceServer struct {
}

func (*UnimplementedRouteDiscoveryServiceServer) StreamRoutes(srv RouteDiscoveryService_StreamRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRoutes not implemented")
}
func (*UnimplementedRouteDiscoveryServiceServer) DeltaRoutes(srv RouteDiscoveryService_DeltaRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRoutes not implemented")
}
func (*UnimplementedRouteDiscoveryServiceServer) FetchRoutes(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRoutes not implemented")
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_StreamRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).StreamRoutes(&routeDiscoveryServiceStreamRoutesServer{stream})
}

type RouteDiscoveryService_StreamRoutesServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceStreamRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceStreamRoutesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceStreamRoutesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_DeltaRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteDiscoveryServiceServer).DeltaRoutes(&routeDiscoveryServiceDeltaRoutesServer{stream})
}

type RouteDiscoveryService_DeltaRoutesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type routeDiscoveryServiceDeltaRoutesServer struct {
	grpc.ServerStream
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeDiscoveryServiceDeltaRoutesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteDiscoveryService_FetchRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v3alpha.RouteDiscoveryService/FetchRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).FetchRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRoutes",
			Handler:    _RouteDiscoveryService_FetchRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRoutes",
			Handler:       _RouteDiscoveryService_StreamRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRoutes",
			Handler:       _RouteDiscoveryService_DeltaRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/rds.proto",
}

// VirtualHostDiscoveryServiceClient is the client API for VirtualHostDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHostDiscoveryServiceClient interface {
	DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error)
}

type virtualHostDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHostDiscoveryServiceClient(cc *grpc.ClientConn) VirtualHostDiscoveryServiceClient {
	return &virtualHostDiscoveryServiceClient{cc}
}

func (c *virtualHostDiscoveryServiceClient) DeltaVirtualHosts(ctx context.Context, opts ...grpc.CallOption) (VirtualHostDiscoveryService_DeltaVirtualHostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VirtualHostDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.VirtualHostDiscoveryService/DeltaVirtualHosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualHostDiscoveryServiceDeltaVirtualHostsClient{stream}
	return x, nil
}

type VirtualHostDiscoveryService_DeltaVirtualHostsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsClient struct {
	grpc.ClientStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualHostDiscoveryServiceServer is the server API for VirtualHostDiscoveryService service.
type VirtualHostDiscoveryServiceServer interface {
	DeltaVirtualHosts(VirtualHostDiscoveryService_DeltaVirtualHostsServer) error
}

// UnimplementedVirtualHostDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHostDiscoveryServiceServer struct {
}

func (*UnimplementedVirtualHostDiscoveryServiceServer) DeltaVirtualHosts(srv VirtualHostDiscoveryService_DeltaVirtualHostsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaVirtualHosts not implemented")
}

func RegisterVirtualHostDiscoveryServiceServer(s *grpc.Server, srv VirtualHostDiscoveryServiceServer) {
	s.RegisterService(&_VirtualHostDiscoveryService_serviceDesc, srv)
}

func _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VirtualHostDiscoveryServiceServer).DeltaVirtualHosts(&virtualHostDiscoveryServiceDeltaVirtualHostsServer{stream})
}

type VirtualHostDiscoveryService_DeltaVirtualHostsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type virtualHostDiscoveryServiceDeltaVirtualHostsServer struct {
	grpc.ServerStream
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *virtualHostDiscoveryServiceDeltaVirtualHostsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _VirtualHostDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.VirtualHostDiscoveryService",
	HandlerType: (*VirtualHostDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaVirtualHosts",
			Handler:       _VirtualHostDiscoveryService_DeltaVirtualHosts_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/rds.proto",
}

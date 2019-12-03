// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/core/health_check.proto

package envoy_api_v3alpha_core

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/api/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HealthStatus int32

const (
	HealthStatus_UNKNOWN   HealthStatus = 0
	HealthStatus_HEALTHY   HealthStatus = 1
	HealthStatus_UNHEALTHY HealthStatus = 2
	HealthStatus_DRAINING  HealthStatus = 3
	HealthStatus_TIMEOUT   HealthStatus = 4
	HealthStatus_DEGRADED  HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0}
}

type HealthCheck struct {
	Timeout               *duration.Duration    `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Interval              *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	InitialJitter         *duration.Duration    `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	IntervalJitter        *duration.Duration    `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	IntervalJitterPercent uint32                `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	UnhealthyThreshold    *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	HealthyThreshold      *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	AltPort               *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	ReuseConnection       *wrappers.BoolValue   `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker                isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	NoTrafficInterval            *duration.Duration          `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	UnhealthyInterval            *duration.Duration          `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	UnhealthyEdgeInterval        *duration.Duration          `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	HealthyEdgeInterval          *duration.Duration          `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	EventLogPath                 string                      `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	AlwaysLogHealthCheckFailures bool                        `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                    `json:"-"`
	XXX_unrecognized             []byte                      `json:"-"`
	XXX_sizecache                int32                       `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

type HealthCheck_HttpHealthCheck struct {
	Host                   string                  `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Path                   string                  `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Send                   *HealthCheck_Payload    `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	Receive                *HealthCheck_Payload    `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	ServiceName            string                  `protobuf:"bytes,5,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	RequestHeadersToAdd    []*HeaderValueOption    `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	RequestHeadersToRemove []string                `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	ExpectedStatuses       []*v3alpha.Int64Range   `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	CodecClientType        v3alpha.CodecClientType `protobuf:"varint,10,opt,name=codec_client_type,json=codecClientType,proto3,enum=envoy.type.v3alpha.CodecClientType" json:"codec_client_type,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*v3alpha.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetCodecClientType() v3alpha.CodecClientType {
	if m != nil {
		return m.CodecClientType
	}
	return v3alpha.CodecClientType_HTTP1
}

type HealthCheck_TcpHealthCheck struct {
	Send                 *HealthCheck_Payload   `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type HealthCheck_GrpcHealthCheck struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

type HealthCheck_CustomHealthCheck struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_092b49335160c834, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v3alpha.core.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.api.v3alpha.core.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.api.v3alpha.core.HealthCheck.CustomHealthCheck")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/core/health_check.proto", fileDescriptor_092b49335160c834)
}

var fileDescriptor_092b49335160c834 = []byte{
	// 1326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x72, 0xd3, 0x46,
	0x17, 0x8f, 0x6c, 0x27, 0xb6, 0x8f, 0x1d, 0x5b, 0xde, 0x10, 0x22, 0xfc, 0x05, 0x3e, 0xf3, 0x7d,
	0x30, 0x35, 0x94, 0xda, 0xad, 0x03, 0x74, 0xf0, 0x0d, 0x8d, 0x12, 0x83, 0x1d, 0xc0, 0x64, 0x16,
	0xa7, 0x0c, 0xc3, 0x4c, 0xd5, 0x8d, 0xb4, 0xb1, 0x55, 0x14, 0xad, 0xba, 0x5a, 0x1b, 0x7c, 0xd7,
	0xe1, 0x8a, 0x99, 0xde, 0xf5, 0xb2, 0x8f, 0xc0, 0x53, 0xf4, 0xa6, 0x97, 0x7d, 0x8f, 0x3e, 0x42,
	0x27, 0x57, 0x1d, 0xad, 0x64, 0xc7, 0x7f, 0x42, 0x30, 0x9d, 0xde, 0x69, 0xcf, 0x39, 0xbf, 0xdf,
	0x39, 0xbb, 0xfb, 0x3b, 0x67, 0x05, 0x37, 0xa8, 0x3b, 0x60, 0xc3, 0x2a, 0xf1, 0xec, 0xea, 0x60,
	0x8b, 0x38, 0x5e, 0x8f, 0x54, 0x4d, 0xc6, 0x69, 0xb5, 0x47, 0x89, 0x23, 0x7a, 0x86, 0xd9, 0xa3,
	0xe6, 0xab, 0x8a, 0xc7, 0x99, 0x60, 0xe8, 0xa2, 0x0c, 0xad, 0x10, 0xcf, 0xae, 0x44, 0xa1, 0x95,
	0x20, 0xb4, 0x78, 0xf5, 0x03, 0x14, 0x87, 0xc4, 0xa7, 0x21, 0xb4, 0x78, 0x39, 0x0c, 0x11, 0x43,
	0x8f, 0x8e, 0x63, 0x7a, 0x42, 0x78, 0x91, 0xfb, 0xca, 0x19, 0x6e, 0x4e, 0xdc, 0xee, 0x08, 0x7e,
	0xa9, 0xcb, 0x58, 0xd7, 0xa1, 0x55, 0xb9, 0x3a, 0xec, 0x1f, 0x55, 0x89, 0x3b, 0x1c, 0x41, 0x67,
	0x5d, 0x56, 0x9f, 0x13, 0x61, 0x33, 0x37, 0xf2, 0x6f, 0xce, 0xfa, 0x7d, 0xc1, 0xfb, 0xa6, 0xf8,
	0x10, 0xfa, 0x35, 0x27, 0x9e, 0x47, 0xb9, 0x1f, 0xf9, 0xaf, 0xf7, 0x2d, 0x8f, 0xc8, 0x9d, 0x11,
	0xd7, 0x65, 0x42, 0x12, 0xfb, 0xd5, 0x01, 0xe5, 0xbe, 0xcd, 0x5c, 0xdb, 0xed, 0x46, 0x61, 0x1b,
	0x03, 0xe2, 0xd8, 0x16, 0x11, 0xb4, 0x3a, 0xfa, 0x08, 0x1d, 0xff, 0xfb, 0x79, 0x03, 0x32, 0x4d,
	0x79, 0x92, 0x3b, 0xc1, 0x41, 0xa2, 0xfb, 0x90, 0x14, 0xf6, 0x31, 0x65, 0x7d, 0xa1, 0x29, 0x25,
	0xa5, 0x9c, 0xa9, 0x5d, 0xaa, 0x84, 0x15, 0x54, 0x46, 0x15, 0x54, 0x76, 0xa3, 0xfa, 0x75, 0x38,
	0xd1, 0x93, 0xef, 0x95, 0x44, 0x4a, 0xb9, 0xb9, 0x84, 0x47, 0x28, 0xb4, 0x0d, 0x29, 0xdb, 0x15,
	0x94, 0x0f, 0x88, 0xa3, 0xc5, 0x3e, 0x85, 0x61, 0x0c, 0x43, 0xdf, 0x40, 0xce, 0x76, 0x6d, 0x61,
	0x13, 0xc7, 0xf8, 0xc1, 0x16, 0x82, 0x72, 0xed, 0xc2, 0x47, 0x88, 0xf0, 0x6a, 0x04, 0xd8, 0x93,
	0xf1, 0x48, 0x87, 0xfc, 0x88, 0x6d, 0x44, 0x11, 0xff, 0x18, 0x45, 0x6e, 0x84, 0x88, 0x38, 0xee,
	0xc2, 0xc6, 0x0c, 0x87, 0xe1, 0x51, 0x6e, 0x52, 0x57, 0x68, 0xa8, 0xa4, 0x94, 0x57, 0xf1, 0xfa,
	0x34, 0x60, 0x3f, 0x74, 0xa2, 0x27, 0xb0, 0xd6, 0x77, 0x43, 0x71, 0x0e, 0x0d, 0xd1, 0xe3, 0xd4,
	0xef, 0x31, 0xc7, 0xd2, 0x12, 0x32, 0xff, 0xe6, 0x5c, 0xfe, 0x83, 0x96, 0x2b, 0xb6, 0x6a, 0xdf,
	0x12, 0xa7, 0x4f, 0x31, 0x1a, 0x03, 0x3b, 0x23, 0x1c, 0x6a, 0x41, 0x61, 0x9e, 0x6c, 0x79, 0x01,
	0x32, 0x75, 0x8e, 0xea, 0x6b, 0x48, 0x11, 0x47, 0x18, 0x1e, 0xe3, 0x42, 0x5b, 0x59, 0x80, 0x21,
	0x49, 0x1c, 0xb1, 0xcf, 0xb8, 0x40, 0x0d, 0x50, 0x39, 0xed, 0xfb, 0xd4, 0x30, 0x99, 0xeb, 0x52,
	0x33, 0x38, 0x2e, 0x2d, 0x29, 0x09, 0x8a, 0x73, 0x04, 0x3a, 0x63, 0x4e, 0x08, 0xcf, 0x4b, 0xcc,
	0xce, 0x18, 0x82, 0x08, 0x14, 0x82, 0x96, 0x32, 0x26, 0x3b, 0x57, 0x4b, 0x49, 0x9e, 0xad, 0xca,
	0xd9, 0xad, 0x5b, 0x99, 0xd0, 0x66, 0xa5, 0x29, 0x84, 0x37, 0xb1, 0x6e, 0x2e, 0xe1, 0x7c, 0x6f,
	0xda, 0x84, 0xbe, 0x03, 0x55, 0x98, 0x33, 0x19, 0xd2, 0x32, 0x43, 0x6d, 0x91, 0x0c, 0x1d, 0x73,
	0x26, 0x41, 0x4e, 0x4c, 0x59, 0x82, 0x2d, 0x74, 0xb9, 0x67, 0x4e, 0x27, 0xc8, 0x2c, 0xbe, 0x85,
	0x87, 0xdc, 0x33, 0x67, 0xb6, 0xd0, 0x9d, 0x36, 0xa1, 0x2e, 0xac, 0x99, 0x7d, 0x5f, 0xb0, 0xe3,
	0xe9, 0x24, 0xab, 0x32, 0xc9, 0x9d, 0x45, 0x92, 0xec, 0x48, 0xf8, 0x74, 0x9a, 0x82, 0x39, 0x6b,
	0x44, 0xcf, 0x60, 0xcd, 0x65, 0x86, 0xe0, 0xe4, 0xe8, 0xc8, 0x36, 0x8d, 0x71, 0xd3, 0x66, 0x3f,
	0xd6, 0xb4, 0xa9, 0x13, 0x7d, 0xf9, 0xbd, 0x12, 0xbb, 0xb9, 0x84, 0x0b, 0x2e, 0xeb, 0x84, 0xf0,
	0xd6, 0xa8, 0x77, 0x31, 0x9c, 0x8a, 0xf8, 0x94, 0x33, 0xf7, 0x09, 0x9c, 0x63, 0xf8, 0x98, 0xf3,
	0x25, 0x6c, 0x9c, 0x72, 0x52, 0xab, 0x4b, 0x4f, 0x89, 0xf3, 0x8b, 0x13, 0xaf, 0x8f, 0x39, 0x1a,
	0x56, 0x97, 0x8e, 0xc9, 0x9f, 0xc3, 0xfa, 0xd9, 0xd4, 0xea, 0xe2, 0xd4, 0x6b, 0x67, 0x11, 0x5f,
	0x83, 0x1c, 0x1d, 0x50, 0x57, 0x18, 0x0e, 0xeb, 0x1a, 0x1e, 0x11, 0x3d, 0xad, 0x50, 0x52, 0xca,
	0x69, 0x9c, 0x95, 0xd6, 0xc7, 0xac, 0xbb, 0x4f, 0x44, 0x0f, 0x3d, 0x80, 0x12, 0x71, 0x5e, 0x93,
	0xa1, 0x2f, 0xc3, 0x26, 0x6f, 0xdc, 0x38, 0x22, 0xb6, 0xd3, 0xe7, 0xd4, 0xd7, 0xd6, 0x4a, 0x4a,
	0x39, 0x85, 0x37, 0xc3, 0xb8, 0xc7, 0xac, 0x3b, 0x71, 0x89, 0x0f, 0xa2, 0x98, 0xe2, 0x4f, 0x0a,
	0x24, 0xf7, 0xc9, 0xd0, 0x61, 0xc4, 0x42, 0x97, 0x21, 0x21, 0xe8, 0x9b, 0x70, 0x80, 0xa7, 0xf5,
	0xe4, 0x89, 0x9e, 0xe0, 0xb1, 0x92, 0xd2, 0x5c, 0xc2, 0xd2, 0x8c, 0x34, 0x58, 0x39, 0xb4, 0x5d,
	0xc2, 0x87, 0x72, 0x3e, 0x67, 0x9b, 0x4b, 0x38, 0x5a, 0xd7, 0x6f, 0xfd, 0xfa, 0xfb, 0xbb, 0x2b,
	0x9f, 0xc1, 0xf5, 0x09, 0x8d, 0xd5, 0xe6, 0xe5, 0x15, 0xa5, 0xd1, 0x73, 0x90, 0xf4, 0xa2, 0x8c,
	0xf1, 0xbf, 0x74, 0xa5, 0xf8, 0x76, 0x19, 0xf2, 0x33, 0x2d, 0x8a, 0x10, 0x24, 0x7a, 0xcc, 0x8f,
	0x4a, 0xc1, 0xf2, 0x1b, 0xfd, 0x07, 0x12, 0xf2, 0x38, 0x62, 0x53, 0xe5, 0x61, 0x69, 0x44, 0xf7,
	0x21, 0xe1, 0x53, 0xd7, 0x8a, 0xc6, 0xf5, 0xe7, 0x8b, 0xc8, 0x3d, 0xaa, 0x07, 0x4b, 0x20, 0x6a,
	0x40, 0x92, 0x53, 0x93, 0xda, 0x03, 0x1a, 0x8d, 0xdc, 0x4f, 0xe2, 0x18, 0x61, 0xd1, 0x55, 0xc8,
	0xfa, 0x94, 0x0f, 0x6c, 0x93, 0x1a, 0x2e, 0x39, 0xa6, 0x72, 0xe2, 0xa6, 0x71, 0x26, 0xb2, 0xb5,
	0xc9, 0x31, 0x45, 0x36, 0x5c, 0xe4, 0xf4, 0xc7, 0x3e, 0xf5, 0x45, 0x70, 0x6f, 0x16, 0xe5, 0xbe,
	0x21, 0x98, 0x41, 0x2c, 0x4b, 0x5b, 0x29, 0xc5, 0xcb, 0x99, 0xda, 0x8d, 0x73, 0x12, 0x5b, 0x94,
	0xcb, 0x21, 0xf9, 0xd4, 0x93, 0x52, 0x4a, 0x9f, 0xe8, 0x2b, 0xbf, 0x28, 0x71, 0xf5, 0xcf, 0x24,
	0x5e, 0x8b, 0x38, 0xc3, 0x20, 0xbf, 0xc3, 0xb6, 0x2d, 0x0b, 0xdd, 0x83, 0x4b, 0x67, 0xa4, 0xe2,
	0xf4, 0x98, 0x0d, 0xa8, 0x96, 0x2a, 0xc5, 0xcb, 0x69, 0x7c, 0x71, 0x16, 0x87, 0xa5, 0x17, 0x3d,
	0x82, 0x02, 0x7d, 0xe3, 0x51, 0x53, 0x50, 0xcb, 0xf0, 0x05, 0x11, 0x7d, 0x9f, 0xfa, 0x5a, 0x5a,
	0x16, 0x78, 0x25, 0x2a, 0x30, 0xf8, 0xab, 0x19, 0x57, 0xd8, 0x72, 0xc5, 0xdd, 0xdb, 0x38, 0xf8,
	0xb5, 0xc1, 0xea, 0x08, 0xf8, 0x2c, 0xc2, 0xa1, 0x17, 0x50, 0x30, 0x99, 0x45, 0x4d, 0xc3, 0x74,
	0xec, 0x40, 0xda, 0x01, 0x52, 0x83, 0x92, 0x52, 0xce, 0xd5, 0xfe, 0x7f, 0x16, 0xd9, 0x4e, 0x10,
	0xbc, 0x23, 0x63, 0x3b, 0x43, 0x8f, 0xca, 0x96, 0x79, 0xab, 0xc4, 0x54, 0x05, 0xe7, 0xcd, 0x69,
	0x57, 0xfd, 0x76, 0xa0, 0xbd, 0x2a, 0x7c, 0x71, 0xbe, 0xf6, 0x66, 0xf4, 0xb5, 0x97, 0x48, 0x25,
	0xd5, 0x14, 0x4e, 0x07, 0x6f, 0x53, 0xf0, 0x14, 0xd4, 0x8a, 0x7f, 0x28, 0x90, 0x9b, 0x9e, 0xe2,
	0x63, 0x49, 0x29, 0xff, 0x82, 0xa4, 0x62, 0xf2, 0xe0, 0xfe, 0x91, 0xa4, 0xea, 0x5b, 0xc1, 0x0e,
	0x2b, 0x70, 0xeb, 0xfc, 0x1d, 0x4e, 0x17, 0x5f, 0x7c, 0x09, 0x2a, 0xa6, 0x96, 0xed, 0x4f, 0x6e,
	0x48, 0x85, 0xf8, 0x2b, 0x3a, 0x8c, 0x7a, 0x2a, 0xf8, 0xac, 0xdf, 0x09, 0xa8, 0xbf, 0x84, 0xca,
	0xf9, 0xd4, 0xb3, 0x44, 0xc5, 0x77, 0x0a, 0xe4, 0x67, 0x5e, 0xa4, 0x39, 0xe1, 0x2b, 0xf3, 0xc2,
	0xdf, 0x84, 0x34, 0xe9, 0x8b, 0x1e, 0xe3, 0xb6, 0x08, 0x67, 0x48, 0x1a, 0x9f, 0x1a, 0x16, 0xbc,
	0xc8, 0x99, 0xb4, 0xc5, 0xdf, 0x14, 0x28, 0xcc, 0xbd, 0x5b, 0xc1, 0xa8, 0x38, 0x2d, 0x62, 0x62,
	0x54, 0x04, 0x46, 0x74, 0x0f, 0xb2, 0x81, 0xd8, 0xac, 0xe0, 0xaf, 0xe4, 0xc8, 0xee, 0x46, 0x23,
	0xe3, 0xc2, 0xdc, 0xc0, 0xde, 0x76, 0x87, 0xcd, 0x25, 0x9c, 0x91, 0xb1, 0x3b, 0x32, 0xb4, 0x7e,
	0x37, 0xa8, 0xf1, 0x2b, 0xa8, 0x9e, 0x5f, 0xe3, 0x5c, 0x3d, 0xfa, 0x2a, 0x64, 0xc2, 0x64, 0x52,
	0xf9, 0x7b, 0x89, 0x54, 0x4c, 0x8d, 0xe3, 0x95, 0xd0, 0x54, 0xbf, 0x16, 0x90, 0xfe, 0x17, 0x2e,
	0x9f, 0x4b, 0xaa, 0xaf, 0x43, 0x6e, 0x72, 0xca, 0x53, 0x2e, 0x87, 0xe7, 0x5e, 0x22, 0x05, 0x6a,
	0xe6, 0xe6, 0xf7, 0x90, 0x0d, 0x63, 0xc3, 0x8e, 0x43, 0x19, 0x48, 0x1e, 0xb4, 0x1f, 0xb5, 0x9f,
	0x3e, 0x6f, 0xab, 0x4b, 0xc1, 0xa2, 0xd9, 0xd8, 0x7e, 0xdc, 0x69, 0xbe, 0x50, 0x15, 0xb4, 0x0a,
	0xe9, 0x83, 0xf6, 0x68, 0x19, 0x43, 0x59, 0x48, 0xed, 0xe2, 0xed, 0x56, 0xbb, 0xd5, 0x7e, 0xa8,
	0xc6, 0x83, 0xc8, 0x4e, 0xeb, 0x49, 0xe3, 0xe9, 0x41, 0x47, 0x4d, 0x48, 0x57, 0xe3, 0x21, 0xde,
	0xde, 0x6d, 0xec, 0xaa, 0xcb, 0x7a, 0x1d, 0xae, 0xd9, 0x2c, 0xd4, 0x89, 0xc7, 0xd9, 0x9b, 0xe1,
	0x07, 0x94, 0xac, 0xab, 0x13, 0x35, 0xef, 0x07, 0x27, 0xb9, 0xaf, 0x1c, 0xae, 0xc8, 0x23, 0xdd,
	0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x01, 0x87, 0x2a, 0x69, 0x92, 0x0d, 0x00, 0x00,
}

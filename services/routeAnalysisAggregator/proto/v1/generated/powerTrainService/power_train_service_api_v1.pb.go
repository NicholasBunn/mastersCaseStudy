// Code generated by protoc-gen-go. DO NOT EDIT.
// source: powerTrainService/proto/v1/power_train_service_api_v1.proto

package powerTrainService

import (
	fmt "fmt"
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

type ModelTypeEnum int32

const (
	ModelTypeEnum_UNKNOWN   ModelTypeEnum = 0
	ModelTypeEnum_OPENWATER ModelTypeEnum = 1
	ModelTypeEnum_ICE       ModelTypeEnum = 2
)

var ModelTypeEnum_name = map[int32]string{
	0: "UNKNOWN",
	1: "OPENWATER",
	2: "ICE",
}

var ModelTypeEnum_value = map[string]int32{
	"UNKNOWN":   0,
	"OPENWATER": 1,
	"ICE":       2,
}

func (x ModelTypeEnum) String() string {
	return proto.EnumName(ModelTypeEnum_name, int32(x))
}

func (ModelTypeEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{0}
}

type PowerTrainEstimateRequest struct {
	UnixTime              []float64     `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	PortPropMotorSpeed    []float32     `protobuf:"fixed32,2,rep,packed,name=port_prop_motor_speed,json=portPropMotorSpeed,proto3" json:"port_prop_motor_speed,omitempty"`
	StbdPropMotorSpeed    []float32     `protobuf:"fixed32,3,rep,packed,name=stbd_prop_motor_speed,json=stbdPropMotorSpeed,proto3" json:"stbd_prop_motor_speed,omitempty"`
	PropellerPitchPort    []float32     `protobuf:"fixed32,4,rep,packed,name=propeller_pitch_port,json=propellerPitchPort,proto3" json:"propeller_pitch_port,omitempty"`
	PropellerPitchStbd    []float32     `protobuf:"fixed32,5,rep,packed,name=propeller_pitch_stbd,json=propellerPitchStbd,proto3" json:"propeller_pitch_stbd,omitempty"`
	Sog                   []float32     `protobuf:"fixed32,6,rep,packed,name=sog,proto3" json:"sog,omitempty"`
	WindDirectionRelative []float32     `protobuf:"fixed32,7,rep,packed,name=wind_direction_relative,json=windDirectionRelative,proto3" json:"wind_direction_relative,omitempty"`
	WindSpeed             []float32     `protobuf:"fixed32,8,rep,packed,name=wind_speed,json=windSpeed,proto3" json:"wind_speed,omitempty"`
	BeaufortNumber        []uint32      `protobuf:"varint,9,rep,packed,name=beaufort_number,json=beaufortNumber,proto3" json:"beaufort_number,omitempty"`
	WaveDirection         []float32     `protobuf:"fixed32,10,rep,packed,name=wave_direction,json=waveDirection,proto3" json:"wave_direction,omitempty"`
	WaveLength            []float32     `protobuf:"fixed32,11,rep,packed,name=wave_length,json=waveLength,proto3" json:"wave_length,omitempty"`
	ModelType             ModelTypeEnum `protobuf:"varint,12,opt,name=model_type,json=modelType,proto3,enum=powerTrainService.v1.ModelTypeEnum" json:"model_type,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}      `json:"-"`
	XXX_unrecognized      []byte        `json:"-"`
	XXX_sizecache         int32         `json:"-"`
}

func (m *PowerTrainEstimateRequest) Reset()         { *m = PowerTrainEstimateRequest{} }
func (m *PowerTrainEstimateRequest) String() string { return proto.CompactTextString(m) }
func (*PowerTrainEstimateRequest) ProtoMessage()    {}
func (*PowerTrainEstimateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{0}
}

func (m *PowerTrainEstimateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowerTrainEstimateRequest.Unmarshal(m, b)
}
func (m *PowerTrainEstimateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowerTrainEstimateRequest.Marshal(b, m, deterministic)
}
func (m *PowerTrainEstimateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerTrainEstimateRequest.Merge(m, src)
}
func (m *PowerTrainEstimateRequest) XXX_Size() int {
	return xxx_messageInfo_PowerTrainEstimateRequest.Size(m)
}
func (m *PowerTrainEstimateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerTrainEstimateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PowerTrainEstimateRequest proto.InternalMessageInfo

func (m *PowerTrainEstimateRequest) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetPortPropMotorSpeed() []float32 {
	if m != nil {
		return m.PortPropMotorSpeed
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetStbdPropMotorSpeed() []float32 {
	if m != nil {
		return m.StbdPropMotorSpeed
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetPropellerPitchPort() []float32 {
	if m != nil {
		return m.PropellerPitchPort
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetPropellerPitchStbd() []float32 {
	if m != nil {
		return m.PropellerPitchStbd
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetSog() []float32 {
	if m != nil {
		return m.Sog
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetWindDirectionRelative() []float32 {
	if m != nil {
		return m.WindDirectionRelative
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetWindSpeed() []float32 {
	if m != nil {
		return m.WindSpeed
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetBeaufortNumber() []uint32 {
	if m != nil {
		return m.BeaufortNumber
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetWaveDirection() []float32 {
	if m != nil {
		return m.WaveDirection
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetWaveLength() []float32 {
	if m != nil {
		return m.WaveLength
	}
	return nil
}

func (m *PowerTrainEstimateRequest) GetModelType() ModelTypeEnum {
	if m != nil {
		return m.ModelType
	}
	return ModelTypeEnum_UNKNOWN
}

type PowerTrackingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PowerTrackingRequest) Reset()         { *m = PowerTrackingRequest{} }
func (m *PowerTrackingRequest) String() string { return proto.CompactTextString(m) }
func (*PowerTrackingRequest) ProtoMessage()    {}
func (*PowerTrackingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{1}
}

func (m *PowerTrackingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowerTrackingRequest.Unmarshal(m, b)
}
func (m *PowerTrackingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowerTrackingRequest.Marshal(b, m, deterministic)
}
func (m *PowerTrackingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerTrackingRequest.Merge(m, src)
}
func (m *PowerTrackingRequest) XXX_Size() int {
	return xxx_messageInfo_PowerTrackingRequest.Size(m)
}
func (m *PowerTrackingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerTrackingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PowerTrackingRequest proto.InternalMessageInfo

type PowerEstimateResponse struct {
	UnixTime             []float64 `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	PowerEstimate        []float32 `protobuf:"fixed32,2,rep,packed,name=power_estimate,json=powerEstimate,proto3" json:"power_estimate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PowerEstimateResponse) Reset()         { *m = PowerEstimateResponse{} }
func (m *PowerEstimateResponse) String() string { return proto.CompactTextString(m) }
func (*PowerEstimateResponse) ProtoMessage()    {}
func (*PowerEstimateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{2}
}

func (m *PowerEstimateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowerEstimateResponse.Unmarshal(m, b)
}
func (m *PowerEstimateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowerEstimateResponse.Marshal(b, m, deterministic)
}
func (m *PowerEstimateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerEstimateResponse.Merge(m, src)
}
func (m *PowerEstimateResponse) XXX_Size() int {
	return xxx_messageInfo_PowerEstimateResponse.Size(m)
}
func (m *PowerEstimateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerEstimateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PowerEstimateResponse proto.InternalMessageInfo

func (m *PowerEstimateResponse) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *PowerEstimateResponse) GetPowerEstimate() []float32 {
	if m != nil {
		return m.PowerEstimate
	}
	return nil
}

type CostEstimateResponse struct {
	UnixTime             []float64 `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	PowerEstimate        []float32 `protobuf:"fixed32,2,rep,packed,name=power_estimate,json=powerEstimate,proto3" json:"power_estimate,omitempty"`
	CostEstimate         []float32 `protobuf:"fixed32,3,rep,packed,name=cost_estimate,json=costEstimate,proto3" json:"cost_estimate,omitempty"`
	TotalCost            float32   `protobuf:"fixed32,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CostEstimateResponse) Reset()         { *m = CostEstimateResponse{} }
func (m *CostEstimateResponse) String() string { return proto.CompactTextString(m) }
func (*CostEstimateResponse) ProtoMessage()    {}
func (*CostEstimateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{3}
}

func (m *CostEstimateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostEstimateResponse.Unmarshal(m, b)
}
func (m *CostEstimateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostEstimateResponse.Marshal(b, m, deterministic)
}
func (m *CostEstimateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostEstimateResponse.Merge(m, src)
}
func (m *CostEstimateResponse) XXX_Size() int {
	return xxx_messageInfo_CostEstimateResponse.Size(m)
}
func (m *CostEstimateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CostEstimateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CostEstimateResponse proto.InternalMessageInfo

func (m *CostEstimateResponse) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *CostEstimateResponse) GetPowerEstimate() []float32 {
	if m != nil {
		return m.PowerEstimate
	}
	return nil
}

func (m *CostEstimateResponse) GetCostEstimate() []float32 {
	if m != nil {
		return m.CostEstimate
	}
	return nil
}

func (m *CostEstimateResponse) GetTotalCost() float32 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

type PowerEvaluationResponse struct {
	UnixTime             []float64 `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	PowerEstimate        []float32 `protobuf:"fixed32,2,rep,packed,name=power_estimate,json=powerEstimate,proto3" json:"power_estimate,omitempty"`
	PowerActual          []float32 `protobuf:"fixed32,3,rep,packed,name=power_actual,json=powerActual,proto3" json:"power_actual,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PowerEvaluationResponse) Reset()         { *m = PowerEvaluationResponse{} }
func (m *PowerEvaluationResponse) String() string { return proto.CompactTextString(m) }
func (*PowerEvaluationResponse) ProtoMessage()    {}
func (*PowerEvaluationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{4}
}

func (m *PowerEvaluationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowerEvaluationResponse.Unmarshal(m, b)
}
func (m *PowerEvaluationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowerEvaluationResponse.Marshal(b, m, deterministic)
}
func (m *PowerEvaluationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerEvaluationResponse.Merge(m, src)
}
func (m *PowerEvaluationResponse) XXX_Size() int {
	return xxx_messageInfo_PowerEvaluationResponse.Size(m)
}
func (m *PowerEvaluationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerEvaluationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PowerEvaluationResponse proto.InternalMessageInfo

func (m *PowerEvaluationResponse) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *PowerEvaluationResponse) GetPowerEstimate() []float32 {
	if m != nil {
		return m.PowerEstimate
	}
	return nil
}

func (m *PowerEvaluationResponse) GetPowerActual() []float32 {
	if m != nil {
		return m.PowerActual
	}
	return nil
}

type PowerTrackingResponse struct {
	PowerActual          []float32 `protobuf:"fixed32,1,rep,packed,name=power_actual,json=powerActual,proto3" json:"power_actual,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PowerTrackingResponse) Reset()         { *m = PowerTrackingResponse{} }
func (m *PowerTrackingResponse) String() string { return proto.CompactTextString(m) }
func (*PowerTrackingResponse) ProtoMessage()    {}
func (*PowerTrackingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66ce88359bc96983, []int{5}
}

func (m *PowerTrackingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowerTrackingResponse.Unmarshal(m, b)
}
func (m *PowerTrackingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowerTrackingResponse.Marshal(b, m, deterministic)
}
func (m *PowerTrackingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerTrackingResponse.Merge(m, src)
}
func (m *PowerTrackingResponse) XXX_Size() int {
	return xxx_messageInfo_PowerTrackingResponse.Size(m)
}
func (m *PowerTrackingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerTrackingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PowerTrackingResponse proto.InternalMessageInfo

func (m *PowerTrackingResponse) GetPowerActual() []float32 {
	if m != nil {
		return m.PowerActual
	}
	return nil
}

func init() {
	proto.RegisterEnum("powerTrainService.v1.ModelTypeEnum", ModelTypeEnum_name, ModelTypeEnum_value)
	proto.RegisterType((*PowerTrainEstimateRequest)(nil), "powerTrainService.v1.PowerTrainEstimateRequest")
	proto.RegisterType((*PowerTrackingRequest)(nil), "powerTrainService.v1.PowerTrackingRequest")
	proto.RegisterType((*PowerEstimateResponse)(nil), "powerTrainService.v1.PowerEstimateResponse")
	proto.RegisterType((*CostEstimateResponse)(nil), "powerTrainService.v1.CostEstimateResponse")
	proto.RegisterType((*PowerEvaluationResponse)(nil), "powerTrainService.v1.PowerEvaluationResponse")
	proto.RegisterType((*PowerTrackingResponse)(nil), "powerTrainService.v1.PowerTrackingResponse")
}

func init() {
	proto.RegisterFile("powerTrainService/proto/v1/power_train_service_api_v1.proto", fileDescriptor_66ce88359bc96983)
}

var fileDescriptor_66ce88359bc96983 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5f, 0x6f, 0x12, 0x4b,
	0x14, 0xbf, 0x5b, 0x7a, 0xdb, 0x72, 0xf8, 0x73, 0xb9, 0x13, 0x6a, 0x57, 0x8d, 0x11, 0x69, 0x1a,
	0x49, 0x8d, 0x20, 0xd5, 0xf8, 0xa0, 0x4f, 0xb4, 0xf2, 0x60, 0xb4, 0x94, 0x50, 0x4c, 0x13, 0x7d,
	0x98, 0x0c, 0x70, 0xa4, 0x93, 0xee, 0xee, 0xac, 0xb3, 0xb3, 0x5b, 0x79, 0xf3, 0x93, 0x18, 0x3f,
	0x94, 0x1f, 0xc8, 0xcc, 0xec, 0x6e, 0xa1, 0x5d, 0x24, 0x69, 0xa2, 0x6f, 0xf0, 0xfb, 0x73, 0xce,
	0x9c, 0x39, 0x3f, 0x18, 0x78, 0xed, 0x8b, 0x4b, 0x94, 0x43, 0xc9, 0xb8, 0x77, 0x8a, 0x32, 0xe2,
	0x63, 0x6c, 0xf9, 0x52, 0x28, 0xd1, 0x8a, 0xda, 0x2d, 0x43, 0x51, 0xa5, 0x39, 0x1a, 0xc4, 0x24,
	0x65, 0x3e, 0xa7, 0x51, 0xbb, 0x69, 0x34, 0xa4, 0x9a, 0x31, 0x37, 0xa3, 0x76, 0xfd, 0xc7, 0x3a,
	0xdc, 0xed, 0x5f, 0x11, 0xdd, 0x40, 0x71, 0x97, 0x29, 0x1c, 0xe0, 0x97, 0x10, 0x03, 0x45, 0xee,
	0x43, 0x3e, 0xf4, 0xf8, 0x57, 0xaa, 0xb8, 0x8b, 0xb6, 0x55, 0xcb, 0x35, 0xac, 0xc1, 0x96, 0x06,
	0x86, 0xdc, 0x45, 0xd2, 0x86, 0x6d, 0x5f, 0x48, 0x45, 0x7d, 0x29, 0x7c, 0xea, 0x0a, 0x25, 0x24,
	0x0d, 0x7c, 0xc4, 0x89, 0xbd, 0x56, 0xcb, 0x35, 0xd6, 0x06, 0x44, 0x93, 0x7d, 0x29, 0xfc, 0x63,
	0x4d, 0x9d, 0x6a, 0x46, 0x5b, 0x02, 0x35, 0x9a, 0x64, 0x2d, 0xb9, 0xd8, 0xa2, 0xc9, 0x1b, 0x96,
	0x67, 0x50, 0xd5, 0x6a, 0x74, 0x1c, 0x94, 0xd4, 0xe7, 0x6a, 0x7c, 0x4e, 0x75, 0x61, 0x7b, 0x3d,
	0x69, 0x92, 0x72, 0x7d, 0x4d, 0xf5, 0x85, 0x54, 0xcb, 0x1c, 0xba, 0xae, 0xfd, 0xef, 0x32, 0xc7,
	0xa9, 0x1a, 0x4d, 0x48, 0x05, 0x72, 0x81, 0x98, 0xda, 0x1b, 0x46, 0xa0, 0x3f, 0x92, 0x97, 0xb0,
	0x73, 0xc9, 0xbd, 0x09, 0x9d, 0x70, 0x89, 0x63, 0xc5, 0x85, 0x47, 0x25, 0x3a, 0x4c, 0xf1, 0x08,
	0xed, 0x4d, 0xa3, 0xda, 0xd6, 0xf4, 0x9b, 0x94, 0x1d, 0x24, 0x24, 0x79, 0x00, 0x60, 0x7c, 0xf1,
	0x54, 0x5b, 0x46, 0x9a, 0xd7, 0x48, 0x3c, 0xcc, 0x63, 0xf8, 0x6f, 0x84, 0x2c, 0xfc, 0xac, 0xaf,
	0xcd, 0x0b, 0xdd, 0x11, 0x4a, 0x3b, 0x5f, 0xcb, 0x35, 0x4a, 0x83, 0x72, 0x0a, 0xf7, 0x0c, 0x4a,
	0xf6, 0xa0, 0x7c, 0xc9, 0x22, 0x9c, 0xf7, 0xb7, 0xc1, 0xd4, 0x2a, 0x69, 0xf4, 0xaa, 0x2d, 0x79,
	0x08, 0x05, 0x23, 0x73, 0xd0, 0x9b, 0xaa, 0x73, 0xbb, 0x60, 0x34, 0xa0, 0xa1, 0xf7, 0x06, 0x21,
	0x87, 0x00, 0xae, 0x98, 0xa0, 0x43, 0xd5, 0xcc, 0x47, 0xbb, 0x58, 0xb3, 0x1a, 0xe5, 0x83, 0xdd,
	0xe6, 0xb2, 0x24, 0x34, 0x8f, 0xb5, 0x6e, 0x38, 0xf3, 0xb1, 0xeb, 0x85, 0xee, 0x20, 0xef, 0xa6,
	0x5f, 0xeb, 0x77, 0xa0, 0x9a, 0x26, 0x64, 0x7c, 0xc1, 0xbd, 0x69, 0x12, 0x8e, 0xfa, 0x27, 0xd8,
	0x36, 0xf8, 0x3c, 0x34, 0x81, 0x2f, 0xbc, 0x00, 0x57, 0xa7, 0x66, 0x0f, 0xca, 0x71, 0x54, 0x31,
	0xb1, 0x25, 0x71, 0x29, 0xf9, 0x8b, 0xb5, 0xea, 0xdf, 0x2d, 0xa8, 0x1e, 0x89, 0x40, 0xfd, 0x8d,
	0xe2, 0x64, 0x17, 0x4a, 0x63, 0x11, 0xa8, 0xb9, 0x2a, 0x8e, 0x5f, 0x71, 0xbc, 0xd0, 0x50, 0xaf,
	0x52, 0x09, 0xc5, 0x1c, 0xaa, 0x51, 0x7b, 0xbd, 0x66, 0xe9, 0x55, 0x1a, 0x44, 0x9f, 0xab, 0xfe,
	0xcd, 0x82, 0x9d, 0x78, 0xfc, 0x88, 0x39, 0x21, 0x8b, 0x53, 0xf0, 0x07, 0xcf, 0xf8, 0x08, 0x8a,
	0xb1, 0x8c, 0x8d, 0x55, 0xc8, 0x9c, 0xe4, 0x88, 0x05, 0x83, 0x75, 0x0c, 0x54, 0x7f, 0x95, 0x2c,
	0x60, 0xbe, 0x98, 0xa4, 0xff, 0x4d, 0xaf, 0x95, 0xf1, 0xee, 0xbf, 0x80, 0xd2, 0xb5, 0x85, 0x93,
	0x02, 0x6c, 0x7e, 0xe8, 0xbd, 0xeb, 0x9d, 0x9c, 0xf5, 0x2a, 0xff, 0x90, 0x12, 0xe4, 0x4f, 0xfa,
	0xdd, 0xde, 0x59, 0x67, 0xd8, 0x1d, 0x54, 0x2c, 0xb2, 0x09, 0xb9, 0xb7, 0x47, 0xdd, 0xca, 0xda,
	0xc1, 0xcf, 0x1c, 0xfc, 0xdf, 0xbf, 0x19, 0x1e, 0xe2, 0x42, 0xe9, 0x5a, 0x10, 0x48, 0x6b, 0x79,
	0xc2, 0x7e, 0xfb, 0x3f, 0x73, 0xef, 0xc9, 0x0a, 0x43, 0x26, 0x01, 0x17, 0x50, 0x5c, 0x4c, 0xc6,
	0xed, 0xbb, 0xed, 0x2f, 0x37, 0x2c, 0x8d, 0xdb, 0x79, 0x32, 0x5b, 0x7a, 0xc7, 0x64, 0x7f, 0x75,
	0xb7, 0xc5, 0x5f, 0xc8, 0xca, 0xb1, 0x32, 0x4b, 0x9b, 0xa5, 0x79, 0x4a, 0x8e, 0x30, 0xcf, 0xd5,
	0xed, 0x27, 0x7c, 0xba, 0xea, 0x3e, 0x33, 0x79, 0x3d, 0xec, 0x7f, 0xec, 0x49, 0x11, 0x2a, 0xec,
	0x78, 0xcc, 0x99, 0x05, 0x3c, 0xe8, 0x4c, 0xa7, 0x12, 0xa7, 0x4c, 0x09, 0x39, 0x7f, 0x5f, 0xa6,
	0xe8, 0xa1, 0x64, 0x0a, 0x27, 0xad, 0x4c, 0xe9, 0xec, 0xb3, 0x34, 0xda, 0x30, 0xbe, 0xe7, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5a, 0x96, 0x4f, 0xb2, 0x06, 0x00, 0x00,
}
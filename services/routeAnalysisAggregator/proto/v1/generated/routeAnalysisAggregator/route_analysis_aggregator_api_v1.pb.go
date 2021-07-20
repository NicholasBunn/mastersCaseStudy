// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routeAnalysisAggregator/proto/v1/route_analysis_aggregator_api_v1.proto

package routeAnalysisAggregator

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

type HabitibilityRating int32

const (
	HabitibilityRating_UNKNOWN_RATING          HabitibilityRating = 0
	HabitibilityRating_NOT_UNCOMFORTABLE       HabitibilityRating = 1
	HabitibilityRating_SLIGHTLY_UNCOMFORTABLE  HabitibilityRating = 2
	HabitibilityRating_FAIRLY_UNCOMFORTABLE    HabitibilityRating = 3
	HabitibilityRating_UNCOMFORTABLE           HabitibilityRating = 4
	HabitibilityRating_VERY_UNCOMFORTABLE      HabitibilityRating = 5
	HabitibilityRating_EXTREMELY_UNCOMFORTABLE HabitibilityRating = 6
)

var HabitibilityRating_name = map[int32]string{
	0: "UNKNOWN_RATING",
	1: "NOT_UNCOMFORTABLE",
	2: "SLIGHTLY_UNCOMFORTABLE",
	3: "FAIRLY_UNCOMFORTABLE",
	4: "UNCOMFORTABLE",
	5: "VERY_UNCOMFORTABLE",
	6: "EXTREMELY_UNCOMFORTABLE",
}

var HabitibilityRating_value = map[string]int32{
	"UNKNOWN_RATING":          0,
	"NOT_UNCOMFORTABLE":       1,
	"SLIGHTLY_UNCOMFORTABLE":  2,
	"FAIRLY_UNCOMFORTABLE":    3,
	"UNCOMFORTABLE":           4,
	"VERY_UNCOMFORTABLE":      5,
	"EXTREMELY_UNCOMFORTABLE": 6,
}

func (x HabitibilityRating) String() string {
	return proto.EnumName(HabitibilityRating_name, int32(x))
}

func (HabitibilityRating) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f2785769c62a07ab, []int{0}
}

type AnalysisRequest struct {
	UnixTime             []float64 `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	Latitude             []float32 `protobuf:"fixed32,2,rep,packed,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            []float32 `protobuf:"fixed32,3,rep,packed,name=longitude,proto3" json:"longitude,omitempty"`
	Heading              []uint64  `protobuf:"varint,4,rep,packed,name=heading,proto3" json:"heading,omitempty"`
	PropPitch            []int64   `protobuf:"varint,5,rep,packed,name=prop_pitch,json=propPitch,proto3" json:"prop_pitch,omitempty"`
	MotorSpeed           []uint64  `protobuf:"varint,6,rep,packed,name=motor_speed,json=motorSpeed,proto3" json:"motor_speed,omitempty"`
	SOG                  []uint32  `protobuf:"varint,7,rep,packed,name=SOG,proto3" json:"SOG,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AnalysisRequest) Reset()         { *m = AnalysisRequest{} }
func (m *AnalysisRequest) String() string { return proto.CompactTextString(m) }
func (*AnalysisRequest) ProtoMessage()    {}
func (*AnalysisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2785769c62a07ab, []int{0}
}

func (m *AnalysisRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisRequest.Unmarshal(m, b)
}
func (m *AnalysisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisRequest.Marshal(b, m, deterministic)
}
func (m *AnalysisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisRequest.Merge(m, src)
}
func (m *AnalysisRequest) XXX_Size() int {
	return xxx_messageInfo_AnalysisRequest.Size(m)
}
func (m *AnalysisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisRequest proto.InternalMessageInfo

func (m *AnalysisRequest) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *AnalysisRequest) GetLatitude() []float32 {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *AnalysisRequest) GetLongitude() []float32 {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *AnalysisRequest) GetHeading() []uint64 {
	if m != nil {
		return m.Heading
	}
	return nil
}

func (m *AnalysisRequest) GetPropPitch() []int64 {
	if m != nil {
		return m.PropPitch
	}
	return nil
}

func (m *AnalysisRequest) GetMotorSpeed() []uint64 {
	if m != nil {
		return m.MotorSpeed
	}
	return nil
}

func (m *AnalysisRequest) GetSOG() []uint32 {
	if m != nil {
		return m.SOG
	}
	return nil
}

type AnalysisResponse struct {
	UnixTime             []float64          `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	AveragePower         float32            `protobuf:"fixed32,2,opt,name=average_power,json=averagePower,proto3" json:"average_power,omitempty"`
	TotalCost            float32            `protobuf:"fixed32,3,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	PowerUsage           []float32          `protobuf:"fixed32,4,rep,packed,name=power_usage,json=powerUsage,proto3" json:"power_usage,omitempty"`
	RmsVibrationX        []float32          `protobuf:"fixed32,5,rep,packed,name=rms_vibration_x,json=rmsVibrationX,proto3" json:"rms_vibration_x,omitempty"`
	RmsVibrationY        []float32          `protobuf:"fixed32,6,rep,packed,name=rms_vibration_y,json=rmsVibrationY,proto3" json:"rms_vibration_y,omitempty"`
	RmsVibrationZ        []float32          `protobuf:"fixed32,7,rep,packed,name=rms_vibration_z,json=rmsVibrationZ,proto3" json:"rms_vibration_z,omitempty"`
	ComfortLevel         HabitibilityRating `protobuf:"varint,8,opt,name=comfort_level,json=comfortLevel,proto3,enum=routeAnalysisAggregatorAPI.v1.HabitibilityRating" json:"comfort_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AnalysisResponse) Reset()         { *m = AnalysisResponse{} }
func (m *AnalysisResponse) String() string { return proto.CompactTextString(m) }
func (*AnalysisResponse) ProtoMessage()    {}
func (*AnalysisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2785769c62a07ab, []int{1}
}

func (m *AnalysisResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalysisResponse.Unmarshal(m, b)
}
func (m *AnalysisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalysisResponse.Marshal(b, m, deterministic)
}
func (m *AnalysisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalysisResponse.Merge(m, src)
}
func (m *AnalysisResponse) XXX_Size() int {
	return xxx_messageInfo_AnalysisResponse.Size(m)
}
func (m *AnalysisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalysisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnalysisResponse proto.InternalMessageInfo

func (m *AnalysisResponse) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *AnalysisResponse) GetAveragePower() float32 {
	if m != nil {
		return m.AveragePower
	}
	return 0
}

func (m *AnalysisResponse) GetTotalCost() float32 {
	if m != nil {
		return m.TotalCost
	}
	return 0
}

func (m *AnalysisResponse) GetPowerUsage() []float32 {
	if m != nil {
		return m.PowerUsage
	}
	return nil
}

func (m *AnalysisResponse) GetRmsVibrationX() []float32 {
	if m != nil {
		return m.RmsVibrationX
	}
	return nil
}

func (m *AnalysisResponse) GetRmsVibrationY() []float32 {
	if m != nil {
		return m.RmsVibrationY
	}
	return nil
}

func (m *AnalysisResponse) GetRmsVibrationZ() []float32 {
	if m != nil {
		return m.RmsVibrationZ
	}
	return nil
}

func (m *AnalysisResponse) GetComfortLevel() HabitibilityRating {
	if m != nil {
		return m.ComfortLevel
	}
	return HabitibilityRating_UNKNOWN_RATING
}

func init() {
	proto.RegisterEnum("routeAnalysisAggregatorAPI.v1.HabitibilityRating", HabitibilityRating_name, HabitibilityRating_value)
	proto.RegisterType((*AnalysisRequest)(nil), "routeAnalysisAggregatorAPI.v1.AnalysisRequest")
	proto.RegisterType((*AnalysisResponse)(nil), "routeAnalysisAggregatorAPI.v1.AnalysisResponse")
}

func init() {
	proto.RegisterFile("routeAnalysisAggregator/proto/v1/route_analysis_aggregator_api_v1.proto", fileDescriptor_f2785769c62a07ab)
}

var fileDescriptor_f2785769c62a07ab = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x4e, 0xdb, 0x3e,
	0x14, 0xc7, 0x7f, 0x49, 0xf8, 0xd7, 0xf3, 0xa3, 0x50, 0xac, 0x8d, 0x45, 0x30, 0xb4, 0x88, 0x49,
	0x53, 0xb4, 0x8b, 0x56, 0x65, 0x97, 0xbb, 0x0a, 0xa8, 0x94, 0x6a, 0x25, 0x45, 0x6e, 0x60, 0x80,
	0x26, 0x59, 0x6e, 0xeb, 0x05, 0x4b, 0x49, 0x9c, 0x39, 0x6e, 0x06, 0xbb, 0xdc, 0x7b, 0xed, 0x2d,
	0xf6, 0x08, 0x7b, 0x90, 0xc9, 0x86, 0xc2, 0xd6, 0xaa, 0xb0, 0xbb, 0xe4, 0x73, 0x3e, 0x3e, 0xf2,
	0xf7, 0xe8, 0x18, 0xda, 0x52, 0x8c, 0x15, 0x0b, 0x32, 0x9a, 0xdc, 0x14, 0xbc, 0x08, 0xe2, 0x58,
	0xb2, 0x98, 0x2a, 0x21, 0x1b, 0xb9, 0x14, 0x4a, 0x34, 0xca, 0x66, 0xc3, 0x08, 0x84, 0xde, 0x19,
	0x84, 0xde, 0x2b, 0x84, 0xe6, 0x9c, 0x94, 0xcd, 0xba, 0x31, 0xd1, 0xce, 0x9c, 0x46, 0xc1, 0x49,
	0xa7, 0x5e, 0x36, 0x77, 0x7f, 0x5a, 0xb0, 0x3e, 0x29, 0x62, 0xf6, 0x65, 0xcc, 0x0a, 0x85, 0xb6,
	0xa1, 0x32, 0xce, 0xf8, 0x35, 0x51, 0x3c, 0x65, 0xae, 0xe5, 0x39, 0xbe, 0x85, 0x57, 0x34, 0x88,
	0x78, 0xca, 0xd0, 0x16, 0xac, 0x24, 0x54, 0x71, 0x35, 0x1e, 0x31, 0xd7, 0xf6, 0x1c, 0xdf, 0xc6,
	0xf7, 0xff, 0xe8, 0x25, 0x54, 0x12, 0x91, 0xc5, 0xb7, 0x45, 0xc7, 0x14, 0x1f, 0x00, 0x72, 0x61,
	0xf9, 0x8a, 0xd1, 0x11, 0xcf, 0x62, 0x77, 0xc1, 0x73, 0xfc, 0x05, 0x3c, 0xf9, 0x45, 0x3b, 0x00,
	0xb9, 0x14, 0x39, 0xc9, 0xb9, 0x1a, 0x5e, 0xb9, 0x8b, 0x9e, 0xe3, 0x3b, 0xb8, 0xa2, 0xc9, 0x89,
	0x06, 0xe8, 0x15, 0xfc, 0x9f, 0x0a, 0x1d, 0xac, 0xc8, 0x19, 0x1b, 0xb9, 0x4b, 0xe6, 0x30, 0x18,
	0xd4, 0xd7, 0x04, 0xd5, 0xc0, 0xe9, 0xf7, 0xda, 0xee, 0xb2, 0xe7, 0xf8, 0x55, 0xac, 0x3f, 0x77,
	0x7f, 0xd9, 0x50, 0x7b, 0x88, 0x55, 0xe4, 0x22, 0x2b, 0xd8, 0xe3, 0xb9, 0x5e, 0x43, 0x95, 0x96,
	0x4c, 0xd2, 0x98, 0x91, 0x5c, 0x7c, 0x65, 0xd2, 0xb5, 0x3d, 0xcb, 0xb7, 0xf1, 0xea, 0x1d, 0x3c,
	0xd1, 0x4c, 0x5f, 0x54, 0x09, 0x45, 0x13, 0x32, 0x14, 0x85, 0x72, 0x1d, 0x63, 0x54, 0x0c, 0x39,
	0x10, 0x85, 0xd2, 0x17, 0x35, 0x67, 0xc9, 0xb8, 0xa0, 0x31, 0x33, 0x29, 0x6d, 0x0c, 0x06, 0x9d,
	0x6a, 0x82, 0xde, 0xc0, 0xba, 0x4c, 0x0b, 0x52, 0xf2, 0x81, 0xa4, 0x8a, 0x8b, 0x8c, 0x5c, 0x9b,
	0xb4, 0x36, 0xae, 0xca, 0xb4, 0x38, 0x9b, 0xd0, 0xf3, 0x59, 0xef, 0xc6, 0xa4, 0x9e, 0xf2, 0x2e,
	0x66, 0xbd, 0x6f, 0x66, 0x08, 0x53, 0xde, 0x25, 0x3a, 0x83, 0xea, 0x50, 0xa4, 0x9f, 0x85, 0x54,
	0x24, 0x61, 0x25, 0x4b, 0xdc, 0x15, 0xcf, 0xf2, 0xd7, 0xf6, 0x9a, 0xf5, 0x47, 0x97, 0xa3, 0x7e,
	0x44, 0x07, 0x5c, 0xf1, 0x01, 0x4f, 0xb8, 0xba, 0xc1, 0x54, 0xf1, 0x2c, 0xc6, 0xab, 0x77, 0x7d,
	0xba, 0xba, 0xcd, 0xdb, 0x1f, 0x16, 0xa0, 0x59, 0x09, 0x21, 0x58, 0x3b, 0x0d, 0x3f, 0x84, 0xbd,
	0x8f, 0x21, 0xc1, 0x41, 0xd4, 0x09, 0xdb, 0xb5, 0xff, 0xd0, 0x73, 0xd8, 0x08, 0x7b, 0x11, 0x39,
	0x0d, 0x0f, 0x7a, 0xc7, 0x87, 0x3d, 0x1c, 0x05, 0xfb, 0xdd, 0x56, 0xcd, 0x42, 0x5b, 0xb0, 0xd9,
	0xef, 0x76, 0xda, 0x47, 0x51, 0xf7, 0x62, 0xaa, 0x66, 0x23, 0x17, 0x9e, 0x1d, 0x06, 0x1d, 0x3c,
	0x53, 0x71, 0xd0, 0x06, 0x54, 0xff, 0x46, 0x0b, 0x68, 0x13, 0xd0, 0x59, 0x0b, 0x4f, 0xab, 0x8b,
	0x68, 0x1b, 0x5e, 0xb4, 0xce, 0x23, 0xdc, 0x3a, 0x6e, 0xcd, 0xf4, 0x59, 0xda, 0xfb, 0xfe, 0xc7,
	0xf6, 0xf7, 0x99, 0x2c, 0xf9, 0x90, 0x21, 0x01, 0xab, 0xb7, 0x88, 0x61, 0x3d, 0x1c, 0x54, 0x7f,
	0x62, 0x48, 0x53, 0xaf, 0x67, 0xab, 0xf1, 0xcf, 0xfe, 0xed, 0x5a, 0xee, 0x7f, 0xba, 0xbc, 0x7c,
	0xf2, 0xb1, 0xc7, 0x2c, 0x63, 0x92, 0x2a, 0x36, 0x6a, 0xcc, 0x51, 0xdf, 0xcf, 0xe1, 0x83, 0x25,
	0xd3, 0xe3, 0xdd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x7c, 0xaa, 0x06, 0x51, 0x04, 0x00,
	0x00,
}

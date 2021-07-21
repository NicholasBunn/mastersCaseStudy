// Code generated by protoc-gen-go. DO NOT EDIT.
// source: oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto

package oceanWeatherService

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

type ArchiveService int32

const (
	ArchiveService_UNKNOWN      ArchiveService = 0
	ArchiveService_STORMGLASS   ArchiveService = 1
	ArchiveService_ERA5         ArchiveService = 2
	ArchiveService_OBSERVATIONS ArchiveService = 3
)

var ArchiveService_name = map[int32]string{
	0: "UNKNOWN",
	1: "STORMGLASS",
	2: "ERA5",
	3: "OBSERVATIONS",
}

var ArchiveService_value = map[string]int32{
	"UNKNOWN":      0,
	"STORMGLASS":   1,
	"ERA5":         2,
	"OBSERVATIONS": 3,
}

func (x ArchiveService) String() string {
	return proto.EnumName(ArchiveService_name, int32(x))
}

func (ArchiveService) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c5e6f7de30ab5156, []int{0}
}

type OceanWeatherPredictionRequest struct {
	Latitude             []float32 `protobuf:"fixed32,1,rep,packed,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            []float32 `protobuf:"fixed32,2,rep,packed,name=longitude,proto3" json:"longitude,omitempty"`
	UnixTime             []float64 `protobuf:"fixed64,3,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OceanWeatherPredictionRequest) Reset()         { *m = OceanWeatherPredictionRequest{} }
func (m *OceanWeatherPredictionRequest) String() string { return proto.CompactTextString(m) }
func (*OceanWeatherPredictionRequest) ProtoMessage()    {}
func (*OceanWeatherPredictionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e6f7de30ab5156, []int{0}
}

func (m *OceanWeatherPredictionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OceanWeatherPredictionRequest.Unmarshal(m, b)
}
func (m *OceanWeatherPredictionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OceanWeatherPredictionRequest.Marshal(b, m, deterministic)
}
func (m *OceanWeatherPredictionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OceanWeatherPredictionRequest.Merge(m, src)
}
func (m *OceanWeatherPredictionRequest) XXX_Size() int {
	return xxx_messageInfo_OceanWeatherPredictionRequest.Size(m)
}
func (m *OceanWeatherPredictionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OceanWeatherPredictionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OceanWeatherPredictionRequest proto.InternalMessageInfo

func (m *OceanWeatherPredictionRequest) GetLatitude() []float32 {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *OceanWeatherPredictionRequest) GetLongitude() []float32 {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *OceanWeatherPredictionRequest) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

type OceanWeatherHistoryRequest struct {
	Latitude             []float32      `protobuf:"fixed32,1,rep,packed,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            []float32      `protobuf:"fixed32,2,rep,packed,name=longitude,proto3" json:"longitude,omitempty"`
	UnixTime             []float64      `protobuf:"fixed64,3,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	ArchiveService       ArchiveService `protobuf:"varint,4,opt,name=archive_service,json=archiveService,proto3,enum=oceanWeatherServiceAPI.v1.ArchiveService" json:"archive_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OceanWeatherHistoryRequest) Reset()         { *m = OceanWeatherHistoryRequest{} }
func (m *OceanWeatherHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*OceanWeatherHistoryRequest) ProtoMessage()    {}
func (*OceanWeatherHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e6f7de30ab5156, []int{1}
}

func (m *OceanWeatherHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OceanWeatherHistoryRequest.Unmarshal(m, b)
}
func (m *OceanWeatherHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OceanWeatherHistoryRequest.Marshal(b, m, deterministic)
}
func (m *OceanWeatherHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OceanWeatherHistoryRequest.Merge(m, src)
}
func (m *OceanWeatherHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_OceanWeatherHistoryRequest.Size(m)
}
func (m *OceanWeatherHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OceanWeatherHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OceanWeatherHistoryRequest proto.InternalMessageInfo

func (m *OceanWeatherHistoryRequest) GetLatitude() []float32 {
	if m != nil {
		return m.Latitude
	}
	return nil
}

func (m *OceanWeatherHistoryRequest) GetLongitude() []float32 {
	if m != nil {
		return m.Longitude
	}
	return nil
}

func (m *OceanWeatherHistoryRequest) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *OceanWeatherHistoryRequest) GetArchiveService() ArchiveService {
	if m != nil {
		return m.ArchiveService
	}
	return ArchiveService_UNKNOWN
}

type OceanWeatherInformationResponse struct {
	UnixTime             []float64 `protobuf:"fixed64,1,rep,packed,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	WindDirection        []float32 `protobuf:"fixed32,2,rep,packed,name=wind_direction,json=windDirection,proto3" json:"wind_direction,omitempty"`
	WindSpeed            []float32 `protobuf:"fixed32,3,rep,packed,name=wind_speed,json=windSpeed,proto3" json:"wind_speed,omitempty"`
	BeaufortNumber       []uint32  `protobuf:"varint,4,rep,packed,name=beaufort_number,json=beaufortNumber,proto3" json:"beaufort_number,omitempty"`
	SwellDirection       []float32 `protobuf:"fixed32,5,rep,packed,name=swell_direction,json=swellDirection,proto3" json:"swell_direction,omitempty"`
	WaveLength           []float32 `protobuf:"fixed32,6,rep,packed,name=wave_length,json=waveLength,proto3" json:"wave_length,omitempty"`
	SwellHeight          []float32 `protobuf:"fixed32,7,rep,packed,name=swell_height,json=swellHeight,proto3" json:"swell_height,omitempty"`
	SwellFrequency       []float32 `protobuf:"fixed32,8,rep,packed,name=swell_frequency,json=swellFrequency,proto3" json:"swell_frequency,omitempty"`
	SwellPeriod          []float32 `protobuf:"fixed32,9,rep,packed,name=swell_period,json=swellPeriod,proto3" json:"swell_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OceanWeatherInformationResponse) Reset()         { *m = OceanWeatherInformationResponse{} }
func (m *OceanWeatherInformationResponse) String() string { return proto.CompactTextString(m) }
func (*OceanWeatherInformationResponse) ProtoMessage()    {}
func (*OceanWeatherInformationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e6f7de30ab5156, []int{2}
}

func (m *OceanWeatherInformationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OceanWeatherInformationResponse.Unmarshal(m, b)
}
func (m *OceanWeatherInformationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OceanWeatherInformationResponse.Marshal(b, m, deterministic)
}
func (m *OceanWeatherInformationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OceanWeatherInformationResponse.Merge(m, src)
}
func (m *OceanWeatherInformationResponse) XXX_Size() int {
	return xxx_messageInfo_OceanWeatherInformationResponse.Size(m)
}
func (m *OceanWeatherInformationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OceanWeatherInformationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OceanWeatherInformationResponse proto.InternalMessageInfo

func (m *OceanWeatherInformationResponse) GetUnixTime() []float64 {
	if m != nil {
		return m.UnixTime
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetWindDirection() []float32 {
	if m != nil {
		return m.WindDirection
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetWindSpeed() []float32 {
	if m != nil {
		return m.WindSpeed
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetBeaufortNumber() []uint32 {
	if m != nil {
		return m.BeaufortNumber
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetSwellDirection() []float32 {
	if m != nil {
		return m.SwellDirection
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetWaveLength() []float32 {
	if m != nil {
		return m.WaveLength
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetSwellHeight() []float32 {
	if m != nil {
		return m.SwellHeight
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetSwellFrequency() []float32 {
	if m != nil {
		return m.SwellFrequency
	}
	return nil
}

func (m *OceanWeatherInformationResponse) GetSwellPeriod() []float32 {
	if m != nil {
		return m.SwellPeriod
	}
	return nil
}

func init() {
	proto.RegisterEnum("oceanWeatherServiceAPI.v1.ArchiveService", ArchiveService_name, ArchiveService_value)
	proto.RegisterType((*OceanWeatherPredictionRequest)(nil), "oceanWeatherServiceAPI.v1.OceanWeatherPredictionRequest")
	proto.RegisterType((*OceanWeatherHistoryRequest)(nil), "oceanWeatherServiceAPI.v1.OceanWeatherHistoryRequest")
	proto.RegisterType((*OceanWeatherInformationResponse)(nil), "oceanWeatherServiceAPI.v1.OceanWeatherInformationResponse")
}

func init() {
	proto.RegisterFile("oceanWeatherService/proto/v1/ocean_weather_service_api_v1.proto", fileDescriptor_c5e6f7de30ab5156)
}

var fileDescriptor_c5e6f7de30ab5156 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0x6d, 0xc1, 0x5d, 0x18, 0x76, 0xbb, 0x64, 0x4c, 0x4c, 0x45, 0x37, 0x8b, 0x24, 0x46,
	0xf4, 0x02, 0xc2, 0x9a, 0x4d, 0x8c, 0x5e, 0x98, 0x6e, 0x5c, 0x5d, 0xe2, 0x4a, 0x49, 0x8b, 0x6e,
	0xe2, 0x4d, 0x33, 0xd0, 0x43, 0x99, 0xa4, 0xcc, 0xe0, 0x74, 0x5a, 0xe4, 0x0d, 0xbc, 0xf2, 0x5d,
	0x7c, 0x0a, 0xdf, 0xc4, 0xe7, 0x30, 0x9d, 0x82, 0x94, 0x15, 0x37, 0x9b, 0x18, 0x2f, 0xfb, 0x9d,
	0x33, 0xff, 0x7f, 0xe6, 0x9c, 0x39, 0x45, 0xaf, 0xf8, 0x08, 0x08, 0xbb, 0x04, 0x22, 0x27, 0x20,
	0x5c, 0x10, 0x09, 0x1d, 0x41, 0x7b, 0x26, 0xb8, 0xe4, 0xed, 0xa4, 0xd3, 0x56, 0x41, 0x6f, 0x9e,
	0x45, 0xbd, 0x28, 0x0b, 0x7b, 0x64, 0x46, 0xbd, 0xa4, 0xd3, 0x52, 0x59, 0xf8, 0xde, 0x16, 0x01,
	0xab, 0xdf, 0x6d, 0x25, 0x9d, 0x46, 0x82, 0x0e, 0xed, 0x5c, 0xb0, 0x2f, 0xc0, 0xa7, 0x23, 0x49,
	0x39, 0x73, 0xe0, 0x73, 0x0c, 0x91, 0xc4, 0x35, 0x54, 0x0a, 0x89, 0xa4, 0x32, 0xf6, 0xc1, 0xd4,
	0xea, 0x85, 0xa6, 0xee, 0xfc, 0xfe, 0xc6, 0x0f, 0x50, 0x39, 0xe4, 0x2c, 0xc8, 0x82, 0xba, 0x0a,
	0xae, 0x01, 0xbe, 0x8f, 0xca, 0x31, 0xa3, 0x5f, 0x3c, 0x49, 0xa7, 0x60, 0x16, 0xea, 0x85, 0xa6,
	0xe6, 0x94, 0x52, 0x30, 0xa0, 0x53, 0x68, 0xfc, 0xd0, 0x50, 0x2d, 0x6f, 0x7c, 0x4e, 0x23, 0xc9,
	0xc5, 0xe2, 0xff, 0xba, 0x62, 0x07, 0x1d, 0x10, 0x31, 0x9a, 0xd0, 0x04, 0x56, 0x8d, 0x32, 0x8b,
	0x75, 0xad, 0x69, 0x1c, 0x3f, 0x69, 0xfd, 0xb5, 0x45, 0x2d, 0x2b, 0x3b, 0xb1, 0x84, 0x8e, 0x41,
	0x36, 0xbe, 0x1b, 0x3f, 0x75, 0x74, 0x94, 0xbf, 0x49, 0x97, 0x8d, 0xb9, 0x98, 0x92, 0xac, 0x87,
	0xd1, 0x8c, 0xb3, 0xe8, 0x4a, 0x51, 0xda, 0x95, 0xa2, 0x1e, 0x21, 0x63, 0x4e, 0x99, 0xef, 0xf9,
	0x54, 0x80, 0x6a, 0xfd, 0xf2, 0x52, 0xfb, 0x29, 0x7d, 0xbd, 0x82, 0xf8, 0x10, 0x21, 0x95, 0x16,
	0xcd, 0x00, 0x7c, 0x75, 0x33, 0xdd, 0x29, 0xa7, 0xc4, 0x4d, 0x01, 0x7e, 0x8c, 0x0e, 0x86, 0x40,
	0xe2, 0x31, 0x17, 0xd2, 0x63, 0xf1, 0x74, 0x08, 0xc2, 0x2c, 0xd6, 0x0b, 0xcd, 0x7d, 0xc7, 0x58,
	0xe1, 0x9e, 0xa2, 0x69, 0x62, 0x34, 0x87, 0x30, 0xcc, 0xf9, 0xdd, 0x56, 0x62, 0x86, 0xc2, 0x6b,
	0xc3, 0x23, 0x54, 0x99, 0x93, 0x04, 0xbc, 0x10, 0x58, 0x20, 0x27, 0xe6, 0x8e, 0x4a, 0x42, 0x29,
	0xba, 0x50, 0x04, 0x3f, 0x44, 0x7b, 0x99, 0xd2, 0x04, 0x68, 0x30, 0x91, 0xe6, 0xae, 0xca, 0xa8,
	0x28, 0x76, 0xae, 0xd0, 0xda, 0x6c, 0x2c, 0xd2, 0xc9, 0xb2, 0xd1, 0xc2, 0x2c, 0xe5, 0xcc, 0xde,
	0xac, 0xe8, 0x5a, 0x6b, 0x06, 0x82, 0x72, 0xdf, 0x2c, 0xe7, 0xb4, 0xfa, 0x0a, 0x3d, 0xed, 0x22,
	0x63, 0x73, 0x14, 0xb8, 0x82, 0x76, 0x3f, 0xf4, 0xde, 0xf5, 0xec, 0xcb, 0x5e, 0xf5, 0x16, 0x36,
	0x10, 0x72, 0x07, 0xb6, 0xf3, 0xfe, 0xed, 0x85, 0xe5, 0xba, 0x55, 0x0d, 0x97, 0x50, 0xf1, 0xcc,
	0xb1, 0x4e, 0xaa, 0x3a, 0xae, 0xa2, 0x3d, 0xfb, 0xd4, 0x3d, 0x73, 0x3e, 0x5a, 0x83, 0xae, 0xdd,
	0x73, 0xab, 0x85, 0xe3, 0xef, 0x3a, 0xba, 0x63, 0xff, 0x39, 0x70, 0xfc, 0x4d, 0x43, 0x77, 0xb7,
	0xaf, 0x03, 0x7e, 0x7e, 0xcd, 0x0b, 0xb9, 0x76, 0x83, 0x6a, 0x2f, 0x6e, 0x78, 0x72, 0xdb, 0xc3,
	0xf9, 0xaa, 0x6d, 0x16, 0xba, 0x5c, 0x13, 0x7c, 0x72, 0x43, 0xcd, 0xcd, 0xb5, 0xfa, 0x97, 0x52,
	0x4e, 0x07, 0x9f, 0x1c, 0xc1, 0x63, 0x09, 0x16, 0x23, 0xe1, 0x22, 0xa2, 0x91, 0x15, 0x04, 0x02,
	0x02, 0x22, 0xb9, 0x58, 0xff, 0x8b, 0x02, 0x60, 0x20, 0x88, 0x04, 0xbf, 0xbd, 0xc5, 0xe7, 0xe5,
	0x16, 0x36, 0xdc, 0x51, 0x67, 0x9f, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x58, 0x2b, 0x47, 0x93,
	0xe4, 0x04, 0x00, 0x00,
}
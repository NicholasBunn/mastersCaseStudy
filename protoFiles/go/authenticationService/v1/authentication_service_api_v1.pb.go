// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authenticationService/proto/v1/authentication_service_api_v1.proto

package authenticationService

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

type NewUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Permission           string   `protobuf:"bytes,3,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewUserRequest) Reset()         { *m = NewUserRequest{} }
func (m *NewUserRequest) String() string { return proto.CompactTextString(m) }
func (*NewUserRequest) ProtoMessage()    {}
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_509a7ba7a2239f73, []int{0}
}

func (m *NewUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewUserRequest.Unmarshal(m, b)
}
func (m *NewUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewUserRequest.Marshal(b, m, deterministic)
}
func (m *NewUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewUserRequest.Merge(m, src)
}
func (m *NewUserRequest) XXX_Size() int {
	return xxx_messageInfo_NewUserRequest.Size(m)
}
func (m *NewUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewUserRequest proto.InternalMessageInfo

func (m *NewUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *NewUserRequest) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

type LoginAuthRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginAuthRequest) Reset()         { *m = LoginAuthRequest{} }
func (m *LoginAuthRequest) String() string { return proto.CompactTextString(m) }
func (*LoginAuthRequest) ProtoMessage()    {}
func (*LoginAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_509a7ba7a2239f73, []int{1}
}

func (m *LoginAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAuthRequest.Unmarshal(m, b)
}
func (m *LoginAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAuthRequest.Marshal(b, m, deterministic)
}
func (m *LoginAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAuthRequest.Merge(m, src)
}
func (m *LoginAuthRequest) XXX_Size() int {
	return xxx_messageInfo_LoginAuthRequest.Size(m)
}
func (m *LoginAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAuthRequest proto.InternalMessageInfo

func (m *LoginAuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginAuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginAuthResponse struct {
	Permissions          string   `protobuf:"bytes,1,opt,name=permissions,proto3" json:"permissions,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginAuthResponse) Reset()         { *m = LoginAuthResponse{} }
func (m *LoginAuthResponse) String() string { return proto.CompactTextString(m) }
func (*LoginAuthResponse) ProtoMessage()    {}
func (*LoginAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_509a7ba7a2239f73, []int{2}
}

func (m *LoginAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAuthResponse.Unmarshal(m, b)
}
func (m *LoginAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAuthResponse.Marshal(b, m, deterministic)
}
func (m *LoginAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAuthResponse.Merge(m, src)
}
func (m *LoginAuthResponse) XXX_Size() int {
	return xxx_messageInfo_LoginAuthResponse.Size(m)
}
func (m *LoginAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAuthResponse proto.InternalMessageInfo

func (m *LoginAuthResponse) GetPermissions() string {
	if m != nil {
		return m.Permissions
	}
	return ""
}

func (m *LoginAuthResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*NewUserRequest)(nil), "authenticationService.v1.NewUserRequest")
	proto.RegisterType((*LoginAuthRequest)(nil), "authenticationService.v1.LoginAuthRequest")
	proto.RegisterType((*LoginAuthResponse)(nil), "authenticationService.v1.LoginAuthResponse")
}

func init() {
	proto.RegisterFile("authenticationService/proto/v1/authentication_service_api_v1.proto", fileDescriptor_509a7ba7a2239f73)
}

var fileDescriptor_509a7ba7a2239f73 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xff, 0xf9, 0x0b, 0x62, 0xa7, 0x2a, 0xba, 0x20, 0x84, 0x1e, 0xa4, 0xe6, 0x54, 0x14,
	0xb2, 0xac, 0x1e, 0x3d, 0xb5, 0xde, 0x44, 0x3c, 0x54, 0x05, 0xf1, 0x12, 0xd6, 0x38, 0x9a, 0x45,
	0xba, 0xbb, 0xee, 0x6c, 0xb6, 0x9f, 0xd9, 0x6f, 0x21, 0x4d, 0x62, 0x4d, 0x31, 0x05, 0xd1, 0xe3,
	0xbc, 0xdf, 0xec, 0x7b, 0xec, 0x63, 0x60, 0x22, 0x4b, 0x5f, 0xa0, 0xf6, 0x2a, 0x97, 0x5e, 0x19,
	0x7d, 0x83, 0x2e, 0xa8, 0x1c, 0xb9, 0x75, 0xc6, 0x1b, 0x1e, 0x04, 0x5f, 0xc5, 0x19, 0xd5, 0x3c,
	0x93, 0x56, 0x65, 0x41, 0xa4, 0xd5, 0x1a, 0x8b, 0x3b, 0x3d, 0xd2, 0x20, 0x92, 0x02, 0x76, 0xaf,
	0x71, 0x7e, 0x47, 0xe8, 0xa6, 0xf8, 0x56, 0x22, 0x79, 0x36, 0x80, 0xad, 0x92, 0xd0, 0x69, 0x39,
	0xc3, 0x38, 0x1a, 0x46, 0xa3, 0xde, 0x74, 0x39, 0x2f, 0x98, 0x95, 0x44, 0x73, 0xe3, 0x9e, 0xe2,
	0xff, 0x35, 0xfb, 0x9c, 0xd9, 0x21, 0x80, 0x45, 0x37, 0x53, 0x44, 0xca, 0xe8, 0x78, 0xa3, 0xa2,
	0x2d, 0x25, 0xb9, 0x84, 0xbd, 0x2b, 0xf3, 0xa2, 0xf4, 0xb8, 0xf4, 0xc5, 0x1f, 0xb3, 0x92, 0x7b,
	0xd8, 0x6f, 0x79, 0x91, 0x35, 0x9a, 0x90, 0x0d, 0xa1, 0xff, 0x15, 0x47, 0x8d, 0x5f, 0x5b, 0x62,
	0x47, 0xb0, 0x2d, 0xf3, 0x1c, 0x89, 0x32, 0x6f, 0x5e, 0x51, 0x37, 0xb6, 0xfd, 0x5a, 0xbb, 0x5d,
	0x48, 0xa7, 0xef, 0x11, 0x1c, 0x8c, 0xbb, 0xca, 0x62, 0x05, 0xec, 0x5c, 0x38, 0x94, 0x1e, 0x9b,
	0xbe, 0xd8, 0x28, 0x5d, 0xd7, 0x6a, 0xba, 0x5a, 0xe9, 0xe0, 0x64, 0xfd, 0xe6, 0xb7, 0x6f, 0x24,
	0xff, 0xd8, 0x33, 0xf4, 0x96, 0x32, 0x3b, 0xfe, 0xd1, 0xdb, 0xdf, 0xe4, 0x4c, 0xc4, 0x03, 0xef,
	0xbe, 0xad, 0x20, 0xf8, 0x79, 0x27, 0x79, 0xdc, 0xac, 0xee, 0xe9, 0xec, 0x23, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0x90, 0x45, 0xb9, 0x95, 0x02, 0x00, 0x00,
}
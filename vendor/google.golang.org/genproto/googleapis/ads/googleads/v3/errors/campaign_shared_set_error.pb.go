// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/campaign_shared_set_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum_CampaignSharedSetError int32

const (
	// Enum unspecified.
	CampaignSharedSetErrorEnum_UNSPECIFIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 0
	// The received error code is not known in this version.
	CampaignSharedSetErrorEnum_UNKNOWN CampaignSharedSetErrorEnum_CampaignSharedSetError = 1
	// The shared set belongs to another customer and permission isn't granted.
	CampaignSharedSetErrorEnum_SHARED_SET_ACCESS_DENIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 2
)

var CampaignSharedSetErrorEnum_CampaignSharedSetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SHARED_SET_ACCESS_DENIED",
}

var CampaignSharedSetErrorEnum_CampaignSharedSetError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"SHARED_SET_ACCESS_DENIED": 2,
}

func (x CampaignSharedSetErrorEnum_CampaignSharedSetError) String() string {
	return proto.EnumName(CampaignSharedSetErrorEnum_CampaignSharedSetError_name, int32(x))
}

func (CampaignSharedSetErrorEnum_CampaignSharedSetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6b80018c978188cf, []int{0, 0}
}

// Container for enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignSharedSetErrorEnum) Reset()         { *m = CampaignSharedSetErrorEnum{} }
func (m *CampaignSharedSetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetErrorEnum) ProtoMessage()    {}
func (*CampaignSharedSetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b80018c978188cf, []int{0}
}

func (m *CampaignSharedSetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Unmarshal(m, b)
}
func (m *CampaignSharedSetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetErrorEnum.Merge(m, src)
}
func (m *CampaignSharedSetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Size(m)
}
func (m *CampaignSharedSetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.CampaignSharedSetErrorEnum_CampaignSharedSetError", CampaignSharedSetErrorEnum_CampaignSharedSetError_name, CampaignSharedSetErrorEnum_CampaignSharedSetError_value)
	proto.RegisterType((*CampaignSharedSetErrorEnum)(nil), "google.ads.googleads.v3.errors.CampaignSharedSetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/campaign_shared_set_error.proto", fileDescriptor_6b80018c978188cf)
}

var fileDescriptor_6b80018c978188cf = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x5d, 0x05, 0x85, 0xec, 0xe0, 0xe8, 0x41, 0x64, 0x8e, 0x1d, 0xfa, 0x00, 0xe9, 0xa1,
	0xb7, 0x08, 0x42, 0xd6, 0xc6, 0x39, 0x84, 0x3a, 0xcc, 0x36, 0x41, 0x0a, 0x21, 0x2e, 0x21, 0x16,
	0xb6, 0xa4, 0x24, 0x75, 0x0f, 0xe4, 0xd1, 0x47, 0xf1, 0x51, 0x7c, 0x00, 0xcf, 0xd2, 0x66, 0xeb,
	0x69, 0x7a, 0xea, 0x9f, 0x7e, 0xbf, 0xdf, 0xf7, 0x7d, 0xf9, 0xc0, 0xad, 0x32, 0x46, 0x6d, 0x64,
	0xcc, 0x85, 0x8b, 0x7d, 0x6c, 0xd2, 0x2e, 0x89, 0xa5, 0xb5, 0xc6, 0xba, 0x78, 0xcd, 0xb7, 0x15,
	0x2f, 0x95, 0x66, 0xee, 0x8d, 0x5b, 0x29, 0x98, 0x93, 0x35, 0x6b, 0x4b, 0xb0, 0xb2, 0xa6, 0x36,
	0xe1, 0xd8, 0x4b, 0x90, 0x0b, 0x07, 0x3b, 0x1f, 0xee, 0x12, 0xe8, 0xfd, 0xe1, 0xe8, 0xd0, 0xbf,
	0x2a, 0x63, 0xae, 0xb5, 0xa9, 0x79, 0x5d, 0x1a, 0xed, 0xbc, 0x1d, 0x59, 0x30, 0x4c, 0xf7, 0x03,
	0x68, 0xdb, 0x9f, 0xca, 0x9a, 0x34, 0x22, 0xd1, 0xef, 0xdb, 0x68, 0x01, 0x2e, 0x8f, 0x57, 0xc3,
	0x0b, 0xd0, 0x5f, 0xe6, 0x74, 0x4e, 0xd2, 0xd9, 0xdd, 0x8c, 0x64, 0x83, 0x93, 0xb0, 0x0f, 0xce,
	0x97, 0xf9, 0x43, 0xfe, 0xf8, 0x9c, 0x0f, 0x7a, 0xe1, 0x08, 0x5c, 0xd1, 0x7b, 0xfc, 0x44, 0x32,
	0x46, 0xc9, 0x82, 0xe1, 0x34, 0x25, 0x94, 0xb2, 0x8c, 0xe4, 0x0d, 0x1a, 0x4c, 0x7e, 0x7a, 0x20,
	0x5a, 0x9b, 0x2d, 0xfc, 0x7f, 0xf1, 0xc9, 0xf5, 0xf1, 0xd1, 0xf3, 0x66, 0xef, 0x79, 0xef, 0x25,
	0xdb, 0xeb, 0xca, 0x6c, 0xb8, 0x56, 0xd0, 0x58, 0x15, 0x2b, 0xa9, 0xdb, 0x57, 0x1d, 0xee, 0x58,
	0x95, 0xee, 0xaf, 0xb3, 0xde, 0xf8, 0xcf, 0x47, 0x70, 0x3a, 0xc5, 0xf8, 0x33, 0x18, 0x4f, 0x7d,
	0x33, 0x2c, 0x1c, 0xf4, 0xb1, 0x49, 0xab, 0x04, 0xb6, 0x23, 0xdd, 0xd7, 0x01, 0x28, 0xb0, 0x70,
	0x45, 0x07, 0x14, 0xab, 0xa4, 0xf0, 0xc0, 0x77, 0x10, 0xf9, 0xbf, 0x08, 0x61, 0xe1, 0x10, 0xea,
	0x10, 0x84, 0x56, 0x09, 0x42, 0x1e, 0x7a, 0x3d, 0x6b, 0xb7, 0x4b, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xfd, 0x86, 0x7d, 0x8a, 0xf3, 0x01, 0x00, 0x00,
}

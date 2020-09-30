// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/bidding_strategy_status.proto

package enums

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

// The possible statuses of a BiddingStrategy.
type BiddingStrategyStatusEnum_BiddingStrategyStatus int32

const (
	// No value has been specified.
	BiddingStrategyStatusEnum_UNSPECIFIED BiddingStrategyStatusEnum_BiddingStrategyStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	BiddingStrategyStatusEnum_UNKNOWN BiddingStrategyStatusEnum_BiddingStrategyStatus = 1
	// The bidding strategy is enabled.
	BiddingStrategyStatusEnum_ENABLED BiddingStrategyStatusEnum_BiddingStrategyStatus = 2
	// The bidding strategy is removed.
	BiddingStrategyStatusEnum_REMOVED BiddingStrategyStatusEnum_BiddingStrategyStatus = 4
)

var BiddingStrategyStatusEnum_BiddingStrategyStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	4: "REMOVED",
}

var BiddingStrategyStatusEnum_BiddingStrategyStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     4,
}

func (x BiddingStrategyStatusEnum_BiddingStrategyStatus) String() string {
	return proto.EnumName(BiddingStrategyStatusEnum_BiddingStrategyStatus_name, int32(x))
}

func (BiddingStrategyStatusEnum_BiddingStrategyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24139d4c475931a6, []int{0, 0}
}

// Message describing BiddingStrategy statuses.
type BiddingStrategyStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingStrategyStatusEnum) Reset()         { *m = BiddingStrategyStatusEnum{} }
func (m *BiddingStrategyStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyStatusEnum) ProtoMessage()    {}
func (*BiddingStrategyStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_24139d4c475931a6, []int{0}
}

func (m *BiddingStrategyStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyStatusEnum.Unmarshal(m, b)
}
func (m *BiddingStrategyStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyStatusEnum.Marshal(b, m, deterministic)
}
func (m *BiddingStrategyStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyStatusEnum.Merge(m, src)
}
func (m *BiddingStrategyStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyStatusEnum.Size(m)
}
func (m *BiddingStrategyStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus", BiddingStrategyStatusEnum_BiddingStrategyStatus_name, BiddingStrategyStatusEnum_BiddingStrategyStatus_value)
	proto.RegisterType((*BiddingStrategyStatusEnum)(nil), "google.ads.googleads.v3.enums.BiddingStrategyStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/bidding_strategy_status.proto", fileDescriptor_24139d4c475931a6)
}

var fileDescriptor_24139d4c475931a6 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x7d, 0x9d, 0x27, 0x0a, 0xe9, 0xc2, 0xa1, 0xe0, 0xc2, 0x62, 0x17, 0xed, 0x07, 0x24, 0x8b,
	0xec, 0xd2, 0x55, 0xc6, 0xc6, 0x52, 0xd4, 0x69, 0xb1, 0x74, 0x04, 0x19, 0x28, 0xa9, 0x19, 0xc2,
	0x40, 0x9b, 0x94, 0x26, 0x53, 0xf0, 0x77, 0x5c, 0xfa, 0x29, 0x7e, 0x8a, 0x1b, 0x7f, 0x41, 0x92,
	0xcc, 0xcc, 0xaa, 0xba, 0x09, 0xe7, 0xe6, 0xdc, 0x73, 0xee, 0xb9, 0x17, 0x8c, 0xa5, 0xd6, 0x72,
	0x5b, 0x20, 0x2e, 0x0c, 0x0a, 0xd0, 0xa1, 0x23, 0x46, 0x85, 0xaa, 0x76, 0x06, 0x6d, 0x4a, 0x21,
	0x4a, 0x25, 0xd7, 0xc6, 0x1e, 0xb8, 0x2d, 0xe4, 0xdb, 0xda, 0x58, 0x6e, 0x2b, 0x03, 0xf7, 0x07,
	0x6d, 0x75, 0x6f, 0x10, 0x14, 0x90, 0x0b, 0x03, 0x5b, 0x31, 0x3c, 0x62, 0xe8, 0xc5, 0xfd, 0x9b,
	0xc6, 0x7b, 0x5f, 0x22, 0xae, 0x94, 0xb6, 0xdc, 0x96, 0x5a, 0xd5, 0xe2, 0xd1, 0x16, 0x5c, 0x27,
	0xc1, 0x7d, 0x59, 0x9b, 0x2f, 0xbd, 0x37, 0x53, 0xd5, 0x6e, 0x34, 0x07, 0x57, 0x27, 0xc9, 0xde,
	0x25, 0xe8, 0xae, 0xd2, 0xe5, 0x82, 0xdd, 0xce, 0xee, 0x66, 0x6c, 0x12, 0xff, 0xeb, 0x75, 0xc1,
	0xc5, 0x2a, 0xbd, 0x4f, 0xe7, 0xcf, 0x69, 0xdc, 0x71, 0x05, 0x4b, 0x69, 0xf2, 0xc0, 0x26, 0x71,
	0xe4, 0x8a, 0x27, 0xf6, 0x38, 0xcf, 0xd8, 0x24, 0x3e, 0x4b, 0xbe, 0x3b, 0x60, 0xf8, 0xaa, 0x77,
	0xf0, 0xcf, 0xc4, 0x49, 0xff, 0xe4, 0xd0, 0x85, 0xcb, 0xbb, 0xe8, 0xbc, 0x24, 0xb5, 0x58, 0xea,
	0x2d, 0x57, 0x12, 0xea, 0x83, 0x44, 0xb2, 0x50, 0x7e, 0x9b, 0xe6, 0x76, 0xfb, 0xd2, 0xfc, 0x72,
	0xca, 0xb1, 0x7f, 0xdf, 0xa3, 0xff, 0x53, 0x4a, 0x3f, 0xa2, 0xc1, 0x34, 0x58, 0x51, 0x61, 0x60,
	0x80, 0x0e, 0x65, 0x18, 0xba, 0xed, 0xcd, 0x67, 0xc3, 0xe7, 0x54, 0x98, 0xbc, 0xe5, 0xf3, 0x0c,
	0xe7, 0x9e, 0xff, 0x8a, 0x86, 0xe1, 0x93, 0x10, 0x2a, 0x0c, 0x21, 0x6d, 0x07, 0x21, 0x19, 0x26,
	0xc4, 0xf7, 0x6c, 0xce, 0x7d, 0x30, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x44, 0x5f, 0x08, 0x27,
	0xe2, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1beta2/arrow.proto

package storage

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// The IPC format to use when serializing Arrow streams.
type ArrowSerializationOptions_Format int32

const (
	// If unspecied the IPC format as of 0.15 release will be used.
	ArrowSerializationOptions_FORMAT_UNSPECIFIED ArrowSerializationOptions_Format = 0
	// Use the legacy IPC message format as of Apache Arrow Release 0.14.
	ArrowSerializationOptions_ARROW_0_14 ArrowSerializationOptions_Format = 1
	// Use the message format as of Apache Arrow Release 0.15.
	ArrowSerializationOptions_ARROW_0_15 ArrowSerializationOptions_Format = 2
)

var ArrowSerializationOptions_Format_name = map[int32]string{
	0: "FORMAT_UNSPECIFIED",
	1: "ARROW_0_14",
	2: "ARROW_0_15",
}

var ArrowSerializationOptions_Format_value = map[string]int32{
	"FORMAT_UNSPECIFIED": 0,
	"ARROW_0_14":         1,
	"ARROW_0_15":         2,
}

func (x ArrowSerializationOptions_Format) String() string {
	return proto.EnumName(ArrowSerializationOptions_Format_name, int32(x))
}

func (ArrowSerializationOptions_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eee3d978daa08622, []int{2, 0}
}

// Arrow schema as specified in
// https://arrow.apache.org/docs/python/api/datatypes.html
// and serialized to bytes using IPC:
// https://arrow.apache.org/docs/ipc.html.
//
// See code samples on how this message can be deserialized.
type ArrowSchema struct {
	// IPC serialized Arrow schema.
	SerializedSchema     []byte   `protobuf:"bytes,1,opt,name=serialized_schema,json=serializedSchema,proto3" json:"serialized_schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrowSchema) Reset()         { *m = ArrowSchema{} }
func (m *ArrowSchema) String() string { return proto.CompactTextString(m) }
func (*ArrowSchema) ProtoMessage()    {}
func (*ArrowSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee3d978daa08622, []int{0}
}

func (m *ArrowSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowSchema.Unmarshal(m, b)
}
func (m *ArrowSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowSchema.Marshal(b, m, deterministic)
}
func (m *ArrowSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowSchema.Merge(m, src)
}
func (m *ArrowSchema) XXX_Size() int {
	return xxx_messageInfo_ArrowSchema.Size(m)
}
func (m *ArrowSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowSchema proto.InternalMessageInfo

func (m *ArrowSchema) GetSerializedSchema() []byte {
	if m != nil {
		return m.SerializedSchema
	}
	return nil
}

// Arrow RecordBatch.
type ArrowRecordBatch struct {
	// IPC-serialized Arrow RecordBatch.
	SerializedRecordBatch []byte   `protobuf:"bytes,1,opt,name=serialized_record_batch,json=serializedRecordBatch,proto3" json:"serialized_record_batch,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ArrowRecordBatch) Reset()         { *m = ArrowRecordBatch{} }
func (m *ArrowRecordBatch) String() string { return proto.CompactTextString(m) }
func (*ArrowRecordBatch) ProtoMessage()    {}
func (*ArrowRecordBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee3d978daa08622, []int{1}
}

func (m *ArrowRecordBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowRecordBatch.Unmarshal(m, b)
}
func (m *ArrowRecordBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowRecordBatch.Marshal(b, m, deterministic)
}
func (m *ArrowRecordBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowRecordBatch.Merge(m, src)
}
func (m *ArrowRecordBatch) XXX_Size() int {
	return xxx_messageInfo_ArrowRecordBatch.Size(m)
}
func (m *ArrowRecordBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowRecordBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowRecordBatch proto.InternalMessageInfo

func (m *ArrowRecordBatch) GetSerializedRecordBatch() []byte {
	if m != nil {
		return m.SerializedRecordBatch
	}
	return nil
}

// Contains options specific to Arrow Serialization.
type ArrowSerializationOptions struct {
	// The Arrow IPC format to use.
	Format               ArrowSerializationOptions_Format `protobuf:"varint,1,opt,name=format,proto3,enum=google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions_Format" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ArrowSerializationOptions) Reset()         { *m = ArrowSerializationOptions{} }
func (m *ArrowSerializationOptions) String() string { return proto.CompactTextString(m) }
func (*ArrowSerializationOptions) ProtoMessage()    {}
func (*ArrowSerializationOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_eee3d978daa08622, []int{2}
}

func (m *ArrowSerializationOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowSerializationOptions.Unmarshal(m, b)
}
func (m *ArrowSerializationOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowSerializationOptions.Marshal(b, m, deterministic)
}
func (m *ArrowSerializationOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowSerializationOptions.Merge(m, src)
}
func (m *ArrowSerializationOptions) XXX_Size() int {
	return xxx_messageInfo_ArrowSerializationOptions.Size(m)
}
func (m *ArrowSerializationOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowSerializationOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowSerializationOptions proto.InternalMessageInfo

func (m *ArrowSerializationOptions) GetFormat() ArrowSerializationOptions_Format {
	if m != nil {
		return m.Format
	}
	return ArrowSerializationOptions_FORMAT_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions_Format", ArrowSerializationOptions_Format_name, ArrowSerializationOptions_Format_value)
	proto.RegisterType((*ArrowSchema)(nil), "google.cloud.bigquery.storage.v1beta2.ArrowSchema")
	proto.RegisterType((*ArrowRecordBatch)(nil), "google.cloud.bigquery.storage.v1beta2.ArrowRecordBatch")
	proto.RegisterType((*ArrowSerializationOptions)(nil), "google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1beta2/arrow.proto", fileDescriptor_eee3d978daa08622)
}

var fileDescriptor_eee3d978daa08622 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xbf, 0xf9, 0x16, 0x5d, 0x1c, 0xa5, 0x8c, 0x01, 0xff, 0x76, 0x32, 0x20, 0x28, 0x42,
	0x62, 0xeb, 0xcf, 0x42, 0x37, 0xb6, 0xda, 0x4a, 0x45, 0x6d, 0x49, 0x15, 0xc1, 0x4d, 0xc8, 0x4c,
	0x63, 0x3a, 0x30, 0xed, 0xa9, 0x99, 0x54, 0xd1, 0x1b, 0xf0, 0xae, 0xbc, 0x36, 0x99, 0x4c, 0xa4,
	0x75, 0x21, 0x76, 0x33, 0x70, 0xce, 0x3b, 0xcf, 0xc3, 0x09, 0x2f, 0xd4, 0x34, 0xa2, 0xce, 0x14,
	0x4b, 0x32, 0x9c, 0x0e, 0x58, 0x9c, 0xea, 0xe7, 0xa9, 0x32, 0x6f, 0x2c, 0xb7, 0x68, 0xa4, 0x56,
	0xec, 0xa5, 0x16, 0x2b, 0x2b, 0xeb, 0x4c, 0x1a, 0x83, 0xaf, 0x74, 0x62, 0xd0, 0x22, 0xd9, 0x2e,
	0x11, 0xea, 0x10, 0xfa, 0x8d, 0x50, 0x8f, 0x50, 0x8f, 0x44, 0x27, 0xb0, 0xd4, 0x28, 0xa8, 0x7e,
	0x32, 0x54, 0x23, 0x49, 0xf6, 0x60, 0x25, 0x57, 0x26, 0x95, 0x59, 0xfa, 0xae, 0x06, 0x22, 0x77,
	0xcb, 0x8d, 0x60, 0x2b, 0xd8, 0x59, 0xe6, 0xe1, 0x2c, 0x28, 0x7f, 0x8e, 0xae, 0x20, 0x74, 0x2c,
	0x57, 0x09, 0x9a, 0x41, 0x53, 0xda, 0x64, 0x48, 0x8e, 0x61, 0x7d, 0x4e, 0x60, 0x5c, 0x22, 0xe2,
	0x22, 0xf2, 0x9a, 0xd5, 0x59, 0x3c, 0xc7, 0x45, 0x9f, 0x01, 0x6c, 0x96, 0x87, 0xf8, 0x58, 0xda,
	0x14, 0xc7, 0xdd, 0x49, 0xf1, 0xcd, 0x89, 0x80, 0xca, 0x13, 0x9a, 0x91, 0xb4, 0x4e, 0x52, 0xad,
	0x5f, 0xd2, 0x85, 0x5e, 0x47, 0x7f, 0x35, 0xd2, 0xb6, 0xd3, 0x71, 0xaf, 0x8d, 0xce, 0xa0, 0x52,
	0x6e, 0xc8, 0x1a, 0x90, 0x76, 0x97, 0xdf, 0x34, 0xee, 0xc4, 0xfd, 0x6d, 0xbf, 0xd7, 0x3a, 0xef,
	0xb4, 0x3b, 0xad, 0x8b, 0xf0, 0x1f, 0xa9, 0x02, 0x34, 0x38, 0xef, 0x3e, 0x88, 0x7d, 0x51, 0x3b,
	0x0c, 0x83, 0x1f, 0xf3, 0x51, 0xf8, 0xbf, 0xf9, 0x11, 0xc0, 0x6e, 0x82, 0xa3, 0xc5, 0x0e, 0x6b,
	0x82, 0xbb, 0xac, 0x57, 0x34, 0xd5, 0x0b, 0x1e, 0xaf, 0x3d, 0xa4, 0x31, 0x93, 0x63, 0x4d, 0xd1,
	0x68, 0xa6, 0xd5, 0xd8, 0xf5, 0xc8, 0xca, 0x48, 0x4e, 0xd2, 0xfc, 0x8f, 0xf6, 0x4f, 0xfd, 0x1c,
	0x57, 0x1c, 0x78, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0x50, 0xd3, 0xab, 0x00, 0x35, 0x02, 0x00,
	0x00,
}

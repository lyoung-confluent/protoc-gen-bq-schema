// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bq_field.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Message containing options related to BigQuery schema generation
// and management via Protobuf.
//
// TODO: register with protobuf-global-extension-registry@google.com
type BigQueryFieldOptions struct {
	// Flag to specify that a field should be marked as 'REQUIRED' when
	// used to generate schema for BigQuery.
	Require bool `protobuf:"varint,1,opt,name=require" json:"require,omitempty"`
	// Optionally override whatever type is resolved by the schema
	// generator. This is useful, for example, to store epoch timestamps
	// with the underlying 'TIMESTAMP' type, when normally, they would
	// be structured as 'INTEGER' fields.
	TypeOverride string `protobuf:"bytes,2,opt,name=type_override,json=typeOverride" json:"type_override,omitempty"`
	// Optionally omit a field from BigQuery schema.
	Ignore bool `protobuf:"varint,3,opt,name=ignore" json:"ignore,omitempty"`
	// Set the description for a field in BigQuery schema.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
}

func (m *BigQueryFieldOptions) Reset()                    { *m = BigQueryFieldOptions{} }
func (m *BigQueryFieldOptions) String() string            { return proto.CompactTextString(m) }
func (*BigQueryFieldOptions) ProtoMessage()               {}
func (*BigQueryFieldOptions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BigQueryFieldOptions) GetRequire() bool {
	if m != nil {
		return m.Require
	}
	return false
}

func (m *BigQueryFieldOptions) GetTypeOverride() string {
	if m != nil {
		return m.TypeOverride
	}
	return ""
}

func (m *BigQueryFieldOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *BigQueryFieldOptions) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

var E_BigQuery = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*BigQueryFieldOptions)(nil),
	Field:         1021,
	Name:          "gen_bq_schema.big_query",
	Tag:           "bytes,1021,opt,name=big_query,json=bigQuery",
	Filename:      "bq_field.proto",
}

func init() {
	proto.RegisterType((*BigQueryFieldOptions)(nil), "gen_bq_schema.BigQueryFieldOptions")
	proto.RegisterExtension(E_BigQuery)
}

func init() { proto.RegisterFile("bq_field.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0xca, 0xd6, 0x65, 0xce, 0x43, 0x10, 0x09, 0x82, 0x10, 0xdc, 0xa5, 0xa7, 0x0c,
	0xf4, 0xe6, 0x71, 0x07, 0xaf, 0xc3, 0x1e, 0xbd, 0x44, 0xb3, 0xbe, 0xc5, 0x07, 0xb3, 0xaf, 0x7d,
	0x6d, 0x85, 0xfd, 0x15, 0xfe, 0xc5, 0x82, 0x34, 0x6d, 0xa1, 0x82, 0xa7, 0xf0, 0x7e, 0xe1, 0xfb,
	0xf8, 0x7e, 0xf2, 0xca, 0x57, 0xee, 0x80, 0x70, 0xcc, 0x6d, 0xc9, 0xd4, 0x90, 0x5a, 0x05, 0x28,
	0x9c, 0xaf, 0x5c, 0xbd, 0xff, 0x80, 0xcf, 0xf7, 0x5b, 0x13, 0x88, 0xc2, 0x11, 0x36, 0xf1, 0xd3,
	0xb7, 0x87, 0x4d, 0x0e, 0xf5, 0x9e, 0xb1, 0x6c, 0x88, 0xfb, 0xc0, 0xfd, 0xb7, 0x90, 0xd7, 0x5b,
	0x0c, 0x2f, 0x2d, 0xf0, 0xe9, 0xb9, 0x2b, 0xda, 0x95, 0x0d, 0x52, 0x51, 0x2b, 0x2d, 0xe7, 0x0c,
	0x55, 0x8b, 0x0c, 0x5a, 0x18, 0x91, 0x26, 0xd9, 0x78, 0xaa, 0xb5, 0x5c, 0x35, 0xa7, 0x12, 0x1c,
	0x7d, 0x01, 0x33, 0xe6, 0xa0, 0xcf, 0x8c, 0x48, 0x17, 0xd9, 0x65, 0x07, 0x77, 0x03, 0x53, 0x37,
	0x72, 0x86, 0xa1, 0x20, 0x06, 0x7d, 0x1e, 0xd3, 0xc3, 0xa5, 0x8c, 0x5c, 0x8e, 0x1b, 0x90, 0x0a,
	0x7d, 0x11, 0xa3, 0x53, 0xf4, 0xf4, 0x26, 0x17, 0x1e, 0x83, 0xab, 0xba, 0x45, 0xea, 0xce, 0xf6,
	0x06, 0x76, 0x34, 0xb0, 0xd3, 0x91, 0xfa, 0x67, 0x6e, 0x44, 0xba, 0x7c, 0x58, 0xdb, 0x3f, 0xda,
	0xf6, 0x3f, 0xa1, 0x2c, 0xf1, 0x03, 0xdd, 0x26, 0xaf, 0xb3, 0x58, 0x57, 0xfb, 0xfe, 0x7d, 0xfc,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x34, 0x80, 0xf5, 0x47, 0x01, 0x00, 0x00,
}

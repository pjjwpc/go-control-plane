// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/substitution_format_string.proto

package envoy_config_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type SubstitutionFormatString struct {
	// Types that are valid to be assigned to Format:
	//	*SubstitutionFormatString_TextFormat
	//	*SubstitutionFormatString_JsonFormat
	Format               isSubstitutionFormatString_Format `protobuf_oneof:"format"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SubstitutionFormatString) Reset()         { *m = SubstitutionFormatString{} }
func (m *SubstitutionFormatString) String() string { return proto.CompactTextString(m) }
func (*SubstitutionFormatString) ProtoMessage()    {}
func (*SubstitutionFormatString) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8eb96e547a61cc8, []int{0}
}

func (m *SubstitutionFormatString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubstitutionFormatString.Unmarshal(m, b)
}
func (m *SubstitutionFormatString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubstitutionFormatString.Marshal(b, m, deterministic)
}
func (m *SubstitutionFormatString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubstitutionFormatString.Merge(m, src)
}
func (m *SubstitutionFormatString) XXX_Size() int {
	return xxx_messageInfo_SubstitutionFormatString.Size(m)
}
func (m *SubstitutionFormatString) XXX_DiscardUnknown() {
	xxx_messageInfo_SubstitutionFormatString.DiscardUnknown(m)
}

var xxx_messageInfo_SubstitutionFormatString proto.InternalMessageInfo

type isSubstitutionFormatString_Format interface {
	isSubstitutionFormatString_Format()
}

type SubstitutionFormatString_TextFormat struct {
	TextFormat string `protobuf:"bytes,1,opt,name=text_format,json=textFormat,proto3,oneof"`
}

type SubstitutionFormatString_JsonFormat struct {
	JsonFormat *_struct.Struct `protobuf:"bytes,2,opt,name=json_format,json=jsonFormat,proto3,oneof"`
}

func (*SubstitutionFormatString_TextFormat) isSubstitutionFormatString_Format() {}

func (*SubstitutionFormatString_JsonFormat) isSubstitutionFormatString_Format() {}

func (m *SubstitutionFormatString) GetFormat() isSubstitutionFormatString_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *SubstitutionFormatString) GetTextFormat() string {
	if x, ok := m.GetFormat().(*SubstitutionFormatString_TextFormat); ok {
		return x.TextFormat
	}
	return ""
}

func (m *SubstitutionFormatString) GetJsonFormat() *_struct.Struct {
	if x, ok := m.GetFormat().(*SubstitutionFormatString_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubstitutionFormatString) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubstitutionFormatString_TextFormat)(nil),
		(*SubstitutionFormatString_JsonFormat)(nil),
	}
}

func init() {
	proto.RegisterType((*SubstitutionFormatString)(nil), "envoy.config.core.v3.SubstitutionFormatString")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3/substitution_format_string.proto", fileDescriptor_d8eb96e547a61cc8)
}

var fileDescriptor_d8eb96e547a61cc8 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0xcd, 0xa8, 0x55, 0x53, 0x84, 0x32, 0x08, 0x2d, 0xc5, 0xc2, 0xd0, 0x55, 0x71, 0x91,
	0x80, 0x83, 0x17, 0x08, 0x22, 0x6e, 0x84, 0xe2, 0x1c, 0xa0, 0x64, 0xa6, 0x33, 0x43, 0xa4, 0xe6,
	0x0d, 0xc9, 0xcb, 0xd0, 0xee, 0x5c, 0x7b, 0x05, 0x6f, 0xe2, 0x09, 0xdc, 0x7a, 0x1b, 0xe9, 0x4a,
	0x26, 0x69, 0xc5, 0x85, 0xee, 0x02, 0x7f, 0xbe, 0xff, 0x7d, 0xef, 0xd1, 0x9b, 0x52, 0xb7, 0xb0,
	0xe1, 0x05, 0xe8, 0x4a, 0xd5, 0xbc, 0x00, 0x53, 0xf2, 0x36, 0xe5, 0xd6, 0xe5, 0x16, 0x15, 0x3a,
	0x54, 0xa0, 0x17, 0x15, 0x98, 0x67, 0x89, 0x0b, 0x8b, 0x46, 0xe9, 0x9a, 0x35, 0x06, 0x10, 0xe2,
	0x0b, 0x8f, 0xb1, 0x80, 0xb1, 0x0e, 0x63, 0x6d, 0x3a, 0xbe, 0xac, 0x01, 0xea, 0x55, 0xc9, 0xfd,
	0x9f, 0xdc, 0x55, 0xdc, 0xa2, 0x71, 0x05, 0x06, 0x66, 0x3c, 0x71, 0xcb, 0x46, 0x72, 0xa9, 0x35,
	0xa0, 0xec, 0x8a, 0x2d, 0xb7, 0x28, 0xd1, 0xd9, 0x5d, 0x3c, 0x6c, 0xe5, 0x4a, 0x2d, 0x25, 0x96,
	0x7c, 0xff, 0x08, 0xc1, 0xf4, 0x8d, 0xd0, 0x51, 0xf6, 0x4b, 0xe8, 0xce, 0xfb, 0x64, 0x5e, 0x27,
	0xbe, 0xa2, 0x7d, 0x2c, 0xd7, 0xb8, 0x93, 0x1c, 0x91, 0x84, 0xcc, 0xce, 0xc4, 0xc9, 0x56, 0x1c,
	0x99, 0x28, 0x21, 0xf7, 0x07, 0x8f, 0xb4, 0x4b, 0x03, 0x11, 0xdf, 0xd2, 0xfe, 0x93, 0xfd, 0x59,
	0x68, 0x14, 0x25, 0x64, 0xd6, 0xbf, 0x1e, 0xb2, 0x20, 0xcd, 0xf6, 0xd2, 0x2c, 0xf3, 0xd2, 0xe2,
	0x74, 0x2b, 0x8e, 0x5f, 0x49, 0x34, 0xf0, 0x2d, 0x1d, 0x17, 0x5a, 0xc4, 0x39, 0xed, 0x85, 0x82,
	0xf8, 0xf0, 0x4b, 0x10, 0xf1, 0xf0, 0xfe, 0xf2, 0xf1, 0xd9, 0x8b, 0x06, 0x11, 0x9d, 0x2a, 0x60,
	0xfe, 0x2c, 0x8d, 0x81, 0xf5, 0x86, 0xfd, 0x75, 0x21, 0x31, 0xf9, 0x6f, 0x91, 0x79, 0x37, 0x7d,
	0x4e, 0xf2, 0x9e, 0xd7, 0x48, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x25, 0xac, 0xdc, 0x98,
	0x01, 0x00, 0x00,
}
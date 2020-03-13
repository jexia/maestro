// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations/annotations.proto

package annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Service struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Codec                string   `protobuf:"bytes,3,opt,name=codec,proto3" json:"codec,omitempty"`
	Method               string   `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_21dfaf6fd39fa3b7, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Service) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Service) GetCodec() string {
	if m != nil {
		return m.Codec
	}
	return ""
}

func (m *Service) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type HTTP struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTP) Reset()         { *m = HTTP{} }
func (m *HTTP) String() string { return proto.CompactTextString(m) }
func (*HTTP) ProtoMessage()    {}
func (*HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_21dfaf6fd39fa3b7, []int{1}
}

func (m *HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTP.Unmarshal(m, b)
}
func (m *HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTP.Marshal(b, m, deterministic)
}
func (m *HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTP.Merge(m, src)
}
func (m *HTTP) XXX_Size() int {
	return xxx_messageInfo_HTTP.Size(m)
}
func (m *HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_HTTP proto.InternalMessageInfo

func (m *HTTP) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *HTTP) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Service)(nil),
	Field:         50012,
	Name:          "maestro.service",
	Tag:           "bytes,50012,opt,name=service",
	Filename:      "annotations/annotations.proto",
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HTTP)(nil),
	Field:         50011,
	Name:          "maestro.http",
	Tag:           "bytes,50011,opt,name=http",
	Filename:      "annotations/annotations.proto",
}

func init() {
	proto.RegisterType((*Service)(nil), "maestro.Service")
	proto.RegisterType((*HTTP)(nil), "maestro.HTTP")
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("annotations/annotations.proto", fileDescriptor_21dfaf6fd39fa3b7)
}

var fileDescriptor_21dfaf6fd39fa3b7 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xd9, 0x7d, 0xfb, 0x6e, 0x35, 0x22, 0x48, 0x10, 0x29, 0x0b, 0xea, 0xb2, 0x88, 0x78,
	0x4a, 0x41, 0x6f, 0x3d, 0xea, 0xc5, 0xcb, 0xa2, 0xd4, 0x3d, 0x79, 0x6b, 0xd3, 0xd8, 0x46, 0xb6,
	0x9d, 0x90, 0xcc, 0x8a, 0x9f, 0xc0, 0x2f, 0xa9, 0x1f, 0x46, 0x3a, 0x4d, 0x4a, 0xc1, 0xdb, 0x3c,
	0x93, 0x67, 0x7e, 0xf3, 0x27, 0xec, 0xbc, 0xe8, 0x3a, 0xc0, 0x02, 0x35, 0x74, 0x2e, 0x9d, 0xc4,
	0xc2, 0x58, 0x40, 0xe0, 0x71, 0x5b, 0x28, 0x87, 0x16, 0x96, 0xab, 0x1a, 0xa0, 0xde, 0xa9, 0x94,
	0xd2, 0xe5, 0xfe, 0x2d, 0xad, 0x94, 0x93, 0x56, 0x1b, 0x04, 0x3b, 0x58, 0xd7, 0x35, 0x8b, 0x5f,
	0x94, 0xfd, 0xd0, 0x52, 0x71, 0xce, 0xa2, 0x06, 0x1c, 0x26, 0xb3, 0xd5, 0xec, 0xe6, 0x30, 0xa7,
	0x98, 0x2f, 0xd9, 0x01, 0xf9, 0x24, 0xec, 0x92, 0x39, 0xe5, 0x47, 0xcd, 0x4f, 0xd9, 0x7f, 0x09,
	0x95, 0x92, 0xc9, 0x3f, 0x7a, 0x18, 0x04, 0x3f, 0x63, 0x8b, 0x56, 0x61, 0x03, 0x55, 0x12, 0x51,
	0xda, 0xab, 0x75, 0xc6, 0xa2, 0xc7, 0xed, 0xf6, 0xb9, 0x27, 0xaa, 0xae, 0x32, 0xa0, 0xbb, 0xd0,
	0x69, 0xd4, 0x93, 0xda, 0xf9, 0xb4, 0x36, 0xdb, 0xb0, 0xd8, 0xf9, 0x21, 0x2f, 0xc5, 0xb0, 0x92,
	0x08, 0x2b, 0x09, 0x3f, 0xfe, 0x93, 0xa1, 0x0b, 0x24, 0x3f, 0x5f, 0xfd, 0x38, 0x47, 0xb7, 0x27,
	0xc2, 0x1f, 0x21, 0x18, 0xf2, 0xc0, 0xc8, 0x1e, 0x58, 0xd4, 0x20, 0x1a, 0x7e, 0xf1, 0x87, 0xb5,
	0xa1, 0x7e, 0x01, 0xf5, 0xed, 0x51, 0xc7, 0x23, 0xaa, 0xdf, 0x20, 0xa7, 0xe2, 0xfb, 0xeb, 0xd7,
	0xab, 0x5a, 0x63, 0xb3, 0x2f, 0x85, 0x84, 0x36, 0x7d, 0x57, 0x9f, 0xba, 0x48, 0xbd, 0x6d, 0xfa,
	0x23, 0xe5, 0x82, 0xe0, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xef, 0x12, 0x05, 0xb3,
	0x01, 0x00, 0x00,
}

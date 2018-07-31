// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/interface.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Defines a strongly typed variable.
type Variable struct {
	// Variable name. This can be referenced in input/output bindings to pass data between
	// nodes/workflows. Variable name is case-sensitive.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Variable literal type.
	Type *LiteralType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// +optional string describing input variable
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Variable) Reset()         { *m = Variable{} }
func (m *Variable) String() string { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()    {}
func (*Variable) Descriptor() ([]byte, []int) {
	return fileDescriptor_interface_6cf1bf4b3e4bfa47, []int{0}
}
func (m *Variable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variable.Unmarshal(m, b)
}
func (m *Variable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variable.Marshal(b, m, deterministic)
}
func (dst *Variable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variable.Merge(dst, src)
}
func (m *Variable) XXX_Size() int {
	return xxx_messageInfo_Variable.Size(m)
}
func (m *Variable) XXX_DiscardUnknown() {
	xxx_messageInfo_Variable.DiscardUnknown(m)
}

var xxx_messageInfo_Variable proto.InternalMessageInfo

func (m *Variable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variable) GetType() *LiteralType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Variable) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Defines strongly typed inputs and outputs.
type TypedInterface struct {
	Inputs               []*Variable `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*Variable `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TypedInterface) Reset()         { *m = TypedInterface{} }
func (m *TypedInterface) String() string { return proto.CompactTextString(m) }
func (*TypedInterface) ProtoMessage()    {}
func (*TypedInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_interface_6cf1bf4b3e4bfa47, []int{1}
}
func (m *TypedInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedInterface.Unmarshal(m, b)
}
func (m *TypedInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedInterface.Marshal(b, m, deterministic)
}
func (dst *TypedInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedInterface.Merge(dst, src)
}
func (m *TypedInterface) XXX_Size() int {
	return xxx_messageInfo_TypedInterface.Size(m)
}
func (m *TypedInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedInterface.DiscardUnknown(m)
}

var xxx_messageInfo_TypedInterface proto.InternalMessageInfo

func (m *TypedInterface) GetInputs() []*Variable {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TypedInterface) GetOutputs() []*Variable {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*Variable)(nil), "flyteidl.core.Variable")
	proto.RegisterType((*TypedInterface)(nil), "flyteidl.core.TypedInterface")
}

func init() {
	proto.RegisterFile("flyteidl/core/interface.proto", fileDescriptor_interface_6cf1bf4b3e4bfa47)
}

var fileDescriptor_interface_6cf1bf4b3e4bfa47 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0xee, 0xb2, 0xea, 0x14, 0x3d, 0xe4, 0x62, 0x5d, 0x10, 0xca, 0x9e, 0x7a, 0x31,
	0xd1, 0xfa, 0x0f, 0xbc, 0x09, 0x9e, 0x8a, 0x78, 0xf0, 0x96, 0xb6, 0xd3, 0x3a, 0x90, 0x4d, 0x42,
	0x3a, 0x3d, 0xf4, 0xdf, 0x4b, 0xa3, 0x15, 0xeb, 0xc1, 0x5b, 0xc8, 0xfb, 0xde, 0x9b, 0xc7, 0x83,
	0xdb, 0xce, 0x4c, 0x8c, 0xd4, 0x1a, 0xd5, 0xb8, 0x80, 0x8a, 0x2c, 0x63, 0xe8, 0x74, 0x83, 0xd2,
	0x07, 0xc7, 0x4e, 0x5c, 0x2e, 0xb2, 0x9c, 0xe5, 0xc3, 0xcd, 0x9a, 0xe6, 0xc9, 0xe3, 0xf0, 0x45,
	0x1e, 0x3d, 0x9c, 0xbf, 0xe9, 0x40, 0xba, 0x36, 0x28, 0x04, 0xec, 0xac, 0x3e, 0x61, 0x96, 0xe4,
	0x49, 0x71, 0x51, 0xc5, 0xb7, 0x90, 0xb0, 0x9b, 0xf1, 0x6c, 0x93, 0x27, 0x45, 0x5a, 0x1e, 0xe4,
	0x2a, 0x58, 0xbe, 0x10, 0x63, 0xd0, 0xe6, 0x75, 0xf2, 0x58, 0x45, 0x4e, 0xe4, 0x90, 0xb6, 0x38,
	0x34, 0x81, 0x3c, 0x93, 0xb3, 0xd9, 0x36, 0x46, 0xfd, 0xfe, 0x3a, 0x32, 0x5c, 0xcd, 0x7c, 0xfb,
	0xbc, 0x74, 0x16, 0x0a, 0xf6, 0x64, 0xfd, 0xc8, 0x43, 0x96, 0xe4, 0xdb, 0x22, 0x2d, 0xaf, 0xff,
	0x5c, 0x59, 0x0a, 0x56, 0xdf, 0x98, 0x78, 0x80, 0x33, 0x37, 0x72, 0x74, 0x6c, 0xfe, 0x77, 0x2c,
	0xdc, 0x53, 0xf9, 0x7e, 0xdf, 0x13, 0x7f, 0x8c, 0xb5, 0x6c, 0xdc, 0x49, 0x99, 0xa9, 0x63, 0xf5,
	0x33, 0x4a, 0x8f, 0x56, 0xf9, 0xfa, 0xae, 0x77, 0x6a, 0xb5, 0x53, 0xbd, 0x8f, 0x13, 0x3d, 0x7e,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x8c, 0xe2, 0xca, 0x6d, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/datacatalog/datacatalog.proto

package datacatalog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateDatasetRequest struct {
	Dataset              *Dataset `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDatasetRequest) Reset()         { *m = CreateDatasetRequest{} }
func (m *CreateDatasetRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDatasetRequest) ProtoMessage()    {}
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{0}
}
func (m *CreateDatasetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatasetRequest.Unmarshal(m, b)
}
func (m *CreateDatasetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatasetRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDatasetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatasetRequest.Merge(dst, src)
}
func (m *CreateDatasetRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDatasetRequest.Size(m)
}
func (m *CreateDatasetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatasetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatasetRequest proto.InternalMessageInfo

func (m *CreateDatasetRequest) GetDataset() *Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type CreateDatasetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDatasetResponse) Reset()         { *m = CreateDatasetResponse{} }
func (m *CreateDatasetResponse) String() string { return proto.CompactTextString(m) }
func (*CreateDatasetResponse) ProtoMessage()    {}
func (*CreateDatasetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{1}
}
func (m *CreateDatasetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatasetResponse.Unmarshal(m, b)
}
func (m *CreateDatasetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatasetResponse.Marshal(b, m, deterministic)
}
func (dst *CreateDatasetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatasetResponse.Merge(dst, src)
}
func (m *CreateDatasetResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDatasetResponse.Size(m)
}
func (m *CreateDatasetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatasetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatasetResponse proto.InternalMessageInfo

type GetDatasetRequest struct {
	Dataset              *DatasetID `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetDatasetRequest) Reset()         { *m = GetDatasetRequest{} }
func (m *GetDatasetRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatasetRequest) ProtoMessage()    {}
func (*GetDatasetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{2}
}
func (m *GetDatasetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatasetRequest.Unmarshal(m, b)
}
func (m *GetDatasetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatasetRequest.Marshal(b, m, deterministic)
}
func (dst *GetDatasetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatasetRequest.Merge(dst, src)
}
func (m *GetDatasetRequest) XXX_Size() int {
	return xxx_messageInfo_GetDatasetRequest.Size(m)
}
func (m *GetDatasetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatasetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatasetRequest proto.InternalMessageInfo

func (m *GetDatasetRequest) GetDataset() *DatasetID {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type GetDatasetResponse struct {
	Dataset              *Dataset `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatasetResponse) Reset()         { *m = GetDatasetResponse{} }
func (m *GetDatasetResponse) String() string { return proto.CompactTextString(m) }
func (*GetDatasetResponse) ProtoMessage()    {}
func (*GetDatasetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{3}
}
func (m *GetDatasetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatasetResponse.Unmarshal(m, b)
}
func (m *GetDatasetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatasetResponse.Marshal(b, m, deterministic)
}
func (dst *GetDatasetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatasetResponse.Merge(dst, src)
}
func (m *GetDatasetResponse) XXX_Size() int {
	return xxx_messageInfo_GetDatasetResponse.Size(m)
}
func (m *GetDatasetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatasetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatasetResponse proto.InternalMessageInfo

func (m *GetDatasetResponse) GetDataset() *Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type GetArtifactRequest struct {
	Dataset *DatasetID `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Types that are valid to be assigned to QueryHandle:
	//	*GetArtifactRequest_ArtifactId
	//	*GetArtifactRequest_TagName
	QueryHandle          isGetArtifactRequest_QueryHandle `protobuf_oneof:"query_handle"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{4}
}
func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(dst, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetDataset() *DatasetID {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type isGetArtifactRequest_QueryHandle interface {
	isGetArtifactRequest_QueryHandle()
}

type GetArtifactRequest_ArtifactId struct {
	ArtifactId string `protobuf:"bytes,2,opt,name=artifact_id,json=artifactId,proto3,oneof"`
}

type GetArtifactRequest_TagName struct {
	TagName string `protobuf:"bytes,3,opt,name=tag_name,json=tagName,proto3,oneof"`
}

func (*GetArtifactRequest_ArtifactId) isGetArtifactRequest_QueryHandle() {}

func (*GetArtifactRequest_TagName) isGetArtifactRequest_QueryHandle() {}

func (m *GetArtifactRequest) GetQueryHandle() isGetArtifactRequest_QueryHandle {
	if m != nil {
		return m.QueryHandle
	}
	return nil
}

func (m *GetArtifactRequest) GetArtifactId() string {
	if x, ok := m.GetQueryHandle().(*GetArtifactRequest_ArtifactId); ok {
		return x.ArtifactId
	}
	return ""
}

func (m *GetArtifactRequest) GetTagName() string {
	if x, ok := m.GetQueryHandle().(*GetArtifactRequest_TagName); ok {
		return x.TagName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetArtifactRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetArtifactRequest_OneofMarshaler, _GetArtifactRequest_OneofUnmarshaler, _GetArtifactRequest_OneofSizer, []interface{}{
		(*GetArtifactRequest_ArtifactId)(nil),
		(*GetArtifactRequest_TagName)(nil),
	}
}

func _GetArtifactRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetArtifactRequest)
	// query_handle
	switch x := m.QueryHandle.(type) {
	case *GetArtifactRequest_ArtifactId:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ArtifactId)
	case *GetArtifactRequest_TagName:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TagName)
	case nil:
	default:
		return fmt.Errorf("GetArtifactRequest.QueryHandle has unexpected type %T", x)
	}
	return nil
}

func _GetArtifactRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetArtifactRequest)
	switch tag {
	case 2: // query_handle.artifact_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.QueryHandle = &GetArtifactRequest_ArtifactId{x}
		return true, err
	case 3: // query_handle.tag_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.QueryHandle = &GetArtifactRequest_TagName{x}
		return true, err
	default:
		return false, nil
	}
}

func _GetArtifactRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetArtifactRequest)
	// query_handle
	switch x := m.QueryHandle.(type) {
	case *GetArtifactRequest_ArtifactId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ArtifactId)))
		n += len(x.ArtifactId)
	case *GetArtifactRequest_TagName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.TagName)))
		n += len(x.TagName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GetArtifactResponse struct {
	Artifact             *Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{5}
}
func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (dst *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(dst, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type CreateArtifactRequest struct {
	Artifact             *Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateArtifactRequest) Reset()         { *m = CreateArtifactRequest{} }
func (m *CreateArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactRequest) ProtoMessage()    {}
func (*CreateArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{6}
}
func (m *CreateArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactRequest.Unmarshal(m, b)
}
func (m *CreateArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactRequest.Marshal(b, m, deterministic)
}
func (dst *CreateArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactRequest.Merge(dst, src)
}
func (m *CreateArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactRequest.Size(m)
}
func (m *CreateArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactRequest proto.InternalMessageInfo

func (m *CreateArtifactRequest) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type CreateArtifactResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateArtifactResponse) Reset()         { *m = CreateArtifactResponse{} }
func (m *CreateArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactResponse) ProtoMessage()    {}
func (*CreateArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{7}
}
func (m *CreateArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactResponse.Unmarshal(m, b)
}
func (m *CreateArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactResponse.Marshal(b, m, deterministic)
}
func (dst *CreateArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactResponse.Merge(dst, src)
}
func (m *CreateArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactResponse.Size(m)
}
func (m *CreateArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactResponse proto.InternalMessageInfo

type AddTagRequest struct {
	Tag                  *Tag     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTagRequest) Reset()         { *m = AddTagRequest{} }
func (m *AddTagRequest) String() string { return proto.CompactTextString(m) }
func (*AddTagRequest) ProtoMessage()    {}
func (*AddTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{8}
}
func (m *AddTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTagRequest.Unmarshal(m, b)
}
func (m *AddTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTagRequest.Marshal(b, m, deterministic)
}
func (dst *AddTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTagRequest.Merge(dst, src)
}
func (m *AddTagRequest) XXX_Size() int {
	return xxx_messageInfo_AddTagRequest.Size(m)
}
func (m *AddTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTagRequest proto.InternalMessageInfo

func (m *AddTagRequest) GetTag() *Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

type AddTagResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTagResponse) Reset()         { *m = AddTagResponse{} }
func (m *AddTagResponse) String() string { return proto.CompactTextString(m) }
func (*AddTagResponse) ProtoMessage()    {}
func (*AddTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{9}
}
func (m *AddTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTagResponse.Unmarshal(m, b)
}
func (m *AddTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTagResponse.Marshal(b, m, deterministic)
}
func (dst *AddTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTagResponse.Merge(dst, src)
}
func (m *AddTagResponse) XXX_Size() int {
	return xxx_messageInfo_AddTagResponse.Size(m)
}
func (m *AddTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddTagResponse proto.InternalMessageInfo

type Dataset struct {
	Id                   *DatasetID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata             *Metadata  `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{10}
}
func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (dst *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(dst, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

func (m *Dataset) GetId() *DatasetID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Dataset) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DatasetID struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Domain               string   `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasetID) Reset()         { *m = DatasetID{} }
func (m *DatasetID) String() string { return proto.CompactTextString(m) }
func (*DatasetID) ProtoMessage()    {}
func (*DatasetID) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{11}
}
func (m *DatasetID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetID.Unmarshal(m, b)
}
func (m *DatasetID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetID.Marshal(b, m, deterministic)
}
func (dst *DatasetID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetID.Merge(dst, src)
}
func (m *DatasetID) XXX_Size() int {
	return xxx_messageInfo_DatasetID.Size(m)
}
func (m *DatasetID) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetID.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetID proto.InternalMessageInfo

func (m *DatasetID) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *DatasetID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasetID) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DatasetID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Artifact struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Dataset              *DatasetID      `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
	Data                 []*ArtifactData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Metadata             *Metadata       `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{12}
}
func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (dst *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(dst, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Artifact) GetDataset() *DatasetID {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func (m *Artifact) GetData() []*ArtifactData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Artifact) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ArtifactData struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                *core.Literal `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ArtifactData) Reset()         { *m = ArtifactData{} }
func (m *ArtifactData) String() string { return proto.CompactTextString(m) }
func (*ArtifactData) ProtoMessage()    {}
func (*ArtifactData) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{13}
}
func (m *ArtifactData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactData.Unmarshal(m, b)
}
func (m *ArtifactData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactData.Marshal(b, m, deterministic)
}
func (dst *ArtifactData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactData.Merge(dst, src)
}
func (m *ArtifactData) XXX_Size() int {
	return xxx_messageInfo_ArtifactData.Size(m)
}
func (m *ArtifactData) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactData.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactData proto.InternalMessageInfo

func (m *ArtifactData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactData) GetValue() *core.Literal {
	if m != nil {
		return m.Value
	}
	return nil
}

type Tag struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ArtifactId           string     `protobuf:"bytes,2,opt,name=artifact_id,json=artifactId,proto3" json:"artifact_id,omitempty"`
	Dataset              *DatasetID `protobuf:"bytes,3,opt,name=dataset,proto3" json:"dataset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{14}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetArtifactId() string {
	if m != nil {
		return m.ArtifactId
	}
	return ""
}

func (m *Tag) GetDataset() *DatasetID {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type Metadata struct {
	KeyMap               map[string]string `protobuf:"bytes,1,rep,name=key_map,json=keyMap,proto3" json:"key_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_datacatalog_afc9e3f0c2661bf5, []int{15}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetKeyMap() map[string]string {
	if m != nil {
		return m.KeyMap
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDatasetRequest)(nil), "datacatalog.CreateDatasetRequest")
	proto.RegisterType((*CreateDatasetResponse)(nil), "datacatalog.CreateDatasetResponse")
	proto.RegisterType((*GetDatasetRequest)(nil), "datacatalog.GetDatasetRequest")
	proto.RegisterType((*GetDatasetResponse)(nil), "datacatalog.GetDatasetResponse")
	proto.RegisterType((*GetArtifactRequest)(nil), "datacatalog.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "datacatalog.GetArtifactResponse")
	proto.RegisterType((*CreateArtifactRequest)(nil), "datacatalog.CreateArtifactRequest")
	proto.RegisterType((*CreateArtifactResponse)(nil), "datacatalog.CreateArtifactResponse")
	proto.RegisterType((*AddTagRequest)(nil), "datacatalog.AddTagRequest")
	proto.RegisterType((*AddTagResponse)(nil), "datacatalog.AddTagResponse")
	proto.RegisterType((*Dataset)(nil), "datacatalog.Dataset")
	proto.RegisterType((*DatasetID)(nil), "datacatalog.DatasetID")
	proto.RegisterType((*Artifact)(nil), "datacatalog.Artifact")
	proto.RegisterType((*ArtifactData)(nil), "datacatalog.ArtifactData")
	proto.RegisterType((*Tag)(nil), "datacatalog.Tag")
	proto.RegisterType((*Metadata)(nil), "datacatalog.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "datacatalog.Metadata.KeyMapEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataCatalogClient is the client API for DataCatalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataCatalogClient interface {
	CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error)
	GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*GetDatasetResponse, error)
	CreateArtifact(ctx context.Context, in *CreateArtifactRequest, opts ...grpc.CallOption) (*CreateArtifactResponse, error)
	GetArtifact(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*GetArtifactResponse, error)
	AddTag(ctx context.Context, in *AddTagRequest, opts ...grpc.CallOption) (*AddTagResponse, error)
}

type dataCatalogClient struct {
	cc *grpc.ClientConn
}

func NewDataCatalogClient(cc *grpc.ClientConn) DataCatalogClient {
	return &dataCatalogClient{cc}
}

func (c *dataCatalogClient) CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error) {
	out := new(CreateDatasetResponse)
	err := c.cc.Invoke(ctx, "/datacatalog.DataCatalog/CreateDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataCatalogClient) GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*GetDatasetResponse, error) {
	out := new(GetDatasetResponse)
	err := c.cc.Invoke(ctx, "/datacatalog.DataCatalog/GetDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataCatalogClient) CreateArtifact(ctx context.Context, in *CreateArtifactRequest, opts ...grpc.CallOption) (*CreateArtifactResponse, error) {
	out := new(CreateArtifactResponse)
	err := c.cc.Invoke(ctx, "/datacatalog.DataCatalog/CreateArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataCatalogClient) GetArtifact(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*GetArtifactResponse, error) {
	out := new(GetArtifactResponse)
	err := c.cc.Invoke(ctx, "/datacatalog.DataCatalog/GetArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataCatalogClient) AddTag(ctx context.Context, in *AddTagRequest, opts ...grpc.CallOption) (*AddTagResponse, error) {
	out := new(AddTagResponse)
	err := c.cc.Invoke(ctx, "/datacatalog.DataCatalog/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataCatalogServer is the server API for DataCatalog service.
type DataCatalogServer interface {
	CreateDataset(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error)
	GetDataset(context.Context, *GetDatasetRequest) (*GetDatasetResponse, error)
	CreateArtifact(context.Context, *CreateArtifactRequest) (*CreateArtifactResponse, error)
	GetArtifact(context.Context, *GetArtifactRequest) (*GetArtifactResponse, error)
	AddTag(context.Context, *AddTagRequest) (*AddTagResponse, error)
}

func RegisterDataCatalogServer(s *grpc.Server, srv DataCatalogServer) {
	s.RegisterService(&_DataCatalog_serviceDesc, srv)
}

func _DataCatalog_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServer).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacatalog.DataCatalog/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServer).CreateDataset(ctx, req.(*CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataCatalog_GetDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServer).GetDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacatalog.DataCatalog/GetDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServer).GetDataset(ctx, req.(*GetDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataCatalog_CreateArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServer).CreateArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacatalog.DataCatalog/CreateArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServer).CreateArtifact(ctx, req.(*CreateArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataCatalog_GetArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServer).GetArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacatalog.DataCatalog/GetArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServer).GetArtifact(ctx, req.(*GetArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataCatalog_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCatalogServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datacatalog.DataCatalog/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCatalogServer).AddTag(ctx, req.(*AddTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataCatalog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datacatalog.DataCatalog",
	HandlerType: (*DataCatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDataset",
			Handler:    _DataCatalog_CreateDataset_Handler,
		},
		{
			MethodName: "GetDataset",
			Handler:    _DataCatalog_GetDataset_Handler,
		},
		{
			MethodName: "CreateArtifact",
			Handler:    _DataCatalog_CreateArtifact_Handler,
		},
		{
			MethodName: "GetArtifact",
			Handler:    _DataCatalog_GetArtifact_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _DataCatalog_AddTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/datacatalog/datacatalog.proto",
}

func init() {
	proto.RegisterFile("flyteidl/datacatalog/datacatalog.proto", fileDescriptor_datacatalog_afc9e3f0c2661bf5)
}

var fileDescriptor_datacatalog_afc9e3f0c2661bf5 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x53, 0xd3, 0x40,
	0x18, 0xc6, 0x4d, 0x53, 0xfb, 0xe7, 0x0d, 0x74, 0x70, 0x05, 0x8c, 0xc1, 0x91, 0x12, 0x67, 0x18,
	0x0e, 0x1a, 0x14, 0x2e, 0xca, 0x0d, 0x01, 0x05, 0xb5, 0x0e, 0x93, 0x61, 0x9c, 0xf1, 0xd4, 0x79,
	0xed, 0x2e, 0x31, 0x36, 0x4d, 0x4a, 0xb2, 0x65, 0x26, 0x47, 0xbf, 0x84, 0x9f, 0xc0, 0xb3, 0x9f,
	0xd1, 0x49, 0xb2, 0x09, 0x9b, 0x10, 0xb0, 0x72, 0xdb, 0x3f, 0xef, 0xfe, 0xfa, 0xbc, 0xcf, 0x66,
	0x9f, 0xc2, 0xe6, 0xb9, 0x17, 0x73, 0xe6, 0x52, 0x6f, 0x9b, 0x22, 0xc7, 0x11, 0x72, 0xf4, 0x02,
	0x47, 0x1e, 0x5b, 0xd3, 0x30, 0xe0, 0x01, 0xd1, 0xa4, 0x25, 0xe3, 0x49, 0x71, 0x68, 0x14, 0x84,
	0x6c, 0xdb, 0x73, 0x39, 0x0b, 0xd1, 0x8b, 0xb2, 0x52, 0xf3, 0x1d, 0x2c, 0x1f, 0x84, 0x0c, 0x39,
	0x3b, 0x44, 0x8e, 0x11, 0xe3, 0x36, 0xbb, 0x98, 0xb1, 0x88, 0x13, 0x0b, 0xda, 0x34, 0x5b, 0xd1,
	0x95, 0xbe, 0xb2, 0xa5, 0xed, 0x2c, 0x5b, 0xf2, 0xef, 0xe4, 0xd5, 0x79, 0x91, 0xf9, 0x08, 0x56,
	0x2a, 0x9c, 0x68, 0x1a, 0xf8, 0x11, 0x33, 0x8f, 0xe0, 0xc1, 0x7b, 0xc6, 0x2b, 0xf4, 0x97, 0x55,
	0xfa, 0x6a, 0x1d, 0xfd, 0xe4, 0xf0, 0x8a, 0x7f, 0x08, 0x44, 0xc6, 0x64, 0xf0, 0xff, 0x56, 0xf9,
	0x4b, 0x49, 0x31, 0xfb, 0x21, 0x77, 0xcf, 0x71, 0x74, 0x77, 0x39, 0x64, 0x03, 0x34, 0x14, 0x90,
	0xa1, 0x4b, 0xf5, 0x46, 0x5f, 0xd9, 0xea, 0x1e, 0xdf, 0xb3, 0x21, 0x5f, 0x3c, 0xa1, 0x64, 0x0d,
	0x3a, 0x1c, 0x9d, 0xa1, 0x8f, 0x13, 0xa6, 0xab, 0x62, 0xbf, 0xcd, 0xd1, 0xf9, 0x8c, 0x13, 0xf6,
	0xb6, 0x07, 0x0b, 0x17, 0x33, 0x16, 0xc6, 0xc3, 0xef, 0xe8, 0x53, 0x8f, 0x99, 0xc7, 0xf0, 0xb0,
	0xa4, 0x4b, 0xf4, 0xf7, 0x0a, 0x3a, 0x39, 0x51, 0x28, 0x5b, 0x29, 0x29, 0x2b, 0x0e, 0x14, 0x65,
	0xe6, 0x87, 0xfc, 0x22, 0xaa, 0x4d, 0xde, 0x81, 0xa5, 0xc3, 0x6a, 0x95, 0x25, 0x6e, 0x75, 0x17,
	0x16, 0xf7, 0x29, 0x3d, 0x43, 0x27, 0xa7, 0x9b, 0xa0, 0x72, 0x74, 0x04, 0x78, 0xa9, 0x04, 0x4e,
	0xaa, 0x92, 0x4d, 0x73, 0x09, 0x7a, 0xf9, 0x21, 0x81, 0xa1, 0xd0, 0x16, 0xe6, 0x92, 0x4d, 0x68,
	0xb8, 0xf4, 0x1f, 0xf6, 0x37, 0x5c, 0x9a, 0xb4, 0x31, 0x61, 0x1c, 0x93, 0x82, 0xd4, 0xf6, 0x6a,
	0x1b, 0x03, 0xb1, 0x69, 0x17, 0x65, 0xe6, 0x18, 0xba, 0x05, 0x83, 0xe8, 0xd0, 0x9e, 0x86, 0xc1,
	0x0f, 0x26, 0x5c, 0xe8, 0xda, 0xf9, 0x94, 0x10, 0x68, 0xa6, 0x97, 0x95, 0x5e, 0xa6, 0x9d, 0x8e,
	0xc9, 0x2a, 0xb4, 0x68, 0x30, 0x41, 0xd7, 0xcf, 0xae, 0xd0, 0x16, 0xb3, 0x84, 0x72, 0xc9, 0xc2,
	0xc8, 0x0d, 0x7c, 0xbd, 0x99, 0x51, 0xc4, 0xd4, 0xfc, 0xa3, 0x40, 0x27, 0xb7, 0x8b, 0xf4, 0x8a,
	0xa6, 0xba, 0xa9, 0x78, 0xe9, 0x43, 0x6b, 0xcc, 0xf7, 0xa1, 0xbd, 0x80, 0x66, 0xda, 0xaa, 0xda,
	0x57, 0xb7, 0xb4, 0x9d, 0xc7, 0xb5, 0x37, 0x96, 0x1c, 0xb3, 0xd3, 0xb2, 0x92, 0x3b, 0xcd, 0xf9,
	0xdc, 0x39, 0x85, 0x05, 0x19, 0x54, 0xd8, 0xa0, 0x48, 0x36, 0x3c, 0x87, 0xfb, 0x97, 0xe8, 0xcd,
	0x58, 0xa1, 0x3a, 0xcf, 0x14, 0x2b, 0xc9, 0x14, 0xeb, 0x53, 0x96, 0x29, 0x76, 0x56, 0x64, 0x7a,
	0xa0, 0x9e, 0xa1, 0x53, 0x0b, 0x5a, 0xaf, 0x79, 0x37, 0xa5, 0x57, 0x23, 0x39, 0xa4, 0xce, 0x97,
	0x0c, 0x3f, 0x15, 0xe8, 0xe4, 0x6d, 0x91, 0x3d, 0x68, 0x8f, 0x59, 0x3c, 0x9c, 0xe0, 0x54, 0x57,
	0x52, 0xc7, 0x36, 0x6a, 0xdb, 0xb7, 0x3e, 0xb2, 0x78, 0x80, 0xd3, 0x23, 0x9f, 0x87, 0xb1, 0xdd,
	0x1a, 0xa7, 0x13, 0xe3, 0x0d, 0x68, 0xd2, 0x32, 0x59, 0x02, 0x75, 0xcc, 0x62, 0xa1, 0x3e, 0x19,
	0x92, 0x65, 0xd9, 0x85, 0xae, 0xe8, 0x76, 0xaf, 0xf1, 0x5a, 0xd9, 0xf9, 0xad, 0x82, 0x96, 0x48,
	0x3b, 0xc8, 0x7e, 0x87, 0x7c, 0x81, 0xc5, 0x52, 0x1a, 0x92, 0xb2, 0x8c, 0xba, 0xc4, 0x35, 0xcc,
	0xdb, 0x4a, 0x44, 0x1e, 0x0c, 0x00, 0xae, 0x52, 0x90, 0x3c, 0x2d, 0x9d, 0xb8, 0x96, 0xb2, 0xc6,
	0xfa, 0x8d, 0xfb, 0x02, 0xf7, 0x15, 0x7a, 0xe5, 0xf7, 0x4d, 0xea, 0x44, 0x54, 0x82, 0xc4, 0x78,
	0x76, 0x6b, 0x8d, 0x40, 0x9f, 0x82, 0x26, 0x05, 0x1a, 0xb9, 0x26, 0xa5, 0x0a, 0xed, 0xdf, 0x5c,
	0x20, 0x88, 0xfb, 0xd0, 0xca, 0xd2, 0x83, 0x18, 0xe5, 0x57, 0x20, 0xe7, 0x90, 0xb1, 0x56, 0xbb,
	0x97, 0x21, 0xbe, 0xb5, 0xd2, 0xff, 0xbc, 0xdd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x99,
	0xc2, 0x71, 0x48, 0x07, 0x00, 0x00,
}
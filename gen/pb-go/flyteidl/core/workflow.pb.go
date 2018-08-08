// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/workflow.proto

package core // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Defines a condition and the execution unit that should be executed if the condition is satisfied.
type IfBlock struct {
	Condition            *BooleanExpression `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	ThenNode             *Node              `protobuf:"bytes,2,opt,name=then_node,json=thenNode,proto3" json:"then_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IfBlock) Reset()         { *m = IfBlock{} }
func (m *IfBlock) String() string { return proto.CompactTextString(m) }
func (*IfBlock) ProtoMessage()    {}
func (*IfBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{0}
}
func (m *IfBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfBlock.Unmarshal(m, b)
}
func (m *IfBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfBlock.Marshal(b, m, deterministic)
}
func (dst *IfBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfBlock.Merge(dst, src)
}
func (m *IfBlock) XXX_Size() int {
	return xxx_messageInfo_IfBlock.Size(m)
}
func (m *IfBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IfBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IfBlock proto.InternalMessageInfo

func (m *IfBlock) GetCondition() *BooleanExpression {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *IfBlock) GetThenNode() *Node {
	if m != nil {
		return m.ThenNode
	}
	return nil
}

// Defines a series of if/else blocks. The first branch whose condition evaluates to true is the one to execute.
// If no conditions were satisfied, the else_node will be executed.
type IfElseBlock struct {
	// +required. First condition to evaluate.
	Case *IfBlock `protobuf:"bytes,1,opt,name=case,proto3" json:"case,omitempty"`
	// +optional. Additional branches to evaluate.
	Other []*IfBlock `protobuf:"bytes,2,rep,name=other,proto3" json:"other,omitempty"`
	// +required.
	//
	// Types that are valid to be assigned to Default:
	//	*IfElseBlock_ElseNode
	//	*IfElseBlock_Error
	Default              isIfElseBlock_Default `protobuf_oneof:"default"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IfElseBlock) Reset()         { *m = IfElseBlock{} }
func (m *IfElseBlock) String() string { return proto.CompactTextString(m) }
func (*IfElseBlock) ProtoMessage()    {}
func (*IfElseBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{1}
}
func (m *IfElseBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfElseBlock.Unmarshal(m, b)
}
func (m *IfElseBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfElseBlock.Marshal(b, m, deterministic)
}
func (dst *IfElseBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfElseBlock.Merge(dst, src)
}
func (m *IfElseBlock) XXX_Size() int {
	return xxx_messageInfo_IfElseBlock.Size(m)
}
func (m *IfElseBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_IfElseBlock.DiscardUnknown(m)
}

var xxx_messageInfo_IfElseBlock proto.InternalMessageInfo

func (m *IfElseBlock) GetCase() *IfBlock {
	if m != nil {
		return m.Case
	}
	return nil
}

func (m *IfElseBlock) GetOther() []*IfBlock {
	if m != nil {
		return m.Other
	}
	return nil
}

type isIfElseBlock_Default interface {
	isIfElseBlock_Default()
}

type IfElseBlock_ElseNode struct {
	ElseNode *Node `protobuf:"bytes,3,opt,name=else_node,json=elseNode,proto3,oneof"`
}

type IfElseBlock_Error struct {
	Error *Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*IfElseBlock_ElseNode) isIfElseBlock_Default() {}

func (*IfElseBlock_Error) isIfElseBlock_Default() {}

func (m *IfElseBlock) GetDefault() isIfElseBlock_Default {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *IfElseBlock) GetElseNode() *Node {
	if x, ok := m.GetDefault().(*IfElseBlock_ElseNode); ok {
		return x.ElseNode
	}
	return nil
}

func (m *IfElseBlock) GetError() *Error {
	if x, ok := m.GetDefault().(*IfElseBlock_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*IfElseBlock) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _IfElseBlock_OneofMarshaler, _IfElseBlock_OneofUnmarshaler, _IfElseBlock_OneofSizer, []interface{}{
		(*IfElseBlock_ElseNode)(nil),
		(*IfElseBlock_Error)(nil),
	}
}

func _IfElseBlock_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*IfElseBlock)
	// default
	switch x := m.Default.(type) {
	case *IfElseBlock_ElseNode:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ElseNode); err != nil {
			return err
		}
	case *IfElseBlock_Error:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("IfElseBlock.Default has unexpected type %T", x)
	}
	return nil
}

func _IfElseBlock_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*IfElseBlock)
	switch tag {
	case 3: // default.else_node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.Default = &IfElseBlock_ElseNode{msg}
		return true, err
	case 4: // default.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Default = &IfElseBlock_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _IfElseBlock_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*IfElseBlock)
	// default
	switch x := m.Default.(type) {
	case *IfElseBlock_ElseNode:
		s := proto.Size(x.ElseNode)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *IfElseBlock_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// BranchNode is a special node that alter the flow of the workflow graph. It allows the control flow to branch at
// runtime based on a series of conditions that get evaluated on various parameters (e.g. inputs, primtives).
type BranchNode struct {
	// +required
	IfElse               *IfElseBlock `protobuf:"bytes,2,opt,name=if_else,json=ifElse,proto3" json:"if_else,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BranchNode) Reset()         { *m = BranchNode{} }
func (m *BranchNode) String() string { return proto.CompactTextString(m) }
func (*BranchNode) ProtoMessage()    {}
func (*BranchNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{2}
}
func (m *BranchNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchNode.Unmarshal(m, b)
}
func (m *BranchNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchNode.Marshal(b, m, deterministic)
}
func (dst *BranchNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchNode.Merge(dst, src)
}
func (m *BranchNode) XXX_Size() int {
	return xxx_messageInfo_BranchNode.Size(m)
}
func (m *BranchNode) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchNode.DiscardUnknown(m)
}

var xxx_messageInfo_BranchNode proto.InternalMessageInfo

func (m *BranchNode) GetIfElse() *IfElseBlock {
	if m != nil {
		return m.IfElse
	}
	return nil
}

// Refers to the task that the Node is to execute.
type TaskNode struct {
	// Types that are valid to be assigned to Reference:
	//	*TaskNode_ReferenceId
	Reference            isTaskNode_Reference `protobuf_oneof:"reference"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskNode) Reset()         { *m = TaskNode{} }
func (m *TaskNode) String() string { return proto.CompactTextString(m) }
func (*TaskNode) ProtoMessage()    {}
func (*TaskNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{3}
}
func (m *TaskNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskNode.Unmarshal(m, b)
}
func (m *TaskNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskNode.Marshal(b, m, deterministic)
}
func (dst *TaskNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskNode.Merge(dst, src)
}
func (m *TaskNode) XXX_Size() int {
	return xxx_messageInfo_TaskNode.Size(m)
}
func (m *TaskNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskNode.DiscardUnknown(m)
}

var xxx_messageInfo_TaskNode proto.InternalMessageInfo

type isTaskNode_Reference interface {
	isTaskNode_Reference()
}

type TaskNode_ReferenceId struct {
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3,oneof"`
}

func (*TaskNode_ReferenceId) isTaskNode_Reference() {}

func (m *TaskNode) GetReference() isTaskNode_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *TaskNode) GetReferenceId() string {
	if x, ok := m.GetReference().(*TaskNode_ReferenceId); ok {
		return x.ReferenceId
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskNode) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskNode_OneofMarshaler, _TaskNode_OneofUnmarshaler, _TaskNode_OneofSizer, []interface{}{
		(*TaskNode_ReferenceId)(nil),
	}
}

func _TaskNode_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskNode)
	// reference
	switch x := m.Reference.(type) {
	case *TaskNode_ReferenceId:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ReferenceId)
	case nil:
	default:
		return fmt.Errorf("TaskNode.Reference has unexpected type %T", x)
	}
	return nil
}

func _TaskNode_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskNode)
	switch tag {
	case 1: // reference.reference_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Reference = &TaskNode_ReferenceId{x}
		return true, err
	default:
		return false, nil
	}
}

func _TaskNode_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskNode)
	// reference
	switch x := m.Reference.(type) {
	case *TaskNode_ReferenceId:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ReferenceId)))
		n += len(x.ReferenceId)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Refers to a the workflow the node is to execute.
type WorkflowNode struct {
	// Types that are valid to be assigned to Reference:
	//	*WorkflowNode_LaunchplanRef
	//	*WorkflowNode_SubWorkflowRef
	Reference            isWorkflowNode_Reference `protobuf_oneof:"reference"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WorkflowNode) Reset()         { *m = WorkflowNode{} }
func (m *WorkflowNode) String() string { return proto.CompactTextString(m) }
func (*WorkflowNode) ProtoMessage()    {}
func (*WorkflowNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{4}
}
func (m *WorkflowNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNode.Unmarshal(m, b)
}
func (m *WorkflowNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNode.Marshal(b, m, deterministic)
}
func (dst *WorkflowNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNode.Merge(dst, src)
}
func (m *WorkflowNode) XXX_Size() int {
	return xxx_messageInfo_WorkflowNode.Size(m)
}
func (m *WorkflowNode) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNode.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNode proto.InternalMessageInfo

type isWorkflowNode_Reference interface {
	isWorkflowNode_Reference()
}

type WorkflowNode_LaunchplanRef struct {
	LaunchplanRef string `protobuf:"bytes,1,opt,name=launchplan_ref,json=launchplanRef,proto3,oneof"`
}

type WorkflowNode_SubWorkflowRef struct {
	SubWorkflowRef string `protobuf:"bytes,2,opt,name=sub_workflow_ref,json=subWorkflowRef,proto3,oneof"`
}

func (*WorkflowNode_LaunchplanRef) isWorkflowNode_Reference() {}

func (*WorkflowNode_SubWorkflowRef) isWorkflowNode_Reference() {}

func (m *WorkflowNode) GetReference() isWorkflowNode_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *WorkflowNode) GetLaunchplanRef() string {
	if x, ok := m.GetReference().(*WorkflowNode_LaunchplanRef); ok {
		return x.LaunchplanRef
	}
	return ""
}

func (m *WorkflowNode) GetSubWorkflowRef() string {
	if x, ok := m.GetReference().(*WorkflowNode_SubWorkflowRef); ok {
		return x.SubWorkflowRef
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WorkflowNode) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WorkflowNode_OneofMarshaler, _WorkflowNode_OneofUnmarshaler, _WorkflowNode_OneofSizer, []interface{}{
		(*WorkflowNode_LaunchplanRef)(nil),
		(*WorkflowNode_SubWorkflowRef)(nil),
	}
}

func _WorkflowNode_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WorkflowNode)
	// reference
	switch x := m.Reference.(type) {
	case *WorkflowNode_LaunchplanRef:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.LaunchplanRef)
	case *WorkflowNode_SubWorkflowRef:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.SubWorkflowRef)
	case nil:
	default:
		return fmt.Errorf("WorkflowNode.Reference has unexpected type %T", x)
	}
	return nil
}

func _WorkflowNode_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WorkflowNode)
	switch tag {
	case 1: // reference.launchplan_ref
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Reference = &WorkflowNode_LaunchplanRef{x}
		return true, err
	case 2: // reference.sub_workflow_ref
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Reference = &WorkflowNode_SubWorkflowRef{x}
		return true, err
	default:
		return false, nil
	}
}

func _WorkflowNode_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WorkflowNode)
	// reference
	switch x := m.Reference.(type) {
	case *WorkflowNode_LaunchplanRef:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.LaunchplanRef)))
		n += len(x.LaunchplanRef)
	case *WorkflowNode_SubWorkflowRef:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.SubWorkflowRef)))
		n += len(x.SubWorkflowRef)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Defines extra information about the Node.
type NodeMetadata struct {
	// A friendly name for the Node
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The overall timeout of a task.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Number of retries per task.
	Retries              *RetryStrategy `protobuf:"bytes,5,opt,name=retries,proto3" json:"retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NodeMetadata) Reset()         { *m = NodeMetadata{} }
func (m *NodeMetadata) String() string { return proto.CompactTextString(m) }
func (*NodeMetadata) ProtoMessage()    {}
func (*NodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{5}
}
func (m *NodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMetadata.Unmarshal(m, b)
}
func (m *NodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMetadata.Marshal(b, m, deterministic)
}
func (dst *NodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMetadata.Merge(dst, src)
}
func (m *NodeMetadata) XXX_Size() int {
	return xxx_messageInfo_NodeMetadata.Size(m)
}
func (m *NodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMetadata proto.InternalMessageInfo

func (m *NodeMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeMetadata) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *NodeMetadata) GetRetries() *RetryStrategy {
	if m != nil {
		return m.Retries
	}
	return nil
}

// Links a variable to an alias.
type Alias struct {
	// Must match one of the output variable names on a node.
	Var string `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
	// A workflow-level unique alias that downstream nodes can refer to in their input.
	Alias                string   `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{6}
}
func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (dst *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(dst, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetVar() string {
	if m != nil {
		return m.Var
	}
	return ""
}

func (m *Alias) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

// A Workflow graph Node. One unit of execution in the graph. Each node can be linked to a Task, a Workflow or a branch
// node.
type Node struct {
	// A workflow-level unique identifier that identifies this node in the workflow. "inputs" and "outputs" are reserved
	// node ids that cannot be used by other nodes.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra metadata about the node.
	Metadata *NodeMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Specifies how to bind the underlying interface's inputs. All required inputs specified in the underlying interface
	// must be fullfilled.
	Inputs []*Binding `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// +optional Specifies execution depdendency for this node ensuring it will only get scheduled to run after all its
	// upstream nodes have completed. This node will have an implicit depdendency on any node that appears in inputs
	// field.
	UpstreamNodeIds []string `protobuf:"bytes,4,rep,name=upstream_node_ids,json=upstreamNodeIds,proto3" json:"upstream_node_ids,omitempty"`
	// +optional. A node can define aliases for a subset of its outputs. This is particularly useful if different nodes
	// need to conform to the same interface (e.g. all branches in a branch node). Downstream nodes must refer to this
	// nodes outputs using the alias if one's specified.
	OutputAliases []*Alias `protobuf:"bytes,5,rep,name=output_aliases,json=outputAliases,proto3" json:"output_aliases,omitempty"`
	// Information about the target to execute in this node.
	//
	// Types that are valid to be assigned to Target:
	//	*Node_TaskNode
	//	*Node_WorkflowNode
	//	*Node_BranchNode
	Target               isNode_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{7}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetMetadata() *NodeMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetInputs() []*Binding {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Node) GetUpstreamNodeIds() []string {
	if m != nil {
		return m.UpstreamNodeIds
	}
	return nil
}

func (m *Node) GetOutputAliases() []*Alias {
	if m != nil {
		return m.OutputAliases
	}
	return nil
}

type isNode_Target interface {
	isNode_Target()
}

type Node_TaskNode struct {
	TaskNode *TaskNode `protobuf:"bytes,6,opt,name=task_node,json=taskNode,proto3,oneof"`
}

type Node_WorkflowNode struct {
	WorkflowNode *WorkflowNode `protobuf:"bytes,7,opt,name=workflow_node,json=workflowNode,proto3,oneof"`
}

type Node_BranchNode struct {
	BranchNode *BranchNode `protobuf:"bytes,8,opt,name=branch_node,json=branchNode,proto3,oneof"`
}

func (*Node_TaskNode) isNode_Target() {}

func (*Node_WorkflowNode) isNode_Target() {}

func (*Node_BranchNode) isNode_Target() {}

func (m *Node) GetTarget() isNode_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Node) GetTaskNode() *TaskNode {
	if x, ok := m.GetTarget().(*Node_TaskNode); ok {
		return x.TaskNode
	}
	return nil
}

func (m *Node) GetWorkflowNode() *WorkflowNode {
	if x, ok := m.GetTarget().(*Node_WorkflowNode); ok {
		return x.WorkflowNode
	}
	return nil
}

func (m *Node) GetBranchNode() *BranchNode {
	if x, ok := m.GetTarget().(*Node_BranchNode); ok {
		return x.BranchNode
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Node) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Node_OneofMarshaler, _Node_OneofUnmarshaler, _Node_OneofSizer, []interface{}{
		(*Node_TaskNode)(nil),
		(*Node_WorkflowNode)(nil),
		(*Node_BranchNode)(nil),
	}
}

func _Node_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Node)
	// target
	switch x := m.Target.(type) {
	case *Node_TaskNode:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TaskNode); err != nil {
			return err
		}
	case *Node_WorkflowNode:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WorkflowNode); err != nil {
			return err
		}
	case *Node_BranchNode:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BranchNode); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Node.Target has unexpected type %T", x)
	}
	return nil
}

func _Node_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Node)
	switch tag {
	case 6: // target.task_node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaskNode)
		err := b.DecodeMessage(msg)
		m.Target = &Node_TaskNode{msg}
		return true, err
	case 7: // target.workflow_node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WorkflowNode)
		err := b.DecodeMessage(msg)
		m.Target = &Node_WorkflowNode{msg}
		return true, err
	case 8: // target.branch_node
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BranchNode)
		err := b.DecodeMessage(msg)
		m.Target = &Node_BranchNode{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Node_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Node)
	// target
	switch x := m.Target.(type) {
	case *Node_TaskNode:
		s := proto.Size(x.TaskNode)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Node_WorkflowNode:
		s := proto.Size(x.WorkflowNode)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Node_BranchNode:
		s := proto.Size(x.BranchNode)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Metadata for the entire workflow.
type WorkflowMetadata struct {
	ExecutionRole        string   `protobuf:"bytes,1,opt,name=execution_role,json=executionRole,proto3" json:"execution_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowMetadata) Reset()         { *m = WorkflowMetadata{} }
func (m *WorkflowMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowMetadata) ProtoMessage()    {}
func (*WorkflowMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{8}
}
func (m *WorkflowMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowMetadata.Unmarshal(m, b)
}
func (m *WorkflowMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowMetadata.Marshal(b, m, deterministic)
}
func (dst *WorkflowMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowMetadata.Merge(dst, src)
}
func (m *WorkflowMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowMetadata.Size(m)
}
func (m *WorkflowMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowMetadata proto.InternalMessageInfo

func (m *WorkflowMetadata) GetExecutionRole() string {
	if m != nil {
		return m.ExecutionRole
	}
	return ""
}

// Flyte Workflow Structure that encapsulates task, branch and subworkflow nodes to form a statically analyzable,
// directed acyclic graph.
type WorkflowTemplate struct {
	// This is an autogenerated id by the system. The id is globally unique across the system.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Extra metadata about the workflow.
	Metadata *WorkflowMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Defines a strongly typed interface for the Workflow. This can include some optional parameters.
	Interface *TypedInterface `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	// A list of nodes. In addition, "globals" is a special reserved node id that can be used to consume workflow inputs.
	Nodes []*Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// A list of output bindings that specify how to construct workflow outputs. Bindings can pull node outputs or
	// specify literals. All workflow outputs specified in the interface field must be bound in order for the workflow
	// to be validated. A workflow has an implicit depdendency on all of its nodes to execute successfully in order to
	// bind final outputs.
	Outputs []*Binding `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// +optional A catch-all node. This node is executed whenever the execution engine determines the workflow has failed.
	// The interface of this node must match the Workflow interface with an additional input named "error" of type
	// pb.lyft.flyte.core.Error.
	FailureNode          *Node    `protobuf:"bytes,6,opt,name=failure_node,json=failureNode,proto3" json:"failure_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowTemplate) Reset()         { *m = WorkflowTemplate{} }
func (m *WorkflowTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkflowTemplate) ProtoMessage()    {}
func (*WorkflowTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_44bb7d2878a19d29, []int{9}
}
func (m *WorkflowTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTemplate.Unmarshal(m, b)
}
func (m *WorkflowTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTemplate.Marshal(b, m, deterministic)
}
func (dst *WorkflowTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTemplate.Merge(dst, src)
}
func (m *WorkflowTemplate) XXX_Size() int {
	return xxx_messageInfo_WorkflowTemplate.Size(m)
}
func (m *WorkflowTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTemplate proto.InternalMessageInfo

func (m *WorkflowTemplate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WorkflowTemplate) GetMetadata() *WorkflowMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkflowTemplate) GetInterface() *TypedInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *WorkflowTemplate) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *WorkflowTemplate) GetOutputs() []*Binding {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *WorkflowTemplate) GetFailureNode() *Node {
	if m != nil {
		return m.FailureNode
	}
	return nil
}

func init() {
	proto.RegisterType((*IfBlock)(nil), "flyteidl.core.IfBlock")
	proto.RegisterType((*IfElseBlock)(nil), "flyteidl.core.IfElseBlock")
	proto.RegisterType((*BranchNode)(nil), "flyteidl.core.BranchNode")
	proto.RegisterType((*TaskNode)(nil), "flyteidl.core.TaskNode")
	proto.RegisterType((*WorkflowNode)(nil), "flyteidl.core.WorkflowNode")
	proto.RegisterType((*NodeMetadata)(nil), "flyteidl.core.NodeMetadata")
	proto.RegisterType((*Alias)(nil), "flyteidl.core.Alias")
	proto.RegisterType((*Node)(nil), "flyteidl.core.Node")
	proto.RegisterType((*WorkflowMetadata)(nil), "flyteidl.core.WorkflowMetadata")
	proto.RegisterType((*WorkflowTemplate)(nil), "flyteidl.core.WorkflowTemplate")
}

func init() {
	proto.RegisterFile("flyteidl/core/workflow.proto", fileDescriptor_workflow_44bb7d2878a19d29)
}

var fileDescriptor_workflow_44bb7d2878a19d29 = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x4e, 0xf3, 0xef, 0x93, 0x1f, 0xca, 0xb0, 0x82, 0x6c, 0xd9, 0x85, 0xca, 0x08, 0xb1, 0x54,
	0x60, 0x57, 0x5d, 0xa9, 0x08, 0x75, 0x85, 0xd4, 0x88, 0x4a, 0xc9, 0x05, 0x5c, 0x0c, 0x2b, 0x21,
	0x71, 0x13, 0x4d, 0xec, 0xe3, 0x64, 0xd4, 0x89, 0xc7, 0x1a, 0x8f, 0x69, 0x23, 0x1e, 0x82, 0x17,
	0xe1, 0x86, 0x77, 0xe1, 0x81, 0x90, 0xc7, 0x1e, 0xa7, 0x71, 0x93, 0xbd, 0x9b, 0xf1, 0xf9, 0xce,
	0xdf, 0x77, 0xbe, 0x33, 0x86, 0x57, 0x91, 0xd8, 0x6a, 0xe4, 0xa1, 0xf0, 0x03, 0xa9, 0xd0, 0x7f,
	0x90, 0xea, 0x3e, 0x12, 0xf2, 0xc1, 0x4b, 0x94, 0xd4, 0x92, 0x8c, 0xac, 0xd5, 0xcb, 0xad, 0x67,
	0xaf, 0xf7, 0xc1, 0x3c, 0xd6, 0xa8, 0x22, 0x16, 0x60, 0x81, 0x3e, 0xab, 0xc5, 0x12, 0x5c, 0xa3,
	0x62, 0x22, 0x2d, 0xad, 0x2f, 0xf7, 0xad, 0x7a, 0x9b, 0xa0, 0x35, 0xd5, 0xe2, 0x06, 0x32, 0x0e,
	0xb9, 0xe6, 0x32, 0x2e, 0xcd, 0x5f, 0xac, 0xa4, 0x5c, 0x09, 0xf4, 0xcd, 0x6d, 0x99, 0x45, 0x7e,
	0x98, 0x29, 0xb6, 0xb3, 0xbb, 0x7f, 0x41, 0x6f, 0x1e, 0x4d, 0x85, 0x0c, 0xee, 0xc9, 0x4f, 0xe0,
	0x54, 0xde, 0x93, 0x93, 0xf3, 0x93, 0x37, 0x83, 0xab, 0x73, 0x6f, 0xaf, 0x09, 0x6f, 0x2a, 0xa5,
	0x40, 0x16, 0xdf, 0x3d, 0x26, 0x0a, 0xd3, 0x94, 0xcb, 0x98, 0xee, 0x5c, 0xc8, 0x25, 0x38, 0x7a,
	0x8d, 0xf1, 0x22, 0x96, 0x21, 0x4e, 0x9a, 0xc6, 0xff, 0x93, 0x9a, 0xff, 0xaf, 0x32, 0x44, 0xda,
	0xcf, 0x51, 0xf9, 0xc9, 0xfd, 0xef, 0x04, 0x06, 0xf3, 0xe8, 0x4e, 0xa4, 0x58, 0x54, 0x70, 0x01,
	0xed, 0x80, 0xa5, 0x58, 0x26, 0xff, 0xb4, 0xe6, 0x5c, 0xd6, 0x49, 0x0d, 0x86, 0x7c, 0x07, 0x1d,
	0xa9, 0xd7, 0xa8, 0x26, 0xcd, 0xf3, 0xd6, 0x07, 0xc0, 0x05, 0x88, 0x5c, 0x81, 0x83, 0x22, 0xc5,
	0xa2, 0xb6, 0xd6, 0xd1, 0xda, 0x66, 0x0d, 0xda, 0xcf, 0x71, 0xf9, 0x39, 0xcf, 0x80, 0x4a, 0x49,
	0x35, 0x69, 0x1b, 0xfc, 0x8b, 0x1a, 0xfe, 0x2e, 0xb7, 0xcd, 0x1a, 0xb4, 0x00, 0x4d, 0x1d, 0xe8,
	0x85, 0x18, 0xb1, 0x4c, 0x68, 0xf7, 0x16, 0x60, 0xaa, 0x58, 0x1c, 0xac, 0x4d, 0x98, 0xb7, 0xd0,
	0xe3, 0xd1, 0x22, 0x8f, 0x5a, 0x92, 0x72, 0xf6, 0xac, 0xd4, 0x8a, 0x01, 0xda, 0xe5, 0xe6, 0xe2,
	0xbe, 0x83, 0xfe, 0x7b, 0x96, 0xde, 0x9b, 0x00, 0x5f, 0xc1, 0x50, 0x61, 0x84, 0x0a, 0xe3, 0x00,
	0x17, 0x3c, 0x34, 0xec, 0x38, 0xb3, 0x06, 0x1d, 0x54, 0x5f, 0xe7, 0xe1, 0x74, 0x00, 0x4e, 0x75,
	0x75, 0x13, 0x18, 0xfe, 0x5e, 0x8a, 0xd1, 0x44, 0xf8, 0x06, 0xc6, 0x82, 0x65, 0x71, 0xb0, 0x4e,
	0x04, 0x8b, 0x17, 0x0a, 0xa3, 0x2a, 0xc6, 0x68, 0xf7, 0x9d, 0x62, 0x44, 0x2e, 0xe0, 0x34, 0xcd,
	0x96, 0x0b, 0xab, 0x64, 0x03, 0x6d, 0x96, 0xd0, 0x71, 0x9a, 0x2d, 0x6d, 0x54, 0x8a, 0xd1, 0x7e,
	0xc6, 0xbf, 0x4f, 0x60, 0x98, 0xa7, 0xfa, 0x05, 0x35, 0x0b, 0x99, 0x66, 0x84, 0x40, 0x3b, 0x66,
	0x9b, 0x62, 0x94, 0x0e, 0x35, 0xe7, 0x9c, 0x09, 0xcd, 0x37, 0x28, 0x33, 0x5d, 0x52, 0xfa, 0xd2,
	0x2b, 0xd4, 0xe9, 0x59, 0x75, 0x7a, 0x3f, 0x97, 0xea, 0xa4, 0x16, 0x49, 0xae, 0xa1, 0xa7, 0x50,
	0x2b, 0x8e, 0xe9, 0xa4, 0x63, 0x9c, 0x5e, 0xd5, 0xe8, 0xa3, 0xa8, 0xd5, 0xf6, 0x37, 0xad, 0x98,
	0xc6, 0xd5, 0x96, 0x5a, 0xb0, 0xeb, 0x43, 0xe7, 0x56, 0x70, 0x96, 0x92, 0x53, 0x68, 0xfd, 0xc9,
	0x54, 0x59, 0x48, 0x7e, 0x24, 0x2f, 0xa0, 0xc3, 0x72, 0x53, 0xd1, 0x1a, 0x2d, 0x2e, 0xee, 0x3f,
	0x2d, 0x68, 0x1b, 0xb6, 0xc6, 0xd0, 0xb4, 0x2c, 0xd3, 0x26, 0x0f, 0xc9, 0x0f, 0xd0, 0xdf, 0x94,
	0x6d, 0x95, 0x13, 0xfc, 0xfc, 0x80, 0x74, 0x6c, 0xe7, 0xb4, 0x02, 0x13, 0x0f, 0xba, 0x3c, 0x4e,
	0x32, 0x9d, 0x4e, 0x5a, 0x07, 0x35, 0x3a, 0xe5, 0x71, 0xc8, 0xe3, 0x15, 0x2d, 0x51, 0xe4, 0x02,
	0x3e, 0xce, 0x92, 0x54, 0x2b, 0x64, 0x1b, 0x23, 0xd4, 0x05, 0x0f, 0xd3, 0x49, 0xfb, 0xbc, 0xf5,
	0xc6, 0xa1, 0x1f, 0x59, 0x43, 0x9e, 0x6a, 0x1e, 0xa6, 0xe4, 0x06, 0xc6, 0x32, 0xd3, 0x49, 0xa6,
	0x17, 0xa6, 0x7a, 0xc3, 0x4e, 0xeb, 0x80, 0x4a, 0x0d, 0x07, 0x74, 0x54, 0x60, 0x6f, 0x0b, 0x28,
	0xb9, 0x06, 0x47, 0xb3, 0xf4, 0xbe, 0xd8, 0x86, 0xae, 0x69, 0xe9, 0xb3, 0x9a, 0x9f, 0x55, 0x5f,
	0xbe, 0x11, 0xda, 0x2a, 0x71, 0x0a, 0xa3, 0x4a, 0x1a, 0xc6, 0xb7, 0x77, 0x90, 0x8e, 0xa7, 0xda,
	0x9b, 0x35, 0xe8, 0xf0, 0xe1, 0xa9, 0x16, 0xdf, 0xc1, 0x60, 0x69, 0x96, 0xa3, 0x88, 0xd0, 0x2f,
	0x85, 0x50, 0x63, 0xa6, 0x5a, 0x9f, 0x59, 0x83, 0xc2, 0xb2, 0xba, 0x4d, 0xfb, 0xd0, 0xd5, 0x4c,
	0xad, 0x50, 0xbb, 0x3f, 0xc2, 0xa9, 0xcd, 0x53, 0x89, 0xee, 0x6b, 0x18, 0xe3, 0x23, 0x06, 0x59,
	0xae, 0xa0, 0x85, 0x92, 0xc2, 0xca, 0x6f, 0x54, 0x7d, 0xa5, 0x52, 0xa0, 0xfb, 0x6f, 0x73, 0xe7,
	0xfb, 0x1e, 0x37, 0x89, 0x60, 0xfa, 0xf9, 0xd4, 0x6f, 0x9e, 0x4d, 0xfd, 0xcb, 0x23, 0x6d, 0x1e,
	0x98, 0xfc, 0x0d, 0x38, 0xd5, 0x03, 0x5f, 0x3e, 0x37, 0xaf, 0xeb, 0x04, 0x6f, 0x13, 0x0c, 0xe7,
	0x16, 0x44, 0x77, 0x78, 0xf2, 0x2d, 0x74, 0x72, 0x6a, 0x8a, 0xd1, 0x1f, 0x79, 0x43, 0x0b, 0x04,
	0xb9, 0x84, 0x5e, 0x31, 0x59, 0x3b, 0xfe, 0x63, 0x12, 0xb3, 0x30, 0x72, 0x0d, 0xc3, 0x88, 0x71,
	0x91, 0x29, 0x7c, 0x3a, 0xfd, 0x83, 0x39, 0x06, 0x25, 0xd0, 0x10, 0x7f, 0xf5, 0xc7, 0xe5, 0x8a,
	0xeb, 0x75, 0xb6, 0xf4, 0x02, 0xb9, 0xf1, 0xc5, 0x36, 0xd2, 0x7e, 0xf5, 0xe3, 0x59, 0x61, 0xec,
	0x27, 0xcb, 0xef, 0x57, 0xd2, 0xdf, 0xfb, 0x17, 0x2d, 0xbb, 0x66, 0xad, 0xdf, 0xfe, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x45, 0x78, 0xfe, 0x45, 0x28, 0x07, 0x00, 0x00,
}

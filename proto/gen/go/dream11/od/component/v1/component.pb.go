// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.28.3
// source: dream11/od/component/v1/component.proto

package v1

import (
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListComponentTypeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Params        map[string]string      `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListComponentTypeRequest) Reset() {
	*x = ListComponentTypeRequest{}
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComponentTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComponentTypeRequest) ProtoMessage() {}

func (x *ListComponentTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComponentTypeRequest.ProtoReflect.Descriptor instead.
func (*ListComponentTypeRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_component_v1_component_proto_rawDescGZIP(), []int{0}
}

func (x *ListComponentTypeRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListComponentTypeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Components    []*v1.Component        `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListComponentTypeResponse) Reset() {
	*x = ListComponentTypeResponse{}
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListComponentTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComponentTypeResponse) ProtoMessage() {}

func (x *ListComponentTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComponentTypeResponse.ProtoReflect.Descriptor instead.
func (*ListComponentTypeResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_component_v1_component_proto_rawDescGZIP(), []int{1}
}

func (x *ListComponentTypeResponse) GetComponents() []*v1.Component {
	if x != nil {
		return x.Components
	}
	return nil
}

type DescribeComponentTypeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ComponentType string                 `protobuf:"bytes,1,opt,name=component_type,json=componentType,proto3" json:"component_type,omitempty"`
	Params        map[string]string      `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeComponentTypeRequest) Reset() {
	*x = DescribeComponentTypeRequest{}
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeComponentTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeComponentTypeRequest) ProtoMessage() {}

func (x *DescribeComponentTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeComponentTypeRequest.ProtoReflect.Descriptor instead.
func (*DescribeComponentTypeRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_component_v1_component_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeComponentTypeRequest) GetComponentType() string {
	if x != nil {
		return x.ComponentType
	}
	return ""
}

func (x *DescribeComponentTypeRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type DescribeComponentTypeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Component     *v1.Component          `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescribeComponentTypeResponse) Reset() {
	*x = DescribeComponentTypeResponse{}
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeComponentTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeComponentTypeResponse) ProtoMessage() {}

func (x *DescribeComponentTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_component_v1_component_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeComponentTypeResponse.ProtoReflect.Descriptor instead.
func (*DescribeComponentTypeResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_component_v1_component_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeComponentTypeResponse) GetComponent() *v1.Component {
	if x != nil {
		return x.Component
	}
	return nil
}

var File_dream11_od_component_v1_component_proto protoreflect.FileDescriptor

var file_dream11_od_component_v1_component_proto_rawDesc = []byte{
	0x0a, 0x27, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x72, 0x65, 0x61, 0x6d,
	0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x55, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e,
	0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xdb, 0x01, 0x0a, 0x1c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a,
	0x1d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x32, 0x9b, 0x02, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x7c, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x88, 0x01,
	0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f,
	0x64, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_component_v1_component_proto_rawDescOnce sync.Once
	file_dream11_od_component_v1_component_proto_rawDescData = file_dream11_od_component_v1_component_proto_rawDesc
)

func file_dream11_od_component_v1_component_proto_rawDescGZIP() []byte {
	file_dream11_od_component_v1_component_proto_rawDescOnce.Do(func() {
		file_dream11_od_component_v1_component_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_component_v1_component_proto_rawDescData)
	})
	return file_dream11_od_component_v1_component_proto_rawDescData
}

var file_dream11_od_component_v1_component_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dream11_od_component_v1_component_proto_goTypes = []any{
	(*ListComponentTypeRequest)(nil),      // 0: dream11.od.component.v1.ListComponentTypeRequest
	(*ListComponentTypeResponse)(nil),     // 1: dream11.od.component.v1.ListComponentTypeResponse
	(*DescribeComponentTypeRequest)(nil),  // 2: dream11.od.component.v1.DescribeComponentTypeRequest
	(*DescribeComponentTypeResponse)(nil), // 3: dream11.od.component.v1.DescribeComponentTypeResponse
	nil,                                   // 4: dream11.od.component.v1.ListComponentTypeRequest.ParamsEntry
	nil,                                   // 5: dream11.od.component.v1.DescribeComponentTypeRequest.ParamsEntry
	(*v1.Component)(nil),                  // 6: dream11.od.dto.v1.Component
}
var file_dream11_od_component_v1_component_proto_depIdxs = []int32{
	4, // 0: dream11.od.component.v1.ListComponentTypeRequest.params:type_name -> dream11.od.component.v1.ListComponentTypeRequest.ParamsEntry
	6, // 1: dream11.od.component.v1.ListComponentTypeResponse.components:type_name -> dream11.od.dto.v1.Component
	5, // 2: dream11.od.component.v1.DescribeComponentTypeRequest.params:type_name -> dream11.od.component.v1.DescribeComponentTypeRequest.ParamsEntry
	6, // 3: dream11.od.component.v1.DescribeComponentTypeResponse.component:type_name -> dream11.od.dto.v1.Component
	0, // 4: dream11.od.component.v1.ComponentService.ListComponentType:input_type -> dream11.od.component.v1.ListComponentTypeRequest
	2, // 5: dream11.od.component.v1.ComponentService.DescribeComponentType:input_type -> dream11.od.component.v1.DescribeComponentTypeRequest
	1, // 6: dream11.od.component.v1.ComponentService.ListComponentType:output_type -> dream11.od.component.v1.ListComponentTypeResponse
	3, // 7: dream11.od.component.v1.ComponentService.DescribeComponentType:output_type -> dream11.od.component.v1.DescribeComponentTypeResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dream11_od_component_v1_component_proto_init() }
func file_dream11_od_component_v1_component_proto_init() {
	if File_dream11_od_component_v1_component_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_component_v1_component_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dream11_od_component_v1_component_proto_goTypes,
		DependencyIndexes: file_dream11_od_component_v1_component_proto_depIdxs,
		MessageInfos:      file_dream11_od_component_v1_component_proto_msgTypes,
	}.Build()
	File_dream11_od_component_v1_component_proto = out.File
	file_dream11_od_component_v1_component_proto_rawDesc = nil
	file_dream11_od_component_v1_component_proto_goTypes = nil
	file_dream11_od_component_v1_component_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: dream11/od/dto/v1/component.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               *int64                 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	ComponentType    *string                `protobuf:"bytes,6,opt,name=component_type,json=componentType,proto3,oneof" json:"component_type,omitempty"`
	ComponentVersion *string                `protobuf:"bytes,7,opt,name=component_version,json=componentVersion,proto3,oneof" json:"component_version,omitempty"`
	CommonSchema     *structpb.Struct       `protobuf:"bytes,8,opt,name=common_schema,json=commonSchema,proto3,oneof" json:"common_schema,omitempty"`
	CommonDefaults   *structpb.Struct       `protobuf:"bytes,9,opt,name=common_defaults,json=commonDefaults,proto3,oneof" json:"common_defaults,omitempty"`
	Flavours         []*Flavour             `protobuf:"bytes,10,rep,name=flavours,proto3" json:"flavours,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_component_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_component_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_component_proto_rawDescGZIP(), []int{0}
}

func (x *Component) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Component) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Component) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Component) GetComponentType() string {
	if x != nil && x.ComponentType != nil {
		return *x.ComponentType
	}
	return ""
}

func (x *Component) GetComponentVersion() string {
	if x != nil && x.ComponentVersion != nil {
		return *x.ComponentVersion
	}
	return ""
}

func (x *Component) GetCommonSchema() *structpb.Struct {
	if x != nil {
		return x.CommonSchema
	}
	return nil
}

func (x *Component) GetCommonDefaults() *structpb.Struct {
	if x != nil {
		return x.CommonDefaults
	}
	return nil
}

func (x *Component) GetFlavours() []*Flavour {
	if x != nil {
		return x.Flavours
	}
	return nil
}

type Flavour struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *string          `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Schema     *structpb.Struct `protobuf:"bytes,2,opt,name=schema,proto3,oneof" json:"schema,omitempty"`
	Defaults   *structpb.Struct `protobuf:"bytes,3,opt,name=defaults,proto3,oneof" json:"defaults,omitempty"`
	Operations []*Operation     `protobuf:"bytes,4,rep,name=operations,proto3" json:"operations,omitempty"`
}

func (x *Flavour) Reset() {
	*x = Flavour{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_component_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flavour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flavour) ProtoMessage() {}

func (x *Flavour) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_component_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flavour.ProtoReflect.Descriptor instead.
func (*Flavour) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_component_proto_rawDescGZIP(), []int{1}
}

func (x *Flavour) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Flavour) GetSchema() *structpb.Struct {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *Flavour) GetDefaults() *structpb.Struct {
	if x != nil {
		return x.Defaults
	}
	return nil
}

func (x *Flavour) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string          `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Schema   *structpb.Struct `protobuf:"bytes,2,opt,name=schema,proto3,oneof" json:"schema,omitempty"`
	Defaults *structpb.Struct `protobuf:"bytes,3,opt,name=defaults,proto3,oneof" json:"defaults,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_component_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_component_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_component_proto_rawDescGZIP(), []int{2}
}

func (x *Operation) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Operation) GetSchema() *structpb.Struct {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *Operation) GetDefaults() *structpb.Struct {
	if x != nil {
		return x.Defaults
	}
	return nil
}

var File_dream11_od_dto_v1_component_proto protoreflect.FileDescriptor

var file_dream11_od_dto_v1_component_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e,
	0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x04, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x06, 0x52, 0x0e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x36, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e,
	0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x52, 0x08,
	0x66, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xf1, 0x01, 0x0a,
	0x07, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x01, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x48, 0x02, 0x52, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e,
	0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0xb5, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x48, 0x01, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a,
	0x08, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x02, 0x52, 0x08, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x0b, 0x0a, 0x09, 0x5f,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f,
	0x64, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_dto_v1_component_proto_rawDescOnce sync.Once
	file_dream11_od_dto_v1_component_proto_rawDescData = file_dream11_od_dto_v1_component_proto_rawDesc
)

func file_dream11_od_dto_v1_component_proto_rawDescGZIP() []byte {
	file_dream11_od_dto_v1_component_proto_rawDescOnce.Do(func() {
		file_dream11_od_dto_v1_component_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_dto_v1_component_proto_rawDescData)
	})
	return file_dream11_od_dto_v1_component_proto_rawDescData
}

var file_dream11_od_dto_v1_component_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dream11_od_dto_v1_component_proto_goTypes = []interface{}{
	(*Component)(nil),             // 0: dream11.od.dto.v1.Component
	(*Flavour)(nil),               // 1: dream11.od.dto.v1.Flavour
	(*Operation)(nil),             // 2: dream11.od.dto.v1.Operation
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 4: google.protobuf.Struct
}
var file_dream11_od_dto_v1_component_proto_depIdxs = []int32{
	3,  // 0: dream11.od.dto.v1.Component.created_at:type_name -> google.protobuf.Timestamp
	3,  // 1: dream11.od.dto.v1.Component.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 2: dream11.od.dto.v1.Component.common_schema:type_name -> google.protobuf.Struct
	4,  // 3: dream11.od.dto.v1.Component.common_defaults:type_name -> google.protobuf.Struct
	1,  // 4: dream11.od.dto.v1.Component.flavours:type_name -> dream11.od.dto.v1.Flavour
	4,  // 5: dream11.od.dto.v1.Flavour.schema:type_name -> google.protobuf.Struct
	4,  // 6: dream11.od.dto.v1.Flavour.defaults:type_name -> google.protobuf.Struct
	2,  // 7: dream11.od.dto.v1.Flavour.operations:type_name -> dream11.od.dto.v1.Operation
	4,  // 8: dream11.od.dto.v1.Operation.schema:type_name -> google.protobuf.Struct
	4,  // 9: dream11.od.dto.v1.Operation.defaults:type_name -> google.protobuf.Struct
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_dream11_od_dto_v1_component_proto_init() }
func file_dream11_od_dto_v1_component_proto_init() {
	if File_dream11_od_dto_v1_component_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dream11_od_dto_v1_component_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Component); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dream11_od_dto_v1_component_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flavour); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dream11_od_dto_v1_component_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_dream11_od_dto_v1_component_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_dream11_od_dto_v1_component_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_dream11_od_dto_v1_component_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_dto_v1_component_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dream11_od_dto_v1_component_proto_goTypes,
		DependencyIndexes: file_dream11_od_dto_v1_component_proto_depIdxs,
		MessageInfos:      file_dream11_od_dto_v1_component_proto_msgTypes,
	}.Build()
	File_dream11_od_dto_v1_component_proto = out.File
	file_dream11_od_dto_v1_component_proto_rawDesc = nil
	file_dream11_od_dto_v1_component_proto_goTypes = nil
	file_dream11_od_dto_v1_component_proto_depIdxs = nil
}

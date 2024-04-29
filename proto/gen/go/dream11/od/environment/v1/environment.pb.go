// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: dream11/od/environment/v1/environment.proto

package v1

import (
	v1 "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListEnvironmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params map[string]string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListEnvironmentRequest) Reset() {
	*x = ListEnvironmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvironmentRequest) ProtoMessage() {}

func (x *ListEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*ListEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{0}
}

func (x *ListEnvironmentRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListEnvironmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environments []*v1.Environment `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *ListEnvironmentResponse) Reset() {
	*x = ListEnvironmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvironmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvironmentResponse) ProtoMessage() {}

func (x *ListEnvironmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvironmentResponse.ProtoReflect.Descriptor instead.
func (*ListEnvironmentResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{1}
}

func (x *ListEnvironmentResponse) GetEnvironments() []*v1.Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

type DescribeEnvironmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvName string            `protobuf:"bytes,1,opt,name=env_name,json=envName,proto3" json:"env_name,omitempty"`
	Params  map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DescribeEnvironmentRequest) Reset() {
	*x = DescribeEnvironmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEnvironmentRequest) ProtoMessage() {}

func (x *DescribeEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*DescribeEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeEnvironmentRequest) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

func (x *DescribeEnvironmentRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type DescribeEnvironmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environment *v1.Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *DescribeEnvironmentResponse) Reset() {
	*x = DescribeEnvironmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeEnvironmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeEnvironmentResponse) ProtoMessage() {}

func (x *DescribeEnvironmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeEnvironmentResponse.ProtoReflect.Descriptor instead.
func (*DescribeEnvironmentResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeEnvironmentResponse) GetEnvironment() *v1.Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

type UpdateEnvironmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvName string           `protobuf:"bytes,1,opt,name=env_name,json=envName,proto3" json:"env_name,omitempty"`
	Data    *structpb.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateEnvironmentRequest) Reset() {
	*x = UpdateEnvironmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEnvironmentRequest) ProtoMessage() {}

func (x *UpdateEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEnvironmentRequest) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

func (x *UpdateEnvironmentRequest) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateEnvironmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEnvironmentResponse) Reset() {
	*x = UpdateEnvironmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEnvironmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEnvironmentResponse) ProtoMessage() {}

func (x *UpdateEnvironmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEnvironmentResponse.ProtoReflect.Descriptor instead.
func (*UpdateEnvironmentResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{5}
}

type CreateEnvironmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvName          string   `protobuf:"bytes,1,opt,name=env_name,json=envName,proto3" json:"env_name,omitempty"`
	Accounts         []string `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	ProvisioningType string   `protobuf:"bytes,3,opt,name=provisioning_type,json=provisioningType,proto3" json:"provisioning_type,omitempty"`
}

func (x *CreateEnvironmentRequest) Reset() {
	*x = CreateEnvironmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnvironmentRequest) ProtoMessage() {}

func (x *CreateEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*CreateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{6}
}

func (x *CreateEnvironmentRequest) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetAccounts() []string {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *CreateEnvironmentRequest) GetProvisioningType() string {
	if x != nil {
		return x.ProvisioningType
	}
	return ""
}

type CreateEnvironmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateEnvironmentResponse) Reset() {
	*x = CreateEnvironmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnvironmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnvironmentResponse) ProtoMessage() {}

func (x *CreateEnvironmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnvironmentResponse.ProtoReflect.Descriptor instead.
func (*CreateEnvironmentResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{7}
}

func (x *CreateEnvironmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteEnvironmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnvName string `protobuf:"bytes,1,opt,name=env_name,json=envName,proto3" json:"env_name,omitempty"`
}

func (x *DeleteEnvironmentRequest) Reset() {
	*x = DeleteEnvironmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEnvironmentRequest) ProtoMessage() {}

func (x *DeleteEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEnvironmentRequest) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

type DeleteEnvironmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteEnvironmentResponse) Reset() {
	*x = DeleteEnvironmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEnvironmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEnvironmentResponse) ProtoMessage() {}

func (x *DeleteEnvironmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_environment_v1_environment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEnvironmentResponse.ProtoReflect.Descriptor instead.
func (*DeleteEnvironmentResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_environment_v1_environment_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEnvironmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_dream11_od_environment_v1_environment_proto protoreflect.FileDescriptor

var file_dream11_od_environment_v1_environment_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64,
	0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31,
	0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x72, 0x65, 0x61,
	0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xcd, 0x01, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x59, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x41, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x1b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x72,
	0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7e, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x35, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa6,
	0x05, 0x0a, 0x12, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d,
	0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x64, 0x72,
	0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x2e, 0x64, 0x72, 0x65, 0x61,
	0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x33, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e,
	0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64,
	0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d,
	0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d,
	0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e,
	0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64,
	0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dream11_od_environment_v1_environment_proto_rawDescOnce sync.Once
	file_dream11_od_environment_v1_environment_proto_rawDescData = file_dream11_od_environment_v1_environment_proto_rawDesc
)

func file_dream11_od_environment_v1_environment_proto_rawDescGZIP() []byte {
	file_dream11_od_environment_v1_environment_proto_rawDescOnce.Do(func() {
		file_dream11_od_environment_v1_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_environment_v1_environment_proto_rawDescData)
	})
	return file_dream11_od_environment_v1_environment_proto_rawDescData
}

var file_dream11_od_environment_v1_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_dream11_od_environment_v1_environment_proto_goTypes = []interface{}{
	(*ListEnvironmentRequest)(nil),      // 0: dream11.od.environment.v1.ListEnvironmentRequest
	(*ListEnvironmentResponse)(nil),     // 1: dream11.od.environment.v1.ListEnvironmentResponse
	(*DescribeEnvironmentRequest)(nil),  // 2: dream11.od.environment.v1.DescribeEnvironmentRequest
	(*DescribeEnvironmentResponse)(nil), // 3: dream11.od.environment.v1.DescribeEnvironmentResponse
	(*UpdateEnvironmentRequest)(nil),    // 4: dream11.od.environment.v1.UpdateEnvironmentRequest
	(*UpdateEnvironmentResponse)(nil),   // 5: dream11.od.environment.v1.UpdateEnvironmentResponse
	(*CreateEnvironmentRequest)(nil),    // 6: dream11.od.environment.v1.CreateEnvironmentRequest
	(*CreateEnvironmentResponse)(nil),   // 7: dream11.od.environment.v1.CreateEnvironmentResponse
	(*DeleteEnvironmentRequest)(nil),    // 8: dream11.od.environment.v1.DeleteEnvironmentRequest
	(*DeleteEnvironmentResponse)(nil),   // 9: dream11.od.environment.v1.DeleteEnvironmentResponse
	nil,                                 // 10: dream11.od.environment.v1.ListEnvironmentRequest.ParamsEntry
	nil,                                 // 11: dream11.od.environment.v1.DescribeEnvironmentRequest.ParamsEntry
	(*v1.Environment)(nil),              // 12: dream11.od.dto.v1.Environment
	(*structpb.Struct)(nil),             // 13: google.protobuf.Struct
}
var file_dream11_od_environment_v1_environment_proto_depIdxs = []int32{
	10, // 0: dream11.od.environment.v1.ListEnvironmentRequest.params:type_name -> dream11.od.environment.v1.ListEnvironmentRequest.ParamsEntry
	12, // 1: dream11.od.environment.v1.ListEnvironmentResponse.environments:type_name -> dream11.od.dto.v1.Environment
	11, // 2: dream11.od.environment.v1.DescribeEnvironmentRequest.params:type_name -> dream11.od.environment.v1.DescribeEnvironmentRequest.ParamsEntry
	12, // 3: dream11.od.environment.v1.DescribeEnvironmentResponse.environment:type_name -> dream11.od.dto.v1.Environment
	13, // 4: dream11.od.environment.v1.UpdateEnvironmentRequest.data:type_name -> google.protobuf.Struct
	0,  // 5: dream11.od.environment.v1.EnvironmentService.ListEnvironment:input_type -> dream11.od.environment.v1.ListEnvironmentRequest
	2,  // 6: dream11.od.environment.v1.EnvironmentService.DescribeEnvironment:input_type -> dream11.od.environment.v1.DescribeEnvironmentRequest
	4,  // 7: dream11.od.environment.v1.EnvironmentService.UpdateEnvironment:input_type -> dream11.od.environment.v1.UpdateEnvironmentRequest
	6,  // 8: dream11.od.environment.v1.EnvironmentService.CreateEnvironment:input_type -> dream11.od.environment.v1.CreateEnvironmentRequest
	8,  // 9: dream11.od.environment.v1.EnvironmentService.DeleteEnvironment:input_type -> dream11.od.environment.v1.DeleteEnvironmentRequest
	1,  // 10: dream11.od.environment.v1.EnvironmentService.ListEnvironment:output_type -> dream11.od.environment.v1.ListEnvironmentResponse
	3,  // 11: dream11.od.environment.v1.EnvironmentService.DescribeEnvironment:output_type -> dream11.od.environment.v1.DescribeEnvironmentResponse
	5,  // 12: dream11.od.environment.v1.EnvironmentService.UpdateEnvironment:output_type -> dream11.od.environment.v1.UpdateEnvironmentResponse
	7,  // 13: dream11.od.environment.v1.EnvironmentService.CreateEnvironment:output_type -> dream11.od.environment.v1.CreateEnvironmentResponse
	9,  // 14: dream11.od.environment.v1.EnvironmentService.DeleteEnvironment:output_type -> dream11.od.environment.v1.DeleteEnvironmentResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_dream11_od_environment_v1_environment_proto_init() }
func file_dream11_od_environment_v1_environment_proto_init() {
	if File_dream11_od_environment_v1_environment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dream11_od_environment_v1_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvironmentRequest); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvironmentResponse); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeEnvironmentRequest); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeEnvironmentResponse); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEnvironmentRequest); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEnvironmentResponse); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnvironmentRequest); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnvironmentResponse); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEnvironmentRequest); i {
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
		file_dream11_od_environment_v1_environment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEnvironmentResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_environment_v1_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dream11_od_environment_v1_environment_proto_goTypes,
		DependencyIndexes: file_dream11_od_environment_v1_environment_proto_depIdxs,
		MessageInfos:      file_dream11_od_environment_v1_environment_proto_msgTypes,
	}.Build()
	File_dream11_od_environment_v1_environment_proto = out.File
	file_dream11_od_environment_v1_environment_proto_rawDesc = nil
	file_dream11_od_environment_v1_environment_proto_goTypes = nil
	file_dream11_od_environment_v1_environment_proto_depIdxs = nil
}

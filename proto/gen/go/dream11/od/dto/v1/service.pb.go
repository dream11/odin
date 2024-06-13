// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: dream11/od/dto/v1/service.proto

package v1

import (
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

type ComponentDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version   string           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Type      string           `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	DependsOn []string         `protobuf:"bytes,4,rep,name=depends_on,json=dependsOn,proto3" json:"depends_on,omitempty"`
	Config    *structpb.Struct `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	Metadata  *structpb.Struct `protobuf:"bytes,6,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
}

func (x *ComponentDefinition) Reset() {
	*x = ComponentDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentDefinition) ProtoMessage() {}

func (x *ComponentDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentDefinition.ProtoReflect.Descriptor instead.
func (*ComponentDefinition) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ComponentDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComponentDefinition) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ComponentDefinition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ComponentDefinition) GetDependsOn() []string {
	if x != nil {
		return x.DependsOn
	}
	return nil
}

func (x *ComponentDefinition) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ComponentDefinition) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ServiceDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version    string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Team       string                 `protobuf:"bytes,3,opt,name=team,proto3" json:"team,omitempty"`
	Components []*ComponentDefinition `protobuf:"bytes,4,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *ServiceDefinition) Reset() {
	*x = ServiceDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceDefinition) ProtoMessage() {}

func (x *ServiceDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceDefinition.ProtoReflect.Descriptor instead.
func (*ServiceDefinition) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceDefinition) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceDefinition) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *ServiceDefinition) GetComponents() []*ComponentDefinition {
	if x != nil {
		return x.Components
	}
	return nil
}

type ComponentProvisioningConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentName  string           `protobuf:"bytes,1,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	DeploymentType string           `protobuf:"bytes,2,opt,name=deployment_type,json=deploymentType,proto3" json:"deployment_type,omitempty"`
	Params         *structpb.Struct `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *ComponentProvisioningConfig) Reset() {
	*x = ComponentProvisioningConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentProvisioningConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentProvisioningConfig) ProtoMessage() {}

func (x *ComponentProvisioningConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentProvisioningConfig.ProtoReflect.Descriptor instead.
func (*ComponentProvisioningConfig) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentProvisioningConfig) GetComponentName() string {
	if x != nil {
		return x.ComponentName
	}
	return ""
}

func (x *ComponentProvisioningConfig) GetDeploymentType() string {
	if x != nil {
		return x.DeploymentType
	}
	return ""
}

func (x *ComponentProvisioningConfig) GetParams() *structpb.Struct {
	if x != nil {
		return x.Params
	}
	return nil
}

type ProvisioningConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentProvisioningConfig []*ComponentProvisioningConfig `protobuf:"bytes,1,rep,name=component_provisioning_config,json=componentProvisioningConfig,proto3" json:"component_provisioning_config,omitempty"`
}

func (x *ProvisioningConfig) Reset() {
	*x = ProvisioningConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisioningConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisioningConfig) ProtoMessage() {}

func (x *ProvisioningConfig) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisioningConfig.ProtoReflect.Descriptor instead.
func (*ProvisioningConfig) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *ProvisioningConfig) GetComponentProvisioningConfig() []*ComponentProvisioningConfig {
	if x != nil {
		return x.ComponentProvisioningConfig
	}
	return nil
}

type ServiceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version     string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedBy   string `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy   string `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tags        string `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`      // JSON encoded string
	Labels      string `protobuf:"bytes,10,opt,name=labels,proto3" json:"labels,omitempty"` // comma seperated string
}

func (x *ServiceMetadata) Reset() {
	*x = ServiceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMetadata) ProtoMessage() {}

func (x *ServiceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMetadata.ProtoReflect.Descriptor instead.
func (*ServiceMetadata) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceMetadata) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceMetadata) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *ServiceMetadata) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *ServiceMetadata) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ServiceMetadata) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ServiceMetadata) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *ServiceMetadata) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

var File_dream11_od_dto_v1_service_proto protoreflect.FileDescriptor

var file_dream11_od_dto_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x9d, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x46, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x72, 0x0a, 0x1d, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e,
	0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x1b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x99, 0x02, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2f, 0x6f, 0x64, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_dto_v1_service_proto_rawDescOnce sync.Once
	file_dream11_od_dto_v1_service_proto_rawDescData = file_dream11_od_dto_v1_service_proto_rawDesc
)

func file_dream11_od_dto_v1_service_proto_rawDescGZIP() []byte {
	file_dream11_od_dto_v1_service_proto_rawDescOnce.Do(func() {
		file_dream11_od_dto_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_dto_v1_service_proto_rawDescData)
	})
	return file_dream11_od_dto_v1_service_proto_rawDescData
}

var file_dream11_od_dto_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dream11_od_dto_v1_service_proto_goTypes = []interface{}{
	(*ComponentDefinition)(nil),         // 0: dream11.od.dto.v1.ComponentDefinition
	(*ServiceDefinition)(nil),           // 1: dream11.od.dto.v1.ServiceDefinition
	(*ComponentProvisioningConfig)(nil), // 2: dream11.od.dto.v1.ComponentProvisioningConfig
	(*ProvisioningConfig)(nil),          // 3: dream11.od.dto.v1.ProvisioningConfig
	(*ServiceMetadata)(nil),             // 4: dream11.od.dto.v1.ServiceMetadata
	(*structpb.Struct)(nil),             // 5: google.protobuf.Struct
}
var file_dream11_od_dto_v1_service_proto_depIdxs = []int32{
	5, // 0: dream11.od.dto.v1.ComponentDefinition.config:type_name -> google.protobuf.Struct
	5, // 1: dream11.od.dto.v1.ComponentDefinition.metadata:type_name -> google.protobuf.Struct
	0, // 2: dream11.od.dto.v1.ServiceDefinition.components:type_name -> dream11.od.dto.v1.ComponentDefinition
	5, // 3: dream11.od.dto.v1.ComponentProvisioningConfig.params:type_name -> google.protobuf.Struct
	2, // 4: dream11.od.dto.v1.ProvisioningConfig.component_provisioning_config:type_name -> dream11.od.dto.v1.ComponentProvisioningConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_dream11_od_dto_v1_service_proto_init() }
func file_dream11_od_dto_v1_service_proto_init() {
	if File_dream11_od_dto_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dream11_od_dto_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentDefinition); i {
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
		file_dream11_od_dto_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceDefinition); i {
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
		file_dream11_od_dto_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentProvisioningConfig); i {
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
		file_dream11_od_dto_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisioningConfig); i {
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
		file_dream11_od_dto_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMetadata); i {
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
	file_dream11_od_dto_v1_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_dto_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dream11_od_dto_v1_service_proto_goTypes,
		DependencyIndexes: file_dream11_od_dto_v1_service_proto_depIdxs,
		MessageInfos:      file_dream11_od_dto_v1_service_proto_msgTypes,
	}.Build()
	File_dream11_od_dto_v1_service_proto = out.File
	file_dream11_od_dto_v1_service_proto_rawDesc = nil
	file_dream11_od_dto_v1_service_proto_goTypes = nil
	file_dream11_od_dto_v1_service_proto_depIdxs = nil
}

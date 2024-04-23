// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: dream11/od/dto/v1/environment.proto

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

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      *int64                 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CreatedAt               *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt               *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	CreatedBy               *string                `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3,oneof" json:"created_by,omitempty"`
	Version                 *int32                 `protobuf:"varint,5,opt,name=version,proto3,oneof" json:"version,omitempty"`
	OrgId                   *int64                 `protobuf:"varint,6,opt,name=org_id,json=orgId,proto3,oneof" json:"org_id,omitempty"`
	ProviderAccountName     *string                `protobuf:"bytes,7,opt,name=provider_account_name,json=providerAccountName,proto3,oneof" json:"provider_account_name,omitempty"`
	Name                    *string                `protobuf:"bytes,8,opt,name=name,proto3,oneof" json:"name,omitempty"`
	ServiceAccountsSnapshot *structpb.Struct       `protobuf:"bytes,9,opt,name=service_accounts_snapshot,json=serviceAccountsSnapshot,proto3,oneof" json:"service_accounts_snapshot,omitempty"`
	Status                  *string                `protobuf:"bytes,10,opt,name=status,proto3,oneof" json:"status,omitempty"`
	AutoDeletionTime        *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=auto_deletion_time,json=autoDeletionTime,proto3,oneof" json:"auto_deletion_time,omitempty"`
	Services                []*ServiceTask         `protobuf:"bytes,12,rep,name=services,proto3" json:"services,omitempty"`
	AccountInformation      []*AccountInformation  `protobuf:"bytes,13,rep,name=account_information,json=accountInformation,proto3" json:"account_information,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_environment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_environment_proto_rawDescGZIP(), []int{0}
}

func (x *Environment) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Environment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Environment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Environment) GetCreatedBy() string {
	if x != nil && x.CreatedBy != nil {
		return *x.CreatedBy
	}
	return ""
}

func (x *Environment) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *Environment) GetOrgId() int64 {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return 0
}

func (x *Environment) GetProviderAccountName() string {
	if x != nil && x.ProviderAccountName != nil {
		return *x.ProviderAccountName
	}
	return ""
}

func (x *Environment) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Environment) GetServiceAccountsSnapshot() *structpb.Struct {
	if x != nil {
		return x.ServiceAccountsSnapshot
	}
	return nil
}

func (x *Environment) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *Environment) GetAutoDeletionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AutoDeletionTime
	}
	return nil
}

func (x *Environment) GetServices() []*ServiceTask {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Environment) GetAccountInformation() []*AccountInformation {
	if x != nil {
		return x.AccountInformation
	}
	return nil
}

type AccountInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderAccountName     string           `protobuf:"bytes,1,opt,name=provider_account_name,json=providerAccountName,proto3" json:"provider_account_name,omitempty"`
	ServiceAccountsSnapshot *structpb.Struct `protobuf:"bytes,2,opt,name=service_accounts_snapshot,json=serviceAccountsSnapshot,proto3" json:"service_accounts_snapshot,omitempty"`
	Status                  string           `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AccountInformation) Reset() {
	*x = AccountInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInformation) ProtoMessage() {}

func (x *AccountInformation) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInformation.ProtoReflect.Descriptor instead.
func (*AccountInformation) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_environment_proto_rawDescGZIP(), []int{1}
}

func (x *AccountInformation) GetProviderAccountName() string {
	if x != nil {
		return x.ProviderAccountName
	}
	return ""
}

func (x *AccountInformation) GetServiceAccountsSnapshot() *structpb.Struct {
	if x != nil {
		return x.ServiceAccountsSnapshot
	}
	return nil
}

func (x *AccountInformation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_dream11_od_dto_v1_environment_proto protoreflect.FileDescriptor

var file_dream11_od_dto_v1_environment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f,
	0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x06,
	0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x37, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x06, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x58, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48,
	0x08, 0x52, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x12, 0x61, 0x75,
	0x74, 0x6f, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x0a, 0x52, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x72,
	0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e,
	0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1c, 0x0a, 0x1a, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x12,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x69, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x72, 0x65, 0x61,
	0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_dto_v1_environment_proto_rawDescOnce sync.Once
	file_dream11_od_dto_v1_environment_proto_rawDescData = file_dream11_od_dto_v1_environment_proto_rawDesc
)

func file_dream11_od_dto_v1_environment_proto_rawDescGZIP() []byte {
	file_dream11_od_dto_v1_environment_proto_rawDescOnce.Do(func() {
		file_dream11_od_dto_v1_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_dto_v1_environment_proto_rawDescData)
	})
	return file_dream11_od_dto_v1_environment_proto_rawDescData
}

var file_dream11_od_dto_v1_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dream11_od_dto_v1_environment_proto_goTypes = []interface{}{
	(*Environment)(nil),           // 0: dream11.od.dto.v1.Environment
	(*AccountInformation)(nil),    // 1: dream11.od.dto.v1.AccountInformation
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 3: google.protobuf.Struct
	(*ServiceTask)(nil),           // 4: dream11.od.dto.v1.ServiceTask
}
var file_dream11_od_dto_v1_environment_proto_depIdxs = []int32{
	2, // 0: dream11.od.dto.v1.Environment.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: dream11.od.dto.v1.Environment.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: dream11.od.dto.v1.Environment.service_accounts_snapshot:type_name -> google.protobuf.Struct
	2, // 3: dream11.od.dto.v1.Environment.auto_deletion_time:type_name -> google.protobuf.Timestamp
	4, // 4: dream11.od.dto.v1.Environment.services:type_name -> dream11.od.dto.v1.ServiceTask
	1, // 5: dream11.od.dto.v1.Environment.account_information:type_name -> dream11.od.dto.v1.AccountInformation
	3, // 6: dream11.od.dto.v1.AccountInformation.service_accounts_snapshot:type_name -> google.protobuf.Struct
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_dream11_od_dto_v1_environment_proto_init() }
func file_dream11_od_dto_v1_environment_proto_init() {
	if File_dream11_od_dto_v1_environment_proto != nil {
		return
	}
	file_dream11_od_dto_v1_service_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dream11_od_dto_v1_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_dream11_od_dto_v1_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInformation); i {
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
	file_dream11_od_dto_v1_environment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_dto_v1_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dream11_od_dto_v1_environment_proto_goTypes,
		DependencyIndexes: file_dream11_od_dto_v1_environment_proto_depIdxs,
		MessageInfos:      file_dream11_od_dto_v1_environment_proto_msgTypes,
	}.Build()
	File_dream11_od_dto_v1_environment_proto = out.File
	file_dream11_od_dto_v1_environment_proto_rawDesc = nil
	file_dream11_od_dto_v1_environment_proto_goTypes = nil
	file_dream11_od_dto_v1_environment_proto_depIdxs = nil
}

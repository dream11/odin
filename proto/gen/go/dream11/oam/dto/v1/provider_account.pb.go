// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: dream11/oam/dto/v1/provider_account.proto

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

type ProviderAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                     string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Provider                 string                    `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Category                 string                    `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Data                     *structpb.Struct          `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Services                 []*ProviderServiceAccount `protobuf:"bytes,5,rep,name=services,proto3" json:"services,omitempty"`
	Default                  bool                      `protobuf:"varint,6,opt,name=default,proto3" json:"default,omitempty"`
	Id                       int64                     `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	LinkedProviderAccountIds []int64                   `protobuf:"varint,8,rep,packed,name=linked_provider_account_ids,json=linkedProviderAccountIds,proto3" json:"linked_provider_account_ids,omitempty"`
}

func (x *ProviderAccount) Reset() {
	*x = ProviderAccount{}
	mi := &file_dream11_oam_dto_v1_provider_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderAccount) ProtoMessage() {}

func (x *ProviderAccount) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_oam_dto_v1_provider_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderAccount.ProtoReflect.Descriptor instead.
func (*ProviderAccount) Descriptor() ([]byte, []int) {
	return file_dream11_oam_dto_v1_provider_account_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProviderAccount) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ProviderAccount) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProviderAccount) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ProviderAccount) GetServices() []*ProviderServiceAccount {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ProviderAccount) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *ProviderAccount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProviderAccount) GetLinkedProviderAccountIds() []int64 {
	if x != nil {
		return x.LinkedProviderAccountIds
	}
	return nil
}

type ProviderServiceAccountEnriched struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *ProviderServiceAccount `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Account *ProviderAccount        `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ProviderServiceAccountEnriched) Reset() {
	*x = ProviderServiceAccountEnriched{}
	mi := &file_dream11_oam_dto_v1_provider_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderServiceAccountEnriched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderServiceAccountEnriched) ProtoMessage() {}

func (x *ProviderServiceAccountEnriched) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_oam_dto_v1_provider_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderServiceAccountEnriched.ProtoReflect.Descriptor instead.
func (*ProviderServiceAccountEnriched) Descriptor() ([]byte, []int) {
	return file_dream11_oam_dto_v1_provider_account_proto_rawDescGZIP(), []int{1}
}

func (x *ProviderServiceAccountEnriched) GetService() *ProviderServiceAccount {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ProviderServiceAccountEnriched) GetAccount() *ProviderAccount {
	if x != nil {
		return x.Account
	}
	return nil
}

type ProviderServiceAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category string           `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Data     *structpb.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Id       int64            `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProviderServiceAccount) Reset() {
	*x = ProviderServiceAccount{}
	mi := &file_dream11_oam_dto_v1_provider_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderServiceAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderServiceAccount) ProtoMessage() {}

func (x *ProviderServiceAccount) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_oam_dto_v1_provider_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderServiceAccount.ProtoReflect.Descriptor instead.
func (*ProviderServiceAccount) Descriptor() ([]byte, []int) {
	return file_dream11_oam_dto_v1_provider_account_proto_rawDescGZIP(), []int{2}
}

func (x *ProviderServiceAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProviderServiceAccount) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProviderServiceAccount) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ProviderServiceAccount) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_dream11_oam_dto_v1_provider_account_proto protoreflect.FileDescriptor

var file_dream11_oam_dto_v1_provider_account_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x64, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x72, 0x65,
	0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x61, 0x6d, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64,
	0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x61, 0x6d, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x1b,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x18, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x1e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x12, 0x44,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x61, 0x6d, 0x2e, 0x64, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e,
	0x6f, 0x61, 0x6d, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31,
	0x31, 0x2f, 0x6f, 0x64, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x61, 0x6d, 0x2f,
	0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_oam_dto_v1_provider_account_proto_rawDescOnce sync.Once
	file_dream11_oam_dto_v1_provider_account_proto_rawDescData = file_dream11_oam_dto_v1_provider_account_proto_rawDesc
)

func file_dream11_oam_dto_v1_provider_account_proto_rawDescGZIP() []byte {
	file_dream11_oam_dto_v1_provider_account_proto_rawDescOnce.Do(func() {
		file_dream11_oam_dto_v1_provider_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_oam_dto_v1_provider_account_proto_rawDescData)
	})
	return file_dream11_oam_dto_v1_provider_account_proto_rawDescData
}

var file_dream11_oam_dto_v1_provider_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dream11_oam_dto_v1_provider_account_proto_goTypes = []any{
	(*ProviderAccount)(nil),                // 0: dream11.oam.dto.v1.ProviderAccount
	(*ProviderServiceAccountEnriched)(nil), // 1: dream11.oam.dto.v1.ProviderServiceAccountEnriched
	(*ProviderServiceAccount)(nil),         // 2: dream11.oam.dto.v1.ProviderServiceAccount
	(*structpb.Struct)(nil),                // 3: google.protobuf.Struct
}
var file_dream11_oam_dto_v1_provider_account_proto_depIdxs = []int32{
	3, // 0: dream11.oam.dto.v1.ProviderAccount.data:type_name -> google.protobuf.Struct
	2, // 1: dream11.oam.dto.v1.ProviderAccount.services:type_name -> dream11.oam.dto.v1.ProviderServiceAccount
	2, // 2: dream11.oam.dto.v1.ProviderServiceAccountEnriched.service:type_name -> dream11.oam.dto.v1.ProviderServiceAccount
	0, // 3: dream11.oam.dto.v1.ProviderServiceAccountEnriched.account:type_name -> dream11.oam.dto.v1.ProviderAccount
	3, // 4: dream11.oam.dto.v1.ProviderServiceAccount.data:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_dream11_oam_dto_v1_provider_account_proto_init() }
func file_dream11_oam_dto_v1_provider_account_proto_init() {
	if File_dream11_oam_dto_v1_provider_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_oam_dto_v1_provider_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dream11_oam_dto_v1_provider_account_proto_goTypes,
		DependencyIndexes: file_dream11_oam_dto_v1_provider_account_proto_depIdxs,
		MessageInfos:      file_dream11_oam_dto_v1_provider_account_proto_msgTypes,
	}.Build()
	File_dream11_oam_dto_v1_provider_account_proto = out.File
	file_dream11_oam_dto_v1_provider_account_proto_rawDesc = nil
	file_dream11_oam_dto_v1_provider_account_proto_goTypes = nil
	file_dream11_oam_dto_v1_provider_account_proto_depIdxs = nil
}

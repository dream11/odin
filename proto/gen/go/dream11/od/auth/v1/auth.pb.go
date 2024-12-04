// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: dream11/od/auth/v1/auth.proto

package v1

import (
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

type GetUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId         string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecretHash string `protobuf:"bytes,2,opt,name=client_secret_hash,json=clientSecretHash,proto3" json:"client_secret_hash,omitempty"`
}

func (x *GetUserTokenRequest) Reset() {
	*x = GetUserTokenRequest{}
	mi := &file_dream11_od_auth_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenRequest) ProtoMessage() {}

func (x *GetUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_auth_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenRequest.ProtoReflect.Descriptor instead.
func (*GetUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_dream11_od_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserTokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetUserTokenRequest) GetClientSecretHash() string {
	if x != nil {
		return x.ClientSecretHash
	}
	return ""
}

type GetUserTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetUserTokenResponse) Reset() {
	*x = GetUserTokenResponse{}
	mi := &file_dream11_od_auth_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResponse) ProtoMessage() {}

func (x *GetUserTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_auth_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResponse.ProtoReflect.Descriptor instead.
func (*GetUserTokenResponse) Descriptor() ([]byte, []int) {
	return file_dream11_od_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_dream11_od_auth_v1_auth_proto protoreflect.FileDescriptor

var file_dream11_od_auth_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x22, 0x60, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0x72, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x27, 0x2e, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x72,
	0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64,
	0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_auth_v1_auth_proto_rawDescOnce sync.Once
	file_dream11_od_auth_v1_auth_proto_rawDescData = file_dream11_od_auth_v1_auth_proto_rawDesc
)

func file_dream11_od_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_dream11_od_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_dream11_od_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_auth_v1_auth_proto_rawDescData)
	})
	return file_dream11_od_auth_v1_auth_proto_rawDescData
}

var file_dream11_od_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dream11_od_auth_v1_auth_proto_goTypes = []any{
	(*GetUserTokenRequest)(nil),  // 0: dream11.od.auth.v1.GetUserTokenRequest
	(*GetUserTokenResponse)(nil), // 1: dream11.od.auth.v1.GetUserTokenResponse
}
var file_dream11_od_auth_v1_auth_proto_depIdxs = []int32{
	0, // 0: dream11.od.auth.v1.AuthService.GetUserToken:input_type -> dream11.od.auth.v1.GetUserTokenRequest
	1, // 1: dream11.od.auth.v1.AuthService.GetUserToken:output_type -> dream11.od.auth.v1.GetUserTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dream11_od_auth_v1_auth_proto_init() }
func file_dream11_od_auth_v1_auth_proto_init() {
	if File_dream11_od_auth_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_auth_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dream11_od_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_dream11_od_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_dream11_od_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_dream11_od_auth_v1_auth_proto = out.File
	file_dream11_od_auth_v1_auth_proto_rawDesc = nil
	file_dream11_od_auth_v1_auth_proto_goTypes = nil
	file_dream11_od_auth_v1_auth_proto_depIdxs = nil
}

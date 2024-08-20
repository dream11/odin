// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: dream11/od/dto/v1/provisioning_type.proto

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

type ProvisioningType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProvisioningType) Reset() {
	*x = ProvisioningType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_provisioning_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisioningType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisioningType) ProtoMessage() {}

func (x *ProvisioningType) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_provisioning_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisioningType.ProtoReflect.Descriptor instead.
func (*ProvisioningType) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_provisioning_type_proto_rawDescGZIP(), []int{0}
}

func (x *ProvisioningType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_dream11_od_dto_v1_provisioning_type_proto protoreflect.FileDescriptor

var file_dream11_od_dto_v1_provisioning_type_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x72, 0x65,
	0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x26,
	0x0a, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x69,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64,
	0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_dto_v1_provisioning_type_proto_rawDescOnce sync.Once
	file_dream11_od_dto_v1_provisioning_type_proto_rawDescData = file_dream11_od_dto_v1_provisioning_type_proto_rawDesc
)

func file_dream11_od_dto_v1_provisioning_type_proto_rawDescGZIP() []byte {
	file_dream11_od_dto_v1_provisioning_type_proto_rawDescOnce.Do(func() {
		file_dream11_od_dto_v1_provisioning_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_dto_v1_provisioning_type_proto_rawDescData)
	})
	return file_dream11_od_dto_v1_provisioning_type_proto_rawDescData
}

var file_dream11_od_dto_v1_provisioning_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dream11_od_dto_v1_provisioning_type_proto_goTypes = []interface{}{
	(*ProvisioningType)(nil), // 0: dream11.od.dto.v1.ProvisioningType
}
var file_dream11_od_dto_v1_provisioning_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dream11_od_dto_v1_provisioning_type_proto_init() }
func file_dream11_od_dto_v1_provisioning_type_proto_init() {
	if File_dream11_od_dto_v1_provisioning_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dream11_od_dto_v1_provisioning_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProvisioningType); i {
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
			RawDescriptor: file_dream11_od_dto_v1_provisioning_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dream11_od_dto_v1_provisioning_type_proto_goTypes,
		DependencyIndexes: file_dream11_od_dto_v1_provisioning_type_proto_depIdxs,
		MessageInfos:      file_dream11_od_dto_v1_provisioning_type_proto_msgTypes,
	}.Build()
	File_dream11_od_dto_v1_provisioning_type_proto = out.File
	file_dream11_od_dto_v1_provisioning_type_proto_rawDesc = nil
	file_dream11_od_dto_v1_provisioning_type_proto_goTypes = nil
	file_dream11_od_dto_v1_provisioning_type_proto_depIdxs = nil
}

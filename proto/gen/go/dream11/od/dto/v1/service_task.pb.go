// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: dream11/od/dto/v1/service_task.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ServiceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *int64                 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3,oneof" json:"created_at,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	CreatedBy  *int64                 `protobuf:"varint,4,opt,name=created_by,json=createdBy,proto3,oneof" json:"created_by,omitempty"`
	OrgId      *int64                 `protobuf:"varint,5,opt,name=org_id,json=orgId,proto3,oneof" json:"org_id,omitempty"`
	Name       *string                `protobuf:"bytes,6,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Version    *string                `protobuf:"bytes,7,opt,name=version,proto3,oneof" json:"version,omitempty"`
	Status     *string                `protobuf:"bytes,8,opt,name=status,proto3,oneof" json:"status,omitempty"`
	EnvId      *int64                 `protobuf:"varint,9,opt,name=env_id,json=envId,proto3,oneof" json:"env_id,omitempty"`
	Components []*ComponentTask       `protobuf:"bytes,10,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *ServiceTask) Reset() {
	*x = ServiceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dream11_od_dto_v1_service_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceTask) ProtoMessage() {}

func (x *ServiceTask) ProtoReflect() protoreflect.Message {
	mi := &file_dream11_od_dto_v1_service_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceTask.ProtoReflect.Descriptor instead.
func (*ServiceTask) Descriptor() ([]byte, []int) {
	return file_dream11_od_dto_v1_service_task_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceTask) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ServiceTask) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ServiceTask) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *ServiceTask) GetCreatedBy() int64 {
	if x != nil && x.CreatedBy != nil {
		return *x.CreatedBy
	}
	return 0
}

func (x *ServiceTask) GetOrgId() int64 {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return 0
}

func (x *ServiceTask) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ServiceTask) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *ServiceTask) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *ServiceTask) GetEnvId() int64 {
	if x != nil && x.EnvId != nil {
		return *x.EnvId
	}
	return 0
}

func (x *ServiceTask) GetComponents() []*ComponentTask {
	if x != nil {
		return x.Components
	}
	return nil
}

var File_dream11_od_dto_v1_service_task_proto protoreflect.FileDescriptor

var file_dream11_od_dto_v1_service_task_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e,
	0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x64, 0x72, 0x65, 0x61, 0x6d,
	0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xff, 0x03, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06, 0x6f,
	0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x05, 0x6f,
	0x72, 0x67, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x06, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x07, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06,
	0x65, 0x6e, 0x76, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x08, 0x52, 0x05,
	0x65, 0x6e, 0x76, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64,
	0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2e, 0x6f, 0x64, 0x2e, 0x64, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x6e,
	0x76, 0x5f, 0x69, 0x64, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x65, 0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x69, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x72, 0x65,
	0x61, 0x6d, 0x31, 0x31, 0x2f, 0x6f, 0x64, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dream11_od_dto_v1_service_task_proto_rawDescOnce sync.Once
	file_dream11_od_dto_v1_service_task_proto_rawDescData = file_dream11_od_dto_v1_service_task_proto_rawDesc
)

func file_dream11_od_dto_v1_service_task_proto_rawDescGZIP() []byte {
	file_dream11_od_dto_v1_service_task_proto_rawDescOnce.Do(func() {
		file_dream11_od_dto_v1_service_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_dream11_od_dto_v1_service_task_proto_rawDescData)
	})
	return file_dream11_od_dto_v1_service_task_proto_rawDescData
}

var file_dream11_od_dto_v1_service_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dream11_od_dto_v1_service_task_proto_goTypes = []interface{}{
	(*ServiceTask)(nil),           // 0: dream11.od.dto.v1.ServiceTask
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*ComponentTask)(nil),         // 2: dream11.od.dto.v1.ComponentTask
}
var file_dream11_od_dto_v1_service_task_proto_depIdxs = []int32{
	1, // 0: dream11.od.dto.v1.ServiceTask.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: dream11.od.dto.v1.ServiceTask.updated_at:type_name -> google.protobuf.Timestamp
	2, // 2: dream11.od.dto.v1.ServiceTask.components:type_name -> dream11.od.dto.v1.ComponentTask
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dream11_od_dto_v1_service_task_proto_init() }
func file_dream11_od_dto_v1_service_task_proto_init() {
	if File_dream11_od_dto_v1_service_task_proto != nil {
		return
	}
	file_dream11_od_dto_v1_component_task_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dream11_od_dto_v1_service_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceTask); i {
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
	file_dream11_od_dto_v1_service_task_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dream11_od_dto_v1_service_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dream11_od_dto_v1_service_task_proto_goTypes,
		DependencyIndexes: file_dream11_od_dto_v1_service_task_proto_depIdxs,
		MessageInfos:      file_dream11_od_dto_v1_service_task_proto_msgTypes,
	}.Build()
	File_dream11_od_dto_v1_service_task_proto = out.File
	file_dream11_od_dto_v1_service_task_proto_rawDesc = nil
	file_dream11_od_dto_v1_service_task_proto_goTypes = nil
	file_dream11_od_dto_v1_service_task_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/ci/Jenkins.proto

package ci

import (
	permissions "github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/permissions"
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

// Configuration for the clouddriver microservice.
type Jenkins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool              `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Masters []*JenkinsMasters `protobuf:"bytes,2,rep,name=masters,proto3" json:"masters,omitempty"`
}

func (x *Jenkins) Reset() {
	*x = Jenkins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jenkins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jenkins) ProtoMessage() {}

func (x *Jenkins) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jenkins.ProtoReflect.Descriptor instead.
func (*Jenkins) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescGZIP(), []int{0}
}

func (x *Jenkins) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Jenkins) GetMasters() []*JenkinsMasters {
	if x != nil {
		return x.Masters
	}
	return nil
}

type JenkinsMasters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permission *permissions.Permissions `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	Address    string                   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Username   string                   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password   string                   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Csrf       bool                     `protobuf:"varint,6,opt,name=csrf,proto3" json:"csrf,omitempty"`
}

func (x *JenkinsMasters) Reset() {
	*x = JenkinsMasters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JenkinsMasters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JenkinsMasters) ProtoMessage() {}

func (x *JenkinsMasters) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JenkinsMasters.ProtoReflect.Descriptor instead.
func (*JenkinsMasters) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescGZIP(), []int{1}
}

func (x *JenkinsMasters) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JenkinsMasters) GetPermission() *permissions.Permissions {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *JenkinsMasters) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *JenkinsMasters) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *JenkinsMasters) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *JenkinsMasters) GetCsrf() bool {
	if x != nil {
		return x.Csrf
	}
	return false
}

var File_proto_deploymentConfigurations_ci_Jenkins_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x69, 0x2f, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x1a, 0x3c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x07, 0x4a, 0x65, 0x6e,
	0x6b, 0x69, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x32,
	0x0a, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x4a, 0x65, 0x6e, 0x6b, 0x69,
	0x6e, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x73, 0x72, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x63, 0x73, 0x72, 0x66, 0x42,
	0x24, 0x5a, 0x22, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescData = file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDesc
)

func file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDescData
}

var file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_deploymentConfigurations_ci_Jenkins_proto_goTypes = []interface{}{
	(*Jenkins)(nil),                 // 0: proto.ci.Jenkins
	(*JenkinsMasters)(nil),          // 1: proto.ci.JenkinsMasters
	(*permissions.Permissions)(nil), // 2: proto.permissions.Permissions
}
var file_proto_deploymentConfigurations_ci_Jenkins_proto_depIdxs = []int32{
	1, // 0: proto.ci.Jenkins.masters:type_name -> proto.ci.JenkinsMasters
	2, // 1: proto.ci.JenkinsMasters.permission:type_name -> proto.permissions.Permissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_ci_Jenkins_proto_init() }
func file_proto_deploymentConfigurations_ci_Jenkins_proto_init() {
	if File_proto_deploymentConfigurations_ci_Jenkins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jenkins); i {
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
		file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JenkinsMasters); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_ci_Jenkins_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_ci_Jenkins_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_ci_Jenkins_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_ci_Jenkins_proto = out.File
	file_proto_deploymentConfigurations_ci_Jenkins_proto_rawDesc = nil
	file_proto_deploymentConfigurations_ci_Jenkins_proto_goTypes = nil
	file_proto_deploymentConfigurations_ci_Jenkins_proto_depIdxs = nil
}
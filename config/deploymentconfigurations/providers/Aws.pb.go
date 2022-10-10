// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.5
// source: proto/deploymentConfigurations/providers/Aws.proto

package providers

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

// Configuration for the clouddriver microservice.
type Aws struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled                bool               `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PrimaryAccount         string             `protobuf:"bytes,2,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	Account                []*AwsAcc          `protobuf:"bytes,3,rep,name=Account,proto3" json:"Account,omitempty"`
	BakeryDefault          *AwsBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefault,proto3" json:"bakeryDefault,omitempty"`
	AccessKey              string             `protobuf:"bytes,5,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	SecretAccessKey        string             `protobuf:"bytes,6,opt,name=secretAccessKey,proto3" json:"secretAccessKey,omitempty"`
	DefaultKeyPairTemplate string             `protobuf:"bytes,7,opt,name=defaultKeyPairTemplate,proto3" json:"defaultKeyPairTemplate,omitempty"`
	DefaultRegions         []string           `protobuf:"bytes,8,rep,name=defaultRegions,proto3" json:"defaultRegions,omitempty"`
	Defaults               *AwsDefaults       `protobuf:"bytes,9,opt,name=defaults,proto3" json:"defaults,omitempty"`
}

func (x *Aws) Reset() {
	*x = Aws{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aws) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aws) ProtoMessage() {}

func (x *Aws) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aws.ProtoReflect.Descriptor instead.
func (*Aws) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Aws_proto_rawDescGZIP(), []int{0}
}

func (x *Aws) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Aws) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *Aws) GetAccount() []*AwsAcc {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Aws) GetBakeryDefault() *AwsBakeryDefaults {
	if x != nil {
		return x.BakeryDefault
	}
	return nil
}

func (x *Aws) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Aws) GetSecretAccessKey() string {
	if x != nil {
		return x.SecretAccessKey
	}
	return ""
}

func (x *Aws) GetDefaultKeyPairTemplate() string {
	if x != nil {
		return x.DefaultKeyPairTemplate
	}
	return ""
}

func (x *Aws) GetDefaultRegions() []string {
	if x != nil {
		return x.DefaultRegions
	}
	return nil
}

func (x *Aws) GetDefaults() *AwsDefaults {
	if x != nil {
		return x.Defaults
	}
	return nil
}

type AwsAcc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AccountId       string   `protobuf:"bytes,2,opt,name=accountId,proto3" json:"accountId,omitempty"`
	AssumeRole      string   `protobuf:"bytes,3,opt,name=assumeRole,proto3" json:"assumeRole,omitempty"`
	ProviderVersion string   `protobuf:"bytes,4,opt,name=providerVersion,proto3" json:"providerVersion,omitempty"`
	Regions         []string `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *AwsAcc) Reset() {
	*x = AwsAcc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsAcc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsAcc) ProtoMessage() {}

func (x *AwsAcc) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsAcc.ProtoReflect.Descriptor instead.
func (*AwsAcc) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Aws_proto_rawDescGZIP(), []int{1}
}

func (x *AwsAcc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AwsAcc) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *AwsAcc) GetAssumeRole() string {
	if x != nil {
		return x.AssumeRole
	}
	return ""
}

func (x *AwsAcc) GetProviderVersion() string {
	if x != nil {
		return x.ProviderVersion
	}
	return ""
}

func (x *AwsAcc) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

type AwsBakeryDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateFile                string   `protobuf:"bytes,1,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	AwsAccessKey                string   `protobuf:"bytes,2,opt,name=awsAccessKey,proto3" json:"awsAccessKey,omitempty"`
	AwsSecretKey                string   `protobuf:"bytes,3,opt,name=awsSecretKey,proto3" json:"awsSecretKey,omitempty"`
	BaseImages                  []string `protobuf:"bytes,4,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	AwsAssociatePublicIpAddress bool     `protobuf:"varint,5,opt,name=awsAssociatePublicIpAddress,proto3" json:"awsAssociatePublicIpAddress,omitempty"`
	DefaultVirtualizationType   string   `protobuf:"bytes,6,opt,name=defaultVirtualizationType,proto3" json:"defaultVirtualizationType,omitempty"`
}

func (x *AwsBakeryDefaults) Reset() {
	*x = AwsBakeryDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsBakeryDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsBakeryDefaults) ProtoMessage() {}

func (x *AwsBakeryDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsBakeryDefaults.ProtoReflect.Descriptor instead.
func (*AwsBakeryDefaults) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Aws_proto_rawDescGZIP(), []int{2}
}

func (x *AwsBakeryDefaults) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

func (x *AwsBakeryDefaults) GetAwsAccessKey() string {
	if x != nil {
		return x.AwsAccessKey
	}
	return ""
}

func (x *AwsBakeryDefaults) GetAwsSecretKey() string {
	if x != nil {
		return x.AwsSecretKey
	}
	return ""
}

func (x *AwsBakeryDefaults) GetBaseImages() []string {
	if x != nil {
		return x.BaseImages
	}
	return nil
}

func (x *AwsBakeryDefaults) GetAwsAssociatePublicIpAddress() bool {
	if x != nil {
		return x.AwsAssociatePublicIpAddress
	}
	return false
}

func (x *AwsBakeryDefaults) GetDefaultVirtualizationType() string {
	if x != nil {
		return x.DefaultVirtualizationType
	}
	return ""
}

type AwsDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IamBaseRole string `protobuf:"bytes,1,opt,name=iamBaseRole,proto3" json:"iamBaseRole,omitempty"`
}

func (x *AwsDefaults) Reset() {
	*x = AwsDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsDefaults) ProtoMessage() {}

func (x *AwsDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsDefaults.ProtoReflect.Descriptor instead.
func (*AwsDefaults) Descriptor() ([]byte, []int) {
	return file_proto_deploymentConfigurations_providers_Aws_proto_rawDescGZIP(), []int{3}
}

func (x *AwsDefaults) GetIamBaseRole() string {
	if x != nil {
		return x.IamBaseRole
	}
	return ""
}

var File_proto_deploymentConfigurations_providers_Aws_proto protoreflect.FileDescriptor

var file_proto_deploymentConfigurations_providers_Aws_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x41, 0x77, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0xa6, 0x03, 0x0a, 0x03, 0x41, 0x77, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x31, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x2e, 0x41, 0x77, 0x73, 0x41, 0x63, 0x63, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x48, 0x0a, 0x0d, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x77, 0x73, 0x42,
	0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x0d, 0x62,
	0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b,
	0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x69, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x77, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x73, 0x52, 0x08, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x9e,
	0x01, 0x0a, 0x06, 0x41, 0x77, 0x73, 0x41, 0x63, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x9f, 0x02, 0x0a, 0x11, 0x41, 0x77, 0x73, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x77, 0x73,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x77, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x40, 0x0a, 0x1b, 0x61, 0x77, 0x73, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x61, 0x77, 0x73, 0x41, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x2f, 0x0a, 0x0b, 0x41, 0x77, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x61, 0x6d, 0x42, 0x61, 0x73, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x61, 0x6d, 0x42, 0x61, 0x73, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deploymentConfigurations_providers_Aws_proto_rawDescOnce sync.Once
	file_proto_deploymentConfigurations_providers_Aws_proto_rawDescData = file_proto_deploymentConfigurations_providers_Aws_proto_rawDesc
)

func file_proto_deploymentConfigurations_providers_Aws_proto_rawDescGZIP() []byte {
	file_proto_deploymentConfigurations_providers_Aws_proto_rawDescOnce.Do(func() {
		file_proto_deploymentConfigurations_providers_Aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deploymentConfigurations_providers_Aws_proto_rawDescData)
	})
	return file_proto_deploymentConfigurations_providers_Aws_proto_rawDescData
}

var file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_deploymentConfigurations_providers_Aws_proto_goTypes = []interface{}{
	(*Aws)(nil),               // 0: proto.providers.Aws
	(*AwsAcc)(nil),            // 1: proto.providers.AwsAcc
	(*AwsBakeryDefaults)(nil), // 2: proto.providers.AwsBakeryDefaults
	(*AwsDefaults)(nil),       // 3: proto.providers.AwsDefaults
}
var file_proto_deploymentConfigurations_providers_Aws_proto_depIdxs = []int32{
	1, // 0: proto.providers.Aws.Account:type_name -> proto.providers.AwsAcc
	2, // 1: proto.providers.Aws.bakeryDefault:type_name -> proto.providers.AwsBakeryDefaults
	3, // 2: proto.providers.Aws.defaults:type_name -> proto.providers.AwsDefaults
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_deploymentConfigurations_providers_Aws_proto_init() }
func file_proto_deploymentConfigurations_providers_Aws_proto_init() {
	if File_proto_deploymentConfigurations_providers_Aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aws); i {
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
		file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsAcc); i {
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
		file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsBakeryDefaults); i {
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
		file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsDefaults); i {
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
			RawDescriptor: file_proto_deploymentConfigurations_providers_Aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_deploymentConfigurations_providers_Aws_proto_goTypes,
		DependencyIndexes: file_proto_deploymentConfigurations_providers_Aws_proto_depIdxs,
		MessageInfos:      file_proto_deploymentConfigurations_providers_Aws_proto_msgTypes,
	}.Build()
	File_proto_deploymentConfigurations_providers_Aws_proto = out.File
	file_proto_deploymentConfigurations_providers_Aws_proto_rawDesc = nil
	file_proto_deploymentConfigurations_providers_Aws_proto_goTypes = nil
	file_proto_deploymentConfigurations_providers_Aws_proto_depIdxs = nil
}

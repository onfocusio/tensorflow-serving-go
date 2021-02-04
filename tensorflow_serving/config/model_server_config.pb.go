// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: tensorflow_serving/config/model_server_config.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	storage_path "github.com/onfocusio/tensorflow-serving-go/tensorflow_serving/sources/storage_path"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The type of model.
// TODO(b/31336131): DEPRECATED.
type ModelType int32

const (
	// Deprecated: Do not use.
	ModelType_MODEL_TYPE_UNSPECIFIED ModelType = 0
	// Deprecated: Do not use.
	ModelType_TENSORFLOW ModelType = 1
	// Deprecated: Do not use.
	ModelType_OTHER ModelType = 2
)

// Enum value maps for ModelType.
var (
	ModelType_name = map[int32]string{
		0: "MODEL_TYPE_UNSPECIFIED",
		1: "TENSORFLOW",
		2: "OTHER",
	}
	ModelType_value = map[string]int32{
		"MODEL_TYPE_UNSPECIFIED": 0,
		"TENSORFLOW":             1,
		"OTHER":                  2,
	}
)

func (x ModelType) Enum() *ModelType {
	p := new(ModelType)
	*p = x
	return p
}

func (x ModelType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModelType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_serving_config_model_server_config_proto_enumTypes[0].Descriptor()
}

func (ModelType) Type() protoreflect.EnumType {
	return &file_tensorflow_serving_config_model_server_config_proto_enumTypes[0]
}

func (x ModelType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModelType.Descriptor instead.
func (ModelType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_model_server_config_proto_rawDescGZIP(), []int{0}
}

// Common configuration for loading a model being served.
type ModelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the model.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of aliases which can be used for addressing this model instead of
	// using the model name. This is an experimental field.
	Alias []string `protobuf:"bytes,9,rep,name=alias,proto3" json:"alias,omitempty"`
	// Base path to the model, excluding the version directory.
	// E.g> for a model at /foo/bar/my_model/123, where 123 is the version, the
	// base path is /foo/bar/my_model.
	//
	// (This can be changed once a model is in serving, *if* the underlying data
	// remains the same. Otherwise there are no guarantees about whether the old
	// or new data will be used for model versions currently loaded.)
	BasePath string `protobuf:"bytes,2,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	// Type of model.
	// TODO(b/31336131): DEPRECATED. Please use 'model_platform' instead.
	//
	// Deprecated: Do not use.
	ModelType ModelType `protobuf:"varint,3,opt,name=model_type,json=modelType,proto3,enum=tensorflow.serving.ModelType" json:"model_type,omitempty"`
	// Type of model (e.g. "tensorflow").
	//
	// (This cannot be changed once a model is in serving.)
	ModelPlatform string `protobuf:"bytes,4,opt,name=model_platform,json=modelPlatform,proto3" json:"model_platform,omitempty"`
	// Version policy for the model indicating which version(s) of the model to
	// load and make available for serving simultaneously.
	// The default option is to serve only the latest version of the model.
	//
	// (This can be changed once a model is in serving.)
	ModelVersionPolicy *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy `protobuf:"bytes,7,opt,name=model_version_policy,json=modelVersionPolicy,proto3" json:"model_version_policy,omitempty"`
	// String labels to associate with versions of the model, allowing inference
	// queries to refer to versions by label instead of number. Multiple labels
	// can map to the same version, but not vice-versa.
	//
	// An envisioned use-case for these labels is canarying tentative versions.
	// For example, one can assign labels "stable" and "canary" to two specific
	// versions. Perhaps initially "stable" is assigned to version 0 and "canary"
	// to version 1. Once version 1 passes canary, one can shift the "stable"
	// label to refer to version 1 (at that point both labels map to the same
	// version -- version 1 -- which is fine). Later once version 2 is ready to
	// canary one can move the "canary" label to version 2. And so on.
	VersionLabels map[string]int64 `protobuf:"bytes,8,rep,name=version_labels,json=versionLabels,proto3" json:"version_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Configures logging requests and responses, to the model.
	//
	// (This can be changed once a model is in serving.)
	LoggingConfig *LoggingConfig `protobuf:"bytes,6,opt,name=logging_config,json=loggingConfig,proto3" json:"logging_config,omitempty"`
}

func (x *ModelConfig) Reset() {
	*x = ModelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_model_server_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelConfig) ProtoMessage() {}

func (x *ModelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_model_server_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelConfig.ProtoReflect.Descriptor instead.
func (*ModelConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_model_server_config_proto_rawDescGZIP(), []int{0}
}

func (x *ModelConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelConfig) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *ModelConfig) GetBasePath() string {
	if x != nil {
		return x.BasePath
	}
	return ""
}

// Deprecated: Do not use.
func (x *ModelConfig) GetModelType() ModelType {
	if x != nil {
		return x.ModelType
	}
	return ModelType_MODEL_TYPE_UNSPECIFIED
}

func (x *ModelConfig) GetModelPlatform() string {
	if x != nil {
		return x.ModelPlatform
	}
	return ""
}

func (x *ModelConfig) GetModelVersionPolicy() *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy {
	if x != nil {
		return x.ModelVersionPolicy
	}
	return nil
}

func (x *ModelConfig) GetVersionLabels() map[string]int64 {
	if x != nil {
		return x.VersionLabels
	}
	return nil
}

func (x *ModelConfig) GetLoggingConfig() *LoggingConfig {
	if x != nil {
		return x.LoggingConfig
	}
	return nil
}

// Static list of models to be loaded for serving.
type ModelConfigList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config []*ModelConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *ModelConfigList) Reset() {
	*x = ModelConfigList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_model_server_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelConfigList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelConfigList) ProtoMessage() {}

func (x *ModelConfigList) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_model_server_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelConfigList.ProtoReflect.Descriptor instead.
func (*ModelConfigList) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_model_server_config_proto_rawDescGZIP(), []int{1}
}

func (x *ModelConfigList) GetConfig() []*ModelConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// ModelServer config.
type ModelServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ModelServer takes either a static file-based model config list or an Any
	// proto representing custom model config that is fetched dynamically at
	// runtime (through network RPC, custom service, etc.).
	//
	// Types that are assignable to Config:
	//	*ModelServerConfig_ModelConfigList
	//	*ModelServerConfig_CustomModelConfig
	Config isModelServerConfig_Config `protobuf_oneof:"config"`
}

func (x *ModelServerConfig) Reset() {
	*x = ModelServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_model_server_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelServerConfig) ProtoMessage() {}

func (x *ModelServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_model_server_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelServerConfig.ProtoReflect.Descriptor instead.
func (*ModelServerConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_model_server_config_proto_rawDescGZIP(), []int{2}
}

func (m *ModelServerConfig) GetConfig() isModelServerConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *ModelServerConfig) GetModelConfigList() *ModelConfigList {
	if x, ok := x.GetConfig().(*ModelServerConfig_ModelConfigList); ok {
		return x.ModelConfigList
	}
	return nil
}

func (x *ModelServerConfig) GetCustomModelConfig() *anypb.Any {
	if x, ok := x.GetConfig().(*ModelServerConfig_CustomModelConfig); ok {
		return x.CustomModelConfig
	}
	return nil
}

type isModelServerConfig_Config interface {
	isModelServerConfig_Config()
}

type ModelServerConfig_ModelConfigList struct {
	ModelConfigList *ModelConfigList `protobuf:"bytes,1,opt,name=model_config_list,json=modelConfigList,proto3,oneof"`
}

type ModelServerConfig_CustomModelConfig struct {
	CustomModelConfig *anypb.Any `protobuf:"bytes,2,opt,name=custom_model_config,json=customModelConfig,proto3,oneof"`
}

func (*ModelServerConfig_ModelConfigList) isModelServerConfig_Config() {}

func (*ModelServerConfig_CustomModelConfig) isModelServerConfig_Config() {}

var File_tensorflow_serving_config_model_server_config_proto protoreflect.FileDescriptor

var file_tensorflow_serving_config_model_server_config_proto_rawDesc = []byte{
	0x0a, 0x33, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x04, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x7d, 0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x4b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x12,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x59, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x48, 0x0a,
	0x0e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x40, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x22,
	0x4a, 0x0a, 0x0f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xb8, 0x01, 0x0a, 0x11,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x51, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x4e, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x16, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a,
	0x02, 0x08, 0x01, 0x12, 0x12, 0x0a, 0x0a, 0x54, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x46, 0x4c, 0x4f,
	0x57, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0d, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x02, 0x1a, 0x02, 0x08, 0x01, 0x42, 0x03, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_config_model_server_config_proto_rawDescOnce sync.Once
	file_tensorflow_serving_config_model_server_config_proto_rawDescData = file_tensorflow_serving_config_model_server_config_proto_rawDesc
)

func file_tensorflow_serving_config_model_server_config_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_config_model_server_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_config_model_server_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_config_model_server_config_proto_rawDescData)
	})
	return file_tensorflow_serving_config_model_server_config_proto_rawDescData
}

var file_tensorflow_serving_config_model_server_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_serving_config_model_server_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_serving_config_model_server_config_proto_goTypes = []interface{}{
	(ModelType)(0),            // 0: tensorflow.serving.ModelType
	(*ModelConfig)(nil),       // 1: tensorflow.serving.ModelConfig
	(*ModelConfigList)(nil),   // 2: tensorflow.serving.ModelConfigList
	(*ModelServerConfig)(nil), // 3: tensorflow.serving.ModelServerConfig
	nil,                       // 4: tensorflow.serving.ModelConfig.VersionLabelsEntry
	(*storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy)(nil), // 5: tensorflow.serving.FileSystemStoragePathSourceConfig.ServableVersionPolicy
	(*LoggingConfig)(nil), // 6: tensorflow.serving.LoggingConfig
	(*anypb.Any)(nil),     // 7: google.protobuf.Any
}
var file_tensorflow_serving_config_model_server_config_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.ModelConfig.model_type:type_name -> tensorflow.serving.ModelType
	5, // 1: tensorflow.serving.ModelConfig.model_version_policy:type_name -> tensorflow.serving.FileSystemStoragePathSourceConfig.ServableVersionPolicy
	4, // 2: tensorflow.serving.ModelConfig.version_labels:type_name -> tensorflow.serving.ModelConfig.VersionLabelsEntry
	6, // 3: tensorflow.serving.ModelConfig.logging_config:type_name -> tensorflow.serving.LoggingConfig
	1, // 4: tensorflow.serving.ModelConfigList.config:type_name -> tensorflow.serving.ModelConfig
	2, // 5: tensorflow.serving.ModelServerConfig.model_config_list:type_name -> tensorflow.serving.ModelConfigList
	7, // 6: tensorflow.serving.ModelServerConfig.custom_model_config:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_config_model_server_config_proto_init() }
func file_tensorflow_serving_config_model_server_config_proto_init() {
	if File_tensorflow_serving_config_model_server_config_proto != nil {
		return
	}
	file_tensorflow_serving_config_logging_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_config_model_server_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelConfig); i {
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
		file_tensorflow_serving_config_model_server_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelConfigList); i {
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
		file_tensorflow_serving_config_model_server_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelServerConfig); i {
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
	file_tensorflow_serving_config_model_server_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ModelServerConfig_ModelConfigList)(nil),
		(*ModelServerConfig_CustomModelConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_config_model_server_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_config_model_server_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_config_model_server_config_proto_depIdxs,
		EnumInfos:         file_tensorflow_serving_config_model_server_config_proto_enumTypes,
		MessageInfos:      file_tensorflow_serving_config_model_server_config_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_config_model_server_config_proto = out.File
	file_tensorflow_serving_config_model_server_config_proto_rawDesc = nil
	file_tensorflow_serving_config_model_server_config_proto_goTypes = nil
	file_tensorflow_serving_config_model_server_config_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: tensorflow_serving/apis/predict.proto

package apis

import (
	proto "github.com/golang/protobuf/proto"
	tensor_go_proto "github.com/onfocusio/tensorflow-serving-go/v2/tensorflow/go/core/framework/tensor_go_proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'inputs' field.
	Inputs map[string]*tensor_go_proto.TensorProto `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is stored in the SavedModel export as a prediction
	// SignatureDef under the 'outputs' field.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter,proto3" json:"output_filter,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_predict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_predict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_predict_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *PredictRequest) GetInputs() map[string]*tensor_go_proto.TensorProto {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *PredictRequest) GetOutputFilter() []string {
	if x != nil {
		return x.OutputFilter
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective Model Specification used to process PredictRequest.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Output tensors.
	Outputs map[string]*tensor_go_proto.TensorProto `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_predict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_predict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_predict_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *PredictResponse) GetOutputs() map[string]*tensor_go_proto.TensorProto {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_tensorflow_serving_apis_predict_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_predict_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x26, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x46, 0x0a, 0x06, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x52, 0x0a, 0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf0, 0x01, 0x0a, 0x0f, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4a, 0x0a, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x1a, 0x53, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x03, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_predict_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_predict_proto_rawDescData = file_tensorflow_serving_apis_predict_proto_rawDesc
)

func file_tensorflow_serving_apis_predict_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_predict_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_predict_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_predict_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_predict_proto_rawDescData
}

var file_tensorflow_serving_apis_predict_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_serving_apis_predict_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),              // 0: tensorflow.serving.PredictRequest
	(*PredictResponse)(nil),             // 1: tensorflow.serving.PredictResponse
	nil,                                 // 2: tensorflow.serving.PredictRequest.InputsEntry
	nil,                                 // 3: tensorflow.serving.PredictResponse.OutputsEntry
	(*ModelSpec)(nil),                   // 4: tensorflow.serving.ModelSpec
	(*tensor_go_proto.TensorProto)(nil), // 5: tensorflow.TensorProto
}
var file_tensorflow_serving_apis_predict_proto_depIdxs = []int32{
	4, // 0: tensorflow.serving.PredictRequest.model_spec:type_name -> tensorflow.serving.ModelSpec
	2, // 1: tensorflow.serving.PredictRequest.inputs:type_name -> tensorflow.serving.PredictRequest.InputsEntry
	4, // 2: tensorflow.serving.PredictResponse.model_spec:type_name -> tensorflow.serving.ModelSpec
	3, // 3: tensorflow.serving.PredictResponse.outputs:type_name -> tensorflow.serving.PredictResponse.OutputsEntry
	5, // 4: tensorflow.serving.PredictRequest.InputsEntry.value:type_name -> tensorflow.TensorProto
	5, // 5: tensorflow.serving.PredictResponse.OutputsEntry.value:type_name -> tensorflow.TensorProto
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_predict_proto_init() }
func file_tensorflow_serving_apis_predict_proto_init() {
	if File_tensorflow_serving_apis_predict_proto != nil {
		return
	}
	file_tensorflow_serving_apis_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_predict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_tensorflow_serving_apis_predict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
			RawDescriptor: file_tensorflow_serving_apis_predict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_predict_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_predict_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_predict_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_predict_proto = out.File
	file_tensorflow_serving_apis_predict_proto_rawDesc = nil
	file_tensorflow_serving_apis_predict_proto_goTypes = nil
	file_tensorflow_serving_apis_predict_proto_depIdxs = nil
}

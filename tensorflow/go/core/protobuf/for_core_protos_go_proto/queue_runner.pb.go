// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: tensorflow/core/protobuf/queue_runner.proto

package for_core_protos_go_proto

import (
	proto "github.com/golang/protobuf/proto"
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

// Protocol buffer representing a QueueRunner.
type QueueRunnerDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Queue name.
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	// A list of enqueue operations.
	EnqueueOpName []string `protobuf:"bytes,2,rep,name=enqueue_op_name,json=enqueueOpName,proto3" json:"enqueue_op_name,omitempty"`
	// The operation to run to close the queue.
	CloseOpName string `protobuf:"bytes,3,opt,name=close_op_name,json=closeOpName,proto3" json:"close_op_name,omitempty"`
	// The operation to run to cancel the queue.
	CancelOpName string `protobuf:"bytes,4,opt,name=cancel_op_name,json=cancelOpName,proto3" json:"cancel_op_name,omitempty"`
	// A list of exception types considered to signal a safely closed queue
	// if raised during enqueue operations.
	QueueClosedExceptionTypes []Code `protobuf:"varint,5,rep,packed,name=queue_closed_exception_types,json=queueClosedExceptionTypes,proto3,enum=tensorflow.error.Code" json:"queue_closed_exception_types,omitempty"`
}

func (x *QueueRunnerDef) Reset() {
	*x = QueueRunnerDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_queue_runner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueRunnerDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueRunnerDef) ProtoMessage() {}

func (x *QueueRunnerDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_queue_runner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueRunnerDef.ProtoReflect.Descriptor instead.
func (*QueueRunnerDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_queue_runner_proto_rawDescGZIP(), []int{0}
}

func (x *QueueRunnerDef) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *QueueRunnerDef) GetEnqueueOpName() []string {
	if x != nil {
		return x.EnqueueOpName
	}
	return nil
}

func (x *QueueRunnerDef) GetCloseOpName() string {
	if x != nil {
		return x.CloseOpName
	}
	return ""
}

func (x *QueueRunnerDef) GetCancelOpName() string {
	if x != nil {
		return x.CancelOpName
	}
	return ""
}

func (x *QueueRunnerDef) GetQueueClosedExceptionTypes() []Code {
	if x != nil {
		return x.QueueClosedExceptionTypes
	}
	return nil
}

var File_tensorflow_core_protobuf_queue_runner_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_queue_runner_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x65, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4f, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f, 0x6f, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x1c, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x19, 0x71, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x42, 0x89, 0x01, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42,
	0x11, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_queue_runner_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_queue_runner_proto_rawDescData = file_tensorflow_core_protobuf_queue_runner_proto_rawDesc
)

func file_tensorflow_core_protobuf_queue_runner_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_queue_runner_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_queue_runner_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_queue_runner_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_queue_runner_proto_rawDescData
}

var file_tensorflow_core_protobuf_queue_runner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_queue_runner_proto_goTypes = []interface{}{
	(*QueueRunnerDef)(nil), // 0: tensorflow.QueueRunnerDef
	(Code)(0),              // 1: tensorflow.error.Code
}
var file_tensorflow_core_protobuf_queue_runner_proto_depIdxs = []int32{
	1, // 0: tensorflow.QueueRunnerDef.queue_closed_exception_types:type_name -> tensorflow.error.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_queue_runner_proto_init() }
func file_tensorflow_core_protobuf_queue_runner_proto_init() {
	if File_tensorflow_core_protobuf_queue_runner_proto != nil {
		return
	}
	file_tensorflow_core_protobuf_error_codes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_queue_runner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueRunnerDef); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_queue_runner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_queue_runner_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_queue_runner_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_queue_runner_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_queue_runner_proto = out.File
	file_tensorflow_core_protobuf_queue_runner_proto_rawDesc = nil
	file_tensorflow_core_protobuf_queue_runner_proto_goTypes = nil
	file_tensorflow_core_protobuf_queue_runner_proto_depIdxs = nil
}

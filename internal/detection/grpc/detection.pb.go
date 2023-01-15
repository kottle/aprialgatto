// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: detection.proto

package grpc

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

type DetectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object string `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *DetectReq) Reset() {
	*x = DetectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectReq) ProtoMessage() {}

func (x *DetectReq) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectReq.ProtoReflect.Descriptor instead.
func (*DetectReq) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{0}
}

func (x *DetectReq) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

type DetectRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DetectRes) Reset() {
	*x = DetectRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectRes) ProtoMessage() {}

func (x *DetectRes) ProtoReflect() protoreflect.Message {
	mi := &file_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectRes.ProtoReflect.Descriptor instead.
func (*DetectRes) Descriptor() ([]byte, []int) {
	return file_detection_proto_rawDescGZIP(), []int{1}
}

var File_detection_proto protoreflect.FileDescriptor

var file_detection_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x23, 0x0a, 0x09, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x32, 0x5a, 0x0a, 0x10, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_detection_proto_rawDescOnce sync.Once
	file_detection_proto_rawDescData = file_detection_proto_rawDesc
)

func file_detection_proto_rawDescGZIP() []byte {
	file_detection_proto_rawDescOnce.Do(func() {
		file_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_detection_proto_rawDescData)
	})
	return file_detection_proto_rawDescData
}

var file_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_detection_proto_goTypes = []interface{}{
	(*DetectReq)(nil), // 0: rpc.detection.DetectReq
	(*DetectRes)(nil), // 1: rpc.detection.DetectRes
}
var file_detection_proto_depIdxs = []int32{
	0, // 0: rpc.detection.DetectionService.DetectedObject:input_type -> rpc.detection.DetectReq
	1, // 1: rpc.detection.DetectionService.DetectedObject:output_type -> rpc.detection.DetectRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_detection_proto_init() }
func file_detection_proto_init() {
	if File_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectReq); i {
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
		file_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectRes); i {
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
			RawDescriptor: file_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_detection_proto_goTypes,
		DependencyIndexes: file_detection_proto_depIdxs,
		MessageInfos:      file_detection_proto_msgTypes,
	}.Build()
	File_detection_proto = out.File
	file_detection_proto_rawDesc = nil
	file_detection_proto_goTypes = nil
	file_detection_proto_depIdxs = nil
}

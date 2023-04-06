// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: api/st_grpc/st_grpc.proto

package st_grpc

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

type LiveMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Author    string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LiveMessage) Reset() {
	*x = LiveMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_st_grpc_st_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveMessage) ProtoMessage() {}

func (x *LiveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_st_grpc_st_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveMessage.ProtoReflect.Descriptor instead.
func (*LiveMessage) Descriptor() ([]byte, []int) {
	return file_api_st_grpc_st_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *LiveMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LiveMessage) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *LiveMessage) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *LiveMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *LiveMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_st_grpc_st_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_st_grpc_st_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_api_st_grpc_st_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *StreamRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

var File_api_st_grpc_st_grpc_proto protoreflect.FileDescriptor

var file_api_st_grpc_st_grpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x32, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x58, 0x0a, 0x12, 0x4c, 0x69, 0x76, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x16, 0x2e, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x72, 0x74, 0x68, 0x75, 0x72, 0x34, 0x30, 0x34, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x74, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_st_grpc_st_grpc_proto_rawDescOnce sync.Once
	file_api_st_grpc_st_grpc_proto_rawDescData = file_api_st_grpc_st_grpc_proto_rawDesc
)

func file_api_st_grpc_st_grpc_proto_rawDescGZIP() []byte {
	file_api_st_grpc_st_grpc_proto_rawDescOnce.Do(func() {
		file_api_st_grpc_st_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_st_grpc_st_grpc_proto_rawDescData)
	})
	return file_api_st_grpc_st_grpc_proto_rawDescData
}

var file_api_st_grpc_st_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_st_grpc_st_grpc_proto_goTypes = []interface{}{
	(*LiveMessage)(nil),   // 0: st_grpc.LiveMessage
	(*StreamRequest)(nil), // 1: st_grpc.StreamRequest
}
var file_api_st_grpc_st_grpc_proto_depIdxs = []int32{
	1, // 0: st_grpc.LiveMessageService.StreamMessages:input_type -> st_grpc.StreamRequest
	0, // 1: st_grpc.LiveMessageService.StreamMessages:output_type -> st_grpc.LiveMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_st_grpc_st_grpc_proto_init() }
func file_api_st_grpc_st_grpc_proto_init() {
	if File_api_st_grpc_st_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_st_grpc_st_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveMessage); i {
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
		file_api_st_grpc_st_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
			RawDescriptor: file_api_st_grpc_st_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_st_grpc_st_grpc_proto_goTypes,
		DependencyIndexes: file_api_st_grpc_st_grpc_proto_depIdxs,
		MessageInfos:      file_api_st_grpc_st_grpc_proto_msgTypes,
	}.Build()
	File_api_st_grpc_st_grpc_proto = out.File
	file_api_st_grpc_st_grpc_proto_rawDesc = nil
	file_api_st_grpc_st_grpc_proto_goTypes = nil
	file_api_st_grpc_st_grpc_proto_depIdxs = nil
}
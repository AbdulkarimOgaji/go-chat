// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: rpc_receive_message.proto

package pb

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

type ReceiveMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *ReceiveMessageRequest) Reset() {
	*x = ReceiveMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_receive_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessageRequest) ProtoMessage() {}

func (x *ReceiveMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_receive_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessageRequest.ProtoReflect.Descriptor instead.
func (*ReceiveMessageRequest) Descriptor() ([]byte, []int) {
	return file_rpc_receive_message_proto_rawDescGZIP(), []int{0}
}

func (x *ReceiveMessageRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ReceiveMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *ReceiveMessageResponse) Reset() {
	*x = ReceiveMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_receive_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessageResponse) ProtoMessage() {}

func (x *ReceiveMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_receive_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessageResponse.ProtoReflect.Descriptor instead.
func (*ReceiveMessageResponse) Descriptor() ([]byte, []int) {
	return file_rpc_receive_message_proto_rawDescGZIP(), []int{1}
}

func (x *ReceiveMessageResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

var File_rpc_receive_message_proto protoreflect.FileDescriptor

var file_rpc_receive_message_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x15, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62,
	0x64, 0x75, 0x6c, 0x6b, 0x61, 0x72, 0x69, 0x6d, 0x6f, 0x67, 0x61, 0x6a, 0x69, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_receive_message_proto_rawDescOnce sync.Once
	file_rpc_receive_message_proto_rawDescData = file_rpc_receive_message_proto_rawDesc
)

func file_rpc_receive_message_proto_rawDescGZIP() []byte {
	file_rpc_receive_message_proto_rawDescOnce.Do(func() {
		file_rpc_receive_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_receive_message_proto_rawDescData)
	})
	return file_rpc_receive_message_proto_rawDescData
}

var file_rpc_receive_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_receive_message_proto_goTypes = []interface{}{
	(*ReceiveMessageRequest)(nil),  // 0: pb.ReceiveMessageRequest
	(*ReceiveMessageResponse)(nil), // 1: pb.ReceiveMessageResponse
	(*Chat)(nil),                   // 2: pb.Chat
}
var file_rpc_receive_message_proto_depIdxs = []int32{
	2, // 0: pb.ReceiveMessageResponse.chat:type_name -> pb.Chat
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_receive_message_proto_init() }
func file_rpc_receive_message_proto_init() {
	if File_rpc_receive_message_proto != nil {
		return
	}
	file_chat_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_receive_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessageRequest); i {
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
		file_rpc_receive_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessageResponse); i {
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
			RawDescriptor: file_rpc_receive_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_receive_message_proto_goTypes,
		DependencyIndexes: file_rpc_receive_message_proto_depIdxs,
		MessageInfos:      file_rpc_receive_message_proto_msgTypes,
	}.Build()
	File_rpc_receive_message_proto = out.File
	file_rpc_receive_message_proto_rawDesc = nil
	file_rpc_receive_message_proto_goTypes = nil
	file_rpc_receive_message_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: rpc_get_my_rooms.proto

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

type GetMyRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetMyRoomsRequest) Reset() {
	*x = GetMyRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_my_rooms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyRoomsRequest) ProtoMessage() {}

func (x *GetMyRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_my_rooms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyRoomsRequest.ProtoReflect.Descriptor instead.
func (*GetMyRoomsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_get_my_rooms_proto_rawDescGZIP(), []int{0}
}

func (x *GetMyRoomsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetMyRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *GetMyRoomsResponse) Reset() {
	*x = GetMyRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_get_my_rooms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyRoomsResponse) ProtoMessage() {}

func (x *GetMyRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_get_my_rooms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetMyRoomsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_get_my_rooms_proto_rawDescGZIP(), []int{1}
}

func (x *GetMyRoomsResponse) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

var File_rpc_get_my_rooms_proto protoreflect.FileDescriptor

var file_rpc_get_my_rooms_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x79, 0x5f, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a, 0x72, 0x6f,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x52,
	0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x64, 0x75, 0x6c, 0x6b, 0x61,
	0x72, 0x69, 0x6d, 0x6f, 0x67, 0x61, 0x6a, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x68, 0x61, 0x74,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_get_my_rooms_proto_rawDescOnce sync.Once
	file_rpc_get_my_rooms_proto_rawDescData = file_rpc_get_my_rooms_proto_rawDesc
)

func file_rpc_get_my_rooms_proto_rawDescGZIP() []byte {
	file_rpc_get_my_rooms_proto_rawDescOnce.Do(func() {
		file_rpc_get_my_rooms_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_get_my_rooms_proto_rawDescData)
	})
	return file_rpc_get_my_rooms_proto_rawDescData
}

var file_rpc_get_my_rooms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_get_my_rooms_proto_goTypes = []interface{}{
	(*GetMyRoomsRequest)(nil),  // 0: pb.GetMyRoomsRequest
	(*GetMyRoomsResponse)(nil), // 1: pb.GetMyRoomsResponse
	(*Room)(nil),               // 2: pb.Room
}
var file_rpc_get_my_rooms_proto_depIdxs = []int32{
	2, // 0: pb.GetMyRoomsResponse.room:type_name -> pb.Room
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_get_my_rooms_proto_init() }
func file_rpc_get_my_rooms_proto_init() {
	if File_rpc_get_my_rooms_proto != nil {
		return
	}
	file_room_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_get_my_rooms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyRoomsRequest); i {
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
		file_rpc_get_my_rooms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyRoomsResponse); i {
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
			RawDescriptor: file_rpc_get_my_rooms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_get_my_rooms_proto_goTypes,
		DependencyIndexes: file_rpc_get_my_rooms_proto_depIdxs,
		MessageInfos:      file_rpc_get_my_rooms_proto_msgTypes,
	}.Build()
	File_rpc_get_my_rooms_proto = out.File
	file_rpc_get_my_rooms_proto_rawDesc = nil
	file_rpc_get_my_rooms_proto_goTypes = nil
	file_rpc_get_my_rooms_proto_depIdxs = nil
}

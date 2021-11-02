// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: chittyclient.proto

package chittyclient

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

type Client_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text    string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Lamport int64  `protobuf:"varint,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Client_Message) Reset() {
	*x = Client_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittyclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_Message) ProtoMessage() {}

func (x *Client_Message) ProtoReflect() protoreflect.Message {
	mi := &file_chittyclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_Message.ProtoReflect.Descriptor instead.
func (*Client_Message) Descriptor() ([]byte, []int) {
	return file_chittyclient_proto_rawDescGZIP(), []int{0}
}

func (x *Client_Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Client_Message) GetLamport() int64 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

func (x *Client_Message) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Client_UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lamport int64  `protobuf:"varint,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *Client_UserName) Reset() {
	*x = Client_UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittyclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client_UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_UserName) ProtoMessage() {}

func (x *Client_UserName) ProtoReflect() protoreflect.Message {
	mi := &file_chittyclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_UserName.ProtoReflect.Descriptor instead.
func (*Client_UserName) Descriptor() ([]byte, []int) {
	return file_chittyclient_proto_rawDescGZIP(), []int{1}
}

func (x *Client_UserName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Client_UserName) GetLamport() int64 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type Client_ResponseCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description *string `protobuf:"bytes,2,opt,name=description,proto3,oneof" json:"description,omitempty"`
}

func (x *Client_ResponseCode) Reset() {
	*x = Client_ResponseCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chittyclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client_ResponseCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_ResponseCode) ProtoMessage() {}

func (x *Client_ResponseCode) ProtoReflect() protoreflect.Message {
	mi := &file_chittyclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_ResponseCode.ProtoReflect.Descriptor instead.
func (*Client_ResponseCode) Descriptor() ([]byte, []int) {
	return file_chittyclient_proto_rawDescGZIP(), []int{2}
}

func (x *Client_ResponseCode) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Client_ResponseCode) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

var File_chittyclient_proto protoreflect.FileDescriptor

var file_chittyclient_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x60, 0x0a, 0x13, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb3, 0x01, 0x0a, 0x0c,
	0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x09,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x36, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4a, 0x6f, 0x69, 0x6e,
	0x12, 0x10, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x10, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x75, 0x6b, 0x73, 0x2d, 0x69, 0x74, 0x75, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x32, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x74, 0x79, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chittyclient_proto_rawDescOnce sync.Once
	file_chittyclient_proto_rawDescData = file_chittyclient_proto_rawDesc
)

func file_chittyclient_proto_rawDescGZIP() []byte {
	file_chittyclient_proto_rawDescOnce.Do(func() {
		file_chittyclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_chittyclient_proto_rawDescData)
	})
	return file_chittyclient_proto_rawDescData
}

var file_chittyclient_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chittyclient_proto_goTypes = []interface{}{
	(*Client_Message)(nil),      // 0: Client_Message
	(*Client_UserName)(nil),     // 1: Client_UserName
	(*Client_ResponseCode)(nil), // 2: Client_ResponseCode
}
var file_chittyclient_proto_depIdxs = []int32{
	0, // 0: ChittyClient.Broadcast:input_type -> Client_Message
	1, // 1: ChittyClient.AnnounceJoin:input_type -> Client_UserName
	1, // 2: ChittyClient.AnnounceLeave:input_type -> Client_UserName
	2, // 3: ChittyClient.Broadcast:output_type -> Client_ResponseCode
	2, // 4: ChittyClient.AnnounceJoin:output_type -> Client_ResponseCode
	2, // 5: ChittyClient.AnnounceLeave:output_type -> Client_ResponseCode
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chittyclient_proto_init() }
func file_chittyclient_proto_init() {
	if File_chittyclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chittyclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client_Message); i {
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
		file_chittyclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client_UserName); i {
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
		file_chittyclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client_ResponseCode); i {
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
	file_chittyclient_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chittyclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chittyclient_proto_goTypes,
		DependencyIndexes: file_chittyclient_proto_depIdxs,
		MessageInfos:      file_chittyclient_proto_msgTypes,
	}.Build()
	File_chittyclient_proto = out.File
	file_chittyclient_proto_rawDesc = nil
	file_chittyclient_proto_goTypes = nil
	file_chittyclient_proto_depIdxs = nil
}

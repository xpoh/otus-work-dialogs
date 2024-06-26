// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: dialogs/v1/dialogs.proto

package dialogs

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

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFrom string `protobuf:"bytes,1,opt,name=user_from,json=userFrom,proto3" json:"user_from,omitempty"`
	UserTo   string `protobuf:"bytes,2,opt,name=user_to,json=userTo,proto3" json:"user_to,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialogs_v1_dialogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dialogs_v1_dialogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_dialogs_v1_dialogs_proto_rawDescGZIP(), []int{0}
}

func (x *SendRequest) GetUserFrom() string {
	if x != nil {
		return x.UserFrom
	}
	return ""
}

func (x *SendRequest) GetUserTo() string {
	if x != nil {
		return x.UserTo
	}
	return ""
}

func (x *SendRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialogs_v1_dialogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dialogs_v1_dialogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_dialogs_v1_dialogs_proto_rawDescGZIP(), []int{1}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserFrom string `protobuf:"bytes,1,opt,name=user_from,json=userFrom,proto3" json:"user_from,omitempty"`
	UserTo   string `protobuf:"bytes,2,opt,name=user_to,json=userTo,proto3" json:"user_to,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialogs_v1_dialogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dialogs_v1_dialogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_dialogs_v1_dialogs_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetUserFrom() string {
	if x != nil {
		return x.UserFrom
	}
	return ""
}

func (x *ListRequest) GetUserTo() string {
	if x != nil {
		return x.UserTo
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dialogs_v1_dialogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dialogs_v1_dialogs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_dialogs_v1_dialogs_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

var File_dialogs_v1_dialogs_proto protoreflect.FileDescriptor

var file_dialogs_v1_dialogs_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x57, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72,
	0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x43, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x89, 0x01, 0x0a, 0x0d, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x53, 0x65,
	0x6e, 0x64, 0x12, 0x17, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x17, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x70, 0x6f, 0x68, 0x2f, 0x6f, 0x74, 0x75, 0x73, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dialogs_v1_dialogs_proto_rawDescOnce sync.Once
	file_dialogs_v1_dialogs_proto_rawDescData = file_dialogs_v1_dialogs_proto_rawDesc
)

func file_dialogs_v1_dialogs_proto_rawDescGZIP() []byte {
	file_dialogs_v1_dialogs_proto_rawDescOnce.Do(func() {
		file_dialogs_v1_dialogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_dialogs_v1_dialogs_proto_rawDescData)
	})
	return file_dialogs_v1_dialogs_proto_rawDescData
}

var file_dialogs_v1_dialogs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dialogs_v1_dialogs_proto_goTypes = []interface{}{
	(*SendRequest)(nil),  // 0: dialogs.v1.SendRequest
	(*SendResponse)(nil), // 1: dialogs.v1.SendResponse
	(*ListRequest)(nil),  // 2: dialogs.v1.ListRequest
	(*ListResponse)(nil), // 3: dialogs.v1.ListResponse
}
var file_dialogs_v1_dialogs_proto_depIdxs = []int32{
	0, // 0: dialogs.v1.DialogService.Send:input_type -> dialogs.v1.SendRequest
	2, // 1: dialogs.v1.DialogService.List:input_type -> dialogs.v1.ListRequest
	1, // 2: dialogs.v1.DialogService.Send:output_type -> dialogs.v1.SendResponse
	3, // 3: dialogs.v1.DialogService.List:output_type -> dialogs.v1.ListResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dialogs_v1_dialogs_proto_init() }
func file_dialogs_v1_dialogs_proto_init() {
	if File_dialogs_v1_dialogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dialogs_v1_dialogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_dialogs_v1_dialogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
		file_dialogs_v1_dialogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_dialogs_v1_dialogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_dialogs_v1_dialogs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dialogs_v1_dialogs_proto_goTypes,
		DependencyIndexes: file_dialogs_v1_dialogs_proto_depIdxs,
		MessageInfos:      file_dialogs_v1_dialogs_proto_msgTypes,
	}.Build()
	File_dialogs_v1_dialogs_proto = out.File
	file_dialogs_v1_dialogs_proto_rawDesc = nil
	file_dialogs_v1_dialogs_proto_goTypes = nil
	file_dialogs_v1_dialogs_proto_depIdxs = nil
}

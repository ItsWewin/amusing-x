// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: join.proto

package userservice

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname         string `protobuf:"bytes,1,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Password         string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	AreaCode         string `protobuf:"bytes,3,opt,name=AreaCode,proto3" json:"AreaCode,omitempty"`
	Phone            string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	VerificationCode string `protobuf:"bytes,5,opt,name=VerificationCode,proto3" json:"VerificationCode,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_join_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_join_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_join_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *JoinRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *JoinRequest) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *JoinRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *JoinRequest) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname string `protobuf:"bytes,1,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	AreaCode string `protobuf:"bytes,2,opt,name=AreaCode,proto3" json:"AreaCode,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_join_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_join_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_join_proto_rawDescGZIP(), []int{1}
}

func (x *JoinResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *JoinResponse) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *JoinResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_join_proto protoreflect.FileDescriptor

var file_join_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6a, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x5c, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41,
	0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41,
	0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x27, 0x5a,
	0x25, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_join_proto_rawDescOnce sync.Once
	file_join_proto_rawDescData = file_join_proto_rawDesc
)

func file_join_proto_rawDescGZIP() []byte {
	file_join_proto_rawDescOnce.Do(func() {
		file_join_proto_rawDescData = protoimpl.X.CompressGZIP(file_join_proto_rawDescData)
	})
	return file_join_proto_rawDescData
}

var file_join_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_join_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),  // 0: userservice.JoinRequest
	(*JoinResponse)(nil), // 1: userservice.JoinResponse
}
var file_join_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_join_proto_init() }
func file_join_proto_init() {
	if File_join_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_join_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_join_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
			RawDescriptor: file_join_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_join_proto_goTypes,
		DependencyIndexes: file_join_proto_depIdxs,
		MessageInfos:      file_join_proto_msgTypes,
	}.Build()
	File_join_proto = out.File
	file_join_proto_rawDesc = nil
	file_join_proto_goTypes = nil
	file_join_proto_depIdxs = nil
}

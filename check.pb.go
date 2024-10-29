// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: check.proto

package protocheck

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Check struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // optional identifier for lookup purposes
	Fail string `protobuf:"bytes,2,opt,name=fail,proto3" json:"fail,omitempty"` // optional message to report when the check fails
	Cel  string `protobuf:"bytes,3,opt,name=cel,proto3" json:"cel,omitempty"`   // CEL expression
}

func (x *Check) Reset() {
	*x = Check{}
	mi := &file_check_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Check) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Check) ProtoMessage() {}

func (x *Check) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Check.ProtoReflect.Descriptor instead.
func (*Check) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{0}
}

func (x *Check) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Check) GetFail() string {
	if x != nil {
		return x.Fail
	}
	return ""
}

func (x *Check) GetCel() string {
	if x != nil {
		return x.Cel
	}
	return ""
}

var file_check_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Check)(nil),
		Field:         20030916,
		Name:          "check.message",
		Tag:           "bytes,20030916,opt,name=message",
		Filename:      "check.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Check)(nil),
		Field:         20030916,
		Name:          "check.field",
		Tag:           "bytes,20030916,opt,name=field",
		Filename:      "check.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional check.Check message = 20030916;
	E_Message = &file_check_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional check.Check field = 20030916;
	E_Field = &file_check_proto_extTypes[1]
)

var File_check_proto protoreflect.FileDescriptor

var file_check_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x65, 0x6c, 0x3a, 0x4d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xc4, 0xcb, 0xc6, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x3a, 0x47, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc4, 0xcb, 0xc6,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x69, 0x63,
	0x6b, 0x6c, 0x65, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_check_proto_rawDescOnce sync.Once
	file_check_proto_rawDescData = file_check_proto_rawDesc
)

func file_check_proto_rawDescGZIP() []byte {
	file_check_proto_rawDescOnce.Do(func() {
		file_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_proto_rawDescData)
	})
	return file_check_proto_rawDescData
}

var file_check_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_check_proto_goTypes = []any{
	(*Check)(nil),                       // 0: check.Check
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
}
var file_check_proto_depIdxs = []int32{
	1, // 0: check.message:extendee -> google.protobuf.MessageOptions
	2, // 1: check.field:extendee -> google.protobuf.FieldOptions
	0, // 2: check.message:type_name -> check.Check
	0, // 3: check.field:type_name -> check.Check
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_check_proto_init() }
func file_check_proto_init() {
	if File_check_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_check_proto_goTypes,
		DependencyIndexes: file_check_proto_depIdxs,
		MessageInfos:      file_check_proto_msgTypes,
		ExtensionInfos:    file_check_proto_extTypes,
	}.Build()
	File_check_proto = out.File
	file_check_proto_rawDesc = nil
	file_check_proto_goTypes = nil
	file_check_proto_depIdxs = nil
}

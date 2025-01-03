// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: person.proto

package golang

import (
	_ "github.com/emicklei/protocheck"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Person struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// with per field state checks
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MiddleName    *string                `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	Surname       string                 `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	BirthDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Health        *Person_Health         `protobuf:"bytes,5,opt,name=health,proto3" json:"health,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetMiddleName() string {
	if x != nil && x.MiddleName != nil {
		return *x.MiddleName
	}
	return ""
}

func (x *Person) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Person) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *Person) GetHealth() *Person_Health {
	if x != nil {
		return x.Health
	}
	return nil
}

// second message in same file
type Pet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Pet) Reset() {
	*x = Pet{}
	mi := &file_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Pet) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

// nested message
type Person_Health struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Weight        int32                  `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person_Health) Reset() {
	*x = Person_Health{}
	mi := &file_person_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person_Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Health) ProtoMessage() {}

func (x *Person_Health) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Health.ProtoReflect.Descriptor instead.
func (*Person_Health) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Person_Health) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x1a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x05, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x4a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xa2,
	0xdc, 0xb4, 0x4c, 0x31, 0x12, 0x1a, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20, 0x31,
	0x1a, 0x13, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x29, 0x20, 0x3e, 0x20, 0x31, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x0b, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x47, 0xa2, 0xdc, 0xb4, 0x4c, 0x42, 0x12, 0x24, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x20,
	0x6e, 0x61, 0x6d, 0x65, 0x20, 0x28, 0x69, 0x66, 0x20, 0x73, 0x65, 0x74, 0x29, 0x20, 0x63, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x73,
	0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x56, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0xa2, 0xdc, 0xb4,
	0x4c, 0x37, 0x12, 0x1d, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74,
	0x20, 0x62, 0x65, 0x20, 0x6c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20,
	0x31, 0x1a, 0x16, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x31, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x78, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x3d, 0xa2, 0xdc, 0xb4, 0x4c, 0x38, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x24, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x67, 0x65, 0x74,
	0x46, 0x75, 0x6c, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x28, 0x29, 0x20, 0x3e, 0x20, 0x32, 0x30, 0x30,
	0x30, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x57, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x4d, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x35, 0xa2, 0xdc, 0xb4, 0x4c, 0x30, 0x12, 0x1d, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x6b, 0x67, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x62, 0x65, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x1a, 0x0f, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x3a, 0x5d, 0xa2, 0xdc, 0xb4, 0x4c, 0x58, 0x0a, 0x10, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x6e,
	0x61, 0x6d, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x20,
	0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x22, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x20,
	0x2b, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20,
	0x3e, 0x20, 0x30, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x61, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4d, 0xa2, 0xdc, 0xb4, 0x4c, 0x48, 0x12,
	0x1c, 0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x64, 0x6f, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x61,
	0x74, 0x20, 0x61, 0x72, 0x65, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x6b, 0x69, 0x6e, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x63, 0x61, 0x74,
	0x27, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6b, 0x69, 0x6e, 0x64, 0x20, 0x3d,
	0x3d, 0x20, 0x27, 0x64, 0x6f, 0x67, 0x27, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x69, 0x63,
	0x6b, 0x6c, 0x65, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x3b, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_person_proto_goTypes = []any{
	(*Person)(nil),                // 0: golang.Person
	(*Pet)(nil),                   // 1: golang.Pet
	(*Person_Health)(nil),         // 2: golang.Person.Health
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_person_proto_depIdxs = []int32{
	3, // 0: golang.Person.birth_date:type_name -> google.protobuf.Timestamp
	2, // 1: golang.Person.health:type_name -> golang.Person.Health
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	file_person_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}

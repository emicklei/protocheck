// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
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

type Person_Gender int32

const (
	Person_UNKNOWN Person_Gender = 0
	Person_MALE    Person_Gender = 1
	Person_FEMALE  Person_Gender = 2
)

// Enum value maps for Person_Gender.
var (
	Person_Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	Person_Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
	}
)

func (x Person_Gender) Enum() *Person_Gender {
	p := new(Person_Gender)
	*p = x
	return p
}

func (x Person_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_person_proto_enumTypes[0].Descriptor()
}

func (Person_Gender) Type() protoreflect.EnumType {
	return &file_person_proto_enumTypes[0]
}

func (x Person_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_Gender.Descriptor instead.
func (Person_Gender) EnumDescriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0, 0}
}

type Person struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// with per field state checks
	Name       string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MiddleName *string                `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3,oneof" json:"middle_name,omitempty"`
	Surname    string                 `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	BirthDate  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	Health     *Person_Health         `protobuf:"bytes,5,opt,name=health,proto3" json:"health,omitempty"`
	// oneof with field checks
	// test with https://playcel.undistro.io/
	//
	// Types that are valid to be assigned to Identification:
	//
	//	*Person_Email
	//	*Person_Phone
	Identification isPerson_Identification `protobuf_oneof:"identification"`
	// repeated with field check
	Nicknames []string `protobuf:"bytes,8,rep,name=nicknames,proto3" json:"nicknames,omitempty"`
	// repeated proto message with  check
	Pets []*Pet `protobuf:"bytes,9,rep,name=pets,proto3" json:"pets,omitempty"`
	// map with field check
	Attributes map[string]string `protobuf:"bytes,10,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// map with proto message value with check
	Favorites     map[string]*Pet        `protobuf:"bytes,11,rep,name=favorites,proto3" json:"favorites,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NoCheckDate   *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=no_check_date,json=noCheckDate,proto3" json:"no_check_date,omitempty"`
	Gender        Person_Gender          `protobuf:"varint,13,opt,name=gender,proto3,enum=golang.Person_Gender" json:"gender,omitempty"`
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

func (x *Person) GetIdentification() isPerson_Identification {
	if x != nil {
		return x.Identification
	}
	return nil
}

func (x *Person) GetEmail() string {
	if x != nil {
		if x, ok := x.Identification.(*Person_Email); ok {
			return x.Email
		}
	}
	return ""
}

func (x *Person) GetPhone() string {
	if x != nil {
		if x, ok := x.Identification.(*Person_Phone); ok {
			return x.Phone
		}
	}
	return ""
}

func (x *Person) GetNicknames() []string {
	if x != nil {
		return x.Nicknames
	}
	return nil
}

func (x *Person) GetPets() []*Pet {
	if x != nil {
		return x.Pets
	}
	return nil
}

func (x *Person) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Person) GetFavorites() map[string]*Pet {
	if x != nil {
		return x.Favorites
	}
	return nil
}

func (x *Person) GetNoCheckDate() *timestamppb.Timestamp {
	if x != nil {
		return x.NoCheckDate
	}
	return nil
}

func (x *Person) GetGender() Person_Gender {
	if x != nil {
		return x.Gender
	}
	return Person_UNKNOWN
}

type isPerson_Identification interface {
	isPerson_Identification()
}

type Person_Email struct {
	Email string `protobuf:"bytes,6,opt,name=email,proto3,oneof"`
}

type Person_Phone struct {
	Phone string `protobuf:"bytes,7,opt,name=phone,proto3,oneof"`
}

func (*Person_Email) isPerson_Identification() {}

func (*Person_Phone) isPerson_Identification() {}

// second message in same file
type Pet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Kind          string                 `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
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

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// message without checks
type Group struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_person_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// nested message
type Person_Health struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Weight        int32                  `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	AvgHartRate   float64                `protobuf:"fixed64,2,opt,name=avg_hart_rate,json=avgHartRate,proto3" json:"avg_hart_rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person_Health) Reset() {
	*x = Person_Health{}
	mi := &file_person_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person_Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Health) ProtoMessage() {}

func (x *Person_Health) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[3]
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

func (x *Person_Health) GetAvgHartRate() float64 {
	if x != nil {
		return x.AvgHartRate
	}
	return 0
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x1a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x0f, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
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
	0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x69, 0x64,
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
	0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x7f, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x67, 0xa2, 0xdc, 0xb4, 0x4c,
	0x62, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20,
	0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x1a, 0x45, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73,
	0x28, 0x27, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x5f, 0x25,
	0x2b, 0x2d, 0x5d, 0x2b, 0x40, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e,
	0x2d, 0x5d, 0x2b, 0x2e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0x7b, 0x32, 0x2c, 0x7d,
	0x24, 0x27, 0x29, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x65, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4d, 0xa2, 0xdc, 0xb4,
	0x4c, 0x48, 0x12, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x1a, 0x32, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5b, 0x30, 0x2d,
	0x39, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x24, 0x27, 0x29, 0x48, 0x00, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0xc4, 0x01, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0xa5, 0x01, 0xa2, 0xdc, 0xb4, 0x4c, 0x60, 0x12,
	0x21, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x1a, 0x3b, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x61,
	0x6c, 0x6c, 0x28, 0x78, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x78, 0x29, 0x3e, 0x30, 0x29, 0xa2,
	0xdc, 0xb4, 0x4c, 0x3b, 0x12, 0x18, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x63,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x61,
	0x6c, 0x6c, 0x28, 0x78, 0x2c, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x78, 0x29, 0x3e, 0x30, 0x29, 0x52,
	0x09, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x04, 0x70, 0x65,
	0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x50, 0x65, 0x74, 0x42, 0x38, 0xa2, 0xdc, 0xb4, 0x4c, 0x33, 0x12, 0x1c, 0x61, 0x74,
	0x20, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x50, 0x65, 0x74, 0x20, 0x69,
	0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x1a, 0x13, 0x73, 0x69, 0x7a, 0x65,
	0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x52,
	0x04, 0x70, 0x65, 0x74, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x44, 0xa2, 0xdc, 0xb4, 0x4c,
	0x3f, 0x12, 0x22, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x1a, 0x19, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x29, 0x20, 0x3e, 0x20, 0x30,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x7f, 0x0a, 0x09,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x42,
	0xa2, 0xdc, 0xb4, 0x4c, 0x3d, 0x12, 0x21, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x20,
	0x6f, 0x6e, 0x65, 0x20, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x20, 0x69, 0x73, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x1a, 0x18, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x29, 0x20, 0x3e,
	0x20, 0x30, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a,
	0x0d, 0x6e, 0x6f, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x6e, 0x6f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0xc1, 0x01, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x4d, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x35, 0xa2, 0xdc, 0xb4, 0x4c, 0x30, 0x12, 0x1d,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x6b, 0x67, 0x20, 0x6d, 0x75, 0x73,
	0x74, 0x20, 0x62, 0x65, 0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x1a, 0x0f, 0x74,
	0x68, 0x69, 0x73, 0x2e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x68, 0x0a, 0x0d, 0x61, 0x76, 0x67, 0x5f, 0x68, 0x61,
	0x72, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x44, 0xa2,
	0xdc, 0xb4, 0x4c, 0x3f, 0x12, 0x23, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x20, 0x68, 0x65,
	0x61, 0x72, 0x74, 0x20, 0x72, 0x61, 0x74, 0x65, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65,
	0x20, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x1a, 0x18, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x61, 0x76, 0x67, 0x5f, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x20, 0x3e, 0x20,
	0x30, 0x2e, 0x30, 0x52, 0x0b, 0x61, 0x76, 0x67, 0x48, 0x61, 0x72, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x49, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x3a, 0xbf, 0x01, 0xa2, 0xdc, 0xb4, 0x4c, 0x58, 0x0a,
	0x10, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x22, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x20, 0x2b, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0xa2, 0xdc, 0xb4, 0x4c, 0x5d, 0x0a, 0x1e, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x6c,
	0x61, 0x72, 0x67, 0x65, 0x72, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20, 0x33, 0x30, 0x30, 0x1a, 0x19,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x20, 0x3c, 0x3d, 0x20, 0x33, 0x30, 0x30, 0x42, 0x10, 0x0a, 0x0e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x03,
	0x50, 0x65, 0x74, 0x12, 0x65, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x51, 0xa2, 0xdc, 0xb4, 0x4c, 0x4c, 0x0a, 0x04, 0x70, 0x65, 0x74, 0x31, 0x12, 0x1a,
	0x6f, 0x6e, 0x6c, 0x79, 0x20, 0x64, 0x6f, 0x67, 0x20, 0x6f, 0x72, 0x20, 0x63, 0x61, 0x74, 0x20,
	0x69, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x1a, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x6b, 0x69, 0x6e, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x63, 0x61, 0x74, 0x27, 0x20, 0x7c,
	0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6b, 0x69, 0x6e, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x27,
	0x64, 0x6f, 0x67, 0x27, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x44, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xa2, 0xdc, 0xb4, 0x4c, 0x2b, 0x12,
	0x14, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x63, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x1b, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x46, 0x0a,
	0x02, 0x68, 0x72, 0x42, 0x08, 0x48, 0x52, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a,
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

var file_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_person_proto_goTypes = []any{
	(Person_Gender)(0),            // 0: golang.Person.Gender
	(*Person)(nil),                // 1: golang.Person
	(*Pet)(nil),                   // 2: golang.Pet
	(*Group)(nil),                 // 3: golang.Group
	(*Person_Health)(nil),         // 4: golang.Person.Health
	nil,                           // 5: golang.Person.AttributesEntry
	nil,                           // 6: golang.Person.FavoritesEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_person_proto_depIdxs = []int32{
	7, // 0: golang.Person.birth_date:type_name -> google.protobuf.Timestamp
	4, // 1: golang.Person.health:type_name -> golang.Person.Health
	2, // 2: golang.Person.pets:type_name -> golang.Pet
	5, // 3: golang.Person.attributes:type_name -> golang.Person.AttributesEntry
	6, // 4: golang.Person.favorites:type_name -> golang.Person.FavoritesEntry
	7, // 5: golang.Person.no_check_date:type_name -> google.protobuf.Timestamp
	0, // 6: golang.Person.gender:type_name -> golang.Person.Gender
	2, // 7: golang.Person.FavoritesEntry.value:type_name -> golang.Pet
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	file_person_proto_msgTypes[0].OneofWrappers = []any{
		(*Person_Email)(nil),
		(*Person_Phone)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		EnumInfos:         file_person_proto_enumTypes,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}

package golang

import (
	"fmt"
	"os"

	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/shared"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type postBuilder struct{}

func (postBuilder) PostBuildCheckerData(f *protogen.Field, cd shared.CheckerData) shared.CheckerData {
	cd.IsSetConditionSource = isSetConditionSource(f)
	return cd
}

func (postBuilder) MessageIdent(m *protogen.Message) string {
	return string(m.GoIdent.GoName)
}

func isSetConditionSource(f *protogen.Field) string {
	if f.Desc.IsList() {
		return fmt.Sprintf("typedX.Get%s() != nil", f.GoName)
	}
	if f.Desc.Kind() == protoreflect.MessageKind {
		return fmt.Sprintf("typedX.Get%s() != nil", f.GoName)
	}
	if f.Desc.Kind() == protoreflect.StringKind {
		return fmt.Sprintf("typedX.Get%s() != \"\"", f.GoName)
	}
	if f.Desc.Kind() == protoreflect.Int32Kind ||
		f.Desc.Kind() == protoreflect.Int64Kind ||
		f.Desc.Kind() == protoreflect.Uint32Kind ||
		f.Desc.Kind() == protoreflect.Uint64Kind ||
		f.Desc.Kind() == protoreflect.Sint32Kind ||
		f.Desc.Kind() == protoreflect.Sint64Kind ||
		f.Desc.Kind() == protoreflect.Sfixed32Kind ||
		f.Desc.Kind() == protoreflect.Sfixed64Kind ||
		f.Desc.Kind() == protoreflect.Fixed32Kind ||
		f.Desc.Kind() == protoreflect.Fixed64Kind {
		return fmt.Sprintf("typedX.Get%s() != 0", f.GoName)
	}
	if f.Desc.Kind() == protoreflect.DoubleKind || f.Desc.Kind() == protoreflect.FloatKind {
		return fmt.Sprintf("typedX.Get%s() != 0.0", f.GoName)
	}
	if f.Desc.Kind() == protoreflect.BytesKind {
		return fmt.Sprintf("len(typedX.Get%s()) > 0", f.GoName)
	}
	if f.Desc.Kind() == protoreflect.BoolKind {
		return fmt.Sprintf("typedX.Get%s() != false", f.GoName)
	}
	fmt.Fprintln(os.Stderr, "unsupported field type", f.Desc.Kind())
	return ""
}

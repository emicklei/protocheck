package java

import (
	"fmt"
	"strings"

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
	return strings.ReplaceAll(string(m.GoIdent.GoName), "_", ".") // nested classes
}

func isSetConditionSource(f *protogen.Field) string {

	if f.Desc.HasDefault() {
		return "true;"
	}
	if f.Desc.IsList() || f.Desc.IsMap() {
		return "true;"
	}
	if f.Desc.Kind() == protoreflect.MessageKind {
		return fmt.Sprintf("((%s)x).has%s();", f.Parent.Desc.Name(), f.GoName)
	}
	return "true;"
}

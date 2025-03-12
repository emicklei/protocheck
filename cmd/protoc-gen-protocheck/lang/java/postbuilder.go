package java

import (
	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/shared"
	"google.golang.org/protobuf/compiler/protogen"
)

type postBuilder struct{}

func (postBuilder) PostBuildCheckerData(f *protogen.Field, cd shared.CheckerData) shared.CheckerData {
	return cd
}

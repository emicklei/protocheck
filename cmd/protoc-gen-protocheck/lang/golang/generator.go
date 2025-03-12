package golang

import (
	"log/slog"

	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/shared"
	"google.golang.org/protobuf/compiler/protogen"
)

func Process(p *protogen.Plugin, f *protogen.File) error {
	fd := shared.BuildFileData(f, postBuilder{})
	content, err := generate(fd)
	if err != nil {
		return err
	}
	outName := f.GeneratedFilenamePrefix + ".check.go"
	slog.Debug("writing", "file", outName)
	out := p.NewGeneratedFile(outName, f.GoImportPath)
	out.P(content)

	return nil
}

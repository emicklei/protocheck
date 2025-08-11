package java

import (
	"log/slog"
	"path"

	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/shared"
	"google.golang.org/protobuf/compiler/protogen"
)

func Process(p *protogen.Plugin, f *protogen.File) error {
	fd := shared.BuildFileData(f, postBuilder{}, *f.Proto.Options.JavaPackage)
	if len(fd.Messages) == 0 {
		slog.Info("no checkable messages, skipping file", "file", f.GeneratedFilenamePrefix)
		return nil
	}
	fd.JavaOuterClassname = *f.Proto.Options.JavaOuterClassname
	content, err := generate(fd)
	if err != nil {
		return err
	}
	outName := path.Join(*f.Proto.Options.JavaPackage, *f.Proto.Options.JavaOuterClassname+"Checkers.java")
	slog.Info("writing", "file", outName)
	out := p.NewGeneratedFile(outName, f.GoImportPath)
	out.P(content)

	return nil
}

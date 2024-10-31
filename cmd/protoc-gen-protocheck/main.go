package main

import (
	"flag"
	"log/slog"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(run)
}

// https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/
func run(p *protogen.Plugin) error {
	p.SupportedFeatures = p.SupportedFeatures + 1 // FEATURE_PROTO3_OPTIONAL
	for _, each := range p.Files {
		if !each.Generate {
			continue
		}
		if !importsCheck(each) {
			continue
		}
		slog.Info("destination",
			"import path", each.GoImportPath,
			"pkg", each.GoPackageName,
			"file", each.Desc.Path(),
			"param", p.Request.GetParameter())

		fd := buildFileData(each)
		content, err := generate(fd)
		if err != nil {
			return err
		}
		outName := strings.Replace(each.Desc.Path(), ".proto", ".check.go", -1)
		f := p.NewGeneratedFile(outName, each.GoImportPath)
		f.P(content)
	}
	return nil
}

// importsCheck returns whether the file imports check.proto.
func importsCheck(f *protogen.File) bool {
	for i := 0; i < f.Desc.Imports().Len(); i++ {
		entry := f.Desc.Imports().Get(i)
		if entry.Path() == "check.proto" {
			return true
		}
	}
	return false
}

func buildFileData(f *protogen.File) FileData {
	fd := FileData{
		PkgName: string(f.GoPackageName),
	}
	for _, each := range f.Messages {
		fd.Messages = append(fd.Messages, buildMessageData(each))
	}
	return fd
}

func buildMessageData(m *protogen.Message) MessageData {
	md := MessageData{
		LowercaseMessageName: strings.ToLower(string(m.Desc.Name())),
		MessageName:          string(m.Desc.Name()),
	}
	for _, each := range m.Fields {
		md.Checkers = append(md.Checkers, buildCheckerData(each))
	}
	return md
}

func buildCheckerData(f *protogen.Field) CheckerData {
	return CheckerData{
		FieldName: f.GoName,
	}
}

package golang

import (
	"strings"

	"github.com/emicklei/protocheck"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Process(p *protogen.Plugin, f *protogen.File) error {

	fd := buildFileData(f)
	content, err := generate(fd)
	if err != nil {
		return err
	}
	outName := strings.Replace(f.Desc.Path(), ".proto", ".check.go", -1)
	out := p.NewGeneratedFile(outName, f.GoImportPath)
	out.P(content)

	return nil
}

func buildFileData(f *protogen.File) FileData {
	fd := FileData{
		PkgName: string(f.GoPackageName),
	}
	for _, each := range f.Messages {
		md := buildMessageData(each)
		// file_person_proto_init()
		md.InitFuncName = "file_" +
			strings.ReplaceAll(string(f.Desc.Path()), ".", "_") +
			"_init"
		fd.Messages = append(fd.Messages, md)
	}
	return fd
}

func buildMessageData(m *protogen.Message) MessageData {
	md := MessageData{
		LowercaseMessageName: strings.ToLower(string(m.Desc.Name())),
		MessageName:          string(m.Desc.Name()),
	}
	for _, each := range m.Fields {
		cd, ok := buildCheckerData(each)
		if ok {
			md.Checkers = append(md.Checkers, cd)
		}
	}
	return md
}

func buildCheckerData(f *protogen.Field) (CheckerData, bool) {
	opts := f.Desc.Options()
	if opts == nil {
		return CheckerData{}, false
	}
	fops := opts.(*descriptorpb.FieldOptions)
	if !proto.HasExtension(fops, protocheck.E_Field) {
		return CheckerData{}, false
	}
	ext, ok := proto.GetExtension(fops, protocheck.E_Field).(*protocheck.Check)
	if !ok {
		return CheckerData{}, false
	}
	return CheckerData{
		FieldName: f.GoName,
		ID:        ext.Id,
		Fail:      ext.Fail,
		Expr:      ext.Cel,
	}, true
}

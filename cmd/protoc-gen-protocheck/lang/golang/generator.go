package golang

import (
	"fmt"
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
	// TODO nested messages
	cds, ok := buildMessageCheckerData(m)
	if ok {
		md.MessageCheckers = append(md.MessageCheckers, cds...)
	}
	for _, each := range m.Fields {
		cd, ok := buildFieldCheckerData(each)
		if ok {
			md.FieldCheckers = append(md.FieldCheckers, cd)
		}
	}
	return md
}
func buildMessageCheckerData(m *protogen.Message) ([]CheckerData, bool) {
	opts := m.Desc.Options()
	if opts == nil {
		return []CheckerData{}, false
	}
	mops := opts.(*descriptorpb.MessageOptions)
	if !proto.HasExtension(mops, protocheck.E_Message) {
		return []CheckerData{}, false
	}
	exts, ok := proto.GetExtension(mops, protocheck.E_Message).([]*protocheck.Check)
	if !ok {
		return []CheckerData{}, false
	}
	cds := []CheckerData{}
	for _, each := range exts {
		cds = append(cds, CheckerData{
			Comment: ifEmpty(each.Id, each.Cel),
			ID:      each.Id,
			Fail:    ifEmpty(each.Fail, each.Cel),
			Expr:    each.Cel,
		})
	}
	return cds, true
}

func buildFieldCheckerData(f *protogen.Field) (CheckerData, bool) {
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
		Comment:    f.GoName,
		FieldName:  f.GoName,
		IsOptional: f.Desc.HasOptionalKeyword(),
		ID:         ext.Id,
		Fail:       ifEmpty(ext.Fail, fmt.Sprintf("[%s] is false", ext.Cel)),
		Expr:       ext.Cel,
	}, true
}

func ifEmpty(content, alt string) string {
	if content == "" {
		return alt
	}
	return content
}

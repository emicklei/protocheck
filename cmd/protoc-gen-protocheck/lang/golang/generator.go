package golang

import (
	"fmt"
	"strings"

	"github.com/emicklei/protocheck"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common"
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
	init := "file_" +
		strings.ReplaceAll(string(f.Desc.Path()), ".", "_") +
		"_init"
	for _, each := range f.Messages {
		fd.Messages = addMessageDataTo(f, each, init, fd.Messages)
	}
	return fd
}

func addMessageDataTo(f *protogen.File, msg *protogen.Message, init string, list []MessageData) []MessageData {
	for _, each := range msg.Messages {
		if !each.Desc.IsMapEntry() {
			list = addMessageDataTo(f, each, init, list)
		}
	}
	md := buildMessageData(msg)
	md.InitFuncName = init
	return append(list, md)
}

func buildMessageData(m *protogen.Message) MessageData {
	md := MessageData{
		LowercaseMessageName: strings.ToLower(string(m.Desc.Name())),
		MessageName:          string(m.GoIdent.GoName),
		ObjectTypeName:       string(m.Desc.FullName()),
	}
	containers := []string{}
	cds, ok := buildMessageCheckerData(m)
	if ok {
		md.MessageCheckers = append(md.MessageCheckers, cds...)
	}
	for _, each := range m.Fields {
		// normal or oneof field
		cd, ok := buildFieldCheckerData(each)
		if ok {
			md.FieldCheckers = append(md.FieldCheckers, cd)
		}
		// repeated Message
		if each.Desc.IsList() && each.Desc.Message() != nil {
			containers = append(containers, string(each.GoName))
		}
		// map[?]Message
		if each.Desc.IsMap() && each.Desc.MapValue().Message() != nil {
			containers = append(containers, string(each.GoName))
		}
	}
	md.ContainerFieldNames = containers
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
		// this is needed to escape backslashes in the generated code
		each.Cel = strings.ReplaceAll(each.Cel, "\\", "\\\\")

		// validate the syntax at generation time
		if iss := parseCEL(each.Cel); iss != nil {
			panic(fmt.Sprintf("invalid CEL expression [%s] for message [%s], error [%v]", each.Cel, m.Desc.FullName(), iss))
		}
		cds = append(cds, CheckerData{
			Comment: ifEmpty(each.Id, each.Cel),
			ID:      each.Id,
			Fail:    ifEmpty(each.Fail, each.Cel),
			Expr:    each.Cel,
		})
	}
	return cds, true
}

func parseCEL(expr string) *cel.Issues {
	env, _ := cel.NewEnv()
	_, err := env.ParseSource(common.NewTextSource(expr))
	return err
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
	// this is needed to escape backslashes in the generated code
	ext.Cel = strings.ReplaceAll(ext.Cel, "\\", "\\\\")

	// validate the syntax at generation time
	if iss := parseCEL(ext.Cel); iss != nil {
		panic(fmt.Sprintf("invalid CEL expression [%s] for field [%s], error [%v]", ext.Cel, f.Desc.FullName(), iss))
	}
	oneOfType := ""  // not a field of oneof
	oneOfField := "" // not for a field of oneof
	// optionals are oneof like
	if !f.Desc.HasOptionalKeyword() && f.Oneof != nil {
		// find the matching field
		var found *protogen.Field
		for _, each := range f.Oneof.Fields {
			if each.Desc.Name() == f.Desc.Name() {
				found = each
			}
		}
		if found == nil {
			panic("should have found the matching one of field")
		}
		oneOfType = string(found.GoIdent.GoName)
		oneOfField = f.Oneof.GoName
	}
	return CheckerData{
		Comment:        f.GoName,
		FieldName:      f.GoName,
		IsOptional:     f.Desc.HasOptionalKeyword(),
		OneOfType:      oneOfType,
		OneOfFieldName: oneOfField,
		ID:             ext.Id,
		Fail:           ifEmpty(ext.Fail, fmt.Sprintf("[%s] is false", ext.Cel)),
		Expr:           ext.Cel,
	}, true
}

func ifEmpty(content, alt string) string {
	if content == "" {
		return alt
	}
	return content
}

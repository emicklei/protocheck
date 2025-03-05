package golang

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/emicklei/protocheck"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Process(p *protogen.Plugin, f *protogen.File) error {
	fd := buildFileData(f)
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

func buildFileData(f *protogen.File) FileData {
	fd := FileData{
		PkgName: string(f.GoPackageName),
	}
	init := "file_" +
		dotsAndSlashsToUnderscore(string(f.Desc.Path())) +
		"_init"
	for _, each := range f.Messages {
		fd.Messages = addMessageDataTo(f, each, init, fd.Messages)
	}
	return fd
}

func dotsAndSlashsToUnderscore(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, ".", "_"), "/", "_")
}

func addMessageDataTo(f *protogen.File, msg *protogen.Message, init string, list []MessageData) []MessageData {
	for _, each := range msg.Messages {
		if !each.Desc.IsMapEntry() {
			list = addMessageDataTo(f, each, init, list)
		}
	}
	md := buildMessageData(msg)
	if !md.HasChecker() {
		return list
	}
	md.InitFuncName = init
	return append(list, md)
}

func buildMessageData(m *protogen.Message) MessageData {
	md := MessageData{
		LowercaseMessageName: strings.ToLower(string(m.GoIdent.GoName)),
		MessageName:          string(m.GoIdent.GoName),
		ObjectTypeName:       string(m.Desc.FullName()),
	}
	containers := []string{}
	fieldsWithMessage := []string{}
	cds, ok := buildMessageCheckerData(m)
	if ok {
		md.MessageCheckers = append(md.MessageCheckers, cds...)
	}
	for _, each := range m.Fields {
		// normal or oneof fields
		cds, ok := buildFieldCheckerData(each)
		if ok {
			md.FieldCheckers = append(md.FieldCheckers, cds...)
		}
		// direct Message type, non repeated
		if messageHasChecks(each.Desc.Message()) && !each.Desc.IsList() && !each.Desc.IsMap() {
			fieldsWithMessage = append(fieldsWithMessage, string(each.GoName))
		}
		// repeated Message
		if each.Desc.IsList() && messageHasChecks(each.Desc.Message()) {
			containers = append(containers, string(each.GoName))
		}
		// map[?]Message
		if each.Desc.IsMap() && messageHasChecks(each.Desc.MapValue().Message()) {
			containers = append(containers, string(each.GoName))
		}
	}
	md.ContainerFieldNames = containers
	md.MessageFieldNames = fieldsWithMessage
	return md
}

func messageHasChecks(md protoreflect.MessageDescriptor) bool {
	if md == nil {
		return false
	}
	if messageHasMessageChecks(md) {
		return true
	}
	for i := 0; i < md.Fields().Len(); i++ {
		f := md.Fields().Get(i)
		if fieldHasChecks(f) {
			return true
		}
	}
	return false
}

func messageHasMessageChecks(md protoreflect.MessageDescriptor) bool {
	if md == nil {
		return false
	}
	opts := md.Options()
	if opts == nil {
		return false
	}
	mops, ok := opts.(*descriptorpb.MessageOptions)
	if !ok {
		return false
	}
	if !proto.HasExtension(mops, protocheck.E_Message) {
		return false
	}
	exts, ok := proto.GetExtension(mops, protocheck.E_Message).([]*protocheck.Check)
	return ok && len(exts) > 0
}

func fieldHasChecks(fd protoreflect.FieldDescriptor) bool {
	if fd == nil {
		return false
	}
	opts := fd.Options()
	if opts == nil {
		return false
	}
	fops, ok := opts.(*descriptorpb.FieldOptions)
	if !ok {
		return false
	}
	if !proto.HasExtension(fops, protocheck.E_Field) {
		return false
	}
	return proto.GetExtension(fops, protocheck.E_Field) != nil
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
			abort(fmt.Sprintf("invalid CEL expression [%s] for message [%s], error [%v]", each.Cel, m.Desc.FullName(), iss))
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

func buildFieldCheckerData(f *protogen.Field) (list []CheckerData, ok bool) {
	opts := f.Desc.Options()
	if opts == nil {
		return list, false
	}
	fops := opts.(*descriptorpb.FieldOptions)
	if !proto.HasExtension(fops, protocheck.E_Field) {
		return list, false
	}
	exts, ok := proto.GetExtension(fops, protocheck.E_Field).([]*protocheck.Check)
	if !ok {
		return list, false
	}
	for _, ext := range exts {
		// this is needed to escape backslashes in the generated code
		ext.Cel = strings.ReplaceAll(ext.Cel, "\\", "\\\\")

		// validate the syntax at generation time
		if iss := parseCEL(ext.Cel); iss != nil {
			abort(fmt.Sprintf("invalid CEL expression [%s] for field [%s], error [%v]", ext.Cel, f.Desc.FullName(), iss))
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
				abort("should have found the matching one of field")
			}
			oneOfType = string(found.GoIdent.GoName)
			oneOfField = f.Oneof.GoName
		}
		cd := CheckerData{
			Comment:              f.GoName,
			FieldName:            f.GoName,
			IsOptional:           f.Desc.HasOptionalKeyword(),
			OneOfType:            oneOfType,
			OneOfFieldName:       oneOfField,
			ID:                   ext.Id,
			Fail:                 ifEmpty(ext.Fail, fmt.Sprintf("[%s] is false", ext.Cel)),
			Expr:                 ext.Cel,
			IsSetConditionSource: isSetConditionSource(f),
		}
		list = append(list, cd)
	}
	return list, true
}

func ifEmpty(content, alt string) string {
	if content == "" {
		return alt
	}
	return content
}

func abort(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func isSetRequired(f *protogen.Field) bool {
	return true // !f.Desc.IsList() // TODO in editions2023 ? maybe always true
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

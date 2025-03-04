package golang

import (
	_ "embed"
	"strings"
	"text/template"
)

//go:embed check_template.txt
var checkDataTemplate string

type FileData struct {
	PkgName  string
	Messages []MessageData
}
type MessageData struct {
	InitFuncName         string
	LowercaseMessageName string
	MessageName          string
	ObjectTypeName       string
	MessageCheckers      []CheckerData
	FieldCheckers        []CheckerData
	MessageFieldNames    []string
	ContainerFieldNames  []string
	HasMethodsAvailable  bool // false for proto2,proto3, true for edition2023
}

func (md MessageData) HasChecker() bool {
	return len(md.FieldCheckers) > 0 || len(md.MessageCheckers) > 0
}

type CheckerData struct {
	Comment              string
	FieldName            string // empty for message level
	IsOptional           bool
	OneOfType            string // set for oneof
	OneOfFieldName       string // set for oneof
	ID                   string
	Expr                 string
	Fail                 string
	IsSetFuncRequired    bool
	IsSetConditionSource string
}

func generate(fd FileData) (string, error) {
	tmpl := template.Must(template.New("check").Parse(checkDataTemplate))
	buf := new(strings.Builder)
	if err := tmpl.Execute(buf, fd); err != nil {
		return "", err
	}
	return buf.String(), nil
}

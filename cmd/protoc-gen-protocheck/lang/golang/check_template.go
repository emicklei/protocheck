package golang

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/shared"
)

//go:embed check_template.txt
var checkDataTemplate string

func generate(fd shared.FileData) (string, error) {
	tmpl := template.Must(template.New("check").Parse(checkDataTemplate))
	buf := new(strings.Builder)
	if err := tmpl.Execute(buf, fd); err != nil {
		return "", err
	}
	return buf.String(), nil
}

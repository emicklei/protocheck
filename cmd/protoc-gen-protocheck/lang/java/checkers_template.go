package java

import (
	"strings"
	"text/template"

	_ "embed"

	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/shared"
)

//go:embed checkers_template.txt
var checkersDataTemplate string

func generate(fd shared.FileData) (string, error) {
	tmpl := template.Must(template.New("check").Parse(checkersDataTemplate))
	buf := new(strings.Builder)
	if err := tmpl.Execute(buf, fd); err != nil {
		return "", err
	}
	return buf.String(), nil
}

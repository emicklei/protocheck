package main

import (
	_ "embed"
	"html/template"
	"strings"
)

//go:embed check_template.txt
var checkDataTemplate string

type FileData struct {
	PkgName  string
	Messages []MessageData
}
type MessageData struct {
	LowercaseMessageName string
	MessageName          string
	Checkers             []CheckerData
}
type CheckerData struct {
	FieldName string
	ID        string
	Expr      string
	Fail      string
}

func generate(fd FileData) (string, error) {
	tmpl := template.Must(template.New("check").Parse(checkDataTemplate))
	buf := new(strings.Builder)
	if err := tmpl.Execute(buf, fd); err != nil {
		return "", err
	}
	return buf.String(), nil
}

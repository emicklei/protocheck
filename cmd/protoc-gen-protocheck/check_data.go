package main

import _ "embed"

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

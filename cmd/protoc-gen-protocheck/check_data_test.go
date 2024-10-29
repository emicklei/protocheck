package main

import (
	"os"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	tmpl := template.Must(template.New("check").Parse(checkDataTemplate))
	fd := FileData{
		PkgName: "test",
		Messages: []MessageData{{
			LowercaseMessageName: "person",
			MessageName:          "Person",
			Checkers: []CheckerData{
				{
					FieldName: "name",
					ID:        "name",
					Expr:      `size(this.name) > 3`,
					Fail:      "length must be greater than 3",
				},
			},
		}, {
			LowercaseMessageName: "thing",
			MessageName:          "Thing",
			Checkers: []CheckerData{
				{
					FieldName: "name",
					ID:        "name",
					Expr:      `size(this.name) > 3`,
					Fail:      "length must be greater than 3",
				},
			},
		}},
	}
	if err := tmpl.Execute(os.Stdout, fd); err != nil {
		t.Fatal(err)
	}
}

package main

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	fd := FileData{
		PkgName: "test",
		Messages: []MessageData{{
			LowercaseMessageName: "person",
			MessageName:          "Person",
			Checkers: []CheckerData{
				{
					FieldName: "name",
					ID:        "name-id",
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
					ID:        "name-id",
					Expr:      `size(this.name) > 3`,
					Fail:      "length must be greater than 3",
				},
			},
		}},
	}
	if _, err := generate(fd); err != nil {
		t.Fatal(err)
	}
}

package golang

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	fd := FileData{
		PkgName: "test",
		Messages: []MessageData{{
			InitFuncName:         "person_init",
			LowercaseMessageName: "person",
			MessageName:          "Person",
			MessageCheckers: []CheckerData{
				{
					ID:   "name-id",
					Expr: `size(this.name) > 3`,
					Fail: "length must be greater than 3",
				},
			},
		}, {
			LowercaseMessageName: "thing",
			MessageName:          "Thing",
			FieldCheckers: []CheckerData{
				{
					FieldName: "name",
					ID:        "name-id",
					Expr:      `size(this) > 3`,
					Fail:      "length must be greater than 3",
				},
			},
		}},
	}
	if src, err := generate(fd); err != nil {
		t.Fatal(err)
	} else {
		t.Log(src)
	}
}

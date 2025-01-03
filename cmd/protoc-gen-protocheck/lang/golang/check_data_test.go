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

func TestIsCELParseable(t *testing.T) {
	tests := []struct {
		expr string
		want bool
	}{
		{"size(this.name) > 3", true},
		{"size(this.name > 3", false},
		{"matches(this.name,'^[0-9]$')", true},
		{"this.name.matches('^[0-9]$')", true},
	}
	for _, tt := range tests {
		err := parseCEL(tt.expr)
		if got := err == nil; got != tt.want {
			t.Errorf("isCELParseable() = %v, want %v, error = %v", got, tt.want, err)
		}
	}
}

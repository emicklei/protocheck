package protocheck

import "testing"

func TestWithParentField(t *testing.T) {
	err := CheckError{}
	err = err.WithParentField("parent", "key")
	if err.Path != ".parent[key] " {
		t.Errorf("expected .parent[key], got %s", err.Path)
	}
	err2 := CheckError{}.WithPath("here")
	err2 = err2.WithParentField("parent", "key")
	if err2.Path != "parent[key].here" {
		t.Errorf("expected parent[key].here, got [%s]", err2.Path)
	}
}

func TestReflectIsSet(t *testing.T) {
	c := new(Check)
	if reflectIsSet(c, "Fail") {
		t.Fail()
	}
	c.Fail = "Failed"
	if !reflectIsSet(c, "Fail") {
		t.Fail()
	}
}

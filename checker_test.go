package protocheck

import "testing"

func TestWithParentField(t *testing.T) {
	err := CheckError{}
	err = err.WithParentField("parent", "key")
	if err.Path != ".parent[key]" {
		t.Errorf("expected .parent[key], got %s", err.Path)
	}
	err2 := CheckError{}.WithPath("here")
	err2 = err2.WithParentField("parent", "key")
	if err2.Path != "parent[key].here" {
		t.Errorf("expected parent[key].here, got [%s]", err2.Path)
	}
}

func TestEnabledChecker(t *testing.T) {
	ch := new(Checker).WithEnabledFunc(func(any) bool { return true })
	if ch.enabledFunc == nil {
		t.Fail()
	}
}

func TestForCoverage(t *testing.T) {
	file_check_proto_init()
	c := new(Check)
	c.GetCel()
	c.GetFail()
	c.GetId()
	c.Reset()
	s := ""
	c.Id = s
	c.Fail = s
	c.Cel = s
	c.GetCel()
	c.GetFail()
	c.GetId()
	c.Descriptor()
	_ = c.String()
	c.ProtoMessage()
	_ = c.ProtoReflect()
	var d *Check = nil
	d.GetCel()
	d.GetFail()
	d.GetId()
	_ = d.ProtoReflect()
}

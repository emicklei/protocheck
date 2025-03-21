package protocheck

import "testing"

func TestWithParentField(t *testing.T) {
	err := new(CheckError)
	err = err.WithParentField("parent", "key")
	if err.Path != "parent[key]" {
		t.Errorf("expected parent[key], got %s", err.Path)
	}
	err2 := new(CheckError).WithPath("here")
	err2 = err2.WithParentField("parent", "key")
	if err2.Path != "parent[key].here" {
		t.Errorf("expected parent[key].here, got [%s]", err2.Path)
	}
}

func TestWithPath(t *testing.T) {
	err := new(CheckError)
	err = err.WithPath("path1").WithPath("path2")
	if err.Path != "path2.path1" {
		t.Errorf("expected path2.path1, got %s", err.Path)
	}
}

func TestEnabledChecker(t *testing.T) {
	ch := new(Checker).WithEnabledFunc(func(any) bool { return true })
	if ch.enabledFunc == nil {
		t.Fail()
	}
}

func TestForCheckCoverage(t *testing.T) {
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
func TestForCheckErrorCoverage(t *testing.T) {
	c := new(CheckError)
	c.GetFail()
	c.GetId()
	c.GetPath()
	c.Reset()
	s := ""
	c.Id = s
	c.Fail = s
	c.Path = s
	c.GetFail()
	c.GetId()
	c.GetPath()
	c.Descriptor()
	_ = c.String()
	c.ProtoMessage()
	_ = c.ProtoReflect()
	var d *CheckError = nil
	d.GetPath()
	d.GetFail()
	d.GetId()
	_ = d.ProtoReflect()
}

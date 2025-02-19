package protocheck

import (
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func TestMakeProgram(t *testing.T) {
	env, _ := cel.NewEnv(
		cel.Types(new(Check)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("check.Check"))))
	_, err := MakeProgram(env, "this != null")
	if err != nil {
		t.Error(err)
	}
}
func TestMakeProgramBadExpression(t *testing.T) {
	env, _ := cel.NewEnv(
		cel.Types(new(Check)))
	_, err := MakeProgram(env, "this ==== null")
	if err == nil {
		t.Error("errror expected")
	}
}

func TestEmptyValidator(t *testing.T) {
	mv := NewMessageValidator([]Checker{}, []Checker{})
	ve := mv.Validate(new(Check))
	if len(ve) != 0 {
		t.Errorf("expected 0 errors, got %d", len(ve))
	}
}

func TestEvalChecker(t *testing.T) {
	expr := "size(this.cel) > 0"
	env, err := cel.NewEnv(
		cel.Types(new(Check)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("check.Check"))))
	if err != nil {
		t.Fatal(err)
	}
	prog, err := MakeProgram(env, expr)
	if err != nil {
		t.Fatal(err)
	}
	checker := NewChecker("1", "CEL cannot be empty", expr, "Cel", false, prog)
	{
		result := evalChecker(checker, map[string]interface{}{"this": nil})
		if len(result) != 1 {
			t.Errorf("expected 1 error, got %d", len(result))
		}
	}
	{
		result := evalChecker(checker, map[string]interface{}{"this": &Check{Cel: "1==1"}})
		if len(result) != 0 {
			t.Errorf("expected 0 error, got %d, %v", len(result), result[0])
		}
	}
	mv := NewMessageValidator([]Checker{checker}, []Checker{checker})
	ve := mv.Validate(new(Check))
	if len(ve) != 2 {
		t.Errorf("expected 2 errors, got %d", len(ve))
	}
	if ve[0].Id != "1" {
		t.Errorf("expected id 1, got %s", ve[0].Id)
	}
	if ve[0].Fail != "CEL cannot be empty" {
		t.Errorf("expected CEL cannot be empty, got %s", ve[0].Fail)
	}
	if ve.Error() != `2 errors occurred:
	* CEL cannot be empty
	* CEL cannot be empty
` {
		t.Errorf("expected CEL cannot be empty, got %s", ve.Error())
	}
}

func TestEvalCheckerNotSetButEnabledOptional(t *testing.T) {
	expr := "size(this.cel) > 0"
	env, err := cel.NewEnv(
		cel.Types(new(Check)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("check.Check"))))
	if err != nil {
		t.Fatal(err)
	}
	prog, err := MakeProgram(env, expr)
	if err != nil {
		t.Fatal(err)
	}
	checkerOptional1 := NewChecker("1", "true", expr, "Cel", true, prog)
	checkerOptional1.isSetFunc = func(this interface{}, fieldName string) bool {
		return false
	}
	checkerOptional1.enabledFunc = func(this interface{}) bool {
		return true
	}
	mv := NewMessageValidator([]Checker{}, []Checker{checkerOptional1})
	ve := mv.Validate(new(Check))
	if len(ve) != 0 {
		t.Errorf("expected 0 errors, got %d", len(ve))
	}
}

func TestEvalCheckerNotEnabled(t *testing.T) {
	expr := "size(this.cel) > 0"
	env, err := cel.NewEnv(
		cel.Types(new(Check)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("check.Check"))))
	if err != nil {
		t.Fatal(err)
	}
	prog, err := MakeProgram(env, expr)
	if err != nil {
		t.Fatal(err)
	}
	checkerOptional1 := NewChecker("1", "true", expr, "Cel", true, prog)
	checkerOptional1.enabledFunc = func(this interface{}) bool {
		return false
	}
	mv := NewMessageValidator([]Checker{}, []Checker{checkerOptional1})
	ve := mv.Validate(new(Check))
	if len(ve) != 0 {
		t.Errorf("expected 0 errors, got %d", len(ve))
	}
}

func TestEvalCheckerNonBoolean(t *testing.T) {
	expr := "1+2"
	env, _ := cel.NewEnv(
		cel.Types(new(Check)))
	prog, _ := MakeProgram(env, expr)
	checker := NewChecker("1", "", expr, "Cel", false, prog)
	{
		result := evalChecker(checker, map[string]interface{}{"this": &Check{Cel: "1+2"}})
		if len(result) == 0 {
			t.Errorf("expected 1 error")
		}
	}
}

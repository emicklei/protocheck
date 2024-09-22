package test

import (
	"fmt"
	"testing"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func TestCheckPerson(t *testing.T) {
	env, err := cel.NewEnv(
		cel.Types(&Person{}),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("check.Person")))) // use proto package here
	check(err)
	ast, iss := env.Compile(`size(this.name)`)
	check(iss.Err())
	prg, err := env.Program(ast)
	joe := &Person{Name: "Joe"}
	out, _, err := prg.Eval(map[string]interface{}{
		"this": joe,
	})
	fmt.Println(out.Value())
}

func TestCheckPersonDescription(t *testing.T) {
	env, err := cel.NewEnv(
		cel.Types(&Person{}),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("check.Person")))) // use proto package here
	check(err)
	ast, iss := env.Compile(`size(this.description) > 0`)
	check(iss.Err())
	prg, err := env.Program(ast)
	joe := &Person{Name: "Joe"}
	out, _, err := prg.Eval(map[string]interface{}{
		"this": joe,
	})
	fmt.Println(out.Value())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func TestCheckPerson2(t *testing.T) {
	joe := &Person{
		Name:      "Joe",
		BirthDate: &timestamppb.Timestamp{Seconds: timestamppb.Now().Seconds + 1}}
	t.Log(joe.Validate())
}

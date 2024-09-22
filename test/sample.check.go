package test

import (
	"github.com/emicklei/protocheck"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var PersonValidator protocheck.MessageValidator

func init() {
	// ensure proto_init (idempotent) is called first.
	file_sample_proto_init()

	env, err := cel.NewEnv(
		cel.Types(&Person{}),
		cel.Types(&timestamppb.Timestamp{}),
		cel.Declarations(
			decls.NewVar("now", decls.NewObjectType("google.protobuf.Timestamp")),
			decls.NewVar("this", decls.NewObjectType("check.Person"))))
	if err != nil {
		panic(err)
	}
	chs := []protocheck.Checker{}
	{
		ast, iss := env.Compile(`size(this.name) > 3`)
		if err = iss.Err(); err != nil {
			panic(err)
		}
		prg, err := env.Program(ast)
		if err != nil {
			panic(err)
		}
		chs = append(chs, protocheck.NewChecker("", "length must be greater than 3", prg))
	}
	PersonValidator = protocheck.NewMessageValidator(chs)
}

func (x *Person) Validate() error {
	return PersonValidator.Validate(x)
}

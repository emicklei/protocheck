package test

import (
	sync "sync"

	"github.com/emicklei/protocheck"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	personValidator     protocheck.MessageValidator
	personValidatorOnce sync.Once
)

func file_person_check_proto_init() {
	// ensure proto_init (idempotent) is called first.
	file_person_proto_init()

	{ // check.Person
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
		{ // name
			ast, iss := env.Compile(`size(this.name) > 3`)
			if err = iss.Err(); err != nil {
				panic(err)
			}
			prg, err := env.Program(ast)
			if err != nil {
				panic(err)
			}
			chs = append(chs, protocheck.NewChecker("", "length must be greater than 3", `size(this.name) > 3`, prg))
		}
		personValidator = protocheck.NewMessageValidator(chs)
	}
}

// Validate checks the validity of the Person message.
// Returns an error if the validation fails.
func (x *Person) Validate() error {
	personValidatorOnce.Do(file_person_check_proto_init)
	return personValidator.Validate(x)
}

// Code generated by protoc-gen-protocheck. DO NOT EDIT.

package golang

import (
	"sync"

	"github.com/emicklei/protocheck"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	personValidator     protocheck.MessageValidator
	personValidatorOnce sync.Once
	petValidator        protocheck.MessageValidator
	petValidatorOnce    sync.Once
)

func file_person_check_proto_init() {
	// ensure proto_init (idempotent) is called first.
	file_person_proto_init()
	{ // golang.Person
		env, err := cel.NewEnv(
			cel.Types(new(Person)),
			cel.Types(new(timestamppb.Timestamp)),
			cel.Declarations(
				decls.NewVar("now", decls.NewObjectType("google.protobuf.Timestamp")),
				decls.NewVar("this", decls.NewObjectType("golang.Person"))))
		if err != nil {
			panic(err)
		}
		chs := []protocheck.Checker{}
		{ // Name
			ast, iss := env.Compile(`size(this.name) > 1`)
			if err = iss.Err(); err != nil {
				panic(err)
			}
			prg, err := env.Program(ast)
			if err != nil {
				panic(err)
			}
			chs = append(chs, protocheck.NewChecker("", "", `size(this.name) > 1`, prg))
		}
		personValidator = protocheck.NewMessageValidator(chs)
	}
	{ // golang.Person
		env, err := cel.NewEnv(
			cel.Types(new(Person)),
			cel.Types(new(timestamppb.Timestamp)),
			cel.Declarations(
				decls.NewVar("now", decls.NewObjectType("google.protobuf.Timestamp")),
				decls.NewVar("this", decls.NewObjectType("golang.Person"))))
		if err != nil {
			panic(err)
		}
		chs := []protocheck.Checker{}
		{ // Description
			ast, iss := env.Compile(`size(this.description) > 0`)
			if err = iss.Err(); err != nil {
				panic(err)
			}
			prg, err := env.Program(ast)
			if err != nil {
				panic(err)
			}
			chs = append(chs, protocheck.NewChecker("", "description cannot be empty", `size(this.description) > 0`, prg))
		}
		personValidator = protocheck.NewMessageValidator(chs)
	}
	{ // golang.Person
		env, err := cel.NewEnv(
			cel.Types(new(Person)),
			cel.Types(new(timestamppb.Timestamp)),
			cel.Declarations(
				decls.NewVar("now", decls.NewObjectType("google.protobuf.Timestamp")),
				decls.NewVar("this", decls.NewObjectType("golang.Person"))))
		if err != nil {
			panic(err)
		}
		chs := []protocheck.Checker{}
		{ // BirthDate
			ast, iss := env.Compile(`this.birth_date < now`)
			if err = iss.Err(); err != nil {
				panic(err)
			}
			prg, err := env.Program(ast)
			if err != nil {
				panic(err)
			}
			chs = append(chs, protocheck.NewChecker("check_birth_date", "", `this.birth_date < now`, prg))
		}
		personValidator = protocheck.NewMessageValidator(chs)
	}
}

// Validate checks the validity of the Person message.
// Returns an error if the validation fails.
func (x *Person) Validate(opts ...cel.EnvOption) error {
	personValidatorOnce.Do(file_person_check_proto_init)
	return personValidator.Validate(x, opts...)
}
func file_pet_check_proto_init() {
	// ensure proto_init (idempotent) is called first.
	file_person_proto_init()
	{ // golang.Pet
		env, err := cel.NewEnv(
			cel.Types(new(Pet)),
			cel.Types(new(timestamppb.Timestamp)),
			cel.Declarations(
				decls.NewVar("now", decls.NewObjectType("google.protobuf.Timestamp")),
				decls.NewVar("this", decls.NewObjectType("golang.Pet"))))
		if err != nil {
			panic(err)
		}
		chs := []protocheck.Checker{}
		{ // Kind
			ast, iss := env.Compile(`this.kind == 'cat' || this.kind == 'dog' `)
			if err = iss.Err(); err != nil {
				panic(err)
			}
			prg, err := env.Program(ast)
			if err != nil {
				panic(err)
			}
			chs = append(chs, protocheck.NewChecker("", "", `this.kind == 'cat' || this.kind == 'dog' `, prg))
		}
		petValidator = protocheck.NewMessageValidator(chs)
	}
}

// Validate checks the validity of the Pet message.
// Returns an error if the validation fails.
func (x *Pet) Validate(opts ...cel.EnvOption) error {
	petValidatorOnce.Do(file_pet_check_proto_init)
	return petValidator.Validate(x, opts...)
}

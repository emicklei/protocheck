package protocheck

import (
	"fmt"
	"strings"

	"github.com/google/cel-go/cel"
)

// CheckError captures a failed check.
type CheckError struct {
	Id   string
	Fail string
	Err  error
}

// ValidationError is a collection of CheckError.
type ValidationError []CheckError

// Error implements error
func (v ValidationError) Error() string {
	b := new(strings.Builder)
	plural := ""
	if len(v) > 1 {
		plural = "s"
	}
	fmt.Fprintf(b, "%d error%s occurred:\n", len(v), plural)
	for _, each := range v {
		fail := each.Fail
		if fail == "" {
			fail = "id=" + each.Id
		}
		if each.Err == nil {
			fmt.Fprintf(b, "\t* %s\n", fail)
		} else {
			fmt.Fprintf(b, "\t* %s:%s\n", fail, each.Err.Error())
		}
	}
	return b.String()
}

// Checker performs one check using a CEL program.
type Checker struct {
	check   *Check
	program cel.Program
}

// NewChecker creates a Checker
func NewChecker(id string, message string, cel string, program cel.Program) Checker {
	return Checker{
		check: &Check{
			Id:   id,
			Fail: message,
			Cel:  cel,
		},
		program: program,
	}
}

type MessageValidator struct {
	checkers []Checker
}

func NewMessageValidator(chs []Checker) MessageValidator {
	return MessageValidator{checkers: chs}
}

func (m MessageValidator) Validate(this any) ValidationError {
	var result ValidationError
	env := map[string]interface{}{"this": this}
	for _, each := range m.checkers {
		out, _, err := each.program.Eval(env)
		if err != nil {
			result = append(result, CheckError{Id: each.check.Id, Fail: each.check.Fail, Err: err})
		} else {
			valid, ok := out.Value().(bool)
			if !ok {
				err := fmt.Errorf("expected check expression [%s] to return a boolean, got [%T]", each.check.Cel, out.Value())
				result = append(result, CheckError{Id: each.check.Id, Fail: each.check.Fail, Err: err})
			} else {
				if !valid {
					result = append(result, CheckError{Id: each.check.Id, Fail: each.check.Fail})
				}
			}
		}
	}
	return result

}

// MakeProgram creates a cel Program for a given expression and environment.
func MakeProgram(env *cel.Env, expression string) (cel.Program, error) {
	ast, iss := env.Compile(expression)
	if err := iss.Err(); err != nil {
		return nil, fmt.Errorf("protocheck: failed to compile CEL expression: %w", err)
	}
	prg, err := env.Program(ast)
	if err != nil {
		return nil, fmt.Errorf("protocheck: failed to make CEL program: %w", err)
	}
	return prg, nil
}

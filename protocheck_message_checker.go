package protocheck

import (
	"errors"
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/hashicorp/go-multierror"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ValidationError struct {
	Id         string
	Message    string
	Expression string
}

func (v ValidationError) Error() string {
	if v.Id == "" {
		return v.Message
	}
	return v.Message
}

type Check struct {
	Id         string
	Message    string
	Expression string
}

type Checker struct {
	id      string
	message string
	program cel.Program
}

func NewChecker(id string, message string, program cel.Program) Checker {
	return Checker{
		id:      id,
		message: message,
		program: program,
	}
}

func (c Checker) String() string {
	return fmt.Sprintf("id: %s, message: %s", c.id, c.message)
}

type MessageValidator struct {
	checkers []Checker
}

func NewMessageValidator(chs []Checker) MessageValidator {
	return MessageValidator{checkers: chs}
}

func (m MessageValidator) Validate(this any) error {
	var result error
	env := map[string]interface{}{
		"this": this,
		"now":  timestamppb.Now(),
	}
	for _, each := range m.checkers {
		out, _, err := each.program.Eval(env)
		if err != nil {
			result = multierror.Append(result, err)
		} else {
			valid, ok := out.Value().(bool)
			if !ok {
				result = multierror.Append(result,
					fmt.Errorf("expected check expression [%s] to return a boolean, got [%T]", each, out.Value()))
			} else {
				if !valid {
					result = multierror.Append(result, errors.New(each.message))
				}
			}
		}
	}
	return result

}

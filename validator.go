package protocheck

import (
	"fmt"

	"github.com/google/cel-go/cel"
)

type ValidationOption byte

const AllFields ValidationOption = 1
const FieldsSetOnly ValidationOption = 2

// MessageValidator holds a collection of checkers to validate a message.
type MessageValidator struct {
	fieldCheckers   []Checker
	messageCheckers []Checker
}

// Validator is an interface that can be implemented by a message to validate itself.
type Validator interface {
	Validate() error
}

// NewMessageValidator creates a MessageValidator using two collections of checkers
func NewMessageValidator(messageCheckers, fieldCheckers []Checker) MessageValidator {
	return MessageValidator{fieldCheckers: fieldCheckers, messageCheckers: messageCheckers}
}

// Validate runs all message and field checkers with the message.
// Always returns a ValidationError which can be empty (no failed checks)
// With options you can control what to validations to skip.
func (m MessageValidator) Validate(this any, options ...ValidationOption) (result ValidationError) {
	optionsMask := optionsMask(options)
	env := map[string]interface{}{"this": this}
	for _, each := range m.messageCheckers {
		result = append(result, evalChecker(each, env)...)
	}
	// are we done?
	if len(m.fieldCheckers) == 0 {
		return result
	}
	for _, each := range m.fieldCheckers {
		// check is enabled
		if each.enabledFunc != nil && !each.enabledFunc(this) {
			continue
		}
		if each.isSetFunc != nil {
			// is it set?
			isSet := each.isSetFunc(this, each.fieldName)
			// skip unset?
			if !isSet && hasOption(optionsMask, FieldsSetOnly) {
				continue
			}
			if !isSet && each.isOptional {
				continue
			}
		}
		result = append(result, evalChecker(each, env)...)
	}
	return

}

func optionsMask(options []ValidationOption) byte {
	var mask byte
	for _, each := range options {
		mask |= byte(each)
	}
	return mask
}

func hasOption(mask byte, option ValidationOption) bool {
	return mask&byte(option) != 0
}

func evalChecker(each Checker, env map[string]interface{}) (result []CheckError) {
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
	return
}

// MakeProgram creates a cel Program for a given expression and environment.
func MakeProgram(env *cel.Env, expression string) (cel.Program, error) {
	ast, iss := env.Compile(expression)
	if err := iss.Err(); err != nil {
		return nil, fmt.Errorf("protocheck: failed to compile CEL expression: %w", err)
	}
	return env.Program(ast)
}

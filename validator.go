package protocheck

import (
	"fmt"
	"slices"

	"github.com/google/cel-go/cel"
)

type ValidationOption byte

// Only if a field is set (non-zero value for proto3) then validate its content.
const FieldsSetOnly ValidationOption = 1

// MessageValidator holds a collection of checkers to validate a message.
type MessageValidator struct {
	fieldCheckers   []Checker
	messageCheckers []Checker
}

// Validator is an interface that can be implemented by a message to validate itself.
type Validator interface {
	Validate(options ...ValidationOption) ValidationError
}

// NewMessageValidator creates a MessageValidator using two collections of checkers
func NewMessageValidator(messageCheckers, fieldCheckers []Checker) MessageValidator {
	return MessageValidator{fieldCheckers: fieldCheckers, messageCheckers: messageCheckers}
}

// Validate runs all message and field checkers with the message.
// Always returns a ValidationError (error) which can be empty (no failed checks)
// With options you can control what to validations to skip.
func (m MessageValidator) Validate(this any, options ...ValidationOption) (result ValidationError) {
	env := map[string]any{"this": this}
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
		// is it set
		isSet := each.isSetFunc(this, each.fieldName)
		// skip unset?
		if !isSet && slices.Contains(options, FieldsSetOnly) {
			continue
		}
		// if not set but optional, skip
		if !isSet && each.isOptional {
			continue
		}
		result = append(result, evalChecker(each, env)...)
	}
	return
}

func evalChecker(each Checker, env map[string]interface{}) (result []*CheckError) {
	out, _, err := each.program.Eval(env)
	if err != nil {
		result = append(result, &CheckError{Id: each.check.Id, Fail: each.check.Fail, Path: each.fieldName})
	} else {
		valid, ok := out.Value().(bool)
		if !ok {
			result = append(result, &CheckError{
				Id:   "exception",
				Fail: fmt.Sprintf("invalid validation expression detected for %s", each.fieldName),
				Path: each.fieldName},
			)
		}
		if !valid {
			result = append(result, &CheckError{Id: each.check.Id, Fail: each.check.Fail, Path: each.fieldName})
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

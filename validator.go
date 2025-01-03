package protocheck

import (
	"fmt"
	reflect "reflect"

	"github.com/google/cel-go/cel"
)

type MessageValidator struct {
	fieldCheckers   []Checker
	messageCheckers []Checker
}

type Validator interface {
	Validate() ValidationError
}

func NewMessageValidator(messageCheckers, fieldCheckers []Checker) MessageValidator {
	return MessageValidator{fieldCheckers: fieldCheckers, messageCheckers: messageCheckers}
}

func (m MessageValidator) Validate(this any) ValidationError {
	var result ValidationError
	env := map[string]interface{}{"this": this}
	for _, each := range m.messageCheckers {
		result = append(result, evalChecker(each, env)...)
	}
	rv := reflect.ValueOf(this)
	for _, each := range m.fieldCheckers {
		methodName := "Get" + each.fieldName
		method := rv.MethodByName(methodName)
		if !method.IsValid() {
			panic(fmt.Errorf("method [%s] not found in %T", methodName, this))
		}
		getValue := method.Call([]reflect.Value{})
		if len(getValue) == 0 && each.isOptional {
			continue
		}
		result = append(result, evalChecker(each, env)...)
	}
	return result

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
	prg, err := env.Program(ast)
	if err != nil {
		return nil, fmt.Errorf("protocheck: failed to make CEL program: %w", err)
	}
	return prg, nil
}

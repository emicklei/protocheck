package protocheck

import (
	"fmt"
	reflect "reflect"
	"strings"

	"github.com/google/cel-go/cel"
)

// CheckError captures a failed check.
type CheckError struct {
	Path string // path to the field that failed, if empty then the field is part of the root message
	Id   string // optional id of the check that failed
	Fail string // optional message to display
	Err  error  // set if the check failed due to an error
}

// WithParentField returns a new CheckError with the parent field prepended to the path.
func (c CheckError) WithParentField(parent string, key any) CheckError {
	path := fmt.Sprintf("%s[%v] ", parent, key)
	if c.Path == "" {
		c.Path = "." + path
	} else {
		c.Path = strings.TrimRight(path, " ") + "." + c.Path
	}
	return c
}

// WithPath returns a new CheckError with the path set.
func (c CheckError) WithPath(path string) CheckError {
	c.Path = path
	return c
}

// Checker performs one check using a CEL program.
type Checker struct {
	check       *Check
	program     cel.Program
	fieldName   string // for field level checks
	isOptional  bool
	enabledFunc func(message any) bool // if set then call it to see if check is enabled
	isSetFunc   func(message any, fieldName string) bool
}

// NewChecker creates a Checker
func NewChecker(id string, fail string, cel string, fieldName string, isOptional bool, program cel.Program) Checker {
	return Checker{
		check: &Check{
			Id:   id,
			Fail: fail,
			Cel:  cel,
		},
		fieldName:  fieldName,
		isOptional: isOptional,
		program:    program,
		isSetFunc:  reflectIsSet,
	}
}

// for proto3
func reflectIsSet(message any, fieldName string) bool {
	rv := reflect.ValueOf(message)
	methodName := "Get" + fieldName
	method := rv.MethodByName(methodName)
	if !method.IsValid() {
		return false
	}
	getValue := method.Call([]reflect.Value{})
	return len(getValue) != 0
}

func (c Checker) WithEnabledFunc(enabledFunc func(any) bool) Checker {
	c.enabledFunc = enabledFunc
	return c
}
func (c Checker) WithIsSetFunc(fun func(message any, fieldName string) bool) Checker {
	c.isSetFunc = fun
	return c
}

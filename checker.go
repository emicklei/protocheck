package protocheck

import (
	"fmt"
	"strings"

	"github.com/google/cel-go/cel"
)

// WithParentField returns a new CheckError with the parent field prepended to the path.
func (c *CheckError) WithParentField(parent string, key any) *CheckError {
	path := fmt.Sprintf("%s[%v]", parent, key)
	if c.Path == "" {
		path = "." + path
	} else {
		path = strings.TrimRight(path, " ") + "." + c.Path
	}
	return &CheckError{Id: c.Id, Fail: c.Fail, Path: path}
}

// WithPath returns a new CheckError with the path set.
func (c *CheckError) WithPath(path string) *CheckError {
	return &CheckError{Id: c.Id, Fail: c.Fail, Path: path}
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
	}
}

func (c Checker) WithEnabledFunc(enabledFunc func(any) bool) Checker {
	c.enabledFunc = enabledFunc
	return c
}
func (c Checker) WithIsSetFunc(fun func(message any, fieldName string) bool) Checker {
	c.isSetFunc = fun
	return c
}

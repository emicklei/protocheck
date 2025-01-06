package protocheck

import (
	"fmt"
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

// WithPath returns a new CheckError with the parent field prepended to the path.
func (c CheckError) WithPath(path string) CheckError {
	c.Path = path
	return c
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
			fmt.Fprintf(b, "\t* %s%s\n", each.Path, fail)
		} else {
			fmt.Fprintf(b, "\t* %s%s err=%s\n", each.Path, fail, each.Err.Error())
		}
	}
	return b.String()
}

// Checker performs one check using a CEL program.
type Checker struct {
	check       *Check
	program     cel.Program
	fieldName   string // for field level checks
	isOptional  bool
	enabledFunc func(any) bool // if set then call it to see if check is enabled
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

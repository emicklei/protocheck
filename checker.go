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
			fmt.Fprintf(b, "\t* %s err=%s\n", fail, each.Err.Error())
		}
	}
	return b.String()
}

// Checker performs one check using a CEL program.
type Checker struct {
	check      *Check
	program    cel.Program
	fieldName  string // for field level checks
	isOptional bool
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

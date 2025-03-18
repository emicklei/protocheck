package protocheck

import (
	"fmt"
	"strings"
)

// ValidationError is a collection of *CheckError.
type ValidationError []*CheckError

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
		path := each.Path
		if path != "" {
			path += " "
		}
		fmt.Fprintf(b, "\t* %s%s\n", path, fail)
	}
	return b.String()
}

var emptyValidationError = ValidationError{} // TODO is this an optimization?

// AsValidationError converts an error or nil to a valid ValidationError (so could be empty).
func AsValidationError(err error) ValidationError {
	if err == nil {
		return emptyValidationError
	}
	if ve, ok := err.(ValidationError); ok {
		return ve
	}
	return emptyValidationError
}

package protocheck

import (
	"errors"
	"testing"
)

func TestAsValidationError(t *testing.T) {
	var nilerr error = nil
	if got := AsValidationError(nilerr); len(got) != 0 {
		t.Errorf("expected emptyValidationError, got %v", got)
	} else {
		msg := got.Error()
		if msg != "0 error occurred:\n" {
			t.Errorf("expected msg, got [%v]", msg)
		}
	}
	var emptyerr error = ValidationError{}
	if got := AsValidationError(emptyerr); len(got) != 0 {
		t.Errorf("expected emptyValidationError, got %v", got)
	} else {
		msg := got.Error()
		if msg != "0 error occurred:\n" {
			t.Errorf("expected msg, got [%v]", msg)
		}
	}
	var ve error = ValidationError{CheckError{}}
	if got := AsValidationError(ve); len(got) == 0 {
		t.Errorf("expected emptyValidationError, got %v", got)
	} else {
		msg := got.Error()
		if msg != "1 error occurred:\n\t* id=\n" {
			t.Errorf("expected msg, got [%v]", msg)
		}
	}
	othererr := errors.New("other")
	if got := AsValidationError(othererr); len(got) != 0 {
		t.Errorf("expected emptyValidationError, got %v", got)
	} else {
		msg := got.Error()
		if msg != "0 error occurred:\n" {
			t.Errorf("expected msg, got [%v]", msg)
		}
	}

}

func TestValidationErrorError(t *testing.T) {
	ve := ValidationError{CheckError{
		Id:  "1",
		Err: errors.New("error"),
	}}
	msg := ve.Error()
	if msg != "1 error occurred:\n\t* id=1 err=error\n" {
		t.Errorf("expected msg, got [%v]", msg)
	}
}

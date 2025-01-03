package golang

import (
	"testing"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func TestCheckPerson(t *testing.T) {
	empty := ""
	joe := &Person{
		Name:       empty,
		MiddleName: &empty,
		BirthDate:  &timestamppb.Timestamp{Seconds: 1},
		Health:     &Person_Health{Weight: 0}}
	err := joe.Validate()
	t.Log(err)

	for _, each := range err {
		t.Logf("%#v", each)
	}
}

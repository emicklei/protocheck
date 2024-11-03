package golang

import (
	"testing"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func TestCheckPerson(t *testing.T) {
	joe := &Person{
		Name:      "",
		BirthDate: &timestamppb.Timestamp{Seconds: 1}}
	err := joe.Validate()
	t.Log(err)
}

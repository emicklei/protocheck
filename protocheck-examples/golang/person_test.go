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
		Name:      "Joe",
		BirthDate: &timestamppb.Timestamp{Seconds: timestamppb.Now().Seconds + 1}}
	t.Log(joe.Validate())
}

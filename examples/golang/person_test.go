package golang

import (
	"testing"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func createValidPerson() *Person {
	notempty := "test"
	p := &Person{
		Name:       notempty,
		Surname:    notempty,
		MiddleName: &notempty,
		BirthDate:  &timestamppb.Timestamp{Seconds: 200 * 365 * 60 * 60 * 60},
		Health:     &Person_Health{Weight: 0}}
	p.Identification = &Person_Email{Email: "a.b@here.com"}
	p.Attributes = map[string]string{notempty: notempty}
	p.Favorites = map[string]*Pet{notempty: {Kind: "cat", Name: notempty}}
	p.Nicknames = []string{notempty}
	return p
}

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
func TestCheckPersonMapWithPet(t *testing.T) {
	p := createValidPerson()
	err := p.Validate()
	if len(err) != 0 {
		t.Fatal(err)
	}
}

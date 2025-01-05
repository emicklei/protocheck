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
	p.Pets = append(p.Pets, &Pet{Kind: "dog", Name: notempty})
	p.Attributes = map[string]string{notempty: notempty}
	p.Favorites = map[string]*Pet{notempty: {Kind: "cat", Name: notempty}}
	p.Nicknames = []string{notempty}
	return p
}

func TestCheckPersonMapWithInvalidPet(t *testing.T) {
	p := createValidPerson()
	p.Favorites["test"].Kind = "spider"
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Id, "pet1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidEmail(t *testing.T) {
	p := createValidPerson()
	p.Identification = &Person_Email{Email: "invalid"}
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Id, "email"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

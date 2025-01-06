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
		Health:     &Person_Health{Weight: 1}}
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
func TestCheckPersonInvalidPet(t *testing.T) {
	p := createValidPerson()
	p.Pets[0].Kind = "spider"
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Id, "pet1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidPetMapValue(t *testing.T) {
	p := createValidPerson()
	p.Favorites["test"].Kind = "spider"
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Id, "pet1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	t.Log(err)
}
func TestCheckPersonInvalidName(t *testing.T) {
	p := createValidPerson()
	p.Name = "?"
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Fail, "name must be longer than 1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidBirthdate(t *testing.T) {
	p := createValidPerson()
	p.BirthDate = &timestamppb.Timestamp{Seconds: 0}
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Fail, "[this.birth_date.getFullYear() > 2000] is false"; got != want {
		t.Errorf("got {%[1]v:%[1]T} want [%[2]v:%[2]T]", got, want)
	}
}

func TestCheckPersonInvalidHealth(t *testing.T) {
	p := createValidPerson()
	p.Health.Weight = -1
	err := p.Validate()
	if len(err) != 1 {
		t.Fatal(err)
	}
	if got, want := err[0].Fail, "weight in kg must be positive"; got != want {
		t.Errorf("got {%[1]v:%[1]T} want [%[2]v:%[2]T]", got, want)
	}
}

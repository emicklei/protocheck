package golang

import (
	"testing"

	"github.com/emicklei/protocheck"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func createValidPerson() *Person {
	notempty := "test"
	p := &Person{
		Name:       notempty,
		Surname:    notempty,
		MiddleName: &notempty,
		BirthDate:  &timestamppb.Timestamp{Seconds: 200 * 365 * 60 * 60 * 60},
		Health:     &Person_Health{Weight: 1, AvgHartRate: 60.0}}
	p.Identification = &Person_Email{Email: "a.b@here.com"}
	p.Pets = append(p.Pets, &Pet{Kind: "dog", Name: notempty})
	p.Attributes = map[string]string{notempty: notempty}
	p.Favorites = map[string]*Pet{notempty: {Kind: "cat", Name: notempty}}
	p.Nicknames = []string{notempty}
	return p
}

// go test -bench=. -count=5 -run=^# -cpuprofile=cpu.prof -memprofile=mem.prof
func BenchmarkValidation(t *testing.B) {
	p := createValidPerson()
	for i := 0; i < t.N; i++ {
		_ = p.Validate()
	}
}

func TestCheckPersonMapWithInvalidPet(t *testing.T) {
	p := createValidPerson()
	p.Favorites["test"].Kind = "spider"

	ve := p.Validate()
	if len(ve) == 0 {
		t.Fatal("expected error")
	}
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Id, "pet1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidEmail(t *testing.T) {
	p := createValidPerson()
	p.Identification = &Person_Email{Email: "invalid"}
	ve := p.Validate()
	if len(ve) == 0 {
		t.Fatal("expected error")
	}
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Id, "email"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	if got, want := ve[0].Path, "Email"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidPet(t *testing.T) {
	p := createValidPerson()
	p.Pets[0].Kind = "spider"
	ve := p.Validate()
	if len(ve) == 0 {
		t.Fatal("expected error")
	}
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Id, "pet1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidPetMapValue(t *testing.T) {
	p := createValidPerson()
	p.Favorites["test"].Kind = "spider"
	ve := p.Validate()
	if len(ve) == 0 {
		t.Fatal("expected error")
	}
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Id, "pet1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
	t.Log(ve)
}
func TestCheckPersonInvalidName(t *testing.T) {
	p := createValidPerson()
	p.Name = "?"
	ve := p.Validate()
	if got, want := ve[0].Fail, "name must be longer than 1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidBirthdate(t *testing.T) {
	p := createValidPerson()
	p.BirthDate = &timestamppb.Timestamp{Seconds: 0}
	ve := p.Validate()
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Fail, "[this.birth_date.getFullYear() > 2000] is false"; got != want {
		t.Errorf("got {%[1]v:%[1]T} want [%[2]v:%[2]T]", got, want)
	}
}

func TestCheckPersonInvalidHealthWeight(t *testing.T) {
	p := createValidPerson()
	p.Health.Weight = -1
	ve := p.Validate()
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Fail, "weight in kg must be positive"; got != want {
		t.Errorf("got {%[1]v:%[1]T} want [%[2]v:%[2]T]", got, want)
	}
	if got, want := ve[0].Path, "Health.Weight"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}

func TestCheckPersonInvalidHealthHeartRate(t *testing.T) {
	p := createValidPerson()
	p.Health.AvgHartRate = -1
	ve := p.Validate()
	if len(ve) != 1 {
		t.Fatal(ve)
	}
	if got, want := ve[0].Fail, "average heart rate must be positive"; got != want {
		t.Errorf("got {%[1]v:%[1]T} want [%[2]v:%[2]T]", got, want)
	}
}

func TestValidateSetFieldsOn(t *testing.T) {
	notempty := "test"
	p := &Person{
		Name:       notempty,
		Nicknames:  nil, // mark not set
		Pets:       nil, // mark not set
		Attributes: nil, // mark not set
		Favorites:  nil} // mark not set
	err := p.Validate(protocheck.FieldsSetOnly)
	if err != nil {
		t.Error("unexpected", err)
	}
}

package golang

import (
	"testing"

	"github.com/emicklei/protocheck"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func createValidPerson() *Person {
	notempty := "test"
	p := new(Person)
	p.SetName(notempty)
	p.SetSurname(notempty)
	p.SetMiddleName(notempty)
	p.SetBirthDate(&timestamppb.Timestamp{Seconds: 200 * 365 * 60 * 60 * 60})

	h := new(Person_Health)
	h.SetWeight(1)
	h.SetAvgHartRate(60.0)
	p.SetHealth(h)

	// either email or phone
	p.SetEmail("a.b@here.com")

	pet := new(Pet)
	pet.SetKind("dog")
	pet.SetName(notempty)
	pets := []*Pet{pet}
	p.SetPets(pets)

	p.SetAttributes(map[string]string{notempty: notempty})

	cat := new(Pet)
	cat.SetKind("cat")
	cat.SetName(notempty)
	p.SetFavorites(map[string]*Pet{notempty: cat})

	p.SetNicknames([]string{notempty})
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
	p.GetFavorites()["test"].SetKind("spider")

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
	p.SetEmail("invalid")
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
	p.GetPets()[0].SetKind("spider")
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
	p.GetFavorites()["test"].SetKind("spider")
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
	p.SetName("?")
	ve := p.Validate()
	if got, want := ve[0].Fail, "name must be longer than 1"; got != want {
		t.Errorf("got [%[1]v:%[1]T] want [%[2]v:%[2]T]", got, want)
	}
}
func TestCheckPersonInvalidBirthdate(t *testing.T) {
	p := createValidPerson()
	p.SetBirthDate(&timestamppb.Timestamp{Seconds: 0})
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
	p.GetHealth().SetWeight(-1)
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
	p.GetHealth().SetAvgHartRate(-1)
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
	p := new(Person)
	p.SetName(notempty)

	err := p.Validate(protocheck.FieldsSetOnly)
	if err != nil {
		t.Error("unexpected", err)
	}
}

// Code generated by protoc-gen-protocheck. DO NOT EDIT.

package golang

import (
	"sync"
	"fmt"
	"log/slog"

	"github.com/emicklei/protocheck"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

var (
	person_healthValidator     protocheck.MessageValidator
	person_healthValidatorOnce sync.Once
	personValidator            protocheck.MessageValidator
	personValidatorOnce        sync.Once
	petValidator               protocheck.MessageValidator
	petValidatorOnce           sync.Once
)

func file_person_health_check_proto_init() error {
	// ensure proto_init (idempotent) is called first.
	file_person_proto_init()
	env, err := cel.NewEnv(
		cel.Types(new(Person_Health)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("golang.Person.Health"))))
	if err != nil {
		return fmt.Errorf("cel.NewEnv failed: %w", err)
	}
	messageCheckers := []protocheck.Checker{}
	fieldCheckers := []protocheck.Checker{}
	{ // Weight
		expr := `this.weight > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "weight in kg must be positive", expr, "Weight", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person_Health)
				return typedX.GetWeight() != 0
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // AvgHartRate
		expr := `this.avg_hart_rate > 0.0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "average heart rate must be positive", expr, "AvgHartRate", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person_Health)
				return typedX.GetAvgHartRate() != 0.0
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	person_healthValidator = protocheck.NewMessageValidator(messageCheckers, fieldCheckers)
	return nil
}

// Validate checks the validity of the Person_Health message.
// Returns a non-empty error if the validation fails, nil otherwise.
func (x *Person_Health) Validate(options ...protocheck.ValidationOption) (list protocheck.ValidationError) {
	if x == nil {
		return nil
	}
	person_healthValidatorOnce.Do(func() {
		if err := file_person_health_check_proto_init(); err != nil {
			slog.Error("checkers initialization failed", "err", err)
		}
	})
	if verrs := person_healthValidator.Validate(x, options...); verrs != nil {
		list = append(list, verrs...)
	}
	if len(list) == 0 {
		return nil
	}
	return list
}
func file_person_check_proto_init() error {
	// ensure proto_init (idempotent) is called first.
	file_person_proto_init()
	env, err := cel.NewEnv(
		cel.Types(new(Person)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("golang.Person"))))
	if err != nil {
		return fmt.Errorf("cel.NewEnv failed: %w", err)
	}
	messageCheckers := []protocheck.Checker{}
	{ // person_invariant
		if prg, err := protocheck.MakeProgram(env, `size(this.name + this.surname) > 0`); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("person_invariant", "name and surname cannot be empty", `size(this.name + this.surname) > 0`, "", false, prg)
			messageCheckers = append(messageCheckers, ch)
		}
	}
	{ // person_health_weight_invariant
		if prg, err := protocheck.MakeProgram(env, `this.health.weight <= 300`); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("person_health_weight_invariant", "weight cannot be larger than 300", `this.health.weight <= 300`, "", false, prg)
			messageCheckers = append(messageCheckers, ch)
		}
	}
	fieldCheckers := []protocheck.Checker{}
	{ // Name
		expr := `size(this.name) > 1`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "name must be longer than 1", expr, "Name", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetName() != ""
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // MiddleName
		expr := `size(this.middle_name) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "middle name (if set) cannot be empty", expr, "MiddleName", true, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetMiddleName() != ""
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Surname
		expr := `size(this.surname) > 1`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "surname must be longer than 1", expr, "Surname", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetSurname() != ""
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // BirthDate
		expr := `this.birth_date.getFullYear() > 2000`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("check_birth_date", "[this.birth_date.getFullYear() > 2000] is false", expr, "BirthDate", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetBirthDate() != nil
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Email
		expr := `this.email.matches('^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$')`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("email", "email is not valid", expr, "Email", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.HasEmail()
			})
			ch = ch.WithEnabledFunc(func(x any) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.HasIdentification()
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Phone
		expr := `this.phone.matches('^[0-9]{3}-[0-9]{3}-[0-9]{4}$')`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "phone is not valid", expr, "Phone", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.HasPhone()
			})
			ch = ch.WithEnabledFunc(func(x any) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.HasIdentification()
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Nicknames
		expr := `size(this.nicknames) > 0 && this.nicknames.all(x,size(x)>0)`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one nickname is required", expr, "Nicknames", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetNicknames() != nil
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Nicknames
		expr := `this.nicknames.all(x,size(x)>0)`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "nickname cannot be empty", expr, "Nicknames", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetNicknames() != nil
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Pets
		expr := `size(this.pets) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one Pet is required", expr, "Pets", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetPets() != nil
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Attributes
		expr := `size(this.attributes) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one attribute is required", expr, "Attributes", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetAttributes() != nil
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Favorites
		expr := `size(this.favorites) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one favorite is required", expr, "Favorites", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				return typedX.GetFavorites() != nil
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	personValidator = protocheck.NewMessageValidator(messageCheckers, fieldCheckers)
	return nil
}

// Validate checks the validity of the Person message.
// Returns a non-empty error if the validation fails, nil otherwise.
func (x *Person) Validate(options ...protocheck.ValidationOption) (list protocheck.ValidationError) {
	if x == nil {
		return nil
	}
	personValidatorOnce.Do(func() {
		if err := file_person_check_proto_init(); err != nil {
			slog.Error("checkers initialization failed", "err", err)
		}
	})
	if verrs := personValidator.Validate(x, options...); verrs != nil {
		list = append(list, verrs...)
	}
	// Health
	if verrs := x.GetHealth().Validate(options...); verrs != nil {
		for _, each := range verrs {
			list = append(list, each.WithPath("Health"))
		}
	}
	// Pets
	for key, msg := range x.GetPets() {
		if verrs := msg.Validate(options...); verrs != nil {
			for _, each := range verrs {
				list = append(list, each.WithParentField("Pets", key))
			}
		}
	}
	// Favorites
	for key, msg := range x.GetFavorites() {
		if verrs := msg.Validate(options...); verrs != nil {
			for _, each := range verrs {
				list = append(list, each.WithParentField("Favorites", key))
			}
		}
	}
	if len(list) == 0 {
		return nil
	}
	return list
}
func file_pet_check_proto_init() error {
	// ensure proto_init (idempotent) is called first.
	file_person_proto_init()
	env, err := cel.NewEnv(
		cel.Types(new(Pet)),
		cel.Declarations(
			decls.NewVar("this", decls.NewObjectType("golang.Pet"))))
	if err != nil {
		return fmt.Errorf("cel.NewEnv failed: %w", err)
	}
	messageCheckers := []protocheck.Checker{}
	fieldCheckers := []protocheck.Checker{}
	{ // Kind
		expr := `this.kind == 'cat' || this.kind == 'dog'`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("pet1", "only dog or cat is allowed", expr, "Kind", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Pet)
				return typedX.GetKind() != ""
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Name
		expr := `size(this.name) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "name cannot be empty", expr, "Name", false, prg)
			ch = ch.WithIsSetFunc(func(x any, _ string) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Pet)
				return typedX.GetName() != ""
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	petValidator = protocheck.NewMessageValidator(messageCheckers, fieldCheckers)
	return nil
}

// Validate checks the validity of the Pet message.
// Returns a non-empty error if the validation fails, nil otherwise.
func (x *Pet) Validate(options ...protocheck.ValidationOption) (list protocheck.ValidationError) {
	if x == nil {
		return nil
	}
	petValidatorOnce.Do(func() {
		if err := file_pet_check_proto_init(); err != nil {
			slog.Error("checkers initialization failed", "err", err)
		}
	})
	if verrs := petValidator.Validate(x, options...); verrs != nil {
		list = append(list, verrs...)
	}
	if len(list) == 0 {
		return nil
	}
	return list
}

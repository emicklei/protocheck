// Code generated by protoc-gen-protocheck. DO NOT EDIT.

package golang

import (
	"fmt"
	"log/slog"
	"sync"

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
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	person_healthValidator = protocheck.NewMessageValidator(messageCheckers, fieldCheckers)
	return nil
}

// Validate checks the validity of the Person_Health message.
// Returns an error if the validation fails.
func (x *Person_Health) Validate() protocheck.ValidationError {
	if x == nil {
		return protocheck.ValidationError{}
	}
	person_healthValidatorOnce.Do(func() {
		if err := file_person_health_check_proto_init(); err != nil {
			slog.Error("checkers initialization failed", "err", err)
		}
	})
	ve := person_healthValidator.Validate(x)
	return ve
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
	fieldCheckers := []protocheck.Checker{}
	{ // Name
		expr := `size(this.name) > 1`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "name must be longer than 1", expr, "Name", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // MiddleName
		expr := `size(this.middle_name) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "middle name (if set) cannot be empty", expr, "MiddleName", true, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Surname
		expr := `size(this.surname) > 1`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "surname must be longer than 1", expr, "Surname", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // BirthDate
		expr := `this.birth_date.getFullYear() > 2000`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("check_birth_date", "[this.birth_date.getFullYear() > 2000] is false", expr, "BirthDate", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Email
		expr := `this.email.matches('^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$')`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("email", "email is not valid", expr, "Email", false, prg)
			ch = ch.WithEnabledFunc(func(x any) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				_, ok := typedX.Identification.(*Person_Email)
				return ok
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
			ch = ch.WithEnabledFunc(func(x any) bool {
				if x == nil {
					return false
				}
				typedX, _ := x.(*Person)
				_, ok := typedX.Identification.(*Person_Phone)
				return ok
			})
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Nicknames
		expr := `size(this.nicknames) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one nickname is required", expr, "Nicknames", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Pets
		expr := `size(this.pets) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one Pet is required", expr, "Pets", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Attributes
		expr := `size(this.attributes) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one attribute is required", expr, "Attributes", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Favorites
		expr := `size(this.favorites) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "at least one favorite is required", expr, "Favorites", false, prg)
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	personValidator = protocheck.NewMessageValidator(messageCheckers, fieldCheckers)
	return nil
}

// Validate checks the validity of the Person message.
// Returns an error if the validation fails.
func (x *Person) Validate() protocheck.ValidationError {
	if x == nil {
		return protocheck.ValidationError{}
	}
	personValidatorOnce.Do(func() {
		if err := file_person_check_proto_init(); err != nil {
			slog.Error("checkers initialization failed", "err", err)
		}
	})
	ve := personValidator.Validate(x)
	for _, nve := range x.GetHealth().Validate() { // Health
		ve = append(ve, nve.WithPath(".Health"))
	}
	for key, msg := range x.GetPets() { // Pets
		for _, nve := range msg.Validate() {
			ve = append(ve, nve.WithParentField("Pets", key))
		}
	}
	for key, msg := range x.GetFavorites() { // Favorites
		for _, nve := range msg.Validate() {
			ve = append(ve, nve.WithParentField("Favorites", key))
		}
	}
	return ve
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
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	{ // Name
		expr := `size(this.name) > 0`
		if prg, err := protocheck.MakeProgram(env, expr); err != nil {
			return fmt.Errorf("protocheck.MakeProgram failed: %w", err)
		} else {
			ch := protocheck.NewChecker("", "name cannot be empty", expr, "Name", false, prg)
			// ch = ch.WithIsSetFunc(func(message any, fieldName string) bool {
			// 	return message.(*Pet).GetName() != ""
			// })
			fieldCheckers = append(fieldCheckers, ch)
		}
	}
	petValidator = protocheck.NewMessageValidator(messageCheckers, fieldCheckers)
	return nil
}

// Validate checks the validity of the Pet message.
// Returns an error if the validation fails.
func (x *Pet) Validate() protocheck.ValidationError {
	if x == nil {
		return protocheck.ValidationError{}
	}
	petValidatorOnce.Do(func() {
		if err := file_pet_check_proto_init(); err != nil {
			slog.Error("checkers initialization failed", "err", err)
		}
	})
	ve := petValidator.Validate(x)
	return ve
}

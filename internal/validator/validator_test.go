package validator

import "testing"

func TestNewValidator(t *testing.T) {
	v := NewValidator()

	if v.Message != "" {
		t.Fatalf("expected empty; Validator.Message")
	}

	if !v.IsValid() {
		t.Fatalf("expected valid; Validator.IsValid()")
	}
}

func TestValidatorCheck(t *testing.T) {
	v := NewValidator()

	v.Check(false, "error message")

	if v.IsValid() {
		t.Fatalf("expected not valid; Validator.IsValid()")
	}

	if v.Message != "error message" {
		t.Fatalf("expected %v got %v; Validator.Message", "error message", v.Message)
	}

	v.Check(true, "this should not overwrite")

	if v.IsValid() {
		t.Fatalf("expected to be unchanged; Validator.IsValid")
	}

	if v.Message == "this should not overwrite" {
		t.Fatalf("expected to be unchanged; Validator.Message")
	}
}

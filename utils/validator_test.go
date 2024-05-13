package util

import (
	"testing"
)

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,gt=7"`
}

// TestValidator tests the Validator function for various struct validation scenarios
func TestValidator(t *testing.T) {
	t.Run("Test Empty Email and Password", func(t *testing.T) {
		login := Login{Email: "", Password: ""}
		err := Validator(login)
		if err == nil {
			t.Errorf("Expected error for empty fields, got nil")
		}
	})

	t.Run("Test Invalid Email Format", func(t *testing.T) {
		login := Login{Email: "invalid-email", Password: "password123"}
		err := Validator(login)
		if err == nil {
			t.Errorf("Expected error for invalid email format, got nil")
		}
	})

	t.Run("Test Short Password", func(t *testing.T) {
		login := Login{Email: "email@example.com", Password: "short"}
		err := Validator(login)
		if err == nil {
			t.Errorf("Expected error for short password, got nil")
		}
	})

	t.Run("Test Valid Input", func(t *testing.T) {
		login := Login{Email: "email@example.com", Password: "password123"}
		err := Validator(login)
		if err != nil {
			t.Errorf("Expected no error for valid input, got %v", err)
		}
	})
}

package util

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if hashedPassword == "" {
		t.Errorf("Hashed password is empty")
	}
}

func TestComparePassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	// Test with correct password
	err = ComparePassword(hashedPassword, password)
	if err != nil {
		t.Errorf("Password comparison should succeed, got error: %v", err)
	}

	// Test with incorrect password
	err = ComparePassword(hashedPassword, "wrongpassword")
	if err == nil {
		t.Errorf("Password comparison should fail with incorrect password")
	}
}

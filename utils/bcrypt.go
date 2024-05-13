package util

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		// Log the error with details
		logrus.Errorf("Hashing password failed: %v", err)
		// Return wrapped error and empty string for hashed password
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(result), nil
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

package util

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	data := map[string]interface{}{
		"ID":    "123",
		"Email": "test@example.com",
	}
	secretKey := "JWT_SECRET"
	expirationTime := 15 * time.Minute

	token, err := Sign(data, secretKey, expirationTime)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// Verify the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(GodotEnv(secretKey)), nil
	})
	assert.Nil(t, err)
	assert.True(t, parsedToken.Valid)

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, data["ID"], claims["ID"])
	assert.Equal(t, data["Email"], claims["Email"])
	assert.True(t, claims["exp"].(float64) > float64(time.Now().Unix()))
}

func TestVerifyToken(t *testing.T) {
	data := map[string]interface{}{
		"ID":    "123",
		"Email": "test@example.com",
	}
	secretKey := "JWT_SECRET"
	expirationTime := 15 * time.Minute

	token, _ := Sign(data, secretKey, expirationTime)

	// Test valid token
	validToken, err := VerifyToken(token, secretKey)
	assert.Nil(t, err)
	assert.NotNil(t, validToken)

	// Test invalid token
	invalidToken, err := VerifyToken("invalidtoken", secretKey)
	assert.NotNil(t, err)
	assert.Nil(t, invalidToken)
}

func TestDecodeToken(t *testing.T) {
	data := map[string]interface{}{
		"ID":    "123",
		"Email": "test@example.com",
	}
	secretKey := "JWT_SECRET"
	expirationTime := 15 * time.Minute

	tokenString, _ := Sign(data, secretKey, expirationTime)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GodotEnv(secretKey)), nil
	})

	decodedToken, err := DecodeToken(token)
	assert.Nil(t, err)
	assert.Equal(t, "123", decodedToken.Claims.ID)
	assert.Equal(t, "test@example.com", decodedToken.Claims.Email)
}

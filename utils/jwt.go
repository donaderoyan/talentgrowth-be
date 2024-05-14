package util

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID            string
	Email         string
	ExpiredAt     time.Time
	Authorization bool
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(data map[string]interface{}, secretPublicKeyEnvName string, expiredAt time.Duration) (string, error) {

	expirationTime := time.Now().Add(expiredAt * time.Minute).Unix()

	jwtSecretKey := GodotEnv(secretPublicKeyEnvName)

	claims := jwt.MapClaims{}
	claims["exp"] = expirationTime
	claims["authorization"] = true

	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return "", err
	}

	return signedToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	if tokenHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	splitToken := strings.Split(tokenHeader, "Bearer ")
	if len(splitToken) != 2 {
		return nil, fmt.Errorf("malformed token")
	}
	accessToken := strings.TrimSpace(splitToken[1])

	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error("Failed to parse token: ", err.Error())
		return nil, err
	}

	return token, nil
}

func VerifyToken(accessToken, SecretPublicKeyEnvName string) (*jwt.Token, error) {
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error("Failed to parse token: ", err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) (AccessToken, error) {
	var token AccessToken
	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if !ok {
		return token, fmt.Errorf("cannot assert claims from token")
	}

	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(claimsBytes, &token.Claims)
	if err != nil {
		return token, err
	}

	return token, nil
}

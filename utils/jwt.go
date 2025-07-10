package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateAccessToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"role":   strings.ToLower(role),
		"type":   "access",
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func GenerateRefreshToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"role":   strings.ToLower(role),
		"type":   "refresh",
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		log.Printf("ParseToken: failed")
		return nil, nil, err
	}
	return token, claims, nil
}
func IsTokenType(claims jwt.MapClaims, expectedType string) bool {
	if typ, ok := claims["type"].(string); ok {
		return typ == expectedType
	}
	return false
}

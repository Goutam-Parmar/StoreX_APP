package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateAccessToken generates a short-lived access token with userId + role
func GenerateAccessToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"role":   role,
		"type":   "access",
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// GenerateRefreshToken generates a long-lived refresh token with userId + role
func GenerateRefreshToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"role":   role,
		"type":   "refresh",
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken validates and parses JWT, returns claims
func ParseToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	log.Println("✅ ParseToken: starting parse")

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		log.Printf("❌ ParseToken: failed: %v\n", err)
		return nil, nil, err
	}

	log.Println("✅ ParseToken: success")
	return token, claims, nil
}

// IsTokenType checks JWT claim type
func IsTokenType(claims jwt.MapClaims, expectedType string) bool {
	if typ, ok := claims["type"].(string); ok {
		return typ == expectedType
	}
	return false
}

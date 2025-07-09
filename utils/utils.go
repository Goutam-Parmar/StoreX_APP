package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type AuthClaims struct {
	Role   string
	UserID string
}

// ExtractAuthClaims parses JWT from Authorization header and extracts userId + role
func ExtractAuthClaims(header string) (*AuthClaims, error) {
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return nil, fmt.Errorf("missing or malformed Authorization header")
	}

	tokenStr := strings.TrimPrefix(header, "Bearer ")
	_, claims, err := ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	log.Println("✅ ExtractAuthClaims: token parsed")

	role, ok2 := claims["role"].(string)
	userID, ok3 := claims["userId"].(string)

	if !ok2 || !ok3 {
		return nil, fmt.Errorf("invalid claims in token")
	}

	log.Printf("✅ ExtractAuthClaims: userID=%s, role=%s", userID, role)

	return &AuthClaims{
		UserID: userID,
		Role:   role,
	}, nil
}

// Tx is a helper for transaction rollback
func Tx(tx *sql.Tx, err *error) {
	if *err != nil {
		_ = tx.Rollback()
	}
}

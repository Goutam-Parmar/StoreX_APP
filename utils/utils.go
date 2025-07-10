package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

type AuthClaims struct {
	Role   string
	UserID string
}

func ExtractAuthClaims(header string) (*AuthClaims, error) {
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return nil, fmt.Errorf("missing or malformed Authorization header")
	}

	tokenStr := strings.TrimPrefix(header, "Bearer ")
	_, claims, err := ParseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	role, okRole := claims["role"].(string)
	userID, okID := claims["userId"].(string)

	if !okRole || !okID {
		return nil, fmt.Errorf("invalid claims in token")
	}
	role = strings.ToLower(role)

	return &AuthClaims{
		UserID: userID,
		Role:   role,
	}, nil
}

// transaction rollback
func Tx(tx *sql.Tx, err *error) {
	if *err != nil {
		_ = tx.Rollback()
	}
}

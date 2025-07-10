package auth

import (
	"StoreXApp/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const AuthClaimsKey contextKey = "authClaims"

// AuthMiddleware parses JWT and attaches claims to context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole checks role from context claims
func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			claims, ok := r.Context().Value(AuthClaimsKey).(*utils.AuthClaims)
			if !ok {
				http.Error(w, "Forbidden: no auth claims", http.StatusForbidden)
				return
			}
			if strings.ToLower(claims.Role) != strings.ToLower(role) {
				http.Error(w, "Forbidden: role mismatch", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

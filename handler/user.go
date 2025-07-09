package handler

import (
	"StoreXApp/auth"
	"StoreXApp/dbhelper"
	"StoreXApp/models"
	"StoreXApp/utils"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func RegisterUserByEmpManager() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// ✅ Use AuthClaims from context
		claims, ok := r.Context().Value(auth.AuthClaimsKey).(*utils.AuthClaims)
		if !ok {
			http.Error(w, "missing or invalid auth claims", http.StatusUnauthorized)
			return
		}

		if claims.Role != "EmpManager" {
			http.Error(w, "you are not EmpManager", http.StatusForbidden)
			return
		}

		var req models.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		req.CreatedBy = claims.UserID
		req.Email = strings.TrimSpace(strings.ToLower(req.Email))

		if err := dbhelper.CheckRegisterCredentials(&req, w); err != nil {
			return
		}

		resp, err := dbhelper.RegisterUserBYAdmin(&req)
		if err != nil {
			http.Error(w, "failed to register user", http.StatusInternalServerError)
			return
		}

		resp = &models.RegisterUserResponse{
			ID:             req.UserID,
			Fname:          req.Fname,
			Lname:          req.Lname,
			Email:          req.Email,
			PhoneNo:        req.PhoneNo,
			Role:           req.Role,
			EmpType:        req.EmpType,
			CreatedBy:      req.CreatedBy,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func RegisterUserByAdmin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// ✅ Use AuthClaims from context
		claims, ok := r.Context().Value(auth.AuthClaimsKey).(*utils.AuthClaims)
		if !ok {
			http.Error(w, "missing or invalid auth claims", http.StatusUnauthorized)
			return
		}

		if claims.Role != "admin" {
			http.Error(w, "you are not admin", http.StatusForbidden)
			return
		}

		var req models.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}

		req.CreatedBy = claims.UserID
		req.Email = strings.TrimSpace(strings.ToLower(req.Email))

		if err := dbhelper.CheckRegisterCredentials(&req, w); err != nil {
			return
		}

		resp, err := dbhelper.RegisterUserBYAdmin(&req)
		if err != nil {
			http.Error(w, "failed to register user", http.StatusInternalServerError)
			return
		}

		resp = &models.RegisterUserResponse{
			ID:             req.UserID,
			Fname:          req.Fname,
			Lname:          req.Lname,
			Email:          req.Email,
			PhoneNo:        req.PhoneNo,
			Role:           req.Role,
			EmpType:        req.EmpType,
			CreatedBy:      req.CreatedBy,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

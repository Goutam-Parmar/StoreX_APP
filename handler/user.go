package handler

import (
	"StoreXApp/auth"
	"StoreXApp/database"
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
		claims, ok := r.Context().Value(auth.AuthClaimsKey).(*utils.AuthClaims)
		if !ok {
			http.Error(w, "missing or invalid auth claims", http.StatusUnauthorized)
			return
		}
		if strings.ToLower(claims.Role) != "employeemanager" {
			http.Error(w, "you are not EmployeeManager", http.StatusForbidden)
			return
		}
		var req models.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}
		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		req.CreatedBy = claims.UserID
		if err := dbhelper.CheckRegisterCredentials(&req, w); err != nil {
			return
		}
		dbResp, err := dbhelper.RegisterUserBYAdmin(&req)
		if err != nil {
			http.Error(w, "failed to register user", http.StatusInternalServerError)
			return
		}
		req.UserID = dbResp.ID
		resp := &models.RegisterUserResponse{
			ID:             dbResp.ID,
			Fname:          req.Fname,
			Lname:          req.Lname,
			Email:          dbResp.Email,
			PhoneNo:        req.PhoneNo,
			Role:           dbResp.Role,
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

		dbResp, err := dbhelper.RegisterUserBYAdmin(&req)
		if err != nil {
			http.Error(w, "failed to register user", http.StatusInternalServerError)
			return
		}

		req.UserID = dbResp.ID

		resp := &models.RegisterUserResponse{
			ID:             dbResp.ID,
			Fname:          req.Fname,
			Lname:          req.Lname,
			Email:          dbResp.Email,
			PhoneNo:        req.PhoneNo,
			Role:           dbResp.Role,
			EmpType:        req.EmpType,
			CreatedBy:      req.CreatedBy,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if err := dbhelper.CheckLoginCredentials(&req, w); err != nil {
			return
		}
		accessToken, err := utils.GenerateAccessToken(req.UserID, req.Role)
		if err != nil {
			http.Error(w, "Could not generate access token", http.StatusInternalServerError)
			return
		}
		refreshToken, err := utils.GenerateRefreshToken(req.UserID, req.Role)
		if err != nil {
			http.Error(w, "Could not generate refresh token", http.StatusInternalServerError)
			return
		}
		res := models.LoginResponse{
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func RegisterSelf() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var req models.SelfRegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if err := dbhelper.CheckSelfRegisterCredentials(&req, w); err != nil {
			return
		}
		req.Role = "Employee"
		req.EmpType = "FullTime"
		newUser, err := dbhelper.RegisterUser(database.ST, &req)
		if err != nil {
			http.Error(w, "Could not register user", http.StatusInternalServerError)
			return
		}

		res := models.SelfRegisterUserResponse{
			ID:             newUser.ID,
			Fname:          newUser.Fname,
			Lname:          newUser.Lname,
			Email:          newUser.Email,
			Role:           newUser.Role,
			EmpType:        newUser.EmpType,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

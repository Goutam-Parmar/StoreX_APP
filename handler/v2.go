package handler

import (
	"StoreXApp/dbhelper"
	"StoreXApp/models"
	"StoreXApp/utils"
	"encoding/json"
	"net/http"
	"time"
)

func SignupV2() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		type SignupRequest struct {
			Email string `json:"email"`
		}
		var req SignupRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		userID, role, err := dbhelper.SignupOrLoginUser(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		accessToken, err := utils.GenerateAccessToken(userID, role)
		if err != nil {
			http.Error(w, "Could not generate access token", http.StatusInternalServerError)
			return
		}

		refreshToken, err := utils.GenerateRefreshToken(userID, role)
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

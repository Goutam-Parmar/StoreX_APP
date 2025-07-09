package handler

import (
	"StoreXApp/database"
	"StoreXApp/dbhelper"
	"StoreXApp/models"
	"StoreXApp/utils"
	"encoding/json"
	"net/http"
	"time"
)

func CreateLaptopAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth header", http.StatusUnauthorized)
			return
		}
		authClaims, err := utils.ExtractAuthClaims(authHeader)
		if err != nil {
			http.Error(w, " jwt token error", http.StatusUnauthorized)
			return
		}
		var req models.CreateLaptopAssetRequest
		req.CreatedBy = authClaims.UserID
		role := authClaims.Role
		if role != "AssetManager" {
			http.Error(w, "you are not AssetManager", http.StatusUnauthorized)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)

		assetID, laptopID, err := dbhelper.CreateLaptopAsset(tx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreateLaptopAssetResponse{
			AssetID:        assetID,
			LaptopID:       laptopID,
			AssetType:      "laptop",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

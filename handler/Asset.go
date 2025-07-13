package handler

import (
	"StoreXApp/database"
	"StoreXApp/dbhelper"
	"StoreXApp/models"
	"StoreXApp/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func CreateLaptopAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authClaims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.CreateLaptopAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, laptopID, err := dbhelper.CreateLaptopAsset(tx, &req)
		if err != nil {
			log.Printf("CreateLaptopAsset error: %v", err)
			http.Error(w, "Failed to create asset", http.StatusInternalServerError)
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
func CreateMobileAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authClaims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.CreateMobileAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, mobileID, err := dbhelper.CreateMobileAsset(tx, &req)
		if err != nil {
			log.Printf(" CreateMobileAsset error: %v", err)
			http.Error(w, "Failed to create mobile asset", http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreateMobileAssetResponse{
			AssetID:        assetID,
			MobileID:       mobileID,
			AssetType:      "mobile",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func CreateMouseAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authClaims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.CreateMouseAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, mouseID, err := dbhelper.CreateMouseAsset(tx, &req)
		if err != nil {
			http.Error(w, "Failed to create mouse asset", http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreateMouseAssetResponse{
			AssetID:        assetID,
			MouseID:        mouseID,
			AssetType:      "mouse",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func CreateMonitorAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth header", http.StatusUnauthorized)
			return
		}
		authClaims, err := utils.ExtractAuthClaims(authHeader)
		var req models.CreateMonitorAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, monitorID, err := dbhelper.CreateMonitorAsset(tx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreateMonitorAssetResponse{
			AssetID:        assetID,
			MonitorID:      monitorID,
			AssetType:      "monitor",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func CreateHarddiskAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth header", http.StatusUnauthorized)
			return
		}
		authClaims, err := utils.ExtractAuthClaims(authHeader)
		var req models.CreateHarddiskAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, harddiskID, err := dbhelper.CreateHarddiskAsset(tx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreateHarddiskAssetResponse{
			AssetID:        assetID,
			HarddiskID:     harddiskID,
			AssetType:      "harddisk",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func CreatePendriveAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth header", http.StatusUnauthorized)
			return
		}
		authClaims, err := utils.ExtractAuthClaims(authHeader)
		var req models.CreatePendriveAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, pendriveID, err := dbhelper.CreatePendriveAsset(tx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreatePendriveAssetResponse{
			AssetID:        assetID,
			PendriveID:     pendriveID,
			AssetType:      "pendrive",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func CreateAccessoriesAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing auth header", http.StatusUnauthorized)
			return
		}
		authClaims, err := utils.ExtractAuthClaims(authHeader)
		var req models.CreateAccessoriesAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
		tx, err := database.ST.Begin()
		defer utils.Tx(tx, &err)
		assetID, accessoriesID, err := dbhelper.CreateAccessoriesAsset(tx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tx.Commit(); err != nil {
			http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
			return
		}
		res := models.CreateAccessoriesAssetResponse{
			AssetID:        assetID,
			AccessoriesID:  accessoriesID,
			AssetType:      "accessories",
			AssetStatus:    "available",
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func LaptopAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		if err := dbhelper.AssignLaptopAsset(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func MobileAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		err = dbhelper.AssignMobileAsset(&req)
		if err != nil {
			http.Error(w, "Could not assign mobile: "+err.Error(), http.StatusInternalServerError)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func MonitorAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		err = dbhelper.AssignMonitorAsset(&req)
		if err != nil {
			http.Error(w, "Could not assign monitor: "+err.Error(), http.StatusInternalServerError)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func MouseAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		err = dbhelper.AssignMouseAsset(&req)
		if err != nil {
			http.Error(w, "Could not assign mouse: "+err.Error(), http.StatusInternalServerError)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
func HardDiskAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		err = dbhelper.AssignHardDiskAsset(&req)
		if err != nil {
			http.Error(w, "Could not assign mouse: "+err.Error(), http.StatusInternalServerError)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

}
func PendriveAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		err = dbhelper.AssignPendriveAsset(&req)
		if err != nil {
			http.Error(w, "Could not assign mouse: "+err.Error(), http.StatusInternalServerError)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

}
func AcessoriesAssignHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "error in token extraction", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		err = dbhelper.AssignAccessoriesAsset(&req)
		if err != nil {
			http.Error(w, "Could not assign mouse: "+err.Error(), http.StatusInternalServerError)
			return
		}
		res := models.AssignAssetResponse{
			AssetID:        req.AssetID,
			AssignedTo:     req.EmployeeID,
			AssignedBy:     req.AssignedBy,
			AssignedAt:     time.Now().Format(time.RFC3339),
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}

}
func GetAllAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		assets, err := dbhelper.GetAllAssets(ctx)
		if err != nil {
			http.Error(w, "Failed to fetch assets: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(assets)
	}
}

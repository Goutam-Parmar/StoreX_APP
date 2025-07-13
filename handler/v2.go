package handler

import (
	"StoreXApp/dbhelper"
	"StoreXApp/models"
	"StoreXApp/utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
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
func DynamicAssignAssetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.AssignAssetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "failed to extract the created_by for jwt ", http.StatusUnauthorized)
			return
		}
		req.AssignedBy = claims.UserID
		if err := dbhelper.DynamicAssignAsset(&req, w); err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Asset assigned successfully"}`))
	}
}

func EmployeeSearchByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var req models.EmployeeSearchByNameRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "failed to extract the created_by for jwt ", http.StatusUnauthorized)
			return
		}

		if claims.Role == "Employee" {
			http.Error(w, "This is the Protected api , not for the Employee", http.StatusUnauthorized)
		}
		req.Name = strings.TrimSpace(req.Name)
		if req.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}
		if len(req.Name) < 3 {
			http.Error(w, "Name must be at least 3 characters long", http.StatusBadRequest)
			return
		}

		results, err := dbhelper.SearchEmployeeByName(req.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		resp := &models.EmployeeSearchByNameResponse{
			Users:          results,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func EmployeeSearchByPhoneNo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var req models.EmployeeSearchByPhoneNoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "failed to extract the created_by for jwt ", http.StatusUnauthorized)
			return
		}

		if claims.Role == "Employee" {
			http.Error(w, "This is the Protected api , not for the Employee", http.StatusUnauthorized)
		}
		req.PhoneNo = strings.TrimSpace(req.PhoneNo)

		if req.PhoneNo == "" {
			http.Error(w, "Phone number is required", http.StatusBadRequest)
			return
		}

		if len(req.PhoneNo) < 4 || len(req.PhoneNo) > 10 {
			http.Error(w, "Phone number prefix must be between 4 to 10 digits", http.StatusBadRequest)
			return
		}

		results, err := dbhelper.SearchEmployeeByPhoneNo(req.PhoneNo)
		if err != nil {
			http.Error(w, "No employees found with given phone number prefix", http.StatusNotFound)
			return
		}

		resp := &models.EmployeeSearchByPhoneNoResponse{
			Users:          results,
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
func GetDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "failed to extract the created_by for jwt ", http.StatusUnauthorized)
			return
		}

		if claims.Role == "Employee" {
			http.Error(w, "This is the Protected api , not for the Employee", http.StatusUnauthorized)
		}

		result, err := dbhelper.GetDashboardCounts()
		if err != nil {
			http.Error(w, "Failed to get dashboard data", http.StatusInternalServerError)
			return
		}

		result.ResponseTimeMs = float64(time.Since(start).Microseconds()) / 1000.0

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}

func GetAssetInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assetID := mux.Vars(r)["Asset_id"]

		info, err := dbhelper.GetAssetInfo(assetID)
		if err != nil {
			http.Error(w, "Asset not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	}
}
func ChangeRole() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.ChangeRoleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.UserID == "" || req.Role == "" {
			http.Error(w, "UserID and Role are required", http.StatusBadRequest)
			return
		}

		if err := dbhelper.ChangeUserRole(req.UserID, req.Role); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
func DeleteAsset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assetID := mux.Vars(r)["Asset_id"]

		if assetID == "" {
			http.Error(w, "Asset ID is required", http.StatusBadRequest)
			return
		}

		err := dbhelper.DeleteAsset(assetID)
		if err != nil {
			if strings.Contains(err.Error(), "retrieve this asset") {
				http.Error(w, err.Error(), http.StatusConflict) // 409 Conflict
			} else {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func RetrieveAsset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RetrieveAssetRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.EmployeeID == "" || req.AssetID == "" || req.Reason == "" {
			http.Error(w, "employee_id, asset_id, and reason are required", http.StatusBadRequest)
			return
		}

		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Failed to extract user from JWT", http.StatusUnauthorized)
			return
		}
		performedBy := claims.UserID

		err = dbhelper.RetrieveAsset(req.AssetID, req.EmployeeID, req.Reason, performedBy)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
func AssetUnAssignedStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assets, err := dbhelper.GetUnAssignedAssets()
		if err != nil {
			http.Error(w, "Failed to get unassigned assets: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(assets)
	}
}
func AssetAssignedStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assets, err := dbhelper.GetAssignedAssets()
		if err != nil {
			http.Error(w, "Failed to get assigned assets: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(assets)
	}
}

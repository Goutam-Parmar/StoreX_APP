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

func RegisterUserByEmpManager() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		authClaims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		var req models.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request payload", http.StatusBadRequest)
			return
		}
		req.Email = strings.TrimSpace(strings.ToLower(req.Email))
		req.CreatedBy = authClaims.UserID
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
		authClaims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		var req models.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		req.CreatedBy = authClaims.UserID
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
		newUser, err := dbhelper.RegisterUser(&req)
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
func GetAllEmployees() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		employees, err := dbhelper.GetAllEmployees(ctx)
		if err != nil {
			http.Error(w, "Failed to fetch employees: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}
func EmployeeSearchByEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		req.Email = strings.TrimSpace(strings.ToLower(req.Email))

		if req.Email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		result, err := dbhelper.SearchEmployeeByEmail(req.Email)
		if err != nil {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		resp := &models.EmployeeSearchByEmailResponse{
			User: &models.SingleEmployeeData{
				ID:      result.ID,
				Fname:   result.Fname,
				Lname:   result.Lname,
				Email:   result.Email,
				Role:    result.Role,
				EmpType: result.EmpType,
			},
			ResponseTimeMs: float64(time.Since(start).Microseconds()) / 1000.0,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
func DeleteEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.DeleteEmp
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.EmployeeID == "" {
			http.Error(w, "Employee ID is required", http.StatusBadRequest)
			return
		}

		hasAssets, err := dbhelper.CheckEmployeeHasAssets(req.EmployeeID)
		if err != nil {
			http.Error(w, "Failed to check assets: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if hasAssets {
			http.Error(w, "Employee still has assigned assets. Please retrieve them first.", http.StatusBadRequest)
			return
		}

		err = dbhelper.DeleteEmployee(req.EmployeeID)
		if err != nil {
			http.Error(w, "Failed to delete employee: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
func GETAssetTimeLine() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employeeID := mux.Vars(r)["Employee_Id"]
		if employeeID == "" {
			http.Error(w, "Employee_Id is required", http.StatusBadRequest)
			return
		}

		timeline, err := dbhelper.GetEmployeeAssetTimeline(employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timeline)
	}
}
func GetMyDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := utils.ExtractAuthClaims(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Failed to extract user from JWT", http.StatusUnauthorized)
			return
		}
		employeeID := claims.UserID

		resp, err := dbhelper.GetMyDashboard(employeeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

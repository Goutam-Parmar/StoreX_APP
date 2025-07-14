package dbhelper

import (
	"StoreXApp/database"
	"StoreXApp/models"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var ValidRoles = map[string]string{
	"admin":           "Admin",
	"employee":        "Employee",
	"assetmanager":    "AssetManager",
	"employeemanager": "EmployeeManager",
}

func RegisterUserBYAdmin(req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {

	key := strings.ToLower(req.Role)
	roleFormatted, ok := ValidRoles[key]
	if !ok {
		return nil, fmt.Errorf("invalid role: %s", req.Role)
	}

	query := `
		INSERT INTO users (email, fname, lname, role, emp_type, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, email, role
	`
	var res models.RegisterUserResponse
	err := database.ST.QueryRowContext(
		context.Background(),
		query,
		req.Email,
		req.Fname,
		req.Lname,
		roleFormatted,
		req.EmpType,
		req.CreatedBy,
	).Scan(&res.ID, &res.Email, &res.Role)

	if err != nil {
		log.Println("RegisterUserBYAdmin: error while inserting user:", err)
		return nil, err
	}

	return &res, nil
}

func CheckSelfRegisterCredentials(req *models.SelfRegisterUserRequest, w http.ResponseWriter) error {
	emailParts := strings.Split(req.Email, "@")
	if len(emailParts) != 2 || emailParts[1] != "remotestate.com" {
		http.Error(w, "Only @remotestate.com emails are allowed", http.StatusBadRequest)
		return errors.New("invalid email domain")
	}
	localPart := emailParts[0]
	if strings.Contains(localPart, ".") {
		parts := strings.SplitN(localPart, ".", 2)
		req.Fname = strings.Title(parts[0])
		req.Lname = strings.Title(parts[1])
	} else {
		req.Fname = strings.Title(localPart)
		req.Lname = ""
	}
	var exists bool
	err := database.ST.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM users WHERE email = $1 AND is_deleted = FALSE
		)
	`, req.Email).Scan(&exists)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return err
	}

	if exists {
		http.Error(w, "This email is already registered", http.StatusConflict)
		return errors.New("email already exists")
	}

	req.Role = strings.ToLower(req.Role)
	return nil
}

func CheckLoginCredentials(req *models.LoginRequest, w http.ResponseWriter) error {
	emailParts := strings.Split(req.Email, "@")
	if len(emailParts) != 2 || emailParts[1] != "remotestate.com" {
		http.Error(w, "only @remotestate.com emails are allowed", http.StatusUnauthorized)
		return errors.New("invalid email domain")
	}
	localPart := emailParts[0]
	if strings.Contains(localPart, ".") {
		parts := strings.SplitN(localPart, ".", 2)
		req.Fname = strings.Title(parts[0])
		req.Lname = strings.Title(parts[1])
	} else {
		req.Fname = strings.Title(localPart)
		req.Lname = ""
	}
	var userID string
	var role string
	err := database.ST.QueryRow(`
		SELECT id, role
		FROM users
		WHERE email = $1 AND is_deleted = FALSE
		LIMIT 1
	`, req.Email).Scan(&userID, &role)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid email or user not found", http.StatusUnauthorized)
		return err
	}
	req.UserID = userID
	req.Role = strings.ToLower(role)
	return nil
}
func RegisterUser(req *models.SelfRegisterUserRequest) (*models.SelfRegisterUserResponse, error) {
	var id string
	err := database.ST.QueryRow(`
		INSERT INTO users (email, fname, lname, role, emp_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		req.Email,
		req.Fname,
		req.Lname,
		req.Role,
		req.EmpType,
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	return &models.SelfRegisterUserResponse{
		ID:      id,
		Fname:   req.Fname,
		Lname:   req.Lname,
		Email:   req.Email,
		Role:    req.Role,
		EmpType: req.EmpType,
	}, nil
}

// for admin, empmanager, assermanager
func CheckRegisterCredentials(req *models.RegisterUserRequest, w http.ResponseWriter) error {

	emailParts := strings.Split(req.Email, "@")

	if len(emailParts) != 2 || emailParts[1] != "remotestate.com" {
		http.Error(w, "Only @remotestate.com emails are allowed", http.StatusBadRequest)
		return errors.New("invalid email domain")
	}
	nameParts := strings.Split(emailParts[0], ".")
	req.Fname = nameParts[0]
	if len(nameParts) > 1 {
		req.Lname = nameParts[1]
	} else {
		req.Lname = ""
	}

	var exists bool
	err := database.ST.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM users WHERE email = $1 AND is_deleted = FALSE
		)
	`, req.Email).Scan(&exists)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return err
	}
	if exists {
		http.Error(w, "This email is already registered", http.StatusConflict)
		return errors.New("email already exists")
	}
	req.Role = strings.ToLower(req.Role)
	return nil
}
func GetAllEmployees() ([]models.EmployeeResponse, error) {
	query := `
		SELECT 
			u.id, 
			u.fname, 
			u.lname, 
			u.email, 
			u.role, 
			u.emp_type, 
			CASE WHEN at.id IS NOT NULL THEN true ELSE false END AS has_asset_assigned
		FROM 
			users u
		LEFT JOIN 
			asset_timeline at ON u.id = at.assigned_to AND at.status = 'assigned'
		WHERE 
			u.is_deleted = FALSE;
	`

	rows, err := database.ST.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.EmployeeResponse

	for rows.Next() {
		var emp models.EmployeeResponse
		err := rows.Scan(
			&emp.ID,
			&emp.Fname,
			&emp.Lname,
			&emp.Email,
			&emp.Role,
			&emp.EmpType,
			&emp.HasAssetAssigned,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func CheckEmployeeHasAssets(employeeID string) (bool, error) {
	var count int
	query := `
		SELECT COUNT(*)
		FROM asset_timeline
		WHERE assigned_to = $1 AND status = 'assigned'
	`
	err := database.ST.QueryRow(query, employeeID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func DeleteEmployee(employeeID string) error {
	_, err := database.ST.Exec(`
		UPDATE users
		SET is_deleted = TRUE, deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND is_deleted = FALSE
	`, employeeID)
	return err
}
func GetMyDashboard(employeeID string) (*models.MyDashboardResponse, error) {
	var resp models.MyDashboardResponse
	err := database.ST.QueryRow(`
		SELECT COUNT(*)
		FROM asset_timeline
		WHERE assigned_to = $1 AND status = 'assigned'
	`, employeeID).Scan(&resp.TotalAssigned)
	if err != nil {
		return nil, err
	}

	rows, err := database.ST.Query(`
		SELECT a.id, a.brand, a.model, a.asset_type, a.asset_status, at.assigned_at
		FROM asset_timeline at
		JOIN assets a ON at.asset_id = a.id
		WHERE at.assigned_to = $1 AND at.status = 'assigned'
		ORDER BY at.assigned_at DESC
	`, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.MyDashboardAssetEntry
		err := rows.Scan(&item.ID, &item.Brand, &item.Model, &item.AssetType, &item.AssetStatus, &item.AssignedAt)
		if err != nil {
			return nil, err
		}
		resp.Assets = append(resp.Assets, item)
	}

	return &resp, nil
}

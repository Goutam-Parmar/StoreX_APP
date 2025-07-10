package dbhelper

import (
	"StoreXApp/database"
	"StoreXApp/models"
	"context"
	"database/sql"
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
func RegisterUser(db *sql.DB, req *models.SelfRegisterUserRequest) (*models.SelfRegisterUserResponse, error) {
	var id string
	err := db.QueryRow(`
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

package dbhelper

import (
	"StoreXApp/database"
	"StoreXApp/models"
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
)

func RegisterUserBYAdmin(req *models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	query := `
		INSERT INTO users (email, fname, lname, role, job_type, created_by)
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
		req.Role,
		req.EmpType,
		req.CreatedBy,
	).Scan(&res.ID, &res.Email, &res.Role)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
func SelfRegisterUser(req *models.SelfRegisterUserRequest) error {
	query := `
		INSERT INTO users (fname, lname, email, role, emp_type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var newUserID string

	err := database.ST.QueryRowContext(
		context.Background(),
		query,
		req.Fname,
		req.Lname,
		req.Email,
		req.Role,
		req.EmpType,
	).Scan(&newUserID)
	log.Println(newUserID)

	if err != nil {
		return err
	}
	req.UserID = newUserID

	return nil
}
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

	return nil
}

func CheckSelfRegisterCredentials(req *models.SelfRegisterUserRequest, w http.ResponseWriter) error {
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

	return nil
}

func CheckLoginCredentials(req *models.LoginRequest, w http.ResponseWriter) error {
	emailParts := strings.Split(req.Email, "@")
	if len(emailParts) != 2 || emailParts[1] != "remotestate.com" {
		http.Error(w, "only @remotestate.com emails are allowed", http.StatusUnauthorized)
		return errors.New("invalid email domain")
	}

	nameParts := strings.Split(emailParts[0], ".")
	req.Fname = nameParts[0]
	if len(nameParts) > 1 {
		req.Lname = nameParts[1]
	} else {
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
	req.Role = role

	return nil
}

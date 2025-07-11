// dbhelper/signup.go

package dbhelper

import (
	"StoreXApp/database"
	"context"
	"errors"
	"log"
	"strings"
)

func SignupOrLoginUser(email string) (string, string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	parts := strings.Split(email, "@")
	if len(parts) != 2 || parts[1] != "remotestate.com" {
		return "", "", errors.New("only @remotestate.com emails allowed")
	}
	localPart := parts[0]
	var fname, lname string
	if strings.Contains(localPart, ".") {
		sub := strings.SplitN(localPart, ".", 2)
		fname = strings.Title(sub[0])
		lname = strings.Title(sub[1])
	} else {
		fname = strings.Title(localPart)
		lname = ""
	}

	var userID string
	var role string

	err := database.ST.QueryRowContext(context.Background(),
		`SELECT id, role FROM users WHERE email=$1 AND is_deleted=FALSE LIMIT 1`,
		email,
	).Scan(&userID, &role)

	if err == nil {
		return userID, strings.ToLower(role), nil
	}

	// Not found → insert as new Employee
	role = "Employee"
	empType := "FullTime"
	err = database.ST.QueryRowContext(context.Background(),
		`INSERT INTO users (email, fname, lname, role, emp_type)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		email, fname, lname, role, empType,
	).Scan(&userID)

	if err != nil {
		log.Println("SignupOrLoginUser: insert error:", err)
		return "", "", err
	}

	return userID, role, nil
}

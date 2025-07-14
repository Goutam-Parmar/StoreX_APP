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
func DynamicAssignAsset(req *models.AssignAssetRequest, w http.ResponseWriter) error {
	var assetType string
	var currentStatus string

	log.Println("inside DynamicAssignAsset db helper")

	err := database.ST.QueryRow(`
		SELECT asset_type, asset_status 
		FROM assets 
		WHERE id = $1 AND deleted_at IS NULL
	`, req.AssetID).Scan(&assetType, &currentStatus)

	log.Println("just after the asset type query")

	if err != nil {
		log.Println("Asset lookup error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		http.Error(w, "Asset is not available, assigned to someone else", http.StatusInternalServerError)
		return errors.New("asset is not available for assignment")
	}

	log.Println("Asset assigned:", req.AssetID)
	log.Println("Asset type is", assetType)
	log.Println("just before the if else statement")

	if assetType == "laptop" {
		return AssignLaptopAsset(req)
	} else if assetType == "mobile" {
		return AssignMobileAsset(req)
	} else if assetType == "monitor" {
		return AssignMonitorAsset(req)
	} else if assetType == "mouse" {
		return AssignMouseAsset(req)
	} else if assetType == "harddisk" {
		return AssignHardDiskAsset(req)
	} else if assetType == "pendrive" {
		return AssignPendriveAsset(req)
	} else if assetType == "accessories" {
		return AssignAccessoriesAsset(req)
	} else {
		return errors.New("unsupported asset type: " + assetType)
	}
}
func GetAllAssets() ([]models.AssetResponse, error) {
	query := `
		SELECT 
			a.id,
			a.asset_type,
			a.brand,
			a.model,
			at.assigned_to,
			a.asset_status
		FROM 
			assets a
		LEFT JOIN 
			asset_timeline at 
		ON 
			a.id = at.asset_id AND at.status = 'assigned'
		WHERE 
			a.is_active = TRUE AND a.deleted_at IS NULL
	`

	rows, err := database.ST.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []models.AssetResponse

	for rows.Next() {
		var asset models.AssetResponse
		err := rows.Scan(
			&asset.ID,
			&asset.AssetType,
			&asset.Brand,
			&asset.Model,
			&asset.AssignedTo,
			&asset.AssetStatus,
		)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return assets, nil
}

func SearchEmployeeByName(keyword string) ([]models.EmployeeSearchByNameUser, error) {
	query := `
		SELECT id, fname, lname, email, role, emp_type 
		FROM users 
		WHERE is_deleted = FALSE 
		AND fname ILIKE $1
	`

	rows, err := database.ST.Query(query, keyword+"%")
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var results []models.EmployeeSearchByNameUser

	for rows.Next() {
		var emp models.EmployeeSearchByNameUser
		err := rows.Scan(
			&emp.ID,
			&emp.Fname,
			&emp.Lname,
			&emp.Email,
			&emp.Role,
			&emp.EmpType,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		results = append(results, emp)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no employees found for '%s'", keyword)
	}

	return results, nil
}
func SearchEmployeeByEmail(email string) (*models.SingleEmployeeData, error) {
	query := `
		SELECT id, fname, lname, email, role, emp_type
		FROM users
		WHERE is_deleted = FALSE AND email = $1
	`

	row := database.ST.QueryRow(query, email)

	var user models.SingleEmployeeData
	err := row.Scan(
		&user.ID,
		&user.Fname,
		&user.Lname,
		&user.Email,
		&user.Role,
		&user.EmpType,
	)

	if err != nil {
		return nil, fmt.Errorf("employee not found")
	}

	return &user, nil
}
func SearchEmployeeByPhoneNo(prefix string) ([]models.SingleEmployeePhoneData, error) {
	query := `
		SELECT id, fname, lname, email, phone_no, role, emp_type
		FROM users
		WHERE is_deleted = FALSE
		AND phone_no IS NOT NULL
		AND phone_no ILIKE $1
	`

	rows, err := database.ST.Query(query, prefix+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.SingleEmployeePhoneData

	for rows.Next() {
		var emp models.SingleEmployeePhoneData
		err := rows.Scan(
			&emp.ID,
			&emp.Fname,
			&emp.Lname,
			&emp.Email,
			&emp.PhoneNo,
			&emp.Role,
			&emp.EmpType,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, emp)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no employees found for phone prefix '%s'", prefix)
	}

	return results, nil
}
func GetDashboardCounts() (*models.DashboardResponse, error) {
	query := `
		SELECT
			COUNT(*) AS total_assets,
			COUNT(*) FILTER (WHERE asset_status = 'assigned') AS assigned_assets,
			COUNT(*) FILTER (WHERE asset_status = 'available') AS available_assets,
			COUNT(*) FILTER (WHERE asset_status = 'waiting_for_service') AS waiting_for_service_assets,
			COUNT(*) FILTER (WHERE asset_status = 'in_service') AS in_service_assets,
			COUNT(*) FILTER (WHERE asset_status = 'damaged') AS damaged_assets
		FROM assets
		WHERE deleted_at IS NULL
	`

	var resp models.DashboardResponse

	err := database.ST.QueryRow(query).Scan(
		&resp.TotalAssets,
		&resp.AssignedAssets,
		&resp.AvailableAssets,
		&resp.WaitingForServiceAssets,
		&resp.InServiceAssets,
		&resp.DamagedAssets,
	)
	if err != nil {
		log.Println("Error fetching dashboard counts:", err)
		return nil, err
	}

	return &resp, nil
}
func GetAssetInfo(assetID string) (*models.SimpleAssetInfoResponse, error) {
	var res models.SimpleAssetInfoResponse

	err := database.ST.QueryRow(`
	SELECT id, brand, model, asset_type, purchase_price, purchased_date, asset_status
	FROM assets WHERE id = $1 AND deleted_at IS NULL
`, assetID).Scan(
		&res.ID, &res.Brand, &res.Model, &res.AssetType,
		&res.PurchasePrice, &res.PurchasedDate, &res.AssetStatus,
	)
	if err != nil {
		return nil, fmt.Errorf("Asset not found: %v", err)
	}

	switch res.AssetType {
	case "laptop":
		err = database.ST.QueryRow(`
			SELECT processor, ram, storage, os FROM laptop WHERE asset_id = $1
		`, assetID).Scan(
			&res.Processor, &res.Ram, &res.Storage, &res.OS,
		)
	case "mouse":
		err = database.ST.QueryRow(`
			SELECT dpi, connection_type FROM mouse WHERE asset_id = $1
		`, assetID).Scan(
			&res.DPI, &res.ConnectionType,
		)
	case "mobile":
		err = database.ST.QueryRow(`
			SELECT imei, ram, storage FROM mobile WHERE asset_id = $1
		`, assetID).Scan(
			&res.IMEI, &res.Ram, &res.Storage,
		)
	case "monitor":
		err = database.ST.QueryRow(`
			SELECT screen_size, resolution, panel_type FROM monitor WHERE asset_id = $1
		`, assetID).Scan(
			&res.ScreenSize, &res.Resolution, &res.PanelType,
		)
	case "harddisk":
		err = database.ST.QueryRow(`
			SELECT capacity, disk_type FROM harddisk WHERE asset_id = $1
		`, assetID).Scan(
			&res.Capacity, &res.DiskType,
		)
	case "pendrive":
		err = database.ST.QueryRow(`
			SELECT capacity FROM pendrive WHERE asset_id = $1
		`, assetID).Scan(
			&res.Capacity,
		)
	case "accessories":
		err = database.ST.QueryRow(`
			SELECT name, work FROM accessories WHERE asset_id = $1
		`, assetID).Scan(
			&res.Name, &res.Work,
		)
	}

	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, fmt.Errorf("Failed to get extra info: %v", err)
	}

	return &res, nil
}
func ChangeUserRole(userID, role string) error {

	validRoles := map[string]string{
		"admin":           "Admin",
		"employee":        "Employee",
		"assetmanager":    "AssetManager",
		"employeemanager": "EmployeeManager",
	}

	key := strings.ToLower(role)
	formattedRole, ok := validRoles[key]
	if !ok {
		return fmt.Errorf("invalid role: %s", role)
	}

	_, err := database.ST.Exec(`
		UPDATE users
		SET role = $1, updated_at = NOW()
		WHERE id = $2 AND is_deleted = FALSE
	`, formattedRole, userID)
	if err != nil {
		return fmt.Errorf("failed to update role: %v", err)
	}

	return nil
}
func DeleteAsset(assetID string) error {

	var status string
	err := database.ST.QueryRow(`
		SELECT asset_status FROM assets
		WHERE id = $1 AND deleted_at IS NULL
	`, assetID).Scan(&status)
	if err != nil {
		return fmt.Errorf("asset not found: %v", err)
	}

	if status == "assigned" {
		return fmt.Errorf("cannot delete: first retrieve this asset from the employee")
	}

	var assignedCount int
	err = database.ST.QueryRow(`
		SELECT COUNT(*) FROM asset_timeline
		WHERE asset_id = $1 AND status = 'assigned'
	`, assetID).Scan(&assignedCount)
	if err != nil {
		return fmt.Errorf("failed to check asset timeline: %v", err)
	}

	if assignedCount > 0 {
		return fmt.Errorf("cannot delete: first retrieve this asset from the employee")
	}

	_, err = database.ST.Exec(`
		UPDATE assets
		SET deleted_at = NOW(), is_active = FALSE
		WHERE id = $1
	`, assetID)
	if err != nil {
		return fmt.Errorf("failed to delete asset: %v", err)
	}

	return nil
}

func RetrieveAsset(assetID, employeeID, reason, performedBy string) error {
	tx, err := database.ST.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	var assigned bool
	err = tx.QueryRow(`
		SELECT EXISTS (
			SELECT 1 FROM asset_timeline
			WHERE asset_id = $1 AND assigned_to = $2 AND status = 'assigned'
		)
	`, assetID, employeeID).Scan(&assigned)
	if err != nil {
		return fmt.Errorf("failed to check assignment: %v", err)
	}
	if !assigned {
		return fmt.Errorf("asset is not assigned to this employee")
	}

	_, err = tx.Exec(`
		UPDATE asset_timeline
		SET status = 'retrieved',
		    returned_at = NOW(),
		    reason = $1
		WHERE asset_id = $2 AND assigned_to = $3 AND status = 'assigned'
	`, reason, assetID, employeeID)
	if err != nil {
		return fmt.Errorf("failed to update asset_timeline: %v", err)
	}

	_, err = tx.Exec(`
		UPDATE assets
		SET asset_status = 'available',
		    updated_at = NOW()
		WHERE id = $1
	`, assetID)
	if err != nil {
		return fmt.Errorf("failed to update asset status: %v", err)
	}

	_, err = tx.Exec(`
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by, performed_at)
		VALUES ($1, 'assigned', 'available', $2, $3, NOW())
	`, assetID, employeeID, performedBy)
	if err != nil {
		return fmt.Errorf("failed to insert asset_history: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
func GetUnAssignedAssets() ([]models.AssetListResponse, error) {
	query := `
		SELECT id, brand, model, asset_type, category, owned_by, purchase_price, purchased_date
		FROM assets
		WHERE asset_status = 'available'
			AND is_active = TRUE
			AND is_retired = FALSE
			AND deleted_at IS NULL
	`

	rows, err := database.ST.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []models.AssetListResponse

	for rows.Next() {
		var asset models.AssetListResponse
		err := rows.Scan(
			&asset.ID,
			&asset.Brand,
			&asset.Model,
			&asset.AssetType,
			&asset.Category,
			&asset.OwnedBy,
			&asset.PurchasePrice,
			&asset.PurchasedDate,
		)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}
func GetAssignedAssets() ([]models.AssetListResponse, error) {
	query := `
		SELECT id, brand, model, asset_type, category, owned_by, purchase_price, purchased_date
		FROM assets
		WHERE asset_status = 'assigned'
			AND is_active = TRUE
			AND is_retired = FALSE
			AND deleted_at IS NULL
	`

	rows, err := database.ST.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []models.AssetListResponse

	for rows.Next() {
		var asset models.AssetListResponse
		err := rows.Scan(
			&asset.ID,
			&asset.Brand,
			&asset.Model,
			&asset.AssetType,
			&asset.Category,
			&asset.OwnedBy,
			&asset.PurchasePrice,
			&asset.PurchasedDate,
		)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}
func GetEmployeeAssetTimeline(employeeID string) ([]models.EmployeeAssetTimeline, error) {
	rows, err := database.ST.Query(`
    SELECT 
      at.asset_id,
      a.brand,
      a.model,
      a.asset_type,
      at.status,
      at.assigned_at,
      at.returned_at,
      at.reason
    FROM asset_timeline at
    JOIN assets a ON a.id = at.asset_id
    WHERE at.assigned_to = $1
    ORDER BY at.assigned_at DESC
  `, employeeID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timeline []models.EmployeeAssetTimeline

	for rows.Next() {
		var row models.EmployeeAssetTimeline
		err := rows.Scan(
			&row.AssetID,
			&row.Brand,
			&row.Model,
			&row.AssetType,
			&row.Status,
			&row.AssignedAt,
			&row.ReturnedAt,
			&row.Reason,
		)
		if err != nil {
			return nil, err
		}
		timeline = append(timeline, row)
	}

	return timeline, nil
}

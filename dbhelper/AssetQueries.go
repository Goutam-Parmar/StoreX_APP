package dbhelper

import (
	"StoreXApp/database"
	"StoreXApp/models"
	"context"
	"database/sql"
	"errors"
	"log"
)

func CreateLaptopAsset(tx *sql.Tx, req *models.CreateLaptopAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'laptop', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var laptopID string
	err = tx.QueryRow(`
		INSERT INTO laptop (asset_id, processor, ram, storage, os, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`,
		assetID, req.Processor, req.Ram, req.Storage, req.OS, req.CreatedBy,
	).Scan(&laptopID)
	if err != nil {
		return "", "", err
	}

	return assetID, laptopID, nil
}
func CreateMobileAsset(tx *sql.Tx, req *models.CreateMobileAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'mobile', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var mobileID string
	err = tx.QueryRow(`
		INSERT INTO mobile (asset_id, imei, ram, storage, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`,
		assetID, req.IMEI, req.Ram, req.Storage, req.CreatedBy,
	).Scan(&mobileID)
	if err != nil {
		return "", "", err
	}

	return assetID, mobileID, nil
}
func CreateMouseAsset(tx *sql.Tx, req *models.CreateMouseAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'mouse', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var mouseID string
	err = tx.QueryRow(`
		INSERT INTO mouse (asset_id, dpi, connection_type, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		assetID, req.DPI, req.ConnectionType, req.CreatedBy,
	).Scan(&mouseID)
	if err != nil {
		return "", "", err
	}

	return assetID, mouseID, nil
}
func CreateMonitorAsset(tx *sql.Tx, req *models.CreateMonitorAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'monitor', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var monitorID string
	err = tx.QueryRow(`
		INSERT INTO monitor (asset_id, screen_size, resolution, panel_type, created_by)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`,
		assetID, req.ScreenSize, req.Resolution, req.PanelType, req.CreatedBy,
	).Scan(&monitorID)
	if err != nil {
		return "", "", err
	}

	return assetID, monitorID, nil
}
func CreateHarddiskAsset(tx *sql.Tx, req *models.CreateHarddiskAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'harddisk', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var harddiskID string
	err = tx.QueryRow(`
		INSERT INTO harddisk (asset_id, capacity, disk_type, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		assetID, req.Capacity, req.DiskType, req.CreatedBy,
	).Scan(&harddiskID)
	if err != nil {
		return "", "", err
	}

	return assetID, harddiskID, nil
}
func CreatePendriveAsset(tx *sql.Tx, req *models.CreatePendriveAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'pendrive', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var pendriveID string
	err = tx.QueryRow(`
		INSERT INTO pendrive (asset_id, capacity, created_by)
		VALUES ($1, $2, $3)
		RETURNING id
	`,
		assetID, req.Capacity, req.CreatedBy,
	).Scan(&pendriveID)
	if err != nil {
		return "", "", err
	}

	return assetID, pendriveID, nil
}

func CreateAccessoriesAsset(tx *sql.Tx, req *models.CreateAccessoriesAssetRequest) (string, string, error) {
	var assetID string
	err := tx.QueryRow(`
		INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
		VALUES ($1, $2, 'accessories', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
		RETURNING id
	`,
		req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", err
	}

	var accessoriesID string
	err = tx.QueryRow(`
		INSERT INTO accessories (asset_id, name, work, created_by)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`,
		assetID, req.Name, req.Work, req.CreatedBy,
	).Scan(&accessoriesID)
	if err != nil {
		return "", "", err
	}

	return assetID, accessoriesID, nil
}
func AssignMobileAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()
	var currentStatus string
	err := database.ST.QueryRowContext(ctx, `
		SELECT asset_status FROM assets WHERE id = $1 AND asset_type = 'mobile' AND deleted_at IS NULL
	`, req.AssetID).Scan(&currentStatus)

	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("mobile is not available for assignment")
	}

	tx, err := database.ST.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now() WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Transaction commit error:", err)
		return err
	}

	return nil
}

func AssignLaptopAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()

	tx, err := database.ST.BeginTx(ctx, nil) // ✔️ Start transaction
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var currentStatus string
	err = tx.QueryRowContext(ctx, `
		SELECT asset_status FROM assets 
		WHERE id = $1 AND asset_type = 'laptop' AND deleted_at IS NULL
		FOR UPDATE -- 🚩 Locks the row to prevent race conditions
	`, req.AssetID).Scan(&currentStatus)
	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("laptop is not available for assignment")
	}
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now() WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}
	return nil
}
func AssignMonitorAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()
	var currentStatus string
	err := database.ST.QueryRowContext(ctx, `
		SELECT asset_status FROM assets WHERE id = $1 AND asset_type = 'monitor' AND deleted_at IS NULL
	`, req.AssetID).Scan(&currentStatus)

	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("monitor is not available for assignment")
	}

	tx, err := database.ST.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now() WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}

	// ✅ Insert into asset_history for audit trail
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Transaction commit error:", err)
		return err
	}

	return nil
}
func AssignMouseAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()
	var currentStatus string
	err := database.ST.QueryRowContext(ctx, `
		SELECT asset_status FROM assets WHERE id = $1 AND asset_type = 'mouse' AND deleted_at IS NULL
	`, req.AssetID).Scan(&currentStatus)

	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("mouse is not available for assignment")
	}

	tx, err := database.ST.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now() WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Transaction commit error:", err)
		return err
	}

	return nil
}
func AssignHardDiskAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()
	var currentStatus string
	err := database.ST.QueryRowContext(ctx, `
		SELECT asset_status FROM assets WHERE id = $1 AND asset_type = 'harddisk' AND deleted_at IS NULL
	`, req.AssetID).Scan(&currentStatus)

	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("harddisk is not available for assignment")
	}

	tx, err := database.ST.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now() WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}
	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Transaction commit error:", err)
		return err
	}

	return nil
}
func AssignPendriveAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()

	var currentStatus string
	err := database.ST.QueryRowContext(ctx, `
		SELECT asset_status FROM assets 
		WHERE id = $1 AND asset_type = 'pendrive' AND deleted_at IS NULL
	`, req.AssetID).Scan(&currentStatus)

	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("pendrive is not available for assignment")
	}

	tx, err := database.ST.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now()
		WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Transaction commit error:", err)
		return err
	}

	return nil
}
func AssignAccessoriesAsset(req *models.AssignAssetRequest) error {
	ctx := context.Background()

	var currentStatus string
	err := database.ST.QueryRowContext(ctx, `
		SELECT asset_status FROM assets 
		WHERE id = $1 AND asset_type = 'accessories' AND deleted_at IS NULL
	`, req.AssetID).Scan(&currentStatus)

	if err != nil {
		log.Println("Asset check error:", err)
		return errors.New("invalid asset ID")
	}

	if currentStatus != "available" {
		return errors.New("accessories is not available for assignment")
	}

	tx, err := database.ST.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Transaction begin error:", err)
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_timeline (asset_id, assigned_to, assigned_by)
		VALUES ($1, $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("Timeline insert error:", err)
		return err
	}

	_, err = tx.ExecContext(ctx, `
		UPDATE assets SET asset_status = 'assigned', updated_at = now()
		WHERE id = $1
	`, req.AssetID)
	if err != nil {
		log.Println("Assets update error:", err)
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO asset_history (asset_id, old_status, new_status, employee_id, performed_by)
		VALUES ($1, 'available', 'assigned', $2, $3)
	`, req.AssetID, req.EmployeeID, req.AssignedBy)
	if err != nil {
		log.Println("History insert error:", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Transaction commit error:", err)
		return err
	}

	return nil
}

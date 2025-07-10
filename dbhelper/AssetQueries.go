package dbhelper

import (
	"StoreXApp/models"
	"database/sql"
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

//	func CreatePendriveAsset(tx *sql.Tx, req *models.CreatePendriveAssetRequest) (string, string, error) {
//		var assetID string
//		err := tx.QueryRow(`
//			INSERT INTO assets (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
//			VALUES ($1, $2, 'pendrive', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
//			RETURNING id
//		`,
//			req.Brand, req.Model, req.Category, req.OwnedBy, req.PurchasePrice, req.PurchasedDate, req.WarrantyStart, req.WarrantyExpire, req.CreatedBy,
//		).Scan(&assetID)
//		if err != nil {
//			return "", "", err
//		}
//
//		var pendriveID string
//		err = tx.QueryRow(`
//			INSERT INTO pendrive (asset_id, capacity, created_by)
//			VALUES ($1, $2, $3)
//			RETURNING id
//		`,
//			assetID, req.Capacity, req.CreatedBy,
//		).Scan(&pendriveID)
//		if err != nil {
//			return "", "", err
//		}
//
//		return assetID, pendriveID, nil
//	}
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

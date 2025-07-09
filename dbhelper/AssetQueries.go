package dbhelper

import (
	"StoreXApp/models"
	"database/sql"
	"fmt"
)

func CreateLaptopAsset(tx *sql.Tx, req *models.CreateLaptopAssetRequest) (string, string, error) {
	var assetID string
	var laptopID string

	// Insert into assets
	queryAsset := `
        INSERT INTO assets 
        (brand, model, asset_type, category, owned_by, purchase_price, purchased_date, warranty_start, warranty_expire, created_by)
        VALUES ($1, $2, 'laptop', $3, COALESCE($4, 'Remotestate'), $5, $6, $7, $8, $9)
        RETURNING id
    `
	err := tx.QueryRow(
		queryAsset,
		req.Brand,
		req.Model,
		req.Category,
		req.OwnedBy,
		req.PurchasePrice,
		req.PurchasedDate,
		req.WarrantyStart,
		req.WarrantyExpire,
		req.CreatedBy,
	).Scan(&assetID)
	if err != nil {
		return "", "", fmt.Errorf("failed to insert asset: %w", err)
	}

	// Insert into laptop
	queryLaptop := `
        INSERT INTO laptop 
        (asset_id, processor, ram, storage, os, created_by)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	err = tx.QueryRow(
		queryLaptop,
		assetID,
		req.Processor,
		req.Ram,
		req.Storage,
		req.OS,
		req.CreatedBy,
	).Scan(&laptopID)
	if err != nil {
		return "", "", fmt.Errorf("failed to insert laptop: %w", err)
	}

	return assetID, laptopID, nil
}

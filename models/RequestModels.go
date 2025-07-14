package models

type LoginRequest struct {
	Email  string `json:"email"`
	Fname  string `json:"-"`
	Lname  string `json:"-"`
	UserID string `json:"-"`
	Role   string `json:"role"`
}

type RegisterUserRequest struct {
	Email     string `json:"email" `
	PhoneNo   string `json:"phone_no"`
	Role      string `json:"role"`
	EmpType   string `json:"emp_type" `
	CreatedBy string `json:"-"`
	Fname     string `json:"-"`
	Lname     string `json:"-"`
	UserID    string `json:"-"`
}
type SelfRegisterUserRequest struct {
	Email     string `json:"email" `
	PhoneNo   string `json:"phone_no"`
	Role      string `json:"role"`
	EmpType   string `json:"emp_type" `
	CreatedBy string `json:"-"`
	Fname     string `json:"-"`
	Lname     string `json:"-"`
	UserID    string `json:"-"`
}

type CreateLaptopAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`
	// laptop
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Storage   string `json:"storage"`
	OS        string `json:"os"`
}

type CreateMobileAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`
	// Mobile-specific fields
	IMEI    string `json:"imei"`
	Ram     string `json:"ram"`
	Storage string `json:"storage"`
}
type CreateMouseAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`

	// Mouse-specific fields
	DPI            string `json:"dpi"`
	ConnectionType string `json:"connection_type"`
}

// MONITOR
type CreateMonitorAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`

	ScreenSize string `json:"screen_size"`
	Resolution string `json:"resolution"`
	PanelType  string `json:"panel_type"`
}

// HARDDISK
type CreateHarddiskAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`

	Capacity string `json:"capacity"`
	DiskType string `json:"disk_type"`
}

// PENDRIVE
type CreatePendriveAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`

	Capacity string `json:"capacity"`
}

type CreateAccessoriesAssetRequest struct {
	Brand          string   `json:"brand"`
	Model          string   `json:"model"`
	Category       *string  `json:"category"`
	OwnedBy        *string  `json:"owned_by"`
	PurchasePrice  *float64 `json:"purchase_price"`
	PurchasedDate  string   `json:"purchased_date"`
	WarrantyStart  *string  `json:"warranty_start"`
	WarrantyExpire *string  `json:"warranty_expire"`
	CreatedBy      string   `json:"-"`

	Name string  `json:"name"`
	Work *string `json:"work"`
}
type AssignAssetRequest struct {
	AssetID    string `json:"asset_id"`
	EmployeeID string `json:"employee_id"`
	AssignedBy string `json:"-"`
}
type EmployeeSearchByNameRequest struct {
	Name string `json:"name"`
}
type EmployeeSearchByPhoneNoRequest struct {
	PhoneNo string `json:"phone_no"`
}

type ChangeRoleRequest struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type RetrieveAssetRequest struct {
	EmployeeID string `json:"employee_id"`
	AssetID    string `json:"asset_id"`
	Reason     string `json:"reason"`
}
type DeleteEmp struct {
	EmployeeID string `json:"employee_id"`
}

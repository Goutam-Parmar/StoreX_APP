package models

type LoginRequest struct {
	Email  string `json:"email"`
	Fname  string
	Lname  string
	UserID string
	Role   string
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
	PhoneNo   string
	Role      string
	EmpType   string
	CreatedBy string
	Fname     string
	Lname     string
	UserID    string
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

	// Laptop-specific fields
	Processor string `json:"processor"`
	Ram       string `json:"ram"`
	Storage   string `json:"storage"`
	OS        string `json:"os"`
}

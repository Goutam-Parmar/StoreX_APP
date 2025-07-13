package models

import "time"

type LoginResponse struct {
	AccessToken    string  `json:"access_token"`
	RefreshToken   string  `json:"refresh_token"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}

type RegisterUserResponse struct {
	ID             string  `json:"id"`
	Fname          string  `json:"fname"`
	Lname          string  `json:"lname,omitempty"`
	Email          string  `json:"email"`
	PhoneNo        string  `json:"phone_no"`
	Role           string  `json:"role"`
	EmpType        string  `json:"emp_type"`
	CreatedBy      string  `json:"create_time"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}

type CreateLaptopAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	LaptopID       string  `json:"laptop_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}
type SelfRegisterUserResponse struct {
	ID             string  `json:"id"`
	Fname          string  `json:"fname"`
	Lname          string  `json:"lname,omitempty"`
	Email          string  `json:"email"`
	Role           string  `json:"role"`
	EmpType        string  `json:"emp_type"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}
type CreateMobileAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	MobileID       string  `json:"mobile_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}
type CreateMouseAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	MouseID        string  `json:"mouse_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}
type CreateMonitorAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	MonitorID      string  `json:"monitor_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}

type CreateHarddiskAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	HarddiskID     string  `json:"harddisk_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}

type CreatePendriveAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	PendriveID     string  `json:"pendrive_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}

type CreateAccessoriesAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	AccessoriesID  string  `json:"accessories_id"`
	AssetType      string  `json:"asset_type"`
	AssetStatus    string  `json:"asset_status"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}
type AssignAssetResponse struct {
	AssetID        string  `json:"asset_id"`
	AssignedTo     string  `json:"assigned_to"`
	AssignedBy     string  `json:"assigned_by"`
	AssignedAt     string  `json:"assigned_at"`
	ResponseTimeMs float64 `json:"response_time_ms"`
}
type EmployeeResponse struct {
	ID               string `json:"id"`
	Fname            string `json:"first_name"`
	Lname            string `json:"last_name"`
	Email            string `json:"email"`
	Role             string `json:"role"`
	EmpType          string `json:"emp_type"`
	HasAssetAssigned bool   `json:"has_asset_assigned"`
}
type AssetResponse struct {
	ID          string  `json:"id"`
	AssetType   string  `json:"asset_type"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	AssignedTo  *string `json:"assigned_to"` // NULL ho sakta hai
	AssetStatus string  `json:"asset_status"`
}

type EmployeeSearchByNameUser struct {
	ID      string `json:"id"`
	Fname   string `json:"fname"`
	Lname   string `json:"lname"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	EmpType string `json:"emp_type"`
}

type EmployeeSearchByNameResponse struct {
	Users          []EmployeeSearchByNameUser `json:"users"`
	ResponseTimeMs float64                    `json:"response_time_ms"`
}
type SingleEmployeeData struct {
	ID      string `json:"id"`
	Fname   string `json:"fname"`
	Lname   string `json:"lname"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	EmpType string `json:"emp_type"`
}

type EmployeeSearchByEmailResponse struct {
	User           *SingleEmployeeData `json:"user"`
	ResponseTimeMs float64             `json:"response_time_ms"`
}
type SingleEmployeePhoneData struct {
	ID      string `json:"id"`
	Fname   string `json:"fname"`
	Lname   string `json:"lname"`
	Email   string `json:"email"`
	PhoneNo string `json:"phone_no"`
	Role    string `json:"role"`
	EmpType string `json:"emp_type"`
}

type EmployeeSearchByPhoneNoResponse struct {
	Users          []SingleEmployeePhoneData `json:"users"`
	ResponseTimeMs float64                   `json:"response_time_ms"`
}
type DashboardResponse struct {
	TotalAssets             int64   `json:"total_assets"`
	AssignedAssets          int64   `json:"assigned_assets"`
	AvailableAssets         int64   `json:"available_assets"`
	WaitingForServiceAssets int64   `json:"waiting_for_service_assets"`
	InServiceAssets         int64   `json:"in_service_assets"`
	DamagedAssets           int64   `json:"damaged_assets"`
	ResponseTimeMs          float64 `json:"response_time_ms"`
}
type SimpleAssetInfoResponse struct {
	ID            string  `json:"id"`
	Brand         string  `json:"brand"`
	Model         string  `json:"model"`
	AssetType     string  `json:"asset_type"`
	PurchasePrice float64 `json:"purchase_price"`
	PurchasedDate string  `json:"purchased_date"`
	AssetStatus   string  `json:"asset_status"`

	// Optional fields
	Processor      *string `json:"processor,omitempty"`
	Ram            *string `json:"ram,omitempty"`
	Storage        *string `json:"storage,omitempty"`
	OS             *string `json:"os,omitempty"`
	IMEI           *string `json:"imei,omitempty"`
	DPI            *string `json:"dpi,omitempty"`
	ConnectionType *string `json:"connection_type,omitempty"`
	Capacity       *string `json:"capacity,omitempty"`
	DiskType       *string `json:"disk_type,omitempty"`
	ScreenSize     *string `json:"screen_size,omitempty"`
	Resolution     *string `json:"resolution,omitempty"`
	PanelType      *string `json:"panel_type,omitempty"`
	Name           *string `json:"name,omitempty"`
	Work           *string `json:"work,omitempty"`
}
type AssetListResponse struct {
	ID            string   `json:"id"`
	Brand         string   `json:"brand"`
	Model         string   `json:"model"`
	AssetType     string   `json:"asset_type"`
	Category      *string  `json:"category"`
	OwnedBy       *string  `json:"owned_by"`
	PurchasePrice *float64 `json:"purchase_price"`
	PurchasedDate string   `json:"purchased_date"`
}
type EmployeeAssetTimeline struct {
	AssetID    string     `json:"asset_id"`
	Brand      string     `json:"brand"`
	Model      string     `json:"model"`
	AssetType  string     `json:"asset_type"`
	Status     string     `json:"status"`
	AssignedAt time.Time  `json:"assigned_at"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
	Reason     *string    `json:"reason,omitempty"`
}
type MyDashboardResponse struct {
	TotalAssigned int                     `json:"total_assigned"`
	Assets        []MyDashboardAssetEntry `json:"assets"`
}

type MyDashboardAssetEntry struct {
	ID          string `json:"id"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	AssetType   string `json:"asset_type"`
	AssetStatus string `json:"asset_status"`
	AssignedAt  string `json:"assigned_at"`
}

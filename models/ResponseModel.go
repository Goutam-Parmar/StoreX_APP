package models

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

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

package models

// LoginResponse has the token that will returns in the Login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
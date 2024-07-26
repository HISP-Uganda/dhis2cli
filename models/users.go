package models

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName,omitempty"`
	Username    string `json:"username,omitempty"`
}

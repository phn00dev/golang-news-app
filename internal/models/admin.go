package models

import "time"

type Admin struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	AdminRole   string    `json:"admin_role"`
	AdminStatus string    `json:"admin_status"`
	Password    string    `json:"password"`
	LastLogin   time.Time `json:"last_login"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

package models

import (
	"time"

	"github.com/phn00dev/golang-news-app/internal/constants"
)

type Admin struct {
	ID          int                    `json:"id"`
	Username    string                 `json:"username"`
	FullName    string                 `json:"full_name"`
	PhoneNumber string                 `json:"phone_number"`
	AdminRole   constants.ADMIN_ROLE   `json:"admin_role"`
	AdminStatus constants.ADMIN_STATUS `json:"admin_status"`
	Password    string                 `json:"password"`
	LastLogin   time.Time              `json:"last_login"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

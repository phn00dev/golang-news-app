package dto

import (
	"github.com/phn00dev/golang-news-app/internal/constants"
)

type CreateAdminRequest struct {
	Username        string                 `json:"username" validate:"required,min=3,max=50"`
	FullName        string                 `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber     string                 `json:"phone_number" validate:"required,e164"`
	AdminRole       constants.ADMIN_ROLE   `json:"admin_role" validate:"required"`
	AdminStatus     constants.ADMIN_STATUS `json:"admin_status" validate:"required"`
	Password        string                 `json:"password" validate:"required,min=5,max=20"`
	ConfirmPassword string                 `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UpdateAdminRequest struct {
	Username    string                 `json:"username" validate:"required,min=3,max=50"`
	FullName    string                 `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber string                 `json:"phone_number" validate:"required,e164"`
	AdminRole   constants.ADMIN_ROLE   `json:"admin_role" validate:"required"`
	AdminStatus constants.ADMIN_STATUS `json:"admin_status" validate:"required"`
}

type UpdateAdminPasswordRequest struct {
	OldPassword        string `json:"old_password" validate:"required,min=5,max=20"`
	NewPassword        string `json:"new_password" validate:"required,min=5,max=20"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,eqfield=NewPassword"`
}

type LoginAdminRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

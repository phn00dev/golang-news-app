package dto

import (
	"github.com/phn00dev/golang-news-app/internal/constants"
	"github.com/phn00dev/golang-news-app/internal/models"
	timeformat "github.com/phn00dev/golang-news-app/internal/utils/timeFormat"
)

type AdminResponse struct {
	ID          int                    `json:"id"`
	Username    string                 `json:"username"`
	FullName    string                 `json:"full_name"`
	PhoneNumber string                 `json:"phone_number"`
	AdminRole   constants.ADMIN_ROLE   `json:"admin_role"`
	AdminStatus constants.ADMIN_STATUS `json:"admin_status"`
	LastLogin   string                 `json:"last_login"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

func GetAllAdminResponses(admins []models.Admin) []AdminResponse {
	adminResponses := make([]AdminResponse, len(admins))
	for i, admin := range admins {
		adminResponses[i] = GetOneAdminResponse(&admin)
	}
	return adminResponses
}

func GetOneAdminResponse(admin *models.Admin) AdminResponse {
	if admin == nil {
		return AdminResponse{}
	}

	return AdminResponse{
		ID:          admin.ID,
		Username:    admin.Username,
		FullName:    admin.FullName,
		PhoneNumber: admin.PhoneNumber,
		AdminRole:   admin.AdminRole,
		AdminStatus: admin.AdminStatus,
		LastLogin:   timeformat.FormatTime(admin.LastLogin),
		CreatedAt:   timeformat.FormatTime(admin.CreatedAt),
		UpdatedAt:   timeformat.FormatTime(admin.UpdatedAt),
	}
}

type LoginAdminResponse struct {
	Token       string                 `json:"token"`
	Username    string                 `json:"username"`
	FullName    string                 `json:"full_name"`
	AdminRole   constants.ADMIN_ROLE   `json:"admin_role"`
	AdminStatus constants.ADMIN_STATUS `json:"admin_status"`
}

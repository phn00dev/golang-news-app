package dto

import (
	"time"

	"github.com/phn00dev/golang-news-app/internal/constants"
	"github.com/phn00dev/golang-news-app/internal/models"
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

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02-01-2006 15:04:05")
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
		LastLogin:   formatTime(admin.LastLogin),
		CreatedAt:   formatTime(admin.CreatedAt),
		UpdatedAt:   formatTime(admin.UpdatedAt),
	}
}

type LoginAdminResponse struct {
	Token       string                 `json:"token"`
	Username    string                 `json:"username"`
	FullName    string                 `json:"full_name"`
	AdminRole   constants.ADMIN_ROLE   `json:"admin_role"`
	AdminStatus constants.ADMIN_STATUS `json:"admin_status"`
}

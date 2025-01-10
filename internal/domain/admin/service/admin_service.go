package service

import "github.com/phn00dev/golang-news-app/internal/domain/admin/dto"

type AdminService interface {
	GetAllAdmins() ([]dto.AdminResponse, error)
	GetOneAdmin(adminID int) (*dto.AdminResponse, error)
	CreateAdmin(createRequest dto.CreateAdminRequest) error
	UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error
	DeleteAdmin(adminID int) error
}

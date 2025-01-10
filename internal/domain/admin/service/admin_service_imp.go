package service

import (
	"github.com/phn00dev/golang-news-app/internal/domain/admin/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/repository"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return adminServiceImp{
		adminRepo: adminRepo,
	}
}

func (adminService adminServiceImp) GetAllAdmins() ([]dto.AdminResponse, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*dto.AdminResponse, error) {
	panic("admin service imp")
}

func (adminService adminServiceImp) CreateAdmin(createRequest dto.CreateAdminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) UpdateAdmin(updateRequest dto.UpdateAdminRequest) error {
	panic("admin service imp")
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	panic("admin service imp")
}

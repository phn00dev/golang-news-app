package service

import (
	"errors"

	"github.com/phn00dev/golang-news-app/internal/domain/admin/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/repository"
	"github.com/phn00dev/golang-news-app/internal/models"
	"github.com/phn00dev/golang-news-app/internal/utils/password"
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
	admins, err := adminService.adminRepo.GetAll()
	if err != nil {
		return nil, err
	}
	adminResponses := dto.GetAllAdminResponses(admins)
	return adminResponses, nil
}

func (adminService adminServiceImp) GetOneAdmin(adminID int) (*dto.AdminResponse, error) {
	if adminID <= 0 {
		return nil, errors.New("admin id not null")
	}
	admin, err := adminService.adminRepo.GetOneByID(adminID)
	if err != nil {
		return nil, err
	}
	adminResponse := dto.GetOneAdminResponse(admin)
	return &adminResponse, nil
}

func (adminService adminServiceImp) CreateAdmin(createRequest dto.CreateAdminRequest) error {
	if createRequest.Password != createRequest.ConfirmPassword {
		return errors.New("password and confirm password do not match")
	}

	hashPassword, err := password.HashPassword(createRequest.Password)
	if err != nil {
		return err
	}
	newAdmin := models.Admin{
		Username:    createRequest.Username,
		FullName:    createRequest.FullName,
		PhoneNumber: createRequest.PhoneNumber,
		AdminRole:   createRequest.AdminRole,
		AdminStatus: createRequest.AdminStatus,
		Password:    *hashPassword,
	}
	return adminService.adminRepo.Create(newAdmin)
}

func (adminService adminServiceImp) UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error {
	admin, err := adminService.adminRepo.GetOneByID(adminID)
	if err != nil {
		return err
	}

	admin.Username = updateRequest.Username
	admin.FullName = updateRequest.FullName
	admin.PhoneNumber = updateRequest.PhoneNumber
	admin.AdminRole = updateRequest.AdminRole
	admin.AdminStatus = updateRequest.AdminStatus

	return adminService.adminRepo.Update(admin.ID, *admin)
}

func (adminService adminServiceImp) DeleteAdmin(adminID int) error {
	if adminID <= 0 {
		return errors.New("admin ID must be a positive number")
	}
	return adminService.adminRepo.Delete(adminID)
}

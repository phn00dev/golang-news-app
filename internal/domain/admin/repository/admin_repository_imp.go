package repository

import (
	"github.com/phn00dev/golang-news-app/internal/models"
	"gorm.io/gorm"
)

type adminRepositoryImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{
		db: db,
	}
}

func (adminRepo adminRepositoryImp) GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	if err := adminRepo.db.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (adminRepo adminRepositoryImp) GetOneByID(adminID int) (*models.Admin, error) {
	var admin models.Admin
	if err := adminRepo.db.First(&admin, adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (adminRepo adminRepositoryImp) Create(admin models.Admin) error {
	return adminRepo.db.Create(&admin).Error
}

func (adminRepo adminRepositoryImp) Update(adminID int, admin models.Admin) error {
	return adminRepo.db.Model(&models.Admin{}).Where("id=?", adminID).Updates(&admin).Error
}

func (adminRepo adminRepositoryImp) Delete(adminID int) error {
	return adminRepo.db.Delete(&models.Admin{}, adminID).Error
}

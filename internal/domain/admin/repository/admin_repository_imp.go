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
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) GetOneByID(adminID int) (*models.Admin, error) {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Create(admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Update(adminID int, admin models.Admin) error {
	panic("admin repo imp")
}

func (adminRepo adminRepositoryImp) Delete(adminID int) error {
	panic("admin repo imp")
}

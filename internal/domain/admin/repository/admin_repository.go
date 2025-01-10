package repository

import "github.com/phn00dev/golang-news-app/internal/models"

type AdminRepository interface {
	GetAll() ([]models.Admin, error)
	GetOneByID(adminID int) (*models.Admin, error)
	Create(admin models.Admin) error
	Update(adminID int, admin models.Admin) error
	Delete(adminID int) error
}

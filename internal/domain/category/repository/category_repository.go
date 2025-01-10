package repository

import "github.com/phn00dev/golang-news-app/internal/models"

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetOne(categoryID int) (*models.Category, error)
	Create(category models.Category) error
	Update(categoryID int, category models.Category) error
	Delete(categoryID int) error

	// helper functions
	VerifyCategoryName(categoryName string) bool
}

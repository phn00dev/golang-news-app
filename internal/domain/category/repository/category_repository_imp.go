package repository

import (
	"github.com/phn00dev/golang-news-app/internal/models"
	"gorm.io/gorm"
)

type categoryRepositoryImp struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepositoryImp{
		db: db,
	}
}

func (categoryRepo categoryRepositoryImp) GetAll() ([]models.Category, error) {
	panic("category repo imp")
}

func (categoryRepo categoryRepositoryImp) GetOne(categoryID int) (*models.Category, error) {
	panic("category repo imp")
}

func (categoryRepo categoryRepositoryImp) Create(category models.Category) error {
	panic("category repo imp")
}

func (categoryRepo categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	panic("category repo imp")
}

func (categoryRepo categoryRepositoryImp) Delete(categoryID int) error {
	panic("category repo imp")
}

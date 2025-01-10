package repository

import (
	"errors"

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
	var categories []models.Category
	if err := categoryRepo.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (categoryRepo categoryRepositoryImp) GetOne(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := categoryRepo.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (categoryRepo categoryRepositoryImp) Create(category models.Category) error {
	return categoryRepo.db.Create(&category).Error
}

func (categoryRepo categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	return categoryRepo.db.Model(&models.Category{}).Where("id=?", categoryID).Error
}

func (categoryRepo categoryRepositoryImp) Delete(categoryID int) error {
	return categoryRepo.db.Delete(&models.Category{}, categoryID).Error
}

func (categoryRepo categoryRepositoryImp) VerifyCategoryName(categoryName string) bool {
	if err := categoryRepo.db.Where("category_name =?", categoryName).First(&models.Category{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	}
	return false
}

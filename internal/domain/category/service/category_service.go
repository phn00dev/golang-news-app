package service

import (
	"github.com/phn00dev/golang-news-app/internal/domain/category/dto"
)

type CategoryService interface {
	GetAllCategories() ([]dto.CategoryResponse, error)
	GetOneCategory(categoryID int) (*dto.CategoryResponse, error)
	CreateCategory(createRequest dto.CreateCategoryRequest) error
	UpdateCategory(categoryID int, updateRequest dto.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error
}

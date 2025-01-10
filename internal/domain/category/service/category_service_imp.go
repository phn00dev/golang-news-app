package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/category/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/category/repository"
	"github.com/phn00dev/golang-news-app/pkg/config"
)

type categoryServiceImp struct {
	categoryRepo repository.CategoryRepository
	config       *config.Config
}

func NewCategoryService(categoryRepo repository.CategoryRepository, config *config.Config) CategoryService {
	return categoryServiceImp{
		categoryRepo: categoryRepo,
		config:       config,
	}
}

func (categoryService categoryServiceImp) GetAllCategories() ([]dto.CategoryResponse, error) {
	panic("category service imp")
}

func (categoryService categoryServiceImp) GetOneCategory() (dto.CategoryResponse, error) {
	panic("category service imp")
}

func (categoryService categoryServiceImp) CreateCategory(ctx *fiber.Ctx, createRequest dto.CreateCategoryRequest) error {
	panic("category service imp")
}

func (categoryService categoryServiceImp) UpdateCategory(ctx *fiber.Ctx, categoryID int, updateRequest dto.UpdateCategoryRequest) error {
	panic("category service imp")
}

func (categoryService categoryServiceImp) DeleteCategory(categoryID int) error {
	panic("category service imp")
}

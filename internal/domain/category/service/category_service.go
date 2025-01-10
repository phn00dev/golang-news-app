package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/category/dto"
)

type CategoryService interface {
	GetAllCategories() ([]dto.CategoryResponse, error)
	GetOneCategory() (dto.CategoryResponse, error)
	CreateCategory(ctx *fiber.Ctx, createRequest dto.CreateCategoryRequest) error
	UpdateCategory(ctx *fiber.Ctx, categoryID int, updateRequest dto.UpdateCategoryRequest) error
	DeleteCategory(categoryID int) error
}

package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/category/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/category/service"
	"github.com/phn00dev/golang-news-app/internal/utils/response"
	"github.com/phn00dev/golang-news-app/internal/utils/validate"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return categoryHandlerImp{
		categoryService: categoryService,
	}
}

func (categoryHandler categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	categories, err := categoryHandler.categoryService.GetAllCategories()
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admins not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "success", categories)
}

func (categoryHandler categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || categoryID <= 0 {
		return response.Error(ctx, fiber.StatusInternalServerError, "Invalid category ID", err)
	}
	category, err := categoryHandler.categoryService.GetOneCategory(categoryID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "category not found", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", category)
}

func (categoryHandler categoryHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateCategoryRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "body parser error", err)
	}

	// validate data
	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validate error", err)
	}

	if err := categoryHandler.categoryService.CreateCategory(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "creation error", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", nil)
}

func (categoryHandler categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || categoryID <= 0 {
		return response.Error(ctx, fiber.StatusInternalServerError, "Invalid category ID", err)
	}

	var updateRequest dto.UpdateCategoryRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parser error", err)
	}
	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validate error", err)
	}

	if err := categoryHandler.categoryService.UpdateCategory(categoryID, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "updated error", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", nil)
}

func (categoryHandler categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || categoryID <= 0 {
		return response.Error(ctx, fiber.StatusInternalServerError, "Invalid category ID", err)
	}

	if err := categoryHandler.categoryService.DeleteCategory(categoryID); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "delete error", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", nil)
}

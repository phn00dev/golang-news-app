package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/category/service"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
}

func NewCategoryService(categoryService service.CategoryService) CategoryHandler {
	return categoryHandlerImp{
		categoryService: categoryService,
	}
}

func (categoryHandler categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (categoryHandler categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (categoryHandler categoryHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (categoryHandler categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

func (categoryHandler categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("category handler imp")
}

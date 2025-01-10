package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/post/service"
)

type postHandlerImp struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) PostHandler {
	return postHandlerImp{
		postService: postService,
	}
}

func (postHandler postHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("post handler imp")
}

func (postHandler postHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("post handler imp")
}

func (postHandler postHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("post handler imp")
}

func (postHandler postHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("post handler imp")
}

func (postHandler postHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("post handler imp")
}

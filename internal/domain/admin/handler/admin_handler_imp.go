package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/service"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: adminService,
	}
}

func (adminService adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	panic("admin service imp")
}

func (adminService adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	panic("admin service imp")
}

func (adminService adminHandlerImp) Create(ctx *fiber.Ctx) error {
	panic("admin service imp")
}

func (adminService adminHandlerImp) Update(ctx *fiber.Ctx) error {
	panic("admin service imp")
}

func (adminService adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	panic("admin service imp")
}

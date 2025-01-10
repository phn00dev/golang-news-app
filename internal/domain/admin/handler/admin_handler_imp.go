package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/service"
	"github.com/phn00dev/golang-news-app/internal/utils/response"
	"github.com/phn00dev/golang-news-app/internal/utils/validate"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) AdminHandler {
	return adminHandlerImp{
		adminService: adminService,
	}
}

func (adminHandler adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	admins, err := adminHandler.adminService.GetAllAdmins()
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "admins not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "success", admins)
}

func (adminHandler adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	adminID, err := ctx.ParamsInt("id")
	if err != nil || adminID <= 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid admin ID", err)
	}
	admin, err := adminHandler.adminService.GetOneAdmin(adminID)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "admin not found", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", admin)
}

func (adminHandler adminHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateAdminRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parser error", err)
	}
	// validate data

	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validate error", err)
	}

	// validate phone number

	verifyPhoneNumber := validate.ValidatePhoneNumber(createRequest.PhoneNumber)
	if !verifyPhoneNumber {
		return response.Error(ctx, fiber.StatusBadRequest, "error phone number", "incorrect phone number")
	}

	// create admin
	if err := adminHandler.adminService.CreateAdmin(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin creation error", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", nil)

}

func (adminHandler adminHandlerImp) Update(ctx *fiber.Ctx) error {
	adminID, err := ctx.ParamsInt("id")
	if err != nil || adminID <= 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid admin ID", err)
	}

	var updateRequest dto.UpdateAdminRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parser error", err)
	}

	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validation error", err)
	}

	// update admin
	if err := adminHandler.adminService.UpdateAdmin(adminID, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin update error", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", nil)
}

func (adminHandler adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	adminID, err := ctx.ParamsInt("id")
	if err != nil || adminID <= 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid admin ID", err)
	}

	// delete admin
	if err := adminHandler.adminService.DeleteAdmin(adminID); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin deletion error", err)
	}

	return response.Success(ctx, fiber.StatusOK, "success", nil)
}

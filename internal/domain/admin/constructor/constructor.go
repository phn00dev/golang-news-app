package constructor

import (
	"github.com/phn00dev/golang-news-app/internal/domain/admin/handler"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/repository"
	"github.com/phn00dev/golang-news-app/internal/domain/admin/service"
	"gorm.io/gorm"
)

var (
	adminRepo    repository.AdminRepository
	adminService service.AdminService
	AdminHandler handler.AdminHandler
)

func InitAdminRequirementCreator(db *gorm.DB) {
	adminRepo = repository.NewAdminRepository(db)
	adminService = service.NewAdminService(adminRepo)
	AdminHandler = handler.NewAdminHandler(adminService)
}

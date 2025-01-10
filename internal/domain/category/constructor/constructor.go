package constructor

import (
	"github.com/phn00dev/golang-news-app/internal/domain/category/handler"
	"github.com/phn00dev/golang-news-app/internal/domain/category/repository"
	"github.com/phn00dev/golang-news-app/internal/domain/category/service"
	"github.com/phn00dev/golang-news-app/pkg/config"
	"gorm.io/gorm"
)

var (
	categoryRepo    repository.CategoryRepository
	categoryService service.CategoryService
	CategoryHandler handler.CategoryHandler
)

func InitCategoryRequirementCreator(db *gorm.DB, config *config.Config) {
	categoryRepo = repository.NewCategoryRepository(db)
	categoryService = service.NewCategoryService(categoryRepo, config)
	CategoryHandler = handler.NewCategoryService(categoryService)
}

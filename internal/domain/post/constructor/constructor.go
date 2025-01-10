package constructor

import (
	categoryRepository "github.com/phn00dev/golang-news-app/internal/domain/category/repository"
	"github.com/phn00dev/golang-news-app/internal/domain/post/handler"
	"github.com/phn00dev/golang-news-app/internal/domain/post/repository"
	"github.com/phn00dev/golang-news-app/internal/domain/post/service"
	"github.com/phn00dev/golang-news-app/pkg/config"
	"gorm.io/gorm"
)

var (
	postRepo     repository.PostRepository
	categoryRepo categoryRepository.CategoryRepository
	postService  service.PostService
	PostHandler  handler.PostHandler
)

func InitPostRequirementCreator(db *gorm.DB, config *config.Config) {
	postRepo = repository.NewPostRepository(db)
	categoryRepo = categoryRepository.NewCategoryRepository(db)
	postService = service.NewPostService(categoryRepo, postRepo, config)
	PostHandler = handler.NewPostHandler(postService)
}

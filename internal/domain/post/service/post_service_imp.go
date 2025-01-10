package service

import (
	"github.com/gofiber/fiber/v2"
	categoryRepository "github.com/phn00dev/golang-news-app/internal/domain/category/repository"
	"github.com/phn00dev/golang-news-app/internal/domain/post/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/post/repository"
	"github.com/phn00dev/golang-news-app/pkg/config"
)

type postServiceImp struct {
	categoryRepo categoryRepository.CategoryRepository
	postRepo     repository.PostRepository
	config       *config.Config
}

func NewPostService(categoryRepo categoryRepository.CategoryRepository,
	postRepo repository.PostRepository, config *config.Config) PostService {
	return postServiceImp{
		categoryRepo: categoryRepo,
		postRepo:     postRepo,
		config:       config,
	}
}

func (postService postServiceImp) GetAllPosts() ([]dto.PostResponse, error) {
	panic("post service imp")
}

func (postService postServiceImp) GetOnePost(postID int) (*dto.PostResponse, error) {
	panic("post service imp")
}

func (postService postServiceImp) CreatePost(ctx *fiber.Ctx, createRequest dto.CreatePostRequest) error {
	panic("post service imp")
}

func (postService postServiceImp) UpdatePost(ctx *fiber.Ctx, postID int, updateRequest dto.UpdatePostRequest) error {
	panic("post service imp")
}

func (postService postServiceImp) DeletePost(postID int) error {
	panic("post service imp")
}

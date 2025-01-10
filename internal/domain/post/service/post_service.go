package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phn00dev/golang-news-app/internal/domain/post/dto"
)

type PostService interface {
	GetAllPosts() ([]dto.PostResponse, error)
	GetOnePost(postID int) (*dto.PostResponse, error)
	CreatePost(ctx *fiber.Ctx, createRequest dto.CreatePostRequest) error
	UpdatePost(ctx *fiber.Ctx, postID int, updateRequest dto.UpdatePostRequest) error
	DeletePost(postID int) error
}

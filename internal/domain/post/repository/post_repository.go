package repository

import "github.com/phn00dev/golang-news-app/internal/models"

type PostRepository interface {
	GetAll() ([]models.Post, error)
	GetOneByID(postID int) (*models.Post, error)
	Create(post models.Post) error
	Update(postID int, post models.Post) error
	Delete(postd int) error
}

package repository

import (
	"github.com/phn00dev/golang-news-app/internal/models"
	"gorm.io/gorm"
)

type postRepositoryImp struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return postRepositoryImp{
		db: db,
	}
}

func (postRepo postRepositoryImp) GetAll() ([]models.Post, error) {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) GetOneByID(postID int) (*models.Post, error) {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) Create(post models.Post) error {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) Update(postID int, post models.Post) error {
	panic("post repo imp")
}

func (postRepo postRepositoryImp) Delete(postd int) error {
	panic("post repo imp")
}

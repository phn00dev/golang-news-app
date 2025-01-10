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
	var posts []models.Post
	if err := postRepo.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (postRepo postRepositoryImp) GetOneByID(postID int) (*models.Post, error) {
	var post models.Post
	if err := postRepo.db.First(&post, postID).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (postRepo postRepositoryImp) Create(post models.Post) error {
	return postRepo.db.Create(&post).Error
}

func (postRepo postRepositoryImp) Update(postID int, post models.Post) error {
	return postRepo.db.Model(models.Post{}).Where("id=?", postID).Updates(post).Error
}

func (postRepo postRepositoryImp) Delete(postID int) error {
	return postRepo.db.Delete(&models.Post{}, postID).Error
}

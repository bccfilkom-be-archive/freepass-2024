package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (repo *PostRepository) FindById(id uint) (*entity.Post, error) {
	var post entity.Post
	if err := repo.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (repo *PostRepository) FindByCandidateId(candId uint) ([]entity.Post, error) {
	var posts []entity.Post
	if err := repo.db.Where("candidate_id = ?", candId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

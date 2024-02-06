package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type IPostRepository interface {
	CreatePost(post *domain.Posts) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) CreatePost(post *domain.Posts) error {
	tx := r.db.Begin()

	err := r.db.Create(post).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type IPostRepository interface {
	GetPostByCondition(post *domain.Posts, condition string, value any) error
	CreatePost(post *domain.Posts) error
	UpdatePost(post *domain.Posts) error
	DeletePost(post *domain.Posts) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) GetPostByCondition(post *domain.Posts, condition string, value any) error {
	err := r.db.First(post, condition, value).Error
	return err
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

func (r *PostRepository) UpdatePost(post *domain.Posts) error {
	tx := r.db.Begin()

	err := r.db.Save(post).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *PostRepository) DeletePost(post *domain.Posts) error {
	tx := r.db.Begin()

	err := r.db.Delete(post).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

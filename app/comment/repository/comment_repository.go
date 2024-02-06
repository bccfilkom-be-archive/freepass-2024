package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *domain.Comments) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) CreateComment(comment *domain.Comments) error {
	tx := r.db.Begin()

	err := r.db.Create(comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
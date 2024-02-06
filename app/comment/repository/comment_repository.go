package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	GetCommentByCondition(comment *domain.Comments, condition string, value any) error
	CreateComment(comment *domain.Comments) error
	DeleteComment(comment *domain.Comments) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) GetCommentByCondition(comment *domain.Comments, condition string, value any) error {
	err := r.db.First(comment, condition, value).Error
	return err
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

func (r *CommentRepository) DeleteComment(comment *domain.Comments) error {
	tx := r.db.Begin()

	err := r.db.Delete(comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

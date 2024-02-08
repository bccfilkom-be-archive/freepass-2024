package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	GetAllCommentByPostID(comments *[]domain.Comments, postId int) error
	GetCommentByCondition(comment *domain.Comments, condition string, value any) error
	CreateComment(comment *domain.Comments) error
	DeleteComment(comment *domain.Comments) error
	DeleteAllComment(comments *[]domain.Comments) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db}
}

func(r *CommentRepository) GetAllCommentByPostID(comments *[]domain.Comments, postId int) error {
	err := r.db.Model(domain.Comments{}).Preload("User").Find(comments, "post_id = ?", postId).Error
	return err
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
	return nil
}

func (r *CommentRepository) DeleteAllComment(comments *[]domain.Comments) error {
	tx := r.db.Begin()

	err := r.db.Delete(comments).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

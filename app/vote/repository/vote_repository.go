package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type IVoteRepository interface {
	GetVoteByCondition(voted *domain.Votes, condition string, value any) error
	Vote(voted *domain.Votes) error
	DeleteVote(voted *domain.Votes) error
}

type VoteRepository struct {
	db *gorm.DB
}

func NewVoteRepository(db *gorm.DB) *VoteRepository {
	return &VoteRepository{db}
}

func (r *VoteRepository) GetVoteByCondition(voted *domain.Votes, condition string, value any) error {
	err := r.db.First(voted, condition, value).Error
	return err
}

func (r *VoteRepository) Vote(voted *domain.Votes) error {
	tx := r.db.Begin()

	err := r.db.Create(voted).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *VoteRepository) DeleteVote(voted *domain.Votes) error {
	tx := r.db.Begin()

	err := r.db.Delete(voted).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

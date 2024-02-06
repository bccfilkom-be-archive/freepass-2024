package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type IElectionTimeRepository interface {
	SetStartAndEndTime(setElectionTime *domain.ElectionTimes) error
}

type ElectionTimeRepository struct {
	db *gorm.DB
}

func NewElectionTimeRepository(db *gorm.DB) *ElectionTimeRepository {
	return &ElectionTimeRepository{db}
}

func (r *ElectionTimeRepository) SetStartAndEndTime(setElectionTime *domain.ElectionTimes) error {
	tx := r.db.Begin()

	err := r.db.Create(setElectionTime).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"gorm.io/gorm"
	"time"
)

type ElectionPeriodRepository struct {
	db *gorm.DB
}

func NewElectionPeriodRepository(db *gorm.DB) *ElectionPeriodRepository {
	return &ElectionPeriodRepository{db}
}

func (repo *ElectionPeriodRepository) GetPeriod() (time.Time, time.Time, error) {
	var period entity.ElectionPeriod
	if err := repo.db.First(&period, 1).Error; err != nil {
		return time.Time{}, time.Time{}, err
	}
	return period.Start, period.End, nil
}

package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"gorm.io/gorm"
)

type CandidateRepository struct {
	db       *gorm.DB
	UserRepo *UserRepository
}

func NewCandidateRepository(db *gorm.DB, userRepo *UserRepository) *CandidateRepository {
	return &CandidateRepository{db, userRepo}
}

func (repo *CandidateRepository) FindAll() ([]entity.Candidate, error) {
	var candidates []entity.Candidate
	if err := repo.db.Find(&candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}

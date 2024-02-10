package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/util/roles"
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

func (repo *CandidateRepository) Promote(user *entity.User) (uint, error) {
	candidate := entity.Candidate{User: *user}
	return candidate.ID, repo.db.Transaction(func(tx *gorm.DB) error {
		if err := repo.UserRepo.SetRole(user, roles.Candidate); err != nil {
			return err
		}
		if err := repo.db.Create(&candidate).Error; err != nil {
			return err
		}
		return nil
	})
}

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

func (repo *CandidateRepository) FindById(id uint) (*entity.Candidate, error) {
	var candidate entity.Candidate
	if err := repo.db.First(&candidate, id).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (repo *CandidateRepository) FindByUserId(id uint) (*entity.Candidate, error) {
	var candidate entity.Candidate
	if err := repo.db.Model(&entity.Candidate{}).Where("user_id = ?", id).First(&candidate).Error; err != nil {
		return nil, err
	}
	return &candidate, nil
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

func (repo *CandidateRepository) AddVote(candidate *entity.Candidate, voter *entity.User) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		if err := repo.UserRepo.DisableVoter(voter); err != nil {
			return err
		}
		return repo.db.Model(&candidate).Update("vote_count", gorm.Expr("vote_count + ?", 1)).Error
	})
}

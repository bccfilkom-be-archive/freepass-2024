package repository

import (
	"freepass-bcc/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByCondition(user *domain.Users, condition string, value string) error
	CreateUser(user *domain.Users) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByCondition(user *domain.Users, condition string, value string) error {
	err := r.db.Model(domain.Users{}).Where(condition, value).First(user).Error
	return err
}

func (r *UserRepository) CreateUser(user *domain.Users) error {
	tx := r.db.Begin()

	err := r.db.Create(user).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

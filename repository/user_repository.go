package repository

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"github.com/nathakusuma/bcc-be-freepass-2024/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user *entity.User) error {
	return repo.db.Create(&user).Error
}

func (repo *UserRepository) ExistsUsername(username string) (bool, error) {
	var count int64
	err := repo.db.Model(&entity.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserRepository) ExistsId(id uint) (bool, error) {
	var count int64
	err := repo.db.Model(&entity.User{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := repo.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindById(id uint) (*entity.User, error) {
	var user entity.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *entity.User, updates *model.UpdateUserRequest) error {
	return repo.db.Model(user).Updates(updates).Error
}

func (repo *UserRepository) Delete(user *entity.User) error {
	return repo.db.Delete(user).Error
}

func (repo *UserRepository) SetRole(user *entity.User, role string) error {
	return repo.db.Model(&user).Update("role", role).Error
}

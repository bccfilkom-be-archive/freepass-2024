package userRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func FindOne(ID string) (*models.User, error) {
	var user *models.User

	if err := database.DB.First(&user, "id = ?", ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func FindUser(u *schemas.UserLoginRequest) (*models.User, error) {
	var user *models.User

	err := database.DB.First(&user, "email = ?", u.Email).Error
	if user.ID == "" {
		return nil, err
	}
	return user, nil
}

func FindUserAdmin(u *schemas.UserLoginRequest) (*models.User, error) {
	var user *models.User

	err := database.DB.First(&user, "email = ? AND is_admin = true", u.Email).Error
	if user.ID == "" {
		return nil, err
	}
	return user, nil
}

func FindAll() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

package repositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
)

func RegisterUser(user *models.User) (schemas.UserResponse, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return schemas.UserResponse{}, err
	}
	newUser := schemas.UserResponse{Email: user.Email, ID: user.ID, Name: user.Name}
	return newUser, nil
}

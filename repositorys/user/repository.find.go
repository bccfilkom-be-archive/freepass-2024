package repositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
	"github.com/AkbarFikri/freepassBCC-2024/schemas"
	"github.com/AkbarFikri/freepassBCC-2024/utils"

)

func FindUser(u *schemas.UserLoginRequest) (*models.User, error) {
	var user *models.User

	err := database.DB.First(&user, "email = ?", u.Email).Error
	if user.ID == "" {
		return nil, err
	}
	if err := utils.ComparePassword(user.Password, u.Password); err != nil {
		return nil, err
	}
	return user, nil
}

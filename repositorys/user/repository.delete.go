package userRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func DeleteUser(ID string) error {
	if err := database.DB.Delete(&models.User{}, ID).Error; err != nil {
		return err
	}
	return nil
}

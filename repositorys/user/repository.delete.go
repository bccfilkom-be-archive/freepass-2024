package userRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
	profileRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/profile"

)

func DeleteUser(ID string) error {
	if err := profileRepositorys.DeleteProfileUser(ID); err != nil {
		return err
	}
	if err := database.DB.Where("id = ?", ID).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}

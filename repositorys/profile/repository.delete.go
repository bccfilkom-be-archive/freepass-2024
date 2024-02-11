package profileRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
	postRepositorys "github.com/AkbarFikri/freepassBCC-2024/repositorys/post"
)

func DeleteProfileUser(ID string) error {
	if err := postRepositorys.DeleteUserPost(ID); err != nil {
		return err
	}
	if err := database.DB.Where("user_id = ?", ID).Delete(&models.Comment{}).Error; err != nil {
		return err
	}
	if err := database.DB.Where("user_id = ?", ID).Delete(&models.Vote{}).Error; err != nil {
		return err
	}
	if err := database.DB.Where("user_id = ?", ID).Delete(&models.Profile{}).Error; err != nil {
		return err
	}
	return nil
}

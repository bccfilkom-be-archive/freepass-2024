package postRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func DeletePost(ID string) error {
	if err := database.DB.Delete(&models.Post{}, ID).Error; err != nil {
		return err
	}
	return nil
}

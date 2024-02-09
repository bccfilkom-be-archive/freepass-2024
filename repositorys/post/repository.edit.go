package postRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func EditPost(data *models.Post) error {
	if err := database.DB.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

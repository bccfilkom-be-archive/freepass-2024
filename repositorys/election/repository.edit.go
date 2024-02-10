package electionRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func EditElection(data *models.Election) error {
	if err := database.DB.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

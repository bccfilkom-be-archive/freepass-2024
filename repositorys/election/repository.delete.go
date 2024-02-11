package electionRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func DeleteElection(ID string) error {
	if err := database.DB.Model(&models.Election{}).Where("id = ?", ID).Delete(&models.Election{}).Error; err != nil {
		return err
	}
	return nil
}

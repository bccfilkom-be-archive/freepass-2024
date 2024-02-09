package electionRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func FindSpecificElection(ID string) (*models.Election, error) {
	var election *models.Election
	if err := database.DB.First(&election).Where("id = ?", ID).Error; err != nil {
		return election, err
	}
	return election, nil
}

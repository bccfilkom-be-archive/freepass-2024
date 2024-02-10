package electionRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func CreateNewElection(data *models.Election) (*models.Election, error) {
	if err := database.DB.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
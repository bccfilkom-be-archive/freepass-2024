package voteRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func CreateNewVote(data *models.Vote) error {
	if err := database.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

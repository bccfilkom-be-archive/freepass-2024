package postRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func CreateNewPost(data *models.Post) (*models.Post, error) {
	if err := database.DB.Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
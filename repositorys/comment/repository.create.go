package commentRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func CreateComment(data *models.Comment) error {
	if err := database.DB.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

package commentRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func FindOne(ID string) (*models.Comment, error) {
	var comment *models.Comment
	if err := database.DB.Where("id = ?", ID).First(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

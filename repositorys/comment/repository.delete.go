package commentRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func DeleteComment(ID string) error {
	if err := database.DB.Delete(&models.Comment{ID: ID}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCommentPost(ID string) error {
	if err := database.DB.Where("post_id = ?", ID).Delete(&models.Comment{}).Error; err != nil {
		return err
	}
	return nil
}

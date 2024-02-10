package postRepositorys

import (
	"gorm.io/gorm/clause"

	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func DeletePost(ID string) error {
	if err := database.DB.Where("post_id = ?", ID).Delete(&models.Comment{}).Error; err != nil {
		return err
	}
	if err := database.DB.Where("id = ?", ID).Delete(&models.Post{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCandidatePost(ID string) error {
	var Posts []models.Post

	if err := database.DB.Select("id").Where("candidate_id = ?", ID).Find(&Posts).Error; err != nil {
		return err
	}

	ids := []string{}

	for _, b := range Posts {
		ids = append(ids, b.ID)
	}

	idsInterface := make([]interface{}, len(ids))
	for i, id := range ids {
		idsInterface[i] = id
	}

	if err := database.DB.Model(&models.Comment{}).Where("post_id IN ?", idsInterface).Delete(&models.Comment{}).Error; err != nil {
		return err
	}

	if err := database.DB.Select(clause.Associations).Where("candidate_id = ?", ID).Delete(&models.Post{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserPost(ID string) error {
	var Posts []models.Post

	if err := database.DB.Select("id").Where("user_id = ?", ID).Find(&Posts).Error; err != nil {
		return err
	}

	ids := []string{}

	for _, b := range Posts {
		ids = append(ids, b.ID)
	}

	idsInterface := make([]interface{}, len(ids))
	for i, id := range ids {
		idsInterface[i] = id
	}

	if err := database.DB.Model(&models.Comment{}).Where("post_id IN ?", idsInterface).Delete(&models.Comment{}).Error; err != nil {
		return err
	}

	if err := database.DB.Select(clause.Associations).Where("user_id = ?", ID).Delete(&models.Post{}).Error; err != nil {
		return err
	}
	return nil
}

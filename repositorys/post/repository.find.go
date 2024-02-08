package postRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func FindAll() ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func FindAllByElectionID(ID string) ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Where("election_id = ?", ID).Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func FindAllByCandidateID(ID string) ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Where("candidate_id = ?", ID).Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func FindAllByElectionAndCandidate(ElectionID string, CandidateID string) ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Where("election_id = ? AND candidate_id = ?", ElectionID, CandidateID).Find(&posts).Error; err != nil {
		return posts, err
	}
	return posts, nil
}

func FindOne(ID string) (*models.Post, error) {
	var post *models.Post
	if err := database.DB.Where("id = ?", ID).First(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

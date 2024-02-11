package voteRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func FindUserVoteForCandidate(userID string, candidateID string) (*models.Vote, error) {
	var vote *models.Vote
	if err := database.DB.Where("user_id = ? AND candidate_id = ?", userID, candidateID).First(&vote).Error; err != nil {
		return nil, err
	}
	return vote, nil
}
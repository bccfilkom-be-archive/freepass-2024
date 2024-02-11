package voteRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func DeleteCandidateVote(ID string) error {
	if err := database.DB.Where("candidate_id = ?", ID).Delete(&models.Vote{}).Error; err != nil {
		return err
	}
	return nil
}

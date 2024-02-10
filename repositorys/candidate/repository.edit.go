package candidateRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func VoteCandidate(ID string, currentVote int) error {
	if err := database.DB.Model(&models.Candidate{}).Where("id = ?", ID).Update("vote", currentVote+1).Error; err != nil {
		return err
	}
	return nil
}

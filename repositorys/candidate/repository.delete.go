package candidateRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func DeleteCandidate(ID string) error {
	if err := database.DB.Model(&models.CandidateInformation{}).Where("candidate_id = ?", ID).Delete(&models.CandidateInformation{}).Error; err != nil {
		return err
	}
	if err := database.DB.Model(&models.Candidate{}).Where("id = ?", ID).Delete(&models.Candidate{}).Error; err != nil {
		return err
	}
	return nil
}

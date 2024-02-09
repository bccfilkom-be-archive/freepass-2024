package candidateRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func FindSpecificCandidate(ID string) (*models.Candidate, error) {
	var candidate *models.Candidate
	if err := database.DB.First(&candidate).Where("id = ?", ID).Error; err != nil {
		return candidate, err
	}
	return candidate, nil
}

func FindCandidateInformations(ID string) (*models.CandidateInformation, error) {
	var informations *models.CandidateInformation
	if err := database.DB.Where("candidate_id = ?", ID).First(&informations).Error; err != nil {
		return informations, err
	}
	return informations, nil
}

func FindCandidateInElection(userID string, electionID string) (*models.Candidate, error) {
	var candidate *models.Candidate
	if err := database.DB.Where("user_id = ? AND election_id = ?", userID, electionID).First(&candidate).Error; err != nil {
		return nil, err
	}
	return candidate, nil
}

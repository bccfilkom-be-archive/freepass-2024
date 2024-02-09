package candidateRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func CreatCandidate(candidate *models.Candidate) (*models.Candidate, error) {
	if err := database.DB.Create(&candidate).Error; err != nil {
		return nil, err
	}
	return candidate, nil
}

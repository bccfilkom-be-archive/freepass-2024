package candidateRepositorys

import (
	"github.com/lib/pq"

	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func CreatCandidate(candidate *models.Candidate) (*models.Candidate, error) {
	if err := database.DB.Create(&candidate).Error; err != nil {
		return nil, err
	}
	return candidate, nil
}

func CreateCandidateInformation(ID string) error {
	text := []string{"1", "2"}
	information := models.CandidateInformation{Vision: "test", Mision: "test", CandidateID: ID, Experience: pq.StringArray(text), Achievement: pq.StringArray(text)}
	if err := database.DB.Create(&information).Error; err != nil {
		return err
	}
	return nil
}

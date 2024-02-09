package candidateRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func DeleteCandidate(ID string) error {
	if err := database.DB.Delete(&models.Candidate{}, ID).Error; err != nil {
		return err
	}
	return nil
}

package electionRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"

)

func FindSpecificElection(ID string) (*models.Election, error) {
	var election *models.Election
	if err := database.DB.Where("id = ?", ID).Find(&election).Error; err != nil {
		return nil, err
	}
	return election, nil
}

func FindElectionNumber(num int, electionID string) error {
	var candidate *models.Candidate
	if err := database.DB.Where("election_num = ? AND election_id = ?", num, electionID).First(&candidate).Error; err != nil {
		return err
	}
	return nil
}

func FindAllElection() ([]models.Election, error) {
	var elections []models.Election
	if err := database.DB.Find(&elections).Error; err != nil {
		return nil, err
	}
	return elections, nil
}

package database

import "github.com/AkbarFikri/freepassBCC-2024/models"

func Migrate() {
	DB.AutoMigrate(models.Election{})
	DB.AutoMigrate(models.Candidate{}, models.CandidateInformation{}, models.Vote{})
	DB.AutoMigrate(models.User{}, models.Profile{}, models.Post{}, models.Comment{})
}

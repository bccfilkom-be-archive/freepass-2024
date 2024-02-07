package profileRepositorys

import (
	"time"

	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func CreateProfile(ID string) error {
	date := time.Date(2004, 05, 24, 0, 0, 0, 0, time.Local)
	profile := models.Profile{UserID: ID, Fullname: "User", Birthplace: "", Birthdate: date, Address: "", ProfileUrl: ""}
	if err := database.DB.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

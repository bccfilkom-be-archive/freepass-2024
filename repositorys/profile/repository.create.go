package profileRepositorys

import (
	"github.com/AkbarFikri/freepassBCC-2024/database"
	"github.com/AkbarFikri/freepassBCC-2024/models"
)

func CreateProfile(ID string) error {
	profile := models.Profile{UserID: ID, Fullname: "User", Birthplace: "", Birthdate: "2005-01-01 00:00:00", Address: "", ProfileUrl: ""}
	if err := database.DB.Create(&profile).Error; err != nil {
		return err
	}
	return nil
}

package database

import (
	"github.com/nathakusuma/bcc-be-freepass-2024/entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Candidate{},
		&entity.Comment{},
		&entity.Post{},
		&entity.ElectionPeriod{},
	)
	return err
}

package database

import "freepass-bcc/domain"

func Migrate() {
	DB.AutoMigrate(
		&domain.Users{},
		&domain.Posts{},
		&domain.Comments{},
	)
}

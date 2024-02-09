package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID   uint
	Username string
	Body     string
}

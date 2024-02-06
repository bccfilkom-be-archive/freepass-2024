package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Author string
	Title  string
	Body   string
}

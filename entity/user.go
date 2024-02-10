package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:VARCHAR(64); NOT NULL"`
	Username string `gorm:"unique; type:VARCHAR(32); NOT NULL;"`
	Password string `gorm:"type:VARCHAR(128); NOT NULL"`
	Bio      string `gorm:"type:VARCHAR(256)"`
	Role     string `gorm:"type:ENUM('admin', 'candidate', 'user'); NOT NULL; DEFAULT:'user'"`
	CanVote  bool   `gorm:"NOT NULL; DEFAULT:true"`
}

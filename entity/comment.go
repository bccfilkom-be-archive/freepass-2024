package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"NOT NULL"`
	UserID  uint
	PostID  uint `gorm:"NOT NULL"`
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Post    Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

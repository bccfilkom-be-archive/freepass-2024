package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

)

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null; unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	IsAdmin   bool      `gorm:"default:false" json:"isAdmin"`
	Profile   Profile   `gorm:"foreignKey:UserID"`
}

func (user *User) BeforeCreate(c *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now().Local()
	return nil
}

func (user *User) BeforeUpdate(db *gorm.DB) error {
	user.UpdatedAt = time.Now().Local()
	return nil
}

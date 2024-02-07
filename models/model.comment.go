package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
	PostID    string    `gorm:"not null" json:"post_id"`
	UserID    string    `gorm:"not null" json:"user_id"`
	Message   string    `gorm:"not null" json:"message"`
}

func (cm *Comment) BeforeCreate(c *gorm.DB) (err error) {
	cm.ID = uuid.New().String()
	cm.CreatedAt = time.Now().Local()
	return nil
}

func (cm *Comment) BeforeUpdate(db *gorm.DB) error {
	cm.UpdatedAt = time.Now().Local()
	return nil
}

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

)

type Election struct {
	ID            string      `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time   `gorm:"not null" json:"createdAt"`
	UpdatedAt     time.Time   `gorm:"not null" json:"updatedAt"`
	Title         string      `gorm:"not null" json:"title"`
	Description   string      `gorm:"not null" json:"description"`
	StartTime     time.Time   `gorm:"not null" json:"start_time"`
	EndTime       time.Time   `gorm:"not null" json:"end_time"`
	MakeByAdminID string      `gorm:"not null" json:"admin_id"`
	Status        string      `gorm:"default:Pending" json:"status"`
	Candidates    []Candidate `gorm:"foreignKey:ElectionID"`
	Posts         []Post      `gorm:"foreignKey:ElectionID"`
	Votes         []Vote      `gorm:"foreignKey:ElectionID"`
}

func (e *Election) BeforeCreate(c *gorm.DB) (err error) {
	e.ID = uuid.New().String()
	e.CreatedAt = time.Now().Local()
	return nil
}

func (e *Election) BeforeUpdate(db *gorm.DB) error {
	e.UpdatedAt = time.Now().Local()
	return nil
}

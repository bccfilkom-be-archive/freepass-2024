package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Candidate struct {
	ID           string                 `gorm:"primaryKey"`
	CreatedAt    time.Time              `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time              `gorm:"not null" json:"updatedAt"`
	UserID       string                 `gorm:"not null" json:"user_id"`
	ElectionNum  int                    `gorm:"not null" json:"election_number"`
	ElectionID   string                 `gorm:"not null" json:"election_id"`
	Post         []Post                 `gorm:"foreignKey:CandidateID"`
	Informations []CandidateInformation `gorm:"foreignKey:CandidateID"`
}

func (cd *Candidate) BeforeCreate(c *gorm.DB) (err error) {
	cd.ID = uuid.New().String()
	cd.CreatedAt = time.Now().Local()
	return nil
}

func (cd *Candidate) BeforeUpdate(db *gorm.DB) error {
	cd.UpdatedAt = time.Now().Local()
	return nil
}

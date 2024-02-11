package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CandidateInformation struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"not null" json:"updatedAt"`
	CandidateID string         `gorm:"not null" json:"candidate_id"`
	Vision      string         `gorm:"not null" json:"vision"`
	Mision      string         `gorm:"not null" json:"mision"`
	Experience  pq.StringArray `gorm:"type:text[]" json:"experience"`
	Achievement pq.StringArray `gorm:"type:text[]" json:"achivement"`
}

func (cdi *CandidateInformation) BeforeCreate(c *gorm.DB) (err error) {
	cdi.ID = uuid.New().String()
	cdi.CreatedAt = time.Now().Local()
	return nil
}

func (cdi *CandidateInformation) BeforeUpdate(db *gorm.DB) error {
	cdi.UpdatedAt = time.Now().Local()
	return nil
}

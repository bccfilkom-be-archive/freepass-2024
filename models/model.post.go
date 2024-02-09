package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
	UserID      string    `gorm:"not null" json:"user_id"`
	CandidateID string    `gorm:"not null" json:"candidate_id"`
	ElectionID  string    `gorm:"not null" json:"election_id"`
	PictureUrl  string    `gorm:"not null" json:"picture_url"`
	Caption     string    `gorm:"not null" json:"caption"`
	Comments    []Comment `gorm:"foreignKey:PostID"`
}

func (p *Post) BeforeCreate(c *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now().Local()
	return nil
}

func (p *Post) BeforeUpdate(db *gorm.DB) error {
	p.UpdatedAt = time.Now().Local()
	return nil
}

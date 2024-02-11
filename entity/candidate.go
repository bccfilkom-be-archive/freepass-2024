package entity

import (
	"gorm.io/gorm"
	"time"
)

type Candidate struct {
	gorm.Model
	UserID    uint   `gorm:"UNIQUE"`
	VoteCount uint64 `gorm:"NOT NULL; DEFAULT:0;"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Posts     []Post
}

func (c *Candidate) AfterDelete(tx *gorm.DB) error {
	return tx.Model(&Post{}).Where("candidate_id = ?", c.ID).Update("deleted_at", time.Now()).Error
}

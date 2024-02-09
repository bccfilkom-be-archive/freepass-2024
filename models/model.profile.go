package models

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	UserID       string    `gorm:"primaryKey" json:"user_id"`
	CreatedAt    time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"not null" json:"updatedAt"`
	Fullname     string    `gorm:"not null" json:"full_name"`
	Birthplace   string    `gorm:"not null" json:"birth_place"`
	Birthdate    time.Time `gorm:"not null" json:"birth_date"`
	Address      string    `gorm:"not null" json:"address"`
	ProfileUrl   string    `gorm:"not null" json:"profile_url"`
	Posts        []Post    `gorm:"foreignKey:UserID"`
	Comments     []Comment `gorm:"foreignKey:UserID"`
	VoteHistorys []Vote    `gorm:"foreignKey:UserID"`
}

func (p *Profile) BeforeCreate(c *gorm.DB) (err error) {
	p.CreatedAt = time.Now().Local()
	return nil
}

func (p *Profile) BeforeUpdate(db *gorm.DB) error {
	p.UpdatedAt = time.Now().Local()
	return nil
}

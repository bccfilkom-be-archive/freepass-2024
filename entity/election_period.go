package entity

import (
	"gorm.io/gorm"
	"time"
)

type ElectionPeriod struct {
	gorm.Model
	Start time.Time `gorm:"NOT NULL"`
	End   time.Time `gorm:"NOT NULL"`
}

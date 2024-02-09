package models

import (
	"time"

	"gorm.io/gorm"
)

type Election struct {
	gorm.Model
	Year           uint
	StartDate      time.Time
	EndDate        time.Time
	WinnerUsername string
}

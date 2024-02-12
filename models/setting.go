package models

import "time"

type Setting struct {
	ID    int       `json:"id"`
	Dates time.Time `json:"date"`
}

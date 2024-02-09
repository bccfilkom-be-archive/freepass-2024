package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Dates       time.Time `json:"timestamp"`
}

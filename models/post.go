package models

import "time"

type Post struct {
	ID          int       `json:"id"`
	UserID		int 	  `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Dates       time.Time `json:"timestamp"`
}

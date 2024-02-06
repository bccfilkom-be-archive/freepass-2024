package domain

import "time"

type Posts struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Post      string    `json:"post"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type PostRequest struct {
	Post string `json:"post"`
}

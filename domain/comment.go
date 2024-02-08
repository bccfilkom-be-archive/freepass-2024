package domain

import "time"

type Comments struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"-"`
	User      Users     `json:"-"`
}

type CommentRequest struct {
	Comment string `json:"comment"`
}

type CommentResponse struct {
	ID        int       `json:"id"`
	User      string    `json:"user"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

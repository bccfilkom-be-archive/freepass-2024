package domain

import "time"

type Comments struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"-"`
}

type CommentRequest struct {
	Comment   string    `json:"comment"`
}

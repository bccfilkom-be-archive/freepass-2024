package model

import "time"

type GetCommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AddCommentRequest struct {
	Content string `binding:"required"`
}

type AddCommentResponse struct {
	ID uint `json:"id"`
}

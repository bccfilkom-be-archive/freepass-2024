package model

import (
	"time"
)

type CreatePostRequest struct {
	Title   string `binding:"required"`
	Content string
}

type GetPostResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CandidateID uint      `json:"candidate_id"`
	Comments    []GetCommentResponse
}

type UpdatePostRequest struct {
	Title   string
	Content string
}

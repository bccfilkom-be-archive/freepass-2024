package model

import (
	"time"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio"`
}

type RegisterUserResponse struct {
	ID uint `json:"id"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	ID    uint   `json:"id"`
	Token string `json:"token"`
}

type GetUserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Bio       string    `json:"bio"`
	Role      string    `json:"role"`
	CanVote   bool      `json:"can_vote"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

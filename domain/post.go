package domain

import "time"

type Posts struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Post      string     `json:"post"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	Comments  []Comments `json:"-" gorm:"foreignKey:post_id;references:id"`
	User      Users      `json:"-"`
}

type PostRequest struct {
	Post string `json:"post"`
}

type PostResponse struct {
	ID        int    `json:"id"`
	Candidate string `json:"candidate"`
	Post      string `json:"post"`
}

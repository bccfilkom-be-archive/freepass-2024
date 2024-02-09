package domain

import "time"

type Votes struct {
	UserId   int       `json:"user_id" gorm:"primary_key"`
	Choice   string    `json:"choice"`
	VoteTime time.Time `json:"vote_time"`
}

type VoteResponse struct {
	User     string    `json:"user"`
	Choice   string    `json:"choice"`
	VoteTime time.Time `json:"vote_time"`
}

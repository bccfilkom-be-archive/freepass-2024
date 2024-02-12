package models

type Vote struct {
	ID          int `json:"id"`
	UserID      int `json: "user_id"`
	CandidateID int `json:"candidate_id"`
}

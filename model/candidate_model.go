package model

type GetCandidateResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Votes  uint64 `json:"votes"`
}

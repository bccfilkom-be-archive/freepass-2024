package model

type PromoteCandidateResponse struct {
	ID uint `json:"candidate_id"`
}

type GetCandidateResponse struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Votes  uint64 `json:"votes"`
}

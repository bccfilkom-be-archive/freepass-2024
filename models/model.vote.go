package models

type Vote struct {
	ID          int    `gorm:"primaryKey; autoIncrement" json:"id"`
	CandidateID string `gorm:"not null" json:"candidate_id"`
	UserID      string `gorm:"not null" json:"user_id"`
	ElectionID  string `gorm:"not null" json:"election_id"`
}

package models

type Comment struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	PostID  int    `json:"post_id"`
	Comment string `json:"comment"`
}

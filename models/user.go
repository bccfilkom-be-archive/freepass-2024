package models

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"` // Agak rancu karena kalau int tidak bisa + atau - cth +62
	Role        string `json:"role"`
	Token       string `json:"token"`
}

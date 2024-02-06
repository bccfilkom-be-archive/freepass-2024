package domain

import "time"

type Users struct {
	ID        int       `json:"id" gorm:"primary key"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	Role      string    `json:"role" gorm:"type:enum('ADMIN', 'USER', 'CANDIDATE')"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Posts     []Posts   `json:"-" gorm:"foreignKey:user_id;references:id"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

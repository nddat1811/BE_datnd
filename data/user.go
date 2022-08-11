package data

import (
	"time"
)

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Role      int        `json:"role"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserPassword struct {
	Password        string `json:"password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserUpdateAccount struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserUpdatePassword struct {
	CurrentPassword string `json:"current_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

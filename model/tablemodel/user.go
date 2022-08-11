package tablemodel

import (
	"time"
)

type User struct {
	Id        int        `gorm:"column:id;primary_key"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	Phone     string     `gorm:"column:phone"`
	Role      int        `gorm:"column:role"`
	Password  string     `gorm:"column:password"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (b *User) TableName() string {
	return "users"
}

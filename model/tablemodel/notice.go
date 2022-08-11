package tablemodel

import (
	"time"
)

type Notice struct {
	Id        int        `gorm:"column:id;primary_key"`
	Status    int        `gorm:"column:status"`
	Title     string     `gorm:"column:title"`
	Content   string     `gorm:"column:content"`
	UserID    int        `gorm:"column:user_id"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	Due_date  *time.Time `gorm:"column:due_date"`
}

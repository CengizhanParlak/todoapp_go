package model

import "time"

type Todo struct {
	ID         uint      `gorm:"column:ID"`
	Task       string    `gorm:"column:name"`
	Category   string    `gorm:"column:category"`
	TodoStatus int       `gorm:"column:status"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
	CreatedBy  string    `gorm:"column:created_by"`
}

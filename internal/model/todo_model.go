package model

import "time"

type Todo struct {
	ID         uint       `gorm:"column:ID" json:"id"`
	Task       string     `gorm:"column:name" json:"task"`
	Category   string     `gorm:"column:category" json:"category"`
	TodoStatus TodoStatus `gorm:"column:status" json:"todoStatus"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt  time.Time  `gorm:"column:deleted_at" json:"deletedAt"`
	CreatedBy  string     `gorm:"column:created_by" json:"createdBy"`
}

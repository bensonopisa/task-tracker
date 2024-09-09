package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `gorm:"size:30;index" json:"title"`
	Status      string    `gorm:"not null;default:pending;index" json:"status"`
	Description *string   `gorm:"size:150" json:"description"`
	DueDate     time.Time `gorm:"index" json:"due_date"`
	UserID      uint      `json:"user_id"`
}

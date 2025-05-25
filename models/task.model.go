package models

import (
	"time"
)

// Task represents a task in the system
type Task struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"` // Foreign key to User
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"size:255;not null" json:"content"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	Status    string    `gorm:"size:50;not null" json:"status"` // e.g., "To Do", "In Progress", "Done"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for Task model
func (Task) TableName() string {
	return "tasks"
}

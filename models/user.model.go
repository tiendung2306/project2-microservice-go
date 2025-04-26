package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"` // Password không được trả về trong JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for User model
func (User) TableName() string {
	return "users"
}

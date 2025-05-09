package models

import (
	"time"

	"gorm.io/gorm"
)

// Notification represents a notification in the system
type Notification struct {
	gorm.Model
	UserID     uint      `json:"user_id"`
	Email      string    `json:"email"`
	Subject    string    `json:"subject"`
	Content    string    `json:"content"`
	Status     string    `json:"status"` // e.g., "pending", "sent", "failed"
	SentAt     time.Time `json:"sent_at"`
	NotifyType string    `json:"notify_type"` // e.g., "email", "sms", etc.
}

package dto

import (
	"time"
)

// CreateNotificationRequest represents the request data for creating a notification
type CreateNotificationRequest struct {
	UserID     uint   `json:"user_id" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Subject    string `json:"subject" binding:"required"`
	Content    string `json:"content" binding:"required"`
	NotifyType string `json:"notify_type" binding:"required"`
}

// NotificationResponse represents the response data for a notification
type NotificationResponse struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	Email      string    `json:"email"`
	Subject    string    `json:"subject"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
	SentAt     time.Time `json:"sent_at"`
	NotifyType string    `json:"notify_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// SendEmailRequest represents a request to send an email
type SendEmailRequest struct {
	To      string `json:"to" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}

// GetNotificationsResponse represents the response for getting all notifications
type GetNotificationsResponse struct {
	Notifications []NotificationResponse `json:"notifications"`
}

package models

import (
	"time"
)

type RefreshToken struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Token     string    `gorm:"type:text;not null" json:"-"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	IsRevoked bool      `gorm:"default:false" json:"is_revoked"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package config

import "time"

// JWTConfig chứa các cấu hình cho JWT
type JWTConfig struct {
	SecretKey       string
	TokenExpiration time.Duration
}

// NewJWTConfig tạo cấu hình JWT mới
func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:       "your-secret-key-should-be-loaded-from-env", // Nên lấy từ biến môi trường
		TokenExpiration: 24 * time.Hour,                              // Token hết hạn sau 24 giờ
	}
}

package service

import (
	"errors"
	"project2-microservice-go/internal/auth-service/config"
	"project2-microservice-go/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims là cấu trúc claims cho JWT
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type IJWTService interface {
	GenerateToken(user *models.User) (string, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
	GenerateRefreshToken(user *models.User) (*jwt.Token, string, error)
}

type jwtService struct {
	config *config.JWTConfig
}

func NewJWTService(config *config.JWTConfig) IJWTService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) GenerateToken(user *models.User) (string, error) {
	claims := &JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.TokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.SecretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.config.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token không hợp lệ")
	}

	return claims, nil
}

func (s *jwtService) GenerateRefreshToken(user *models.User) (*jwt.Token, string, error) {
	claims := &JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // Refresh token kéo dài 7 ngày
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte(s.config.SecretKey))
	return token, stringToken, err
}

package service

import (
	"project2-microservice-go/errors"
	"project2-microservice-go/internal/auth-service/dto"
	"project2-microservice-go/internal/auth-service/repository"
	"project2-microservice-go/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(request *dto.LoginRequest) (*dto.AuthResponse, error)               // Login method
	Register(request *dto.RegisterRequest) (*dto.AuthResponse, error)         // Register method
	RefreshToken(request *dto.RefreshTokenRequest) (*dto.AuthResponse, error) // RefreshToken method
	ChangePassword(userID uint, request *dto.ChangePasswordRequest) error     // ChangePassword method
}

type authService struct {
	authRepository repository.IAuthRepository
	jwtService     IJWTService
}

func NewAuthService(ar repository.IAuthRepository, js IJWTService) IAuthService {
	return &authService{
		authRepository: ar,
		jwtService:     js,
	}
}

func (s *authService) Login(request *dto.LoginRequest) (*dto.AuthResponse, error) {

	if !s.authRepository.IsEmailExists(request.Email) {
		return nil, errors.NewUnauthorizedError("Email or password is not correct", nil)
	}

	user, err := s.authRepository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get user", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, errors.NewUnauthorizedError("Email or password is not correct", nil)
	}
	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate token", err)
	}
	refreshToken, stringRefreshToken, err := s.jwtService.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate refresh token", err)
	}

	tokenExpiresAt, err := refreshToken.Claims.GetExpirationTime()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get expiration time", err)
	}
	tokenCreatedAt, err := refreshToken.Claims.GetIssuedAt()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get issued at time", err)
	}

	refreshTokenModel := &models.RefreshToken{
		Token:     stringRefreshToken,
		UserID:    user.ID,
		IsRevoked: false,
		ExpiresAt: tokenExpiresAt.Time,
		CreatedAt: tokenCreatedAt.Time,
		UpdatedAt: tokenCreatedAt.Time,
	}

	if err := s.authRepository.SaveRefreshToken(refreshTokenModel); err != nil {
		return nil, errors.NewInternalServerError("Failed to save refresh token", err)
	}

	return &dto.AuthResponse{
		Token:        token,
		RefreshToken: stringRefreshToken,
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
	}, nil
}

func (s *authService) Register(request *dto.RegisterRequest) (*dto.AuthResponse, error) {
	if s.authRepository.IsEmailExists(request.Email) {
		return nil, errors.NewConflictError("Email already exists", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to hash password", err)
	}

	newUser := &models.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.authRepository.CreateUser(newUser); err != nil {
		return nil, errors.NewInternalServerError("Failed to create user", err)
	}

	token, err := s.jwtService.GenerateToken(newUser)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate token", err)
	}
	refreshToken, stringRefreshToken, err := s.jwtService.GenerateRefreshToken(newUser)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate refresh token", err)
	}

	tokenExpiresAt, err := refreshToken.Claims.GetExpirationTime()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get expiration time", err)
	}
	tokenCreatedAt, err := refreshToken.Claims.GetIssuedAt()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get issued at time", err)
	}

	refreshTokenModel := &models.RefreshToken{
		Token:     stringRefreshToken,
		UserID:    newUser.ID,
		IsRevoked: false,
		ExpiresAt: tokenExpiresAt.Time,
		CreatedAt: tokenCreatedAt.Time,
		UpdatedAt: tokenCreatedAt.Time,
	}

	if err := s.authRepository.SaveRefreshToken(refreshTokenModel); err != nil {
		return nil, errors.NewInternalServerError("Failed to save refresh token", err)
	}

	return &dto.AuthResponse{
		Token:        token,
		RefreshToken: stringRefreshToken,
		UserID:       newUser.ID,
		Username:     newUser.Username,
		Email:        newUser.Email,
	}, nil
}

func (s *authService) RefreshToken(request *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {
	refreshToken, err := s.jwtService.ValidateToken(request.RefreshToken)
	if err != nil {
		return nil, errors.NewUnauthorizedError("Invalid refresh token", err)
	}

	user, err := s.authRepository.GetUserByEmail(refreshToken.Email)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get user", err)
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate token", err)
	}
	newRefreshToken, stringRefreshToken, err := s.jwtService.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to generate refresh token", err)
	}

	tokenExpiresAt, err := newRefreshToken.Claims.GetExpirationTime()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get expiration time", err)
	}
	tokenCreatedAt, err := newRefreshToken.Claims.GetIssuedAt()
	if err != nil {
		return nil, errors.NewInternalServerError("Failed to get issued at time", err)
	}

	refreshTokenModel := &models.RefreshToken{
		Token:     stringRefreshToken,
		UserID:    user.ID,
		IsRevoked: false,
		ExpiresAt: tokenExpiresAt.Time,
		CreatedAt: tokenCreatedAt.Time,
		UpdatedAt: tokenCreatedAt.Time,
	}

	if err := s.authRepository.SaveRefreshToken(refreshTokenModel); err != nil {
		return nil, errors.NewInternalServerError("Failed to save refresh token", err)
	}

	return &dto.AuthResponse{
		Token:        token,
		RefreshToken: stringRefreshToken,
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
	}, nil
}

func (s *authService) ChangePassword(userID uint, request *dto.ChangePasswordRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternalServerError("Failed to hash password", err)
	}

	if err := s.authRepository.UpdateUserPassword(userID, string(hashedPassword)); err != nil {
		return errors.NewInternalServerError("Failed to update password", err)
	}

	return nil
}

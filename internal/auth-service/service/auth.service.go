package service

import "project2-microservice-go/internal/auth-service/repository"

type IAuthService interface {
	Login(email, password string) (string, error)    // Login method
	Register(email, password string) (string, error) // Register method
}

type authService struct {
	authRepository repository.IAuthRepository
}

func NewAuthService(ar repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: ar,
	}
}

func (s *authService) Login(email, password string) (string, error) {
	// Logic to handle login
	// This is a placeholder implementation
	if email == "admin@gmail.com" && password == "password" {
		return "Login successful", nil
	}
	return "", nil
}

func (s *authService) Register(email, password string) (string, error) {
	// Logic to handle registration
	// This is a placeholder implementation
	if email != "" && password != "" {
		return "Registration successful", nil
	}
	return "", nil
}

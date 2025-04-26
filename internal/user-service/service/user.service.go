package service

import (
	"project2-microservice-go/internal/user-service/repository"
	"project2-microservice-go/models"
)

type IUserService interface {
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	users, err := us.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

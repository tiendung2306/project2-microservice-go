package service

import (
	"project2-microservice-go/internal/user-service/dto"
	"project2-microservice-go/internal/user-service/repository"
	"project2-microservice-go/models"
	"project2-microservice-go/rabbitmq"
)

type IUserService interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(request *dto.CreateUserRequest) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, request *dto.UpdateUserRequest) (*models.User, error)
	DeleteUser(id string) error
}

type userService struct {
	userRepository repository.IUserRepository
	rabbitmqClient *rabbitmq.RabbitMQ
}

func NewUserService(userRepository repository.IUserRepository, rabbitmqClient *rabbitmq.RabbitMQ) IUserService {
	return &userService{
		userRepository: userRepository,
		rabbitmqClient: rabbitmqClient,
	}
}

func (us *userService) GetAllUsers() ([]models.User, error) {
	users, err := us.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *userService) CreateUser(request *dto.CreateUserRequest) (*models.User, error) {
	user, err := us.userRepository.Create(request)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserByID(id string) (*models.User, error) {
	user, err := us.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) UpdateUser(id string, request *dto.UpdateUserRequest) (*models.User, error) {
	user, err := us.userRepository.Update(id, request)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) DeleteUser(id string) error {
	err := us.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

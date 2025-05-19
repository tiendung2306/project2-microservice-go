package repository

import (
	"project2-microservice-go/internal/user-service/dto"
	"project2-microservice-go/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() ([]models.User, error)
	Create(user *dto.CreateUserRequest) (*models.User, error)
	FindByID(id string) (*models.User, error)
	Update(id string, request *dto.UpdateUserRequest) (*models.User, error)
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := ur.db.Find(&users)
	return users, result.Error
}

func (ur *userRepository) Create(request *dto.CreateUserRequest) (*models.User, error) {
	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	result := ur.db.Create(&user)
	return &user, result.Error
}

func (ur *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	result := ur.db.First(&user, id)
	return &user, result.Error
}

func (ur *userRepository) Update(id string, request *dto.UpdateUserRequest) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if request.Username != "" {
		updates["username"] = request.Username
	}
	if request.Email != "" {
		updates["email"] = request.Email
	}

	if len(updates) == 0 {
		return &user, nil
	}

	if err := ur.db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Delete(id string) error {
	result := ur.db.Delete(&models.User{}, id)
	return result.Error
}

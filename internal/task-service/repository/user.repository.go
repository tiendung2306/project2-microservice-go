package repository

import (
	"project2-microservice-go/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserEmailByID(userID uint) (string, error)
	IsUserExists(userID uint) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetUserEmailByID(userID uint) (string, error) {
	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.Email, nil
}

func (ur *userRepository) IsUserExists(userID uint) (bool, error) {
	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {
		return false, err
	}
	return true, nil
}

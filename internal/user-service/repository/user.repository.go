package repository

import (
	"project2-microservice-go/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	FindAll() ([]models.User, error)
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

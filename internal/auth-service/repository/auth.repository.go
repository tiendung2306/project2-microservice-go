package repository

import (
	"project2-microservice-go/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
	IsEmailExists(email string) bool
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	SaveRefreshToken(token *models.RefreshToken) error
	UpdateUserPassword(userID uint, hashedPassword string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{
		db: db,
	}
}

func (ar *authRepository) IsEmailExists(email string) bool {
	var count int64
	ar.db.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (ar *authRepository) CreateUser(user *models.User) error {
	err := ar.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ar.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *authRepository) SaveRefreshToken(token *models.RefreshToken) error {
	err := ar.db.Create(token).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *authRepository) UpdateUserPassword(userID uint, hashedPassword string) error {
	err := ar.db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
	if err != nil {
		return err
	}
	return nil
}

package database

import (
	"log"
	"project2-microservice-go/models"
)

// AutoMigrate tự động tạo hoặc cập nhật schema database dựa trên các models
func (s *service) AutoMigrate() error {
	log.Println("Creating schema database...")

	err := s.db.AutoMigrate(
		&models.User{},
		&models.Task{},
		&models.RefreshToken{},
	)

	if err != nil {
		log.Printf("Create schema fail: %v", err)
		return err
	}

	log.Println("Create schema successfully")
	return nil
}

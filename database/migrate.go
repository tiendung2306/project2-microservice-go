package database

import (
	"log"
	"project2-microservice-go/models"
)

func (s *service) AutoMigrate() error {
	log.Println("Creating schema database...")

	err := s.db.AutoMigrate(
		&models.User{},
		&models.Task{},
		&models.RefreshToken{},
		&models.Notification{},
	)

	if err != nil {
		log.Printf("Create schema fail: %v", err)
		return err
	}

	log.Println("Create schema successfully")
	return nil
}

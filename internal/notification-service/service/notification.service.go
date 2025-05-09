package service

import (
	"fmt"
	"net/smtp"
	"os"
	"project2-microservice-go/internal/notification-service/dto"
	"project2-microservice-go/internal/notification-service/repository"
)

type INotificationService interface {
	CreateNotification(request *dto.CreateNotificationRequest) (*dto.NotificationResponse, error)
	GetAllNotifications() ([]dto.NotificationResponse, error)
	GetNotificationByID(id uint) (*dto.NotificationResponse, error)
	GetNotificationsByUserID(userID uint) ([]dto.NotificationResponse, error)
	SendEmail(req *dto.SendEmailRequest) error
	UpdateNotificationStatus(id uint, status string) error
	DeleteNotification(id uint) error
}

type notificationService struct {
	notificationRepo repository.INotificationRepository
}

func NewNotificationService(notificationRepo repository.INotificationRepository) INotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

func (ns *notificationService) CreateNotification(request *dto.CreateNotificationRequest) (*dto.NotificationResponse, error) {
	return ns.notificationRepo.CreateNotification(request)
}

func (ns *notificationService) GetAllNotifications() ([]dto.NotificationResponse, error) {
	return ns.notificationRepo.GetAllNotifications()
}

func (ns *notificationService) GetNotificationByID(id uint) (*dto.NotificationResponse, error) {
	return ns.notificationRepo.GetNotificationByID(id)
}

func (ns *notificationService) GetNotificationsByUserID(userID uint) ([]dto.NotificationResponse, error) {
	return ns.notificationRepo.GetNotificationsByUserID(userID)
}

func (ns *notificationService) SendEmail(req *dto.SendEmailRequest) error {
	// Retrieve email configuration from environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpFrom := os.Getenv("SMTP_FROM")

	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" || smtpFrom == "" {
		return fmt.Errorf("missing SMTP configuration")
	}

	// Set up authentication information
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Compose email
	to := []string{req.To}
	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", smtpFrom, req.To, req.Subject, req.Body))

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, to, msg)
	if err != nil {
		return err
	}

	return nil
}

func (ns *notificationService) UpdateNotificationStatus(id uint, status string) error {
	return ns.notificationRepo.UpdateNotificationStatus(id, status)
}

func (ns *notificationService) DeleteNotification(id uint) error {
	return ns.notificationRepo.DeleteNotification(id)
}

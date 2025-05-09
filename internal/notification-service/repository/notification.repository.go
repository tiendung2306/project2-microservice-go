package repository

import (
	"project2-microservice-go/internal/notification-service/dto"
	"project2-microservice-go/models"
	"time"

	"gorm.io/gorm"
)

type INotificationRepository interface {
	CreateNotification(request *dto.CreateNotificationRequest) (*dto.NotificationResponse, error)
	GetAllNotifications() ([]dto.NotificationResponse, error)
	GetNotificationByID(id uint) (*dto.NotificationResponse, error)
	GetNotificationsByUserID(userID uint) ([]dto.NotificationResponse, error)
	UpdateNotificationStatus(id uint, status string) error
	DeleteNotification(id uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) INotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (nr *notificationRepository) CreateNotification(request *dto.CreateNotificationRequest) (*dto.NotificationResponse, error) {
	notification := models.Notification{
		UserID:     request.UserID,
		Email:      request.Email,
		Subject:    request.Subject,
		Content:    request.Content,
		Status:     "pending",
		NotifyType: request.NotifyType,
	}

	result := nr.db.Create(&notification)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.NotificationResponse{
		ID:         notification.ID,
		UserID:     notification.UserID,
		Email:      notification.Email,
		Subject:    notification.Subject,
		Content:    notification.Content,
		Status:     notification.Status,
		SentAt:     notification.SentAt,
		NotifyType: notification.NotifyType,
		CreatedAt:  notification.CreatedAt,
		UpdatedAt:  notification.UpdatedAt,
	}, nil
}

func (nr *notificationRepository) GetAllNotifications() ([]dto.NotificationResponse, error) {
	var notifications []models.Notification
	result := nr.db.Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	var notificationResponses []dto.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, dto.NotificationResponse{
			ID:         notification.ID,
			UserID:     notification.UserID,
			Email:      notification.Email,
			Subject:    notification.Subject,
			Content:    notification.Content,
			Status:     notification.Status,
			SentAt:     notification.SentAt,
			NotifyType: notification.NotifyType,
			CreatedAt:  notification.CreatedAt,
			UpdatedAt:  notification.UpdatedAt,
		})
	}

	return notificationResponses, nil
}

func (nr *notificationRepository) GetNotificationByID(id uint) (*dto.NotificationResponse, error) {
	var notification models.Notification
	result := nr.db.First(&notification, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.NotificationResponse{
		ID:         notification.ID,
		UserID:     notification.UserID,
		Email:      notification.Email,
		Subject:    notification.Subject,
		Content:    notification.Content,
		Status:     notification.Status,
		SentAt:     notification.SentAt,
		NotifyType: notification.NotifyType,
		CreatedAt:  notification.CreatedAt,
		UpdatedAt:  notification.UpdatedAt,
	}, nil
}

func (nr *notificationRepository) GetNotificationsByUserID(userID uint) ([]dto.NotificationResponse, error) {
	var notifications []models.Notification
	result := nr.db.Where("user_id = ?", userID).Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	var notificationResponses []dto.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, dto.NotificationResponse{
			ID:         notification.ID,
			UserID:     notification.UserID,
			Email:      notification.Email,
			Subject:    notification.Subject,
			Content:    notification.Content,
			Status:     notification.Status,
			SentAt:     notification.SentAt,
			NotifyType: notification.NotifyType,
			CreatedAt:  notification.CreatedAt,
			UpdatedAt:  notification.UpdatedAt,
		})
	}

	return notificationResponses, nil
}

func (nr *notificationRepository) UpdateNotificationStatus(id uint, status string) error {
	result := nr.db.Model(&models.Notification{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status": status,
		"sent_at": func() time.Time {
			if status == "sent" {
				return time.Now()
			}
			return time.Time{}
		}(),
	})
	return result.Error
}

func (nr *notificationRepository) DeleteNotification(id uint) error {
	result := nr.db.Delete(&models.Notification{}, id)
	return result.Error
}

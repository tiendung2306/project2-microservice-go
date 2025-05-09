package controller

import (
	"net/http"
	"project2-microservice-go/internal/notification-service/dto"
	"project2-microservice-go/internal/notification-service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationService service.INotificationService
}

func NewNotificationController(notificationService service.INotificationService) *NotificationController {
	return &NotificationController{
		notificationService: notificationService,
	}
}

// CreateNotification handles the creation of a new notification
func (nc *NotificationController) CreateNotification(c *gin.Context) {
	var request dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification, err := nc.notificationService.CreateNotification(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

// GetAllNotifications returns all notifications
func (nc *NotificationController) GetAllNotifications(c *gin.Context) {
	notifications, err := nc.notificationService.GetAllNotifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetNotificationsResponse{Notifications: notifications})
}

// GetNotificationByID returns a notification by ID
func (nc *NotificationController) GetNotificationByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	notification, err := nc.notificationService.GetNotificationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// GetNotificationsByUserID returns all notifications for a specific user
func (nc *NotificationController) GetNotificationsByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	notifications, err := nc.notificationService.GetNotificationsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.GetNotificationsResponse{Notifications: notifications})
}

// SendEmail handles sending an email notification
func (nc *NotificationController) SendEmail(c *gin.Context) {
	var request dto.SendEmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := nc.notificationService.SendEmail(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}

// UpdateNotificationStatus updates the status of a notification
func (nc *NotificationController) UpdateNotificationStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	var request struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = nc.notificationService.UpdateNotificationStatus(uint(id), request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification status updated"})
}

// DeleteNotification deletes a notification
func (nc *NotificationController) DeleteNotification(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	err = nc.notificationService.DeleteNotification(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted"})
}

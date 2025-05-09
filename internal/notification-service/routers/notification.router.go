package routers

import (
	"project2-microservice-go/database"
	"project2-microservice-go/internal/notification-service/controller"
	"project2-microservice-go/internal/notification-service/repository"
	"project2-microservice-go/internal/notification-service/service"
	"project2-microservice-go/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterNotificationRoutes registers notification-related routes
func RegisterNotificationRoutes(router *gin.RouterGroup, jwtMiddleware *middleware.JWTAuthMiddleware) {
	db := database.New()
	notificationRepository := repository.NewNotificationRepository(db.DB())
	notificationService := service.NewNotificationService(notificationRepository)
	notificationController := controller.NewNotificationController(notificationService)
	notificationGroup := router.Group("/notifications", jwtMiddleware.AuthRequired())
	{
		notificationGroup.POST("", notificationController.CreateNotification)
		notificationGroup.GET("", notificationController.GetAllNotifications)
		notificationGroup.GET("/:id", notificationController.GetNotificationByID)
		notificationGroup.GET("/user/:user_id", notificationController.GetNotificationsByUserID)
		notificationGroup.PUT("/:id/status", notificationController.UpdateNotificationStatus)
		notificationGroup.DELETE("/:id", notificationController.DeleteNotification)
	}

	// Email sending endpoint
	router.POST("/sendemail", jwtMiddleware.AuthRequired(), notificationController.SendEmail)
}

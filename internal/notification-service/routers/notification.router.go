package routers

import (
	"fmt"
	"log"
	"project2-microservice-go/database"
	"project2-microservice-go/internal/notification-service/consumer"
	"project2-microservice-go/internal/notification-service/controller"
	"project2-microservice-go/internal/notification-service/repository"
	"project2-microservice-go/internal/notification-service/service"
	"project2-microservice-go/middleware"
	"project2-microservice-go/rabbitmq"

	"github.com/gin-gonic/gin"
)

// RegisterNotificationRoutes registers notification-related routes
func RegisterNotificationRoutes(router *gin.RouterGroup, jwtMiddleware *middleware.JWTAuthMiddleware) {
	// Initialize database
	db := database.New()

	rabbitmqClient, err := rabbitmq.Initialize()
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	// Initialize repositories
	notificationRepository := repository.NewNotificationRepository(db.DB())

	// Initialize services
	notificationService := service.NewNotificationService(notificationRepository)

	// Initialize consumer
	notificationConsumer := consumer.NewNotificationConsumer(notificationService, rabbitmqClient)
	if err := notificationConsumer.Start(); err != nil {
		log.Printf("Failed to start notification consumer: %v", err)
		panic(fmt.Sprintf("Failed to start notification consumer: %v", err))
	}

	// Initialize controller
	notificationController := controller.NewNotificationController(notificationService)

	// Register routes
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

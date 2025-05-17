package routers

import (
	"fmt"
	"log"
	"project2-microservice-go/database"
	"project2-microservice-go/internal/task-service/controller"
	"project2-microservice-go/internal/task-service/repository"
	"project2-microservice-go/internal/task-service/service"
	"project2-microservice-go/middleware"
	"project2-microservice-go/rabbitmq"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.RouterGroup, jwtMiddleware *middleware.JWTAuthMiddleware) {
	db := database.New()
	taskRepository := repository.NewTaskRepository(db.DB())
	userRepository := repository.NewUserRepository(db.DB())

	// Initialize RabbitMQ with retry logic
	var rabbitmqClient *rabbitmq.RabbitMQ
	var err error
	maxRetries := 5
	retryDelay := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		rabbitmqClient, err = rabbitmq.NewRabbitMQ()
		if err == nil {
			log.Println("Successfully connected to RabbitMQ")
			break
		}
		log.Printf("Failed to connect to RabbitMQ (attempt %d/%d): %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		log.Printf("Failed to connect to RabbitMQ after %d attempts: %v", maxRetries, err)
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	taskService := service.NewTaskService(taskRepository, userRepository, rabbitmqClient)
	taskController := controller.NewTaskController(taskService)
	taskGroup := router.Group("/task")          // Group all /task routes
	taskGroup.Use(jwtMiddleware.AuthRequired()) // Apply JWT authentication middleware
	{
		taskGroup.GET("", taskController.GetAllTasks)       // GET /api/task
		taskGroup.POST("", taskController.CreateTask)       // POST /api/task
		taskGroup.GET("/:id", taskController.GetTaskByID)   // GET /api/task/:id
		taskGroup.PUT("/:id", taskController.UpdateTask)    // PUT /api/task/:id
		taskGroup.DELETE("/:id", taskController.DeleteTask) // DELETE /api/task/:id
	}
}

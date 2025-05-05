package routers

import (
	"project2-microservice-go/database"
	"project2-microservice-go/internal/task-service/controller"
	"project2-microservice-go/internal/task-service/repository"
	"project2-microservice-go/internal/task-service/service"
	"project2-microservice-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.RouterGroup, jwtMiddleware *middleware.JWTAuthMiddleware) {
	db := database.New()
	taskRepository := repository.NewTaskRepository(db.DB())
	taskService := service.NewTaskService(taskRepository)
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

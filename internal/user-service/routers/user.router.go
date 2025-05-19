package routers

import (
	"fmt"
	"log"
	"project2-microservice-go/database"
	"project2-microservice-go/internal/user-service/controller"
	"project2-microservice-go/internal/user-service/repository"
	"project2-microservice-go/internal/user-service/service"
	"project2-microservice-go/rabbitmq"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	rabbitmqClient, err := rabbitmq.Initialize()
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}
	db := database.New()
	userRepository := repository.NewUserRepository(db.DB())
	userService := service.NewUserService(userRepository, rabbitmqClient)
	userController := controller.NewUserController(userService)
	userGroup := router.Group("/user") // Group all /user routes
	{
		userGroup.GET("", userController.GetAllUsers)                         // GET /api/user
		userGroup.POST("", userController.CreateUser)                         // POST /api/user
		userGroup.GET("/:id", userController.GetUserByID)                     // GET /api/user/:id
		userGroup.PATCH("/:id", userController.UpdateUser)                    // PATCH /api/user/:id
		userGroup.DELETE("/:id", userController.DeleteUser)                   // DELETE /api/user/:id
		userGroup.POST("/change-password/:id", userController.ChangePassword) // POST /api/user/change-password
	}
}

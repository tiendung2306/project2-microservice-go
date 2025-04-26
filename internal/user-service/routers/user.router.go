package routers

import (
	"project2-microservice-go/database"
	"project2-microservice-go/internal/user-service/controller"
	"project2-microservice-go/internal/user-service/repository"
	"project2-microservice-go/internal/user-service/service"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	db := database.New()
	userRepository := repository.NewUserRepository(db.DB())
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userGroup := router.Group("/user") // Group all /user routes
	{
		userGroup.GET("", userController.GetAllUsers)       // GET /api/user
		userGroup.POST("", userController.CreateUser)       // POST /api/user
		userGroup.GET("/:id", userController.GetUserByID)   // GET /api/user/:id
		userGroup.PUT("/:id", userController.UpdateUser)    // PUT /api/user/:id
		userGroup.DELETE("/:id", userController.DeleteUser) // DELETE /api/user/:id
	}
}

package routers

import (
	"project2-microservice-go/database"
	"project2-microservice-go/internal/auth-service/controller"
	"project2-microservice-go/internal/auth-service/repository"
	"project2-microservice-go/internal/auth-service/service"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	db := database.New()
	authRepository := repository.NewAuthRepository(db.DB())
	userService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(userService)
	authGroup := router.Group("/auth") // Group all /user routes
	{
		authGroup.POST("/login", authController.Login)       // POST /api/auth/login
		authGroup.POST("/register", authController.Register) // POST /api/auth/register
	}
}

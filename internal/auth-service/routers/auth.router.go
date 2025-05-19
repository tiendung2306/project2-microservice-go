package routers

import (
	"project2-microservice-go/database"
	"project2-microservice-go/internal/auth-service/config"
	"project2-microservice-go/internal/auth-service/controller"
	"project2-microservice-go/internal/auth-service/repository"
	"project2-microservice-go/internal/auth-service/service"
	"project2-microservice-go/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	db := database.New()
	authRepository := repository.NewAuthRepository(db.DB())
	jwtService := service.NewJWTService(config.NewJWTConfig())
	authService := service.NewAuthService(authRepository, jwtService)
	authController := controller.NewAuthController(authService)
	jwtMiddleware := middleware.NewJWTAuthMiddleware(jwtService)
	authGroup := router.Group("/auth") // Group all /user routes
	{
		authGroup.POST("/login", authController.Login)                                                  // POST /api/auth/login
		authGroup.POST("/register", authController.Register)                                            // POST /api/auth/register
		authGroup.POST("/refresh-token", authController.RefreshToken)                                   // POST /api/auth/refresh-token
		authGroup.POST("/change-password", jwtMiddleware.AuthRequired(), authController.ChangePassword) // POST /api/auth/change-password
	}
}

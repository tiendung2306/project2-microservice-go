package routers

import (
	"project2-microservice-go/internal/notification-service/controller"

	"github.com/gin-gonic/gin"
)

// RegisterHealthRoutes registers health-related routes
func RegisterHealthRoutes(router *gin.RouterGroup) {
	healthController := controller.NewHealthController()
	healthGroup := router.Group("/health") // Group all /health routes
	{
		healthGroup.GET("", healthController.Health) // GET /api/health
	}
}

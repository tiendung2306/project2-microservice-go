package routes

import (
	"project2-microservice-go/internal/dashboard/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) error {
	serviceHandler, err := handler.NewServiceHandler()
	if err != nil {
		return err
	}

	api := r.Group("/api")
	{
		services := api.Group("/services")
		{
			services.GET("/status", serviceHandler.GetServicesStatus)
			services.POST("/:name/restart", serviceHandler.RestartService)
			services.POST("/:name/stop", serviceHandler.StopService)
			services.POST("/:name/start", serviceHandler.StartService)
		}
		api.GET("/health", serviceHandler.GetHealth)
	}

	return nil
}

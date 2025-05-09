package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// Health return API health
func (hc *HealthController) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "up",
		"service":   "notification-service",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

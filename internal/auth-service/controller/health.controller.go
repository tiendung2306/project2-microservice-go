package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHealthController interface {
	HealthCheck(c *gin.Context)
}

type healthController struct {
}

func NewHealthController() IHealthController {
	return &healthController{}
}

func (h *healthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server is healthy!"})
}

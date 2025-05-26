package handler

import (
	"net/http"

	"project2-microservice-go/internal/dashboard/service"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	serviceControl *service.ServiceControl
}

func NewServiceHandler() (*ServiceHandler, error) {
	control, err := service.NewServiceControl()
	if err != nil {
		return nil, err
	}
	return &ServiceHandler{serviceControl: control}, nil
}

func (h *ServiceHandler) GetServicesStatus(c *gin.Context) {
	services, err := h.serviceControl.GetAllServicesStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func (h *ServiceHandler) RestartService(c *gin.Context) {
	serviceName := c.Param("name")
	err := h.serviceControl.RestartService(serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service đã được khởi động lại thành công"})
}

func (h *ServiceHandler) StopService(c *gin.Context) {
	serviceName := c.Param("name")
	err := h.serviceControl.StopService(serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service đã được dừng thành công"})
}

func (h *ServiceHandler) StartService(c *gin.Context) {
	serviceName := c.Param("name")
	err := h.serviceControl.StartService(serviceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service đã được bật thành công"})
}

func (h *ServiceHandler) GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Dashboard service is running"})
}

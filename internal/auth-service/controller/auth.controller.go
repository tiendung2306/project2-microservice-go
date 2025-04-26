package controller

import (
	"project2-microservice-go/internal/auth-service/service"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	authService service.IAuthService
}

func NewAuthController(as service.IAuthService) IAuthController {
	return &authController{
		authService: as,
	}
}

func (ac *authController) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}
	if password == "" {
		c.JSON(400, gin.H{"error": "Password is required"})
		return
	}
	data, err := ac.authService.Login(email, password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": data})
}

func (ac *authController) Register(c *gin.Context) {
	// Logic to handle registration
	c.JSON(200, gin.H{"message": "Registration successful"})
}

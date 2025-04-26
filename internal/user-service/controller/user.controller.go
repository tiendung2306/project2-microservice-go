package controller

import (
	"project2-microservice-go/internal/user-service/service"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	userService service.IUserService
}

func NewUserController(us service.IUserService) IUserController {
	return &userController{
		userService: us,
	}
}

func (u *userController) GetAllUsers(c *gin.Context) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(200, users)
}

func (u *userController) CreateUser(c *gin.Context) {
	// Logic to create a user
	c.JSON(200, gin.H{"message": "Create user"})
}

func (u *userController) GetUserByID(c *gin.Context) {
	// Logic to get a user by ID
	c.JSON(200, gin.H{"message": "Get user by ID"})
}

func (u *userController) UpdateUser(c *gin.Context) {
	// Logic to update a user
	c.JSON(200, gin.H{"message": "Update user"})
}

func (u *userController) DeleteUser(c *gin.Context) {
	// Logic to delete a user
	c.JSON(200, gin.H{"message": "Delete user"})
}

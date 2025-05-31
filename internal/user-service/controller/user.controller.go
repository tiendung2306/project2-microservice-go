package controller

import (
	"fmt"
	"project2-microservice-go/internal/user-service/dto"
	"project2-microservice-go/internal/user-service/service"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	ChangePassword(c *gin.Context)
	GetMe(c *gin.Context)
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
	var createUserRequest dto.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}
	user, err := u.userService.CreateUser(&createUserRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	userResponse := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(201, userResponse)
}

func (u *userController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := u.userService.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	userResponse := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(200, userResponse)
}

func (u *userController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUserRequest dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}
	user, err := u.userService.UpdateUser(id, &updateUserRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update user " + err.Error()})
		return
	}
	userResponse := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(200, userResponse)
}

func (u *userController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := u.userService.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (u *userController) ChangePassword(c *gin.Context) {
	id := c.Param("id")
	var changePasswordRequest dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&changePasswordRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}
	fmt.Printf("id: %s\n", id)
	fmt.Printf("changePasswordRequest: %+v\n", changePasswordRequest)
}

func (u *userController) GetMe(c *gin.Context) {
	fmt.Printf("GetMe\n")
	user, err := u.userService.GetMe(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get me: " + err.Error()})
		return
	}
	c.JSON(200, user)
}

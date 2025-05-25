package controller

import (
	"project2-microservice-go/internal/task-service/dto"
	"project2-microservice-go/internal/task-service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ITaskController interface {
	CreateTask(c *gin.Context)
	GetAllTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTasksByUserID(c *gin.Context)
}

type taskController struct {
	taskService service.ITaskService
}

func NewTaskController(taskService service.ITaskService) ITaskController {
	return &taskController{
		taskService: taskService,
	}
}

func (t *taskController) CreateTask(c *gin.Context) {
	var CreateTaskRequest dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&CreateTaskRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	task, err := t.taskService.CreateTask(&CreateTaskRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task, " + err.Error()})
		return
	}
	c.JSON(201, task)
}
func (t *taskController) GetAllTasks(c *gin.Context) {
	status := c.Query("status")
	tasks, err := t.taskService.GetAllTasks(status)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(200, tasks)
}
func (t *taskController) GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := t.taskService.GetTaskByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}
func (t *taskController) UpdateTask(c *gin.Context) {
	var UpdateTaskRequest dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&UpdateTaskRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := t.taskService.UpdateTask(uint(id), &UpdateTaskRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update task, " + err.Error()})
		return
	}
	c.JSON(200, task)
}
func (t *taskController) DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	err = t.taskService.DeleteTask(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete task, " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}

func (t *taskController) GetTasksByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	status := c.Query("status")
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	tasks, err := t.taskService.GetTasksByUserID(userID, status)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(200, tasks)
}

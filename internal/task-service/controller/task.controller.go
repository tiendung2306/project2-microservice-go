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
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(201, task)
}
func (t *taskController) GetAllTasks(c *gin.Context) {
	tasks, err := t.taskService.GetAllTasks()
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
	// Implementation for updating a task
}
func (t *taskController) DeleteTask(c *gin.Context) {
	// Implementation for deleting a task
}

package repository

import (
	"project2-microservice-go/internal/task-service/dto"
	"project2-microservice-go/models"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	CreateTask(request *dto.CreateTaskRequest) (*dto.TaskResponse, error)
	GetAllTasks(status string) ([]models.Task, error)
	GetTaskByID(id int) (*dto.TaskResponse, error)
	UpdateTask(id uint, request *dto.UpdateTaskRequest) (models.Task, error)
	DeleteTask(id int) error
	GetTasksByUserID(userID int, status string) ([]models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) CreateTask(request *dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	task := models.Task{
		UserID:    request.UserID,
		Title:     request.Title,
		Content:   request.Content,
		StartDate: request.StartDate,
		DueDate:   request.DueDate,
		Status:    request.Status,
	}

	result := tr.db.Create(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.TaskResponse{
		ID:        task.ID,
		UserID:    task.UserID,
		Title:     task.Title,
		Content:   task.Content,
		StartDate: task.StartDate,
		DueDate:   task.DueDate,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tr *taskRepository) GetAllTasks(status string) ([]models.Task, error) {
	var tasks []models.Task
	var result *gorm.DB
	if status == "" {
		result = tr.db.Find(&tasks)
	} else {
		result = tr.db.Where("status = ?", status).Find(&tasks)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (tr *taskRepository) GetTaskByID(id int) (*dto.TaskResponse, error) {
	var task models.Task
	result := tr.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.TaskResponse{
		ID:        task.ID,
		UserID:    task.UserID,
		Title:     task.Title,
		Content:   task.Content,
		StartDate: task.StartDate,
		DueDate:   task.DueDate,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tr *taskRepository) UpdateTask(id uint, request *dto.UpdateTaskRequest) (models.Task, error) {
	// First get the existing task to preserve UserID
	var existingTask models.Task
	if err := tr.db.First(&existingTask, id).Error; err != nil {
		return models.Task{}, err
	}

	task := models.Task{
		Title:     request.Title,
		Content:   request.Content,
		StartDate: request.StartDate,
		DueDate:   request.DueDate,
		Status:    request.Status,
	}
	result := tr.db.Model(&models.Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return models.Task{
		ID:        id,
		UserID:    existingTask.UserID,
		Title:     task.Title,
		Content:   task.Content,
		StartDate: task.StartDate,
		DueDate:   task.DueDate,
		Status:    task.Status,
		CreatedAt: existingTask.CreatedAt,
		UpdatedAt: existingTask.UpdatedAt,
	}, nil
}

func (tr *taskRepository) DeleteTask(id int) error {
	result := tr.db.Delete(&models.Task{}, id)
	return result.Error
}

func (tr *taskRepository) GetTasksByUserID(userID int, status string) ([]models.Task, error) {
	var tasks []models.Task
	var result *gorm.DB
	if status == "" {
		result = tr.db.Where("user_id = ?", userID).Find(&tasks)
	} else {
		result = tr.db.Where("user_id = ? AND status = ?", userID, status).Find(&tasks)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

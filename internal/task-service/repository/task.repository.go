package repository

import (
	"project2-microservice-go/internal/task-service/dto"
	"project2-microservice-go/models"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	CreateTask(request *dto.CreateTaskRequest) (*dto.TaskResponse, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id int) (*dto.TaskResponse, error)
	UpdateTask(id int, task models.Task) (models.Task, error)
	DeleteTask(id int) error
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

func (tr *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := tr.db.Find(&tasks)
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

func (tr *taskRepository) UpdateTask(id int, task models.Task) (models.Task, error) {
	result := tr.db.Model(&models.Task{}).Where("id = ?", id).Updates(task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func (tr *taskRepository) DeleteTask(id int) error {
	result := tr.db.Delete(&models.Task{}, id)
	return result.Error
}

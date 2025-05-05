package service

import (
	"project2-microservice-go/internal/task-service/dto"
	"project2-microservice-go/internal/task-service/repository"
	"project2-microservice-go/models"
)

type ITaskService interface {
	CreateTask(request *dto.CreateTaskRequest) (*dto.TaskResponse, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id int) (*dto.TaskResponse, error)
	UpdateTask(id int, task models.Task) (models.Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	taskRepository repository.ITaskRepository
}

func NewTaskService(taskRepository repository.ITaskRepository) ITaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (t *taskService) CreateTask(request *dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	// Implementation for creating a task
	createdTask, err := t.taskRepository.CreateTask(request)
	if err != nil {
		return nil, err
	}
	return createdTask, nil
}
func (t *taskService) GetAllTasks() ([]models.Task, error) {
	// Implementation for getting all tasks
	tasks, err := t.taskRepository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (t *taskService) GetTaskByID(id int) (*dto.TaskResponse, error) {
	// Implementation for getting a task by ID
	task, err := t.taskRepository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (t *taskService) UpdateTask(id int, task models.Task) (models.Task, error) {
	// Implementation for updating a task
	updatedTask, err := t.taskRepository.UpdateTask(id, task)
	if err != nil {
		return models.Task{}, err
	}
	return updatedTask, nil
}
func (t *taskService) DeleteTask(id int) error {
	// Implementation for deleting a task
	err := t.taskRepository.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

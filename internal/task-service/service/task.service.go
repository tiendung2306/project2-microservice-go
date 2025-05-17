package service

import (
	"log"
	"project2-microservice-go/errors"
	"project2-microservice-go/internal/task-service/dto"
	"project2-microservice-go/internal/task-service/repository"
	"project2-microservice-go/models"
	"project2-microservice-go/rabbitmq"
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
	userRepository repository.IUserRepository
	rabbitmq       *rabbitmq.RabbitMQ
}

func NewTaskService(taskRepository repository.ITaskRepository, userRepository repository.IUserRepository, rabbitmq *rabbitmq.RabbitMQ) ITaskService {
	return &taskService{
		taskRepository: taskRepository,
		userRepository: userRepository,
		rabbitmq:       rabbitmq,
	}
}

func (t *taskService) CreateTask(request *dto.CreateTaskRequest) (*dto.TaskResponse, error) {
	userExists, err := t.userRepository.IsUserExists(request.UserID)
	if !userExists {
		return nil, errors.NewNotFoundError("user not found", err)
	}
	// Implementation for creating a task
	createdTask, err := t.taskRepository.CreateTask(request)
	if err != nil {
		return nil, err
	}

	// Get user email from user repository
	userEmail, err := t.userRepository.GetUserEmailByID(request.UserID)
	if err != nil {
		log.Printf("Failed to get user email for user ID %d: %v", request.UserID, err)
		return createdTask, nil
	}

	// Send notification
	message := rabbitmq.NotificationMessage{
		Type:    "task",
		Action:  "create",
		UserID:  request.UserID,
		Email:   userEmail,
		Subject: "New Task Created",
		Content: "A new task has been created: " + request.Title,
	}
	err = t.rabbitmq.PublishMessage(message)
	if err != nil {
		log.Printf("Failed to publish message to RabbitMQ: %v", err)
	} else {
		log.Printf("Successfully published message to RabbitMQ: %+v", message)
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

	// Get user email from user repository
	userEmail, err := t.userRepository.GetUserEmailByID(task.UserID)
	if err != nil {
		// Log error but don't fail the request
		// TODO: Add proper logging
		return updatedTask, nil
	}

	// Send notification
	message := rabbitmq.NotificationMessage{
		Type:    "task",
		Action:  "update",
		UserID:  task.UserID,
		Email:   userEmail,
		Subject: "Task Updated",
		Content: "Task has been updated: " + task.Title,
	}
	err = t.rabbitmq.PublishMessage(message)
	if err != nil {
		// Log error but don't fail the request
		// TODO: Add proper logging
	}

	return updatedTask, nil
}

func (t *taskService) DeleteTask(id int) error {
	// Get task info before deleting for notification
	task, err := t.taskRepository.GetTaskByID(id)
	if err != nil {
		return err
	}

	// Implementation for deleting a task
	err = t.taskRepository.DeleteTask(id)
	if err != nil {
		return err
	}

	// Get user email from user repository
	userEmail, err := t.userRepository.GetUserEmailByID(task.UserID)
	if err != nil {
		// Log error but don't fail the request
		// TODO: Add proper logging
		return nil
	}

	// Send notification
	message := rabbitmq.NotificationMessage{
		Type:    "task",
		Action:  "delete",
		UserID:  task.UserID,
		Email:   userEmail,
		Subject: "Task Deleted",
		Content: "Task has been deleted: " + task.Title,
	}
	err = t.rabbitmq.PublishMessage(message)
	if err != nil {
		// Log error but don't fail the request
		// TODO: Add proper logging
	}

	return nil
}

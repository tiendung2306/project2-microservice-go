package dto

import "time"

type CreateTaskRequest struct {
	UserID    uint      `json:"user_id" binding:"required"`
	Title     string    `json:"title" binding:"required,max=255"`
	Content   string    `json:"content" binding:"max=255"`
	StartDate time.Time `json:"start_date" binding:"required"`
	DueDate   time.Time `json:"due_date" binding:"required"`
	Status    string    `json:"status" binding:"required,oneof='To Do' 'In Progress' 'Done'"`
}

type TaskResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTaskRequest struct {
	Title     string    `json:"title" binding:"required,max=255"`
	Content   string    `json:"content" binding:"max=255"`
	StartDate time.Time `json:"start_date" binding:"required"`
	DueDate   time.Time `json:"due_date" binding:"required"`
	Status    string    `json:"status" binding:"required,oneof='To Do' 'In Progress' 'Done'"`
}

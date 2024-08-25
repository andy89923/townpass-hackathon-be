package domain

import (
	"time"
)

type LostItem struct {
    ID        uint `gorm:"primaryKey"`
    LostTime time.Time
    Kind      string
    PropertyName string
    Location   string
    PhoneNumber string
}

type LostItemRepository interface {
    // Create(todo *Todo) (*Todo, error)
	// GetAll() ([]*Todo, error)
	// GetByID(id uint) (*Todo, error)
	// Update(todo *Todo) (*Todo, error)
	// Delete(id uint) error
}

type LostItemService interface {
    // AddNewTodo(*Todo) (*Todo, error)
    // GetAllTodos() ([]*Todo, error)
    // GetTodoById(id uint) (*Todo, error)
    // UpdateTodo(todo *Todo) (*Todo, error)
    // DeleteTodo(id uint) error
}


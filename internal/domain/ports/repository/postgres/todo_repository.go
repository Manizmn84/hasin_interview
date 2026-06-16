package postgres

import "github.com/Manizmn84/hasin_interview/internal/domain/entities"

type TodoRepository interface {
	CreateTodo(todo *entities.Todo) error
	DeleteTodo(todo *entities.Todo) error
	GetTodoByID(id uint) (*entities.Todo, error)
	UpdateTodo(todo *entities.Todo) error
	GetAllTodo() ([]*entities.Todo, error)
}

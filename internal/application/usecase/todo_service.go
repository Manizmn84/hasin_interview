package usecase

import tododto "github.com/Manizmn84/hasin_interview/internal/application/dto/todo"

type TodoService interface {
	CreateTodo(todoReq tododto.TodoCreateRequest) (*tododto.TodoCreateResponse, error)
	DeleteTodo(id uint) error
	GetTodoByID(id uint) (*tododto.GetTodoByIDResponse, error)
	GetAllTodosSortedByNp() ([]tododto.GetTodoByIDResponse, error)
	UpdateTodo(todoUpdate tododto.TodoUpdateRequest) error
}

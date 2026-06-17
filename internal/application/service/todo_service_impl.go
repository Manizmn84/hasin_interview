package service

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	tododto "github.com/Manizmn84/hasin_interview/internal/application/dto/todo"
	"github.com/Manizmn84/hasin_interview/internal/domain/entities"
	"github.com/Manizmn84/hasin_interview/internal/domain/enums"
	"github.com/Manizmn84/hasin_interview/internal/domain/logging"
	"github.com/Manizmn84/hasin_interview/internal/domain/ports"
)

type TodoService struct {
	logger     logging.Logger
	cfg        *bootstrap.Config
	unitOfWork ports.UnitOfWork
}

func NewTodoService(
	cfg *bootstrap.Config,
	unitOfWork ports.UnitOfWork,
) *TodoService {
	logger := logging.NewLogger(cfg)

	return &TodoService{
		logger:     logger,
		cfg:        cfg,
		unitOfWork: unitOfWork,
	}
}

func (ts *TodoService) CreateTodo(todoReq tododto.TodoCreateRequest) (*tododto.TodoCreateResponse, error) {
	todo := &entities.Todo{
		Np:     todoReq.Np,
		Title:  todoReq.Title,
		Dsc:    todoReq.Dsc,
		Status: enums.NotFinish,
	}

	err := ts.unitOfWork.Factory().TodoRepository().CreateTodo(todo)

	if err != nil {
		return &tododto.TodoCreateResponse{}, err
	}

	return &tododto.TodoCreateResponse{
		Title: todo.Title,
	}, nil
}

func (ts *TodoService) DeleteTodo(id uint) error {
	todo, err := ts.unitOfWork.Factory().TodoRepository().GetTodoByID(id)

	if err != nil {
		return err
	}

	err = ts.unitOfWork.Factory().TodoRepository().DeleteTodo(todo)

	if err != nil {
		return err
	}
	return nil
}

func (ts *TodoService) GetTodoByID(id uint) (*tododto.GetTodoByIDResponse, error) {
	todo, err := ts.unitOfWork.Factory().TodoRepository().GetTodoByID(id)

	if err != nil {
		return &tododto.GetTodoByIDResponse{}, err
	}

	res := &tododto.GetTodoByIDResponse{
		ID:     todo.ID,
		Np:     todo.Np,
		Status: todo.Status.String(),
		Title:  todo.Title,
		Dsc:    todo.Dsc,
	}

	return res, nil
}

func (ts *TodoService) GetAllTodosSortedByNp() ([]tododto.GetTodoByIDResponse, error) {
	todos, err := ts.unitOfWork.Factory().TodoRepository().GetAllTodo()

	if err != nil {
		return nil, err
	}

	resTodos := make([]tododto.GetTodoByIDResponse, 0, len(todos))

	for _, todo := range todos {
		resDto := tododto.GetTodoByIDResponse{
			ID:     todo.ID,
			Np:     todo.Np,
			Title:  todo.Title,
			Dsc:    todo.Dsc,
			Status: todo.Status.String(),
		}
		resTodos = append(resTodos, resDto)
	}

	return resTodos, nil
}

func (ts *TodoService) UpdateTodo(todoUpdate tododto.TodoUpdateRequest) error {
	todo, err := ts.unitOfWork.Factory().TodoRepository().GetTodoByID(todoUpdate.ID)

	if err != nil {
		return err
	}
	todo.Title = todoUpdate.Title
	todo.Dsc = todoUpdate.Dsc
	todo.Np = todoUpdate.Np
	todo.Status = enums.TodoStatus(todoUpdate.Status)

	err = ts.unitOfWork.Factory().TodoRepository().UpdateTodo(todo)

	if err != nil {
		return err
	}

	return nil
}

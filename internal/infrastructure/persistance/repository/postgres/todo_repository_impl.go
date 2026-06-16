package postgres

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/domain/entities"
	"gorm.io/gorm"
)

type TodoRepository struct {
	cfg *bootstrap.Config
	db  *gorm.DB
}

func NewTodoRepository(cfg *bootstrap.Config, db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		cfg: cfg,
		db:  db,
	}
}

func (tr *TodoRepository) CreateTodo(todo *entities.Todo) error {
	return tr.db.Create(&todo).Error
}

func (tr *TodoRepository) DeleteTodo(todo *entities.Todo) error {
	return tr.db.Delete(&todo).Error
}

func (tr *TodoRepository) GetTodoByID(id uint) (*entities.Todo, error) {
	var todo entities.Todo

	err := tr.db.Where("id = ?", id).First(&todo).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (tr *TodoRepository) UpdateTodo(todo *entities.Todo) error {
	return tr.db.Save(&todo).Error
}

func (tr *TodoRepository) GetAllTodo() ([]*entities.Todo, error) {
	var todos []*entities.Todo

	err := tr.db.Order("Np DESC").Find(&todos).Error

	if err != nil {
		return nil, err
	}

	return todos, nil

}

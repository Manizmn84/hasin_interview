package entities

import (
	"github.com/Manizmn84/hasin_interview/internal/domain/enums"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Np  float64 
	Title string
	Dsc string
	Status enums.TodoStatus
}
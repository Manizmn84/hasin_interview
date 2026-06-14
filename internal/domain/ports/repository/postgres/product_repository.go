package postgres

import "github.com/Manizmn84/hasin_interview/internal/domain/entities"

type ProductRepository interface {
	CreateProduct(product *entities.Product) error
}

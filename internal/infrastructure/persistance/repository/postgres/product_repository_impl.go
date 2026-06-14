package postgres

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	cfg *bootstrap.Config
	db  *gorm.DB
}

func NewProductRepository(cfg *bootstrap.Config, db *gorm.DB) *ProductRepository {

	return &ProductRepository{
		cfg: cfg,
		db:  db,
	}
}

func (pr *ProductRepository) CreateProduct(product *entities.Product) error {
	return pr.db.Create(&product).Error
}

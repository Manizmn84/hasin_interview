package persistance

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	postgresDomain "github.com/Manizmn84/hasin_interview/internal/domain/ports/repository/postgres"
	"github.com/Manizmn84/hasin_interview/internal/infrastructure/persistance/repository/postgres"
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	cfg *bootstrap.Config
	db  *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB, cfg *bootstrap.Config) *RepositoryFactory {
	return &RepositoryFactory{
		db:  db,
		cfg: cfg,
	}
}

func (rf *RepositoryFactory) ProductRepository() postgresDomain.ProductRepository {
	return postgres.NewProductRepository(rf.cfg, rf.db)
}

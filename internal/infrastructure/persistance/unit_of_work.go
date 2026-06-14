package persistance

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/domain/ports"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db  *gorm.DB
	cfg *bootstrap.Config
}

func NewUnitOfWork(db *gorm.DB, cfg *bootstrap.Config) *UnitOfWork {
	return &UnitOfWork{
		db:  db,
		cfg: cfg,
	}
}

func (u *UnitOfWork) Factory() ports.RepositoryFactory {
	return NewRepositoryFactory(u.db, u.cfg)
}

func (u *UnitOfWork) WithTransaction(fn func(ports.RepositoryFactory) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		factory := NewRepositoryFactory(tx, u.cfg)
		return fn(factory)
	})
}

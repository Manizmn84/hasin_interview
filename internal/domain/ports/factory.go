package ports

import postgresDomain "github.com/Manizmn84/hasin_interview/internal/domain/ports/repository/postgres"

type RepositoryFactory interface {
	ProductRepository() postgresDomain.ProductRepository
	TodoRepository() postgresDomain.TodoRepository
}

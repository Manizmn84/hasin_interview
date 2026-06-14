package service

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	productdto "github.com/Manizmn84/hasin_interview/internal/application/dto/product"
	"github.com/Manizmn84/hasin_interview/internal/domain/entities"
	"github.com/Manizmn84/hasin_interview/internal/domain/logging"
	"github.com/Manizmn84/hasin_interview/internal/domain/ports"
)

type ProductService struct {
	logger     logging.Logger
	cfg        *bootstrap.Config
	unitOfWork ports.UnitOfWork
}

func NewProductService(
	cfg *bootstrap.Config,
	unitOfWork ports.UnitOfWork,
) *ProductService {
	logger := logging.NewLogger(cfg)

	return &ProductService{
		logger:     logger,
		cfg:        cfg,
		unitOfWork: unitOfWork,
	}
}

func (ps *ProductService) Create(req productdto.ProductCreateRequest) (*productdto.ProductCreateResponse, error) {
	product := entities.Product{Name: req.Name}

	err := ps.unitOfWork.Factory().ProductRepository().CreateProduct(&product)

	if err != nil {
		return &productdto.ProductCreateResponse{}, err
	}

	return &productdto.ProductCreateResponse{Message: "the Product Create successfully"}, nil
}

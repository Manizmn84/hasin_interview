package usecase

import productdto "github.com/Manizmn84/hasin_interview/internal/application/dto/product"

type ProductService interface {
	Create(req productdto.ProductCreateRequest) (*productdto.ProductCreateResponse, error)
}

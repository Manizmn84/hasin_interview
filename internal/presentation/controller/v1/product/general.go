package product

import (
	"net/http"

	"github.com/Manizmn84/hasin_interview/bootstrap"
	productdto "github.com/Manizmn84/hasin_interview/internal/application/dto/product"
	"github.com/Manizmn84/hasin_interview/internal/application/usecase"
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/base"
	"github.com/gin-gonic/gin"
)

type ProductGeneralControler struct {
	productService usecase.ProductService
	cfg            *bootstrap.Config
}

func NewProductGeneralControler(cfg *bootstrap.Config, productService usecase.ProductService) *ProductGeneralControler {
	return &ProductGeneralControler{
		cfg:            cfg,
		productService: productService,
	}
}

func (pgc *ProductGeneralControler) SayHello(ctx *gin.Context) {

	base.Response(ctx, http.StatusCreated, "successMessage.createParking", nil)
}

func (pgc *ProductGeneralControler) Create(ctx *gin.Context) {
	type CreateParams struct {
		Name string `json:"name" validate:"required"`
	}

	params := base.Validated[CreateParams](ctx)

	req := productdto.ProductCreateRequest{
		Name: params.Name,
	}
	res, err := pgc.productService.Create(req)

	if err != nil {
		panic(err)
	}

	base.Response(ctx, http.StatusAccepted, res.Message, nil)
}

package http

import (
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/product"
	"github.com/gin-gonic/gin"
)

func SetupGeneralRoutes(router *gin.RouterGroup, productGeneralController *product.ProductGeneralControler) {
	productGroup := router.Group("/product")
	{
		productGroup.GET("", productGeneralController.SayHello)
		productGroup.POST("/Create", productGeneralController.Create)
	}
}

package http

import (
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/product"
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/todo"
	"github.com/gin-gonic/gin"
)

func SetupGeneralRoutes(router *gin.RouterGroup, productGeneralController *product.ProductGeneralControler, todoGeneralController *todo.TodoGeneralController) {
	productGroup := router.Group("/product")
	{
		productGroup.GET("", productGeneralController.SayHello)
		productGroup.POST("/Create", productGeneralController.Create)
	}

	todoGroup := router.Group("/todo")
	{
		todoGroup.POST("/create", todoGeneralController.Create)
		todoGroup.GET("", todoGeneralController.List)
		todoGroup.GET("/:id", todoGeneralController.GetByID)
		todoGroup.PUT("update/:id", todoGeneralController.Update)
		todoGroup.DELETE("delete/:id", todoGeneralController.Delete)
	}
}

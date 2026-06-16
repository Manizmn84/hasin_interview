package routes

import (
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/product"
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/todo"
	"github.com/Manizmn84/hasin_interview/internal/presentation/routes/http/v1"

	"github.com/gin-gonic/gin"
)

func Run(ginEngine *gin.Engine, productGeneralController *product.ProductGeneralControler, todoGeneralController *todo.TodoGeneralController) {
	v1 := ginEngine.Group("/v1")
	{
		http.SetupGeneralRoutes(v1, productGeneralController, todoGeneralController)
		http.SetupCustomerRoutes(v1)
	}
}

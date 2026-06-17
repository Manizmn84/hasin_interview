package main

import (
	"github.com/Manizmn84/hasin_interview/bootstrap"
	"github.com/Manizmn84/hasin_interview/internal/application/service"
	"github.com/Manizmn84/hasin_interview/internal/infrastructure/database"
	"github.com/Manizmn84/hasin_interview/internal/infrastructure/localization"
	"github.com/Manizmn84/hasin_interview/internal/infrastructure/persistance"
	"github.com/Manizmn84/hasin_interview/internal/infrastructure/validation"
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/product"
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/v1/todo"
	"github.com/Manizmn84/hasin_interview/internal/presentation/middleware"
	"github.com/Manizmn84/hasin_interview/internal/presentation/routes"
	"github.com/gin-contrib/cors"

	// "github.com/Manizmn84/hasin_interview/wire"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	cfg := bootstrap.Run()

	db := database.NewPostgresDataBase(&cfg.Env.PrimaryDB)

	if db == nil {
		panic("the db is nil")
	}

	// Repositories
	unitOfWork := persistance.NewUnitOfWork(db, cfg)

	// Services
	productService := service.NewProductService(cfg, unitOfWork)
	todoService := service.NewTodoService(cfg, unitOfWork)

	// Controllers
	productGeneralController := product.NewProductGeneralControler(cfg, productService)
	todoGeneralController := todo.NewTodoGeneralController(cfg, todoService)

	// mid
	recoveryMid := middleware.NewRecoveryMiddleware(cfg.Constant)
	local := localization.NewTranslationService()
	transMid := middleware.TranslatorMiddleware(local)

	// Router
	gin.DisableConsoleColor()
	ginEngine := gin.Default()
	ginEngine.Use(recoveryMid.Recovery)
	ginEngine.Use(transMid)
	ginEngine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("ir_phone", validation.IranianMobileValidator, true)
		val.RegisterValidation("ir_postal", validation.IranianPostalCodeValidator, true)
	}

	routes.Run(ginEngine, productGeneralController, todoGeneralController)

	ginEngine.Run()
}

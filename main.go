package main

import (
	"github/ecommerce/adapter/handlers/auth"
	"github/ecommerce/adapter/handlers/category_handlers"
	"github/ecommerce/adapter/handlers/health_handler"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/cmd"
	"github/ecommerce/domain/health"
	"github/ecommerce/infra/memory"
	"github/ecommerce/infra/postgresql"
	"github/ecommerce/usecase"
)

func main() {

	db := postgresql.GetDB()
	// category dependencies
	catRepo := postgresql.NewCategoryRepo(db)
	catService := usecase.NewCategoryService(catRepo)
	catHandler := category_handlers.NewProductHandler(catService)

	//product dependencies
	repo := postgresql.NewProductRepo(db)
	productservice := usecase.NewProductService(repo)
	productHandler := product_handlers.NewProductHandler(productservice)

	//auth dependencies
	authRepo := postgresql.NewUserRepo(db)
	userService := usecase.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(userService)

	//health dependencies-in-memory
	_health := memory.HealthRepo{
		Health: health.Health{
			Status:  200,
			Message: "OK",
		},
	}
	healthservice := usecase.NewHealthService(&_health)
	healthHandler := health_handler.NewHealthHandler(healthservice)

	cmd.CreateServer(catHandler, productHandler, authHandler, healthHandler)

}

package main

import (
	"github/ecommerce/adapter/handlers/health_handler"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/cmd"
	"github/ecommerce/domain/health"
	"github/ecommerce/infra/memory"
	"github/ecommerce/infra/postgresql"
	"github/ecommerce/usecase"
)

func main() {

	// //product dependencies
	db := postgresql.GetDB()
	repo := postgresql.NewProductRepo(db)
	productservice := usecase.NewProductService(repo)
	productHandler := product_handlers.NewProductHandler(productservice)

	// //health dependencies-in-memory
	_health := memory.HealthRepo{
		Health: health.Health{
			Status:  200,
			Message: "OK",
		},
	}
	healthservice := usecase.NewHealthService(&_health)
	healthHandler := health_handler.NewHealthHandler(healthservice)

	cmd.CreateServer(productHandler, healthHandler)

}

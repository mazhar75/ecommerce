package main

import (
	"github/ecommerce/adapter/handlers/health_handler"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/cmd"
	"github/ecommerce/domain/health"
	"github/ecommerce/domain/product"
	"github/ecommerce/infra/memory"
	"github/ecommerce/usecase"
)

func main() {
	products := memory.ProductRepo{
		ProductList: []product.Product{
			{ProductId: 1, Name: "Mango", Description: "Good quality", Type: "Fruits", Price: 70, ImgUrl: ""},
			{ProductId: 2, Name: "Mouse", Description: "Good quality", Type: "Electronics", Price: 100, ImgUrl: ""},
			{ProductId: 3, Name: "Sharee", Description: "Good quality", Type: "Shopping", Price: 999, ImgUrl: ""},
		},
	}

	_health := memory.HealthRepo{
		Health: health.Health{
			Status:  200,
			Message: "OK",
		},
	}

	productservice := usecase.NewProductService(&products)
	productHandler := product_handlers.NewProductHandler(productservice)

	healthservice := usecase.NewHealthService(&_health)
	healthHandler := health_handler.NewHealthHandler(healthservice)

	cmd.CreateServer(productHandler, healthHandler)
}

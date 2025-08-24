package main

import (
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/cmd"
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

	service := usecase.NewProductService(&products)
	productHandler := product_handlers.NewProductHandler(service)

	cmd.CreateServer(productHandler)
}

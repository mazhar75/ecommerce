package main

import (
	// alias
	"fmt"
	"github/ecommerce/internal/adpter/http/handler"
	"github/ecommerce/internal/adpter/http/middleware"
	"github/ecommerce/internal/domain/product"
	httpadapter "github/ecommerce/internal/framework/http"
	"github/ecommerce/internal/framework/memory"
	"github/ecommerce/internal/usecase"
	"log"
	"net/http"
)

var (
	productHandler *handler.ProductHandler
)

func init() {
	products := memory.ProductRepo{
		ProductList: []product.Product{
			{ProductId: 1, Name: "Mango", Description: "Good quality", Type: "Fruits", Price: 70, ImgUrl: ""},
			{ProductId: 2, Name: "Mouse", Description: "Good quality", Type: "Electronics", Price: 100, ImgUrl: ""},
			{ProductId: 3, Name: "Sharee", Description: "Good quality", Type: "Shopping", Price: 999, ImgUrl: ""},
		},
	}
	service := usecase.NewProductService(&products)
	productHandler = handler.NewProductHandler(service)
}

func main() {
	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	httpadapter.RegisterRoutes(mux, productHandler)
	corsHandler := middleware.CorsMiddleware(mux)
	err := http.ListenAndServe(":9090", corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Handler -> Usecase (Service) -> Port (Interface) -> Repo (Concrete Implementation)
*/

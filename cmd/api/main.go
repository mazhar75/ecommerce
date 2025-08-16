package main

import (
	// alias
	"fmt"
	httpadapter "github/ecommerce/internal/adpter/http"
	"github/ecommerce/internal/adpter/http/handler"
	"github/ecommerce/internal/domain"
	"github/ecommerce/internal/repo/memory"
	"github/ecommerce/internal/usecase"
	"log"
	"net/http"
)

var (
	productHandler *handler.ProductHandler
)

func init() {
	products := memory.ProductRepo{
		ProductList: []domain.Product{
			{Id: 1, Name: "Mango", Type: "Fruits", Price: 70, ImgUrl: ""},
			{Id: 2, Name: "Mouse", Type: "Electronics", Price: 100, ImgUrl: ""},
			{Id: 3, Name: "Sharee", Type: "Shopping", Price: 999, ImgUrl: ""},
		},
	}
	service := usecase.NewProductService(&products)
	productHandler = handler.NewProductHandler(service)
}

func main() {
	fmt.Println("Server starting...")
	mux := http.NewServeMux()
	httpadapter.RegisterRoutes(mux, productHandler)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Handler -> Usecase (Service) -> Port (Interface) -> Repo (Concrete Implementation)
*/

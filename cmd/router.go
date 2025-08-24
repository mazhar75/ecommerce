package cmd

import (
	"github/ecommerce/adapter/handlers/product_handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *product_handlers.ProductHandler) {
	mux.Handle("GET /products", http.HandlerFunc(productHandler.GetProducts))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(productHandler.GetProductById))
	mux.Handle("POST /products", http.HandlerFunc(productHandler.CreateProduct))
}

package cmd

import (
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/middlewares"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *product_handlers.ProductHandler) {
	h := http.HandlerFunc(productHandler.GetProducts)
	h1 := middlewares.MiddlwareTest1(h)
	h2 := middlewares.MiddlewareTest2(h1)
	h3 := middlewares.Logger(h2)

	mux.Handle("GET /products", h3)
	mux.Handle("GET /products/{productId}", http.HandlerFunc(productHandler.GetProductById))
	mux.Handle("POST /products", http.HandlerFunc(productHandler.CreateProduct))
}

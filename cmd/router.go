package cmd

import (
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/middlewares"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *product_handlers.ProductHandler) {

	manager := middlewares.NewManager()
	manager.Use(middlewares.MiddlwareTest1, middlewares.MiddlewareTest2, middlewares.Logger, middlewares.CorsMiddleware)
	mux.Handle("GET /products", manager.With(http.HandlerFunc(productHandler.GetProducts)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(productHandler.GetProductById)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(productHandler.CreateProduct)))
}

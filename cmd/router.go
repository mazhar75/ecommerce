package cmd

import (
	"github/ecommerce/adapter/handlers/health_handler"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/middlewares"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *product_handlers.ProductHandler, healthHandler *health_handler.HealthHandler, manager *middlewares.Manager) {

	mux.Handle("GET /products", manager.With(http.HandlerFunc(productHandler.GetProducts)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(productHandler.GetProductById)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(productHandler.CreateProduct)))
	mux.Handle("GET /health", manager.With(http.HandlerFunc(healthHandler.GetHealth)))
}

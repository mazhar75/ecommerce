package cmd

import (
	"fmt"
	"github/ecommerce/adapter/handlers/health_handler"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/middlewares"
	"log"
	"net/http"
)

// cmd/server.go
func CreateServer(productHandler *product_handlers.ProductHandler, healthHandler *health_handler.HealthHandler) {
	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	manager.Use(middlewares.MiddlwareTest1, middlewares.MiddlewareTest2, middlewares.Logger, middlewares.CorsMiddleware)
	RegisterRoutes(mux, productHandler, healthHandler, manager)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}

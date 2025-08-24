package cmd

import (
	"fmt"
	"github/ecommerce/adapter/handlers/product_handlers"
	"github/ecommerce/middlewares"
	"log"
	"net/http"
)

// cmd/server.go
func CreateServer(productHandler *product_handlers.ProductHandler) {
	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	RegisterRoutes(mux, productHandler)
	corsHandler := middlewares.CorsMiddleware(mux)
	err := http.ListenAndServe(":9090", corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}

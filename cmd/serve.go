package cmd

import (
	"fmt"
	"github/ecommerce/adapter/handlers/product_handlers"
	"log"
	"net/http"
)

// cmd/server.go
func CreateServer(productHandler *product_handlers.ProductHandler) {
	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	RegisterRoutes(mux, productHandler)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}

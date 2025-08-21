package cmd

import (
	"fmt"
	"github/ecommerce/internal/adpter/http/handler/product_handler"
	"github/ecommerce/internal/adpter/http/middleware"
	httpadapter "github/ecommerce/internal/framework/http"
	"log"
	"net/http"
)

// cmd/server.go
func CreateServer(productHandler *product_handler.ProductHandler) {
	fmt.Println("Server starting...at 9090")
	mux := http.NewServeMux()
	httpadapter.RegisterRoutes(mux, productHandler)
	corsHandler := middleware.CorsMiddleware(mux)
	err := http.ListenAndServe(":9090", corsHandler)
	if err != nil {
		log.Fatal(err)
	}
}

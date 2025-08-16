package http

import (
	"github/ecommerce/internal/adpter/http/handler"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *handler.ProductHandler) {
	// GET /products route
	mux.HandleFunc("/products", productHandler.GetProducts)
	mux.HandleFunc("/product", productHandler.GetProductById)
	mux.HandleFunc("/create-product", productHandler.CreateProduct)
}

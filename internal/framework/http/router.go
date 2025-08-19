package http

import (
	"github/ecommerce/internal/adpter/http/handler"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *handler.ProductHandler) {
	mux.Handle("GET /products", http.HandlerFunc(productHandler.GetProducts))
	mux.Handle("GET /product", http.HandlerFunc(productHandler.GetProductById))
	mux.Handle("POST /create-product", http.HandlerFunc(productHandler.CreateProduct))
}

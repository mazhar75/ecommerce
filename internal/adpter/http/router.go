package http

import (
    "net/http"
    "github/ecommerce/internal/adapter/http/handler"
)

func RegisterRoutes(mux *http.ServeMux, productHandler *handler.ProductHandler) {
    // GET /products route
    mux.HandleFunc("/products", productHandler.GetProducts)
}

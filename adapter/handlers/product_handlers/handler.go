package product_handlers

import (
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type ProductHandler struct {
	Service *usecase.ProductService
}

func NewProductHandler(service *usecase.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductById)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct)))
}

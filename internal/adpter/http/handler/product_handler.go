package handler

import (
    "net/http"
    "encoding/json"
    "github/ecommerce/internal/usecase"
)

type ProductHandler struct {
    Service *usecase.ProductService
}

// Constructor
func NewProductHandler(service *usecase.ProductService) *ProductHandler {
    return &ProductHandler{Service: service}
}

// Handler method
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
    products, _ := h.Service.GetAllProducts()
    json.NewEncoder(w).Encode(products)
}

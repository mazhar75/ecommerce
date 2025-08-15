package handler

import (
	"encoding/json"
	"github/ecommerce/internal/usecase"
	"net/http"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products, err := h.Service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

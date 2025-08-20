package handler

import (
	"encoding/json"
	"github/ecommerce/internal/domain/product"
	"github/ecommerce/internal/usecase"
	"net/http"
	"strconv"
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

	products, err := h.Service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}
func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	product, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "No id found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)

}
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var prod product.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&prod)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	err = h.Service.InsertProduct(prod)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resource created successfully"))

}

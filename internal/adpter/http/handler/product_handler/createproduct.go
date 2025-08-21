package product_handler

import (
	"encoding/json"
	"github/ecommerce/internal/domain/product"
	"net/http"
)

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

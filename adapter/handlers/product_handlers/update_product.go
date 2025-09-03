package product_handlers

import (
	"encoding/json"
	"github/ecommerce/domain/product"
	"net/http"
	"strconv"
)

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("productId")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var prod product.Product
	prod.ProductId = id
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&prod)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	err = h.Service.UpdateProduct(prod)
	if err != nil {
		http.Error(w, "invalid product", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resource updated successfully"))

}

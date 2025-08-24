package product_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {

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

	product, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "No id found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)

}

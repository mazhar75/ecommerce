package category_handlers

import (
	"encoding/json"
	"github/ecommerce/domain/product"
	"net/http"
)

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {

	var cat product.Category
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cat)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	cat, err = h.Service.InsertCategory(cat)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resource created successfully"))

}

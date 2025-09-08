package product_handlers

import (
	"net/http"
	"strconv"
)

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("productId")
	pId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id ", http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteProduct(pId)
	if err != nil {
		http.Error(w, "not found product with given id", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Resource deleted successfully"))

}

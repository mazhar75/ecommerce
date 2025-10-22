package cart_item_handler

import (
	"net/http"
	"strconv"
	"strings"
)

// DELETE /cart/{user_id}/remove/{product_id}  → Remove product
func (h *CartItemHandler) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	user_id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	pid, err := strconv.Atoi(parts[4])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteProductFromCart(user_id, pid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cart Item deleted successfully"))

}

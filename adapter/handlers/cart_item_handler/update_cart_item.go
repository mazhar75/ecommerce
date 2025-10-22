package cart_item_handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type updateReqBody struct {
	productId int `json:"product_id"`
	quantity  int `json"quantity"`
}

// PATCH  /cart/{user_id}/update       → Change quantity
func (h *CartItemHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
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
	var product updateReqBody
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	err = h.Service.ChangeQuantity(user_id, product.productId, product.quantity)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cart Item updatated successfully"))
}

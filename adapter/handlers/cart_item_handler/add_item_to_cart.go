package cart_item_handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// POST   /cart/{user_id}/add          → Add product to cart
type addReqBody struct {
	productId int `json:"product_id"`
}

func (h *CartItemHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
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
	var product addReqBody
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	err = h.Service.AddProductToCart(user_id, product.productId)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cart Item added successfully"))
}

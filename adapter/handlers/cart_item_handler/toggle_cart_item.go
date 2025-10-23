package cart_item_handler

import (
	"encoding/json"
	"net/http"
)

type selectReqBody struct {
	CartItemId int `json:"cart_item_id"`
}

// PATCH  /cart/{user_id}/toggle/{product_id}       → Change selection toward place order
func (h *CartItemHandler) ToggleSelection(w http.ResponseWriter, r *http.Request) {

	var cartItem selectReqBody
	err := json.NewDecoder(r.Body).Decode(&cartItem)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	err = h.Service.ToggleProductInCart(cartItem.CartItemId)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Cart Item updatated successfully"))
}

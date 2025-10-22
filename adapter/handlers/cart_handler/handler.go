package cart_handler

import (
	"encoding/json"
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
	"strconv"
)

type CartHandler struct {
	Service *usecase.CartService
}

func NewCartHandler(service *usecase.CartService) *CartHandler {
	return &CartHandler{Service: service}
}
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("userId")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	cartItems, err := h.Service.GetCartByUserId(id)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(cartItems)

}

func (h *CartHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /cart/{userId}",
		manager.With(http.HandlerFunc(h.GetCart),
			middlewares.AuthMiddleware,
		))
}

/*
GET    /cart/{user_id}              → Get a user’s cart
POST   /cart/{user_id}/add          → Add product to cart
PATCH  /cart/{user_id}/update       → Change quantity
DELETE /cart/{user_id}/remove/:pid  → Remove product
*/

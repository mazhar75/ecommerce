package cart_item_handler

import (
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type CartItemHandler struct {
	Service *usecase.CartItemService
}

func NewCartItemHandler(service *usecase.CartItemService) *CartItemHandler {
	return &CartItemHandler{Service: service}
}
func (h *CartItemHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle("POST /cart/{user_id}/add", manager.With(
		http.HandlerFunc(h.AddProductToCart),
		middlewares.AuthMiddleware,
	))

	mux.Handle("PATCH /cart/{user_id}/update", manager.With(
		http.HandlerFunc(h.UpdateQuantity),
		middlewares.AuthMiddleware,
	))
	mux.Handle("DELETE /cart/{user_id}/remove/{product_id}",
		manager.With(
			http.HandlerFunc(h.DeleteFromCart),
			middlewares.AuthMiddleware,
		))
}

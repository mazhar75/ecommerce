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

}

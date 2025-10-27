package checkout_handler

import (
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type CheckoutHandler struct {
	Service *usecase.CheckoutService
}

func NewCheckoutHandler(service *usecase.CheckoutService) *CheckoutHandler {
	return &CheckoutHandler{Service: service}
}
func (h *CheckoutHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /checkout", manager.With(
		http.HandlerFunc(h.CreateCheckout),
		middlewares.AuthMiddleware,
	))
}

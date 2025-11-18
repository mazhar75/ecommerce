package order_handler

import (
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type OrderHandler struct {
	Service *usecase.OrderService
}

func NewOrderHandler(service *usecase.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}
func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /order/place", manager.With(
		http.HandlerFunc(h.AddOrder),
		middlewares.AuthMiddleware,
	))
	mux.Handle("POST /order/changestatus", manager.With(
		http.HandlerFunc(h.ChangeStatus),
		middlewares.AuthMiddleware,
	))
}

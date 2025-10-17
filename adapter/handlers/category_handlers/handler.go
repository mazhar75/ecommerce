package category_handlers

import (
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type CategoryHandler struct {
	Service *usecase.CategoryService
}

func NewProductHandler(service *usecase.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}
func (h *CategoryHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /category", manager.With(
		http.HandlerFunc(h.CreateCategory),
		middlewares.AuthMiddleware,
	))
}

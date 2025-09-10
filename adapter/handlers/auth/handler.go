package auth

import (
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type AuthHandler struct {
	Service *usecase.AuthService
}

func NewAuthHandler(service *usecase.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /auth/register", manager.With(http.HandlerFunc(h.Register)))
}

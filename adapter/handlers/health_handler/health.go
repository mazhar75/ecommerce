package health_handler

import (
	"encoding/json"
	"github/ecommerce/middlewares"
	"github/ecommerce/usecase"
	"net/http"
)

type HealthHandler struct {
	Service *usecase.HealthService
}

func NewHealthHandler(service *usecase.HealthService) *HealthHandler {
	return &HealthHandler{Service: service}
}
func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	health := h.Service.GetHealth()
	json.NewEncoder(w).Encode(health)
}
func (h *HealthHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /health", manager.With(http.HandlerFunc(h.GetHealth)))
}

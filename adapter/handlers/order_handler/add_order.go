package order_handler

import (
	"encoding/json"
	"net/http"
)

type ReqFormat struct {
	UserId int `json:"user_id"`
	CartId int `json:"cart_id"`
}

func (h *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var req ReqFormat
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	ord, err := h.Service.OrderPlacement(req.UserId, req.CartId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.WriteHeader(http.StatusCreated)
	data, err := json.Marshal(ord)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Order Details\n"))
	w.Write(data)
}

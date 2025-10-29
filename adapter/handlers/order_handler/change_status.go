package order_handler

import (
	"encoding/json"
	"net/http"
)

type ReqFormat1 struct {
	OrderId int    `json:"order_id"`
	Status  string `json:"status"`
}

func (h *OrderHandler) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	var req ReqFormat1
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	ord, err := h.Service.UpdateOrderStatus(req.OrderId, req.Status)
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

package checkout_handler

import (
	"encoding/json"
	"net/http"
)

type ReqFormat struct {
	CartId int `json:"cart_id"`
}

func (h *CheckoutHandler) CreateCheckout(w http.ResponseWriter, r *http.Request) {
	var reqbody ReqFormat
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqbody)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	outOfStock, err := h.Service.AddCheckoutfromPgtoRedis(reqbody.CartId)
	if err != nil {
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	data, err := json.Marshal(outOfStock)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

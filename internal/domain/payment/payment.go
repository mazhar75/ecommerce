package payment

type Payment struct {
	PaymentId     int     `json:"payment_id"`
	OrderId       int     `json:"order_id"`
	Status        string  `json:"status"`
	Amount        float64 `json:"anount"`
	TransactionId string  `json:"transaction_id"`
	PaymentMethod string  `json:"payment_method"`
}

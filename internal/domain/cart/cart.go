package cart





type CartItems struct {
	CartItemsId int `json:"cartitems_id"`
	CartId      int `json:"cart_id"`
	ProductId   int `json:"product_id"`
	Quantity    int `json:"quantity"`
}

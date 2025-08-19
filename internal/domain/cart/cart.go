package cart

type Cart struct {
	CartId int `json:"cart_id"`
	UserId int `json:"user_id"`
}
type CartItems struct {
	CartItemsId int `json:"cartitems_id"`
	CartId      int `json:"cart_id"`
	ProductId   int `json:"product_id"`
	Quantity    int `json:"quantity"`
}

package cart

type Carts struct {
	CartId int `json:"cart_id"`
	UserID int `json:"user_id"`
}
type CartItems struct {
	CartItemsId int `json:"cartitems_id"`
	CartId      int `json:"cart_id"`
	ProductId   int `json:"product_id"`
	Quantity    int `json:"quantity"`
}

type CartRepository interface {
	GetCartByUserId(user_id int) Carts
}

type CartItemRepository interface {
	AddProductToCart(user_id int, product_id int) error
	ChangeQuantity(user_id int, product_id int, quantity int) error
	DeleteProductFromCart(user_id int, product_id int) error
}

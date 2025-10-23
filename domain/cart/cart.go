package cart

type Carts struct {
	CartId int `json:"cart_id"`
	UserID int `json:"user_id"`
}
type CartItems struct {
	CartItemsId int  `json:"cartitems_id"`
	CartId      int  `json:"cart_id"`
	ProductId   int  `json:"product_id"`
	Quantity    int  `json:"quantity"`
	IsSelected  bool `json:"is_selected"`
}

type CartRepository interface {
	GetCartByUserId(user_id int) ([]CartItems, error)
}

type CartItemRepository interface {
	AddProductToCart(user_id int, product_id int) error
	ChangeQuantity(user_id int, product_id int, quantity int) error
	DeleteProductFromCart(user_id int, product_id int) error
	// select and unselect from checkout
	ToggleProductInCart(cart_item_id int) error
}

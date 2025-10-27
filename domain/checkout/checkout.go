package checkout

type Checkout struct {
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Quantity  int    `json:"qantity"`
}

type OutOfStock struct {
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
}
type CheckoutPgRepository interface {
	GetAllProductsFromCart(cart_id int) ([]Checkout, []OutOfStock, error)
}
type CheckoutRedisRepository interface {
	InsertAllProductsIntoRedis(products []Checkout) ([]OutOfStock, error)
}

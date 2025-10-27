package checkout

import "time"

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

type CheckoutFinal struct {
	CartId      int        `json:"cart_id"`
	ProductList []Checkout `json:"product_list"`
	TotalPrice  float64    `json:"total_price"`
	ExpireAt    time.Time  `json:"expire_at"`
}
type CheckoutPgRepository interface {
	GetAllProductsFromCart(cart_id int) ([]Checkout, []OutOfStock, error)
}
type CheckoutRedisRepository interface {
	InsertAllProductsIntoRedis(products []Checkout, cart_id int) error
}

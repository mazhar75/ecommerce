package order

type Orders struct {
	OrderId     int     `json:"order_id"`
	UserId      int     `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}
type OrderItems struct {
	OrderItemId int     `json:"order_item_id"`
	OrderId     int     `json:"order_id"`
	ProductId   int     `json:"product_id"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
type OrderRepository interface {
	CheckExistanceOfCheckout(cart_id int) (bool, Orders, []OrderItems, error)
	AddOrder(user_id int, total float64, status string) (Orders, error)
	UpdateStatus(order_id int) (Orders, error)
}
type OrderItemsRepository interface {
	AddItems(order_id int, product_id int, quantity int, price int) error
}

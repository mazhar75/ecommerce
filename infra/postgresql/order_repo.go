package postgresql

import (
	"database/sql"
	"fmt"
	"github/ecommerce/domain/order"
)

type OrderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

var _ order.OrderRepository = &OrderRepo{}

func (r *OrderRepo) AddOrder(user_id int, total float64, status string) (order.Orders, error) {
	query := `insert into orders(user_id,total_amount,status) values($1,$2,$3) returning order_id`
	var order_id int
	err := r.DB.QueryRow(query, user_id, total, status).Scan(&order_id)
	if err != nil {
		fmt.Println(err)
		return order.Orders{}, err
	}
	return order.Orders{
		OrderId:     order_id,
		UserId:      user_id,
		TotalAmount: total,
		Status:      status,
	}, nil

}
func (r *OrderRepo) UpdateStatus(order_id int, status string) (order.Orders, error) {
	query := `update orders set status=$1 where order_id=$2`
	_, err := r.DB.Exec(query, status, order_id)
	if err != nil {
		fmt.Println(err)
		return order.Orders{}, err
	}
	query = `select order_id,product_id,total_amount,status from orders where order_id=$1`
	rows, err := r.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return order.Orders{}, err
	}
	defer rows.Close()
	var ord order.Orders
	err = rows.Scan(&ord.OrderId, &ord.UserId, &ord.TotalAmount, &ord.Status)
	if err != nil {
		fmt.Println(err)
		return order.Orders{}, err
	}
	return ord, nil

}

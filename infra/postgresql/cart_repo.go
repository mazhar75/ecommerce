package postgresql

import (
	"database/sql"
	"github/ecommerce/domain/cart"
)

type CartRepo struct {
	DB *sql.DB
}

var _ cart.CartRepository = &CartRepo{}

func NewCartRepo(db *sql.DB) *CartRepo {
	return &CartRepo{DB: db}
}
func (r *CartRepo) GetCartByUserId(user_id int) ([]cart.CartItems, error) {
	query := `select *
	           from cart
			   join cart_item
			   on cart.cart_id=cart_item.cart_id
			   where cart.user_id=$1`
	rows, err := r.DB.Query(query, user_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var items []cart.CartItems

	for rows.Next() {
		var item cart.CartItems
		err := rows.Scan(&item.CartId, &item.CartItemsId, &item.ProductId, &item.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil

}

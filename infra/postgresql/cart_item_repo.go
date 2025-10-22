package postgresql

import (
	"database/sql"
	"fmt"
	"github/ecommerce/domain/cart"
)

type CartItemRepo struct {
	DB *sql.DB
}

var _ cart.CartItemRepository = &cart.CartItemRepository{}

func NewCartItemRepo(db *sql.DB) *CartItemRepo {
	return &CartItemRepo{DB: db}
}

func (r *CartItemRepo) AddProductToCart(user_id int, product_id int) error {
	var cartItemId, quantity int
	query := `select cart_item.cart_item_id, cart_item.quantity
             from cart_item 
			 join cart on cart_item.cart_id=cart.cart_id 
			 where cart.user_id=$1 and cart_item.product_id=$2`
	err := db.QueryRow(query, user_id, product_id).Scan(&cartItemId, &quantity)
	if err == sql.ErrNoRows {
		query = `select cart_id from cart where user_id=$1`
		var cartId, cartItemId int
		err = db.QueryRow(query, user_id).Scan(&cartId)
		if err != nil {
			fmt.Println(err)
			return err
		}
		quantity := 1
		query = `insert into cart_items(cart_id,product_id,quantity) values($1,$2,$3)`
		err = db.QueryRow(query, cartId, product_id, quantity).Scan(&cartItemId)
		fmt.Println("Product added user:", product_id, user_id)
		return nil

	} else if err != nil {
		return err
	} else {
		query = `update cart_item set quanity=$1 where cart_item_id=$2`
		_, err = db.Exec(query, quantity, cartItemId)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

}
func (r *CartItemRepo) ChangeQuantity(user_id int, product_id int, quantity int) error {
	query := `select cart_id from cart where user_id=$1`
	var cartId int
	err := db.QueryRow(query, user_id).Scan(&cartId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	query = `update cart_item set quantity=$1 where cart_id=$2 and product_id=$3`
	_, err = db.Exec(query, quantity, cartId, product_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func (r *CartItemRepo) DeleteProductFromCart(user_id int, product_id int) error {

}

package postgresql

import (
	"database/sql"
	"fmt"
	"github/ecommerce/domain/cart"
)

type CartItemRepo struct {
	DB *sql.DB
}

var _ cart.CartItemRepository = &CartItemRepo{}

func NewCartItemRepo(db *sql.DB) *CartItemRepo {
	return &CartItemRepo{DB: db}
}

func (r *CartItemRepo) AddProductToCart(user_id int, product_id int) error {
	var cartItemId, quantity int
	query := `select cart_item.cart_item_id, cart_item.quantity
             from cart_item 
			 join cart on cart_item.cart_id=cart.cart_id 
			 where cart.user_id=$1 and cart_item.product_id=$2`
	err := r.DB.QueryRow(query, user_id, product_id).Scan(&cartItemId, &quantity)
	if err == sql.ErrNoRows {
		query = `select cart_id from cart where user_id=$1`
		var cartId, cartItemId int
		err = r.DB.QueryRow(query, user_id).Scan(&cartId)
		if err != nil {
			fmt.Println("Here error 1")
			fmt.Println(err)
			return err
		}
		quantity := 1
		isSelected := false
		query = `insert into cart_item(cart_id,product_id,quantity,is_selected) values($1,$2,$3,$4) returning cart_item_id`
		err = r.DB.QueryRow(query, cartId, product_id, quantity, isSelected).Scan(&cartItemId)
		if err != nil {
			fmt.Println("Here error 2")
			fmt.Println(err)
			return err
		}
		fmt.Println("Product added user:", product_id, user_id)
		return nil

	} else if err != nil {
		fmt.Println("Here error 3")
		return err
	} else {
		query = `update cart_item set quantity=$1 where cart_item_id=$2`
		quantity++
		_, err = r.DB.Exec(query, quantity, cartItemId)
		if err != nil {
			fmt.Println("Here error 4")
			fmt.Println(err)
			return err
		}
		return nil
	}

}
func (r *CartItemRepo) ChangeQuantity(user_id int, product_id int, quantity int) error {
	query := `select cart_id from cart where user_id=$1`
	var cartId int
	err := r.DB.QueryRow(query, user_id).Scan(&cartId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	query = `update cart_item set quantity=$1 where cart_id=$2 and product_id=$3`
	_, err = r.DB.Exec(query, quantity, cartId, product_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func (r *CartItemRepo) DeleteProductFromCart(user_id int, product_id int) error {
	query := `select cart_id from cart where user_id=$1`
	var cartId int
	err := r.DB.QueryRow(query, user_id).Scan(&cartId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	query = `delete from cart_item where cart_id=$1 and product_id=$2`
	_, err = r.DB.Exec(query, cartId, product_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

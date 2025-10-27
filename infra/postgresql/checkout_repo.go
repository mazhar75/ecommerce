package postgresql

import (
	"database/sql"
	"fmt"
	"github/ecommerce/domain/checkout"
	"strings"
)

type productCheckList struct {
	ProductId         int     `json:"product_id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	OrderQuantity     int     `json:"order_quantity"`
	InventoryQuantity int     `json:"inventory_quantity"`
	InventoryReserved int     `json:"inventory_reserved"`
}

type CheckoutRepo struct {
	DB *sql.DB
}

func NewCheckoutRepo(db *sql.DB) *CheckoutRepo {
	return &CheckoutRepo{DB: db}
}

var _ checkout.CheckoutPgRepository = &CheckoutRepo{}

func (r *CheckoutRepo) GetAllProductsFromCart(cart_id int) ([]checkout.Checkout, []checkout.OutOfStock, error) {
	query := `select cart_item.product_id, product.name, product.price, cart_item.quantity,inventory.quantity, inventory.reserved
           from cart_item
		   join product on cart_item.product_id=product.product_id
		   join inventory on product.product_id=inventory.product_id
		   where cart_item.cart_id=$1 and cart_item.is_selected=$2`
	rows, err := r.DB.Query(query, cart_id, true)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer rows.Close()
	var productinfo []productCheckList
	for rows.Next() {
		var p productCheckList
		err = rows.Scan(&p.ProductId, &p.Name, &p.Price, &p.OrderQuantity, &p.InventoryQuantity, &p.InventoryReserved)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		productinfo = append(productinfo, p)
	}
	var productCheckList []checkout.Checkout
	var outOfStockList []checkout.OutOfStock
	var updateInventoryTable []struct {
		ProductId int
		Quantity  int
		Reserved  int
	}
	for _, p := range productinfo {
		if p.InventoryQuantity < p.OrderQuantity {
			outOfStockList = append(outOfStockList,
				checkout.OutOfStock{
					ProductId: p.ProductId,
					Name:      p.Name,
				},
			)
		} else {
			productCheckList = append(productCheckList,
				checkout.Checkout{
					ProductId: p.ProductId,
					Name:      p.Name,
					Price:     int(p.Price),
					Quantity:  p.OrderQuantity,
				},
			)
			updateInventoryTable = append(updateInventoryTable,
				struct {
					ProductId int
					Quantity  int
					Reserved  int
				}{
					ProductId: p.ProductId,
					Quantity:  p.InventoryQuantity - p.OrderQuantity,
					Reserved:  p.InventoryReserved + p.OrderQuantity,
				},
			)
		}
	}

	// Perform batch update on inventory table

	if len(updateInventoryTable) > 0 {
		tx, err := r.DB.Begin()
		if err != nil {
			fmt.Println("Error starting transaction:", err)
			return nil, nil, err
		}

		// Prepare the query for batch update
		valueStrings := make([]string, 0, len(updateInventoryTable))
		args := make([]interface{}, 0, len(updateInventoryTable)*3)
		for i, item := range updateInventoryTable {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
			args = append(args, item.ProductId, item.Quantity, item.Reserved)
		}

		query := fmt.Sprintf(`
         UPDATE inventory
         SET quantity = data.quantity::INTEGER, reserved = data.reserved::INTEGER, updated = NOW()
         FROM (VALUES %s) AS data(product_id, quantity, reserved)
         WHERE inventory.product_id = data.product_id::INTEGER`,
			strings.Join(valueStrings, ","))

		// Execute the batch update
		_, err = tx.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			fmt.Println("Error executing batch update:", err)
			return nil, nil, err
		}

		// Commit the transaction
		err = tx.Commit()
		if err != nil {
			fmt.Println("Error committing transaction:", err)
			return nil, nil, err
		}
	}

	return productCheckList, outOfStockList, nil

}

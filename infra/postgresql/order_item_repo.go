package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"github/ecommerce/domain/order"
	"strings"
)

var ctx context.Context

type OrderItemRepo struct {
	DB *sql.DB
}

func NewOrderItemRepo(db *sql.DB) *OrderItemRepo {
	return &OrderItemRepo{DB: db}
}

var _ order.OrderItemsRepository = &OrderItemRepo{}

func (r *OrderItemRepo) AddItems(orderID int, productIDs []int, quantities []int, prices []float64) error {
	if len(productIDs) != len(quantities) || len(productIDs) != len(prices) {
		return fmt.Errorf("productIDs, quantities, and prices must have the same length")
	}

	// Build the VALUES placeholders like: ($1,$2,$3,$4), ($5,$6,$7,$8), ...
	valueStrings := make([]string, 0, len(productIDs))
	valueArgs := make([]interface{}, 0, len(productIDs)*4)

	argPosition := 1
	for i := range productIDs {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", argPosition, argPosition+1, argPosition+2, argPosition+3))
		valueArgs = append(valueArgs, orderID, productIDs[i], quantities[i], prices[i])
		argPosition += 4
	}

	query := fmt.Sprintf(`INSERT INTO order_items (order_id, product_id, quantity, price) VALUES %s`,
		strings.Join(valueStrings, ","),
	)

	_, err := r.DB.ExecContext(ctx, query, valueArgs...)
	if err != nil {
		return fmt.Errorf("failed to insert order items: %w", err)
	}

	return nil
}

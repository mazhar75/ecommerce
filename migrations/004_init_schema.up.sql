CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    total_amount FLOAT NOT NULL,
    status VARCHAR(100) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW()
);

-- ORDER_ITEMS TABLE
CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(order_id) ON DELETE CASCADE,
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    quantity INT NOT NULL DEFAULT 1,
    price FLOAT NOT NULL
);
-- ORDERS
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_status ON orders(status);

-- ORDER_ITEMS
CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_product_id ON order_items(product_id);
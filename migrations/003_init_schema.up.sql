ALTER TABLE cart_item ADD COLUMN is_selected BOOLEAN DEFAULT FALSE;

CREATE TABLE inventory (
    inventory_id SERIAL PRIMARY KEY,
    product_id INT,
    category_id INT,
    quantity INT NOT NULL DEFAULT 0,
    reserved INT NOT NULL DEFAULT 0,
    updated TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (category_id) REFERENCES category(category_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES product(product_id) ON DELETE CASCADE
);

create index inventory_query_product_id on inventory(product_id,quantity); 
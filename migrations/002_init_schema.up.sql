create TABLE cart(
    cart_id SERIAL PRIMARY KEY,
    user_id INT,
    created_at timestamp DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
       ON DELETE CASCADE
);
create table cart_item(
    cart_item_id SERIAL PRIMARY KEY,
    cart_id INT,
    product_id INT,
    quantity INT,
    created_at timestamp DEFAULT NOW(),
    FOREIGN KEY(cart_id) REFERENCES cart(cart_id)
       ON DELETE CASCADE,
    FOREIGN KEY(product_id) REFERENCES product(product_id)
       ON DELETE CASCADE
);
create index cart_user on cart(cart_id,user_id);
create index cart_items on cart_item(cart_item_id,cart_id);
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    is_verified boolean
);

CREATE TABLE category (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    category_id INT,
    name VARCHAR(100),
    description VARCHAR(255),
    type VARCHAR(100),
    price FLOAT,
    img_url VARCHAR(255),
    FOREIGN KEY (category_id) REFERENCES category(category_id)
);

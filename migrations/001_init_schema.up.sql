CREATE TABLE user (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    email VARCHAR(25) UNIQUE,
    password VARCHAR(50),
    is_verified boolean
);

CREATE TABLE category (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(30)
);

CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    category_id INT,
    name VARCHAR(20),
    description VARCHAR(50),
    type VARCHAR(20),
    price FLOAT,
    img_url VARCHAR(100),
    FOREIGN KEY (category_id) REFERENCES category(category_id)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    is_verified boolean,
    is_shop_owner boolean,
    created_at timestamp DEFAULT NOW(),
    last_login  timestamp DEFAULT NOW()
);

CREATE TABLE category (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    created_at timestamp DEFAULT NOW()
);

CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    category_id INT,
    name VARCHAR(100),
    description VARCHAR(255),
    type VARCHAR(100),
    price FLOAT,
    img_url VARCHAR(255),
    created_at timestamp DEFAULT NOW(),
    FOREIGN KEY (category_id) REFERENCES category(category_id)
     ON DELETE CASCADE
     
);

create index users_mail on users(email);
create index category_id on category(category_id);
create index categoryId_productId on product(product_id,category_id);

CREATE TABLE users (
  id   string PRIMARY KEY,
  name text    NOT NULL,
  bio  text    NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME
);

CREATE TABLE products (
  id   string  PRIMARY KEY,
  name text    NOT NULL
);

CREATE TABLE orders (
  id   string  PRIMARY KEY,
  name text    NOT NULL,
  user_id string,
  product_id string,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(product_id) REFERENCES products(id)
);
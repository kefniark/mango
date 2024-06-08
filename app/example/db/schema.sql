CREATE TABLE users (
  id   UUID PRIMARY KEY,
  name text    NOT NULL,
  bio  text    NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE products (
  id   UUID  PRIMARY KEY,
  name text    NOT NULL
);

CREATE TABLE orders (
  id   UUID  PRIMARY KEY,
  name text    NOT NULL,
  user_id UUID,
  product_id UUID,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(product_id) REFERENCES products(id)
);
-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    username VARCHAR(12) NOT NULL,
    password TEXT NOT NULL,
    age INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    stock INT NOT NULL,
    price INT NOT NULL,
    description TEXT NOT NULL,
    category_id INT NOT NULL,
  CONSTRAINT category_relation
      FOREIGN KEY(category_id) 
    REFERENCES categories(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    checkout BOOLEAN NOT NULL,
    paid BOOLEAN NOT NULL,
    total INT NOT NULL,
	CONSTRAINT product_relation
      FOREIGN KEY(product_id) 
	  REFERENCES products(id),
	CONSTRAINT user_relation
      FOREIGN KEY(user_id) 
	  REFERENCES users(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE histories (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    bank_name VARCHAR(256) NOT NULL,
  CONSTRAINT product_relation
      FOREIGN KEY(product_id) 
    REFERENCES products(id),
  CONSTRAINT user_relation
      FOREIGN KEY(user_id) 
    REFERENCES users(id),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

-- CREATE OR REPLACE FUNCTION trigger_set_timestamp()
-- RETURNS TRIGGER AS $$
-- BEGIN
--   NEW.updated_at = NOW();
--   RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER set_timestamp
-- BEFORE UPDATE ON users
-- FOR EACH ROW
-- EXECUTE PROCEDURE trigger_set_timestamp();

-- CREATE TRIGGER set_timestamp
-- BEFORE UPDATE ON categories
-- FOR EACH ROW
-- EXECUTE PROCEDURE trigger_set_timestamp();

-- CREATE TRIGGER set_timestamp
-- BEFORE UPDATE ON products
-- FOR EACH ROW
-- EXECUTE PROCEDURE trigger_set_timestamp();

-- CREATE TRIGGER set_timestamp
-- BEFORE UPDATE ON carts
-- FOR EACH ROW
-- EXECUTE PROCEDURE trigger_set_timestamp();

-- CREATE TRIGGER set_timestamp
-- BEFORE UPDATE ON histories
-- FOR EACH ROW
-- EXECUTE PROCEDURE trigger_set_timestamp();

-- +migrate StatementEnd
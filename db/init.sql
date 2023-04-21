-- psql -U user_coffee
CREATE DATABASE coffee_delivery_remix_db;
-- CREATE user user_coffee;
ALTER user user_coffee WITH ENCRYPTED PASSWORD '1459';
GRANT ALL privileges ON DATABASE coffee_delivery_remix_db TO user_coffee;
\c coffee_delivery_remix_db;

CREATE OR REPLACE FUNCTION trigger_timestamps() RETURNS TRIGGER AS $$
  BEGIN
    NEW.updated_at = NOW();
    return NEW;
  END;
$$ LANGUAGE plpgsql;

CREATE TABLE product (
  id serial primary key, 
  img varchar, 
  price integer, 
  title varchar, 
  description text, 
  stok integer, 
  categories varchar,
  type varchar,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE cart (
  id serial primary key
);

CREATE TABLE cart_product (
  id serial PRIMARY KEY,
  product_id INT NOT NULL,
  cart_id INT NOT NULL,
  quantity integer,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_cart FOREIGN KEY(cart_id) REFERENCES cart(id) ON DELETE CASCADE,
  CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES product(id) ON DELETE CASCADE
);

-- Create timestamps TRIGGER to all tables that have the column 
DO $$
DECLARE 
  t text;
BEGIN
  FOR t IN 
    SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
  LOOP
    EXECUTE format('
      CREATE TRIGGER timestamps 
        BEFORE UPDATE ON %I
        FOR EACH ROW EXECUTE FUNCTION trigger_timestamps();', t);
  END LOOP; 
END;
$$ LANGUAGE plpgsql;

GRANT ALL privileges ON ALL tables IN SCHEMA public TO user_coffee;

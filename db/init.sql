-- psql -U user_coffee
CREATE DATABASE coffee_delivery_remix_db;
-- CREATE user user_coffee;
ALTER user user_coffee WITH ENCRYPTED PASSWORD '1459';
GRANT ALL privileges ON DATABASE coffee_delivery_remix_db TO user_coffee;
\c coffee_delivery_remix_db;
CREATE TABLE product (
  id serial primary key, 
  img varchar, 
  price integer, 
  title varchar, 
  description text, 
  stok integer, 
  categories varchar,
  type varchar
);
GRANT ALL privileges ON ALL tables IN SCHEMA public TO user_coffee;

CREATE TABLE cart (
  id serial PRIMARY KEY,
  product_id INT NOT NULL,
  quantity integer,
  CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES product(id)
);
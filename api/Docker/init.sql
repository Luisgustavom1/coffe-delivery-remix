-- psql -U user_coffe
CREATE DATABASE api_coffe_delivery_remix;
-- CREATE user user_coffe;
ALTER user user_coffe WITH ENCRYPTED PASSWORD '1459';
GRANT ALL privileges ON DATABASE api_coffe_delivery_remix TO user_coffe;
\c api_coffe_delivery_remix;
CREATE TABLE coffes (id serial primary key, img varchar, price bigint, title varchar, description text, stok bigint, categories varchar);
GRANT ALL privileges ON ALL tables IN SCHEMA public TO user_coffe;

CREATE TABLE cart (
  id serial PRIMARY KEY,
  quantity integer,
  productId integer REFERENCES coffes
);
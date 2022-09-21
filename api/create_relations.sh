psql -U user_coffe
create database api_coffe_delivery_remix; 
create user user_coffe; 
alter user user_coffe with encrypted password '1459'; 
grant all privileges on database api_coffe_delivery_remix to user_coffe; 
\c api_coffe_delivery_remix; 
create table coffes (id serial primary key, img varchar, price bigint, title varchar, description text, stok bigint); 
grant all privileges on all tables in schema public to user_coffe;
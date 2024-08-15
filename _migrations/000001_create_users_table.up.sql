CREATE TABLE IF NOT EXISTS users (
  id varchar(255) not null PRIMARY KEY,
  name varchar(255) not null,
  email varchar(255) not null UNIQUE,
  password varchar(255) not null
);
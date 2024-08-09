CREATE TABLE users (
    id serial PRIMARY KEY,
    email varchar(80) UNIQUE,
    username varchar(80) ,
    "password" varchar(255)
)


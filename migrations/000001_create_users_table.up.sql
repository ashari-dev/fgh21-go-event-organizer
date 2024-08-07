CREATE TABLE users (
    id serial PRIMARY KEY,
    email varchar(80),
    username varchar(80),
    "password" varchar(255)
)
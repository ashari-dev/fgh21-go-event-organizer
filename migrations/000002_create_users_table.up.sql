CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    username VARCHAR(50),
    password VARCHAR(255) NOT NULL,
    role_id INT REFERENCES roles(id)
)

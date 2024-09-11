CREATE TABLE events(
    id SERIAL PRIMARY KEY,
    image VARCHAR(255),
    title VARCHAR(100) NOT NULL,
    "date" VARCHAR(20),
    "description" TEXT,
    location_id INT REFERENCES locations(id),
    created_by INT REFERENCES users(id)
)
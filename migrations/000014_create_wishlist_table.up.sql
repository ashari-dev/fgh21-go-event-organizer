CREATE TABLE wishlist(
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL,
    event_id INT REFERENCES events(id) NOT NULL
)
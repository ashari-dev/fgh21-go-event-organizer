CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    event_id INT REFERENCES events(id),
    payment_method_id INT REFERENCES payment_method(id),
    user_id INT REFERENCES users(id)
)
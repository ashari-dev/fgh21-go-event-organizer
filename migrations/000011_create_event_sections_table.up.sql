CREATE TABLE event_sections(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    price INT,
    quantity INT,
    event_id INT REFERENCES events(id)
)

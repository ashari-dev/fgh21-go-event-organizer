CREATE TABLE event_categories(
    id SERIAL PRIMARY KEY,
    event_id INT REFERENCES events(id), 
    category_id INT REFERENCES categories(id)
)
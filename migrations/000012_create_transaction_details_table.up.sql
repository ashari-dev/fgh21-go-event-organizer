CREATE TABLE transaction_details (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions("id"),
    section_id INT REFERENCES event_sections("id"),
    ticket_qty INT
)
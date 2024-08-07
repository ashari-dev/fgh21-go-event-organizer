CREATE TABLE event_categories(
    id serial PRIMARY KEY,
    events_id int references events(id),
    category_id int references categories(id)
)
CREATE TABLE events(
    id serial PRIMARY KEY,
    "image" varchar(255),
    title varchar(100),
    "date" date,
    "description" text,
    location_id int references locations(id),
    create_by int references users(id)
)
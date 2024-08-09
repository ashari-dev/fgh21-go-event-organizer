CREATE TABLE events(
    id serial PRIMARY KEY,
    "image" varchar(255),
    title varchar(100),
    "date" VARCHAR(50),
    "description" text,
    location_id int references locations(id),
    created_by int references users(id)
)


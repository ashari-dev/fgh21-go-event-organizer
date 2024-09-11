CREATE TABLE profile(
    id serial PRIMARY KEY,
    picture varchar(255),
    full_name varchar(100),
    birth_date VARCHAR,
    gender smallint,
    phone_number varchar(20),
    profession varchar(50),
    nationality_id int references nationalities(id),
    users_id int references users(id)
) 
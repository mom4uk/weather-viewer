CREATE TABLE locations
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255),
    user_id INTEGER,
    latitude DECIMAL,
    longitude DECIMAL,

    CONSTRAINT fk_locations_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);
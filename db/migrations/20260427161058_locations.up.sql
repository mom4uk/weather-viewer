CREATE TABLE locations
(
    id INTEGER GENERATED ALWAYS AS IDENTITY,
    name VARCHAR,
    user_id   INTEGER,
    latitude  DECIMAL,
    longitude DECIMAL
);
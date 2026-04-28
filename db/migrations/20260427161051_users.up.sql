CREATE TABLE users
(
    id       GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login    VARCHAR,
    password VARCHAR,
);
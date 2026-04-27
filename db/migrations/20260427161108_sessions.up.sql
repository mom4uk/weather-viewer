CREATE TABLE sessions
(
    id         INTEGER GENERATED ALWAYS AS IDENTITY,
    user_id    INTEGER,
    expires_at DATETIME
);
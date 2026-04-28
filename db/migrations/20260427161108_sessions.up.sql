CREATE TABLE sessions
(
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INTEGER,
    expires_at TIMESTAMP,

    CONSTRAINT fk_sessions_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);
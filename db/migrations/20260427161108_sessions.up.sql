CREATE TABLE sessions
(
    id UUID PRIMARY KEY,
    user_id INTEGER,
    expires_at TIMESTAMP,

    CONSTRAINT fk_sessions_users
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);
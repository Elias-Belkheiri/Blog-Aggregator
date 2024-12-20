-- +goose Up
CREATE TABLE feeds (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT DEFAULT NULL,
    user_id TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

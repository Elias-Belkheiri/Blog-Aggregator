-- +goose Up
ALTER TABLE users ADD COLUMN apikey VARCHAR(64) DEFAULT encode(sha256(random()::text::bytea), 'hex');
ALTER TABLE users ADD CONSTRAINT users_apikey_unique UNIQUE (apikey);
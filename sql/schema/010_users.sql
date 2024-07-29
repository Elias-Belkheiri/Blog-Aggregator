-- +goose Up
ALTER TABLE users DROP COLUMN apikey;
ALTER TABLE users ADD COLUMN email text NOT NULL UNIQUE;
ALTER TABLE users ADD COLUMN password text NOT NULL;
-- +goose Up
ALTER TABLE users ADD CONSTRAINT username_unique UNIQUE (username); 
-- +goose Up
ALTER TABLE users RENAME COLUMN name TO username;
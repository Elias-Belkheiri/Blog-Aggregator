-- +goose Up
ALTER TABLE users ALTER COLUMN id SET DATA TYPE text;
-- +goose Up
ALTER TABLE feeds ADD CONSTRAINT unique_url_constraint UNIQUE (url);
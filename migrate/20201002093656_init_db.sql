-- +goose Up
CREATE EXTENSION "uuid-ossp";
CREATE TABLE "user" ( name text NOT NULL, created_at timestamp NOT NULL);


-- +goose Down
DROP EXTENSION "uuid-ossp";
DROP TABLE "user";

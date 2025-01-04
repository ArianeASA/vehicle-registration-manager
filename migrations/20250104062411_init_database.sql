-- +goose NO TRANSACTION

-- +goose Up
create schema if not exists prod;

-- +goose Down
drop schema prod;



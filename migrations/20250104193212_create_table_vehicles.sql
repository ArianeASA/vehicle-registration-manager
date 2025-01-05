-- +goose Up
-- +goose StatementBegin
create table if not exists vehicles
(
    id    varchar(36) not null
        constraint vehicle_pkey
            primary key,
    brand varchar(50),
    model varchar(50),
    year  integer,
    color varchar(50),
    price numeric(10, 2)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists vehicles;
-- +goose StatementEnd

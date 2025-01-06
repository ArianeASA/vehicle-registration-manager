-- +goose Up
-- +goose StatementBegin
alter table vehicles
    add status varchar(50);

alter table vehicles
    add license_plate varchar(50);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table vehicles
    drop column status;

alter table vehicles
    drop column license_plate;
-- +goose StatementEnd

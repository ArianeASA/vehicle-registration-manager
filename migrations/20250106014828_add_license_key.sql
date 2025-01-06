-- +goose Up
-- +goose StatementBegin
alter table vehicles
    alter column license_plate set not null;

alter table vehicles
    add constraint license_plate_pk
        unique (license_plate, id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table vehicles
    alter column license_plate drop not null;

alter table vehicles
    drop constraint license_plate_pk;
-- +goose StatementEnd

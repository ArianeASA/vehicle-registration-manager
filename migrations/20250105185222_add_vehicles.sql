-- +goose Up
-- +goose StatementBegin
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Toyota', 'etios', 2022, 'red', 13223.63);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Fiat', 'uno', 2022, 'yellow', 39676.24);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Honda', 'carro', 2023, 'prata', 324343.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Chevrolet', 'cruze', 2022, 'black', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Ford', 'ka', 2022, 'white', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Volkswagen', 'gol', 2000, 'blue', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Renault', 'kwid', 2019, 'green', 56673.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Hyundai', 'hb20', 2004, 'orange', 78232.44);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Nissan', 'kicks', 2022, 'brown', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Peugeot', '208', 2022, 'silver', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Citroen', 'c3', 2022, 'gray', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Mitsubishi', 'lancer', 2022, 'purple', 23423.00);
insert into vehicles (id, brand, model, year, color, price) values (gen_random_uuid(), 'Kia', 'sportage', 2022, 'pink', 23423.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from vehicles;
-- +goose StatementEnd

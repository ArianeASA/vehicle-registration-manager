-- +goose Up
-- +goose StatementBegin

WITH unique_values AS (
    SELECT id,
           upper(SUBSTRING(replace(gen_random_uuid()::text, '-', ''), 29, 6)) ||
            SUBSTRING(replace(random()::text, '.', ''), 3, 4) ||
            '-BR' AS random_string  FROM vehicles
)
UPDATE vehicles
SET license_plate = uv.random_string, status = 'FOR_SALE'
FROM unique_values uv
WHERE vehicles.id = uv.id
  AND vehicles.license_plate IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
UPDATE vehicles SET license_plate = NULL, status = NULL;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE order(
    id SERIAL PRIMARY KEY,
    First_name VARCHAR(100),
    Last_name VARCHAR(100),
    order_date DATE
);
-- +goose StatementEnd
-- +goose Down

drop order
-- +goose StatementBegin
-- +goose StatementEnd
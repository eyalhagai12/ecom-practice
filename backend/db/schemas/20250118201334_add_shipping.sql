-- +goose Up
-- +goose StatementBegin
CREATE TYPE shipping_status AS ENUM ('pending', 'collected', 'shippping' ,'delivered', 'cancelled');


CREATE TABLE IF NOT EXISTS shipping (
    id UUID PRIMARY KEY NOT NULL,
    status shipping_status NOT NULL DEFAULT 'pending',
    order_id UUID REFERENCES "order"(id) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shipping;
DROP TYPE IF EXISTS shipping_status;
-- +goose StatementEnd

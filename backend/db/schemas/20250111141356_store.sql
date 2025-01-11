-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS STORE (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS store;
-- +goose StatementEnd

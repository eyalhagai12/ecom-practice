-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product (
  id UUID PRIMARY KEY NOT NULL,
  name VARCHAR(800) NOT NULL,
  price FLOAT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  store_id UUID REFERENCES store(id) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product;
-- +goose StatementEnd

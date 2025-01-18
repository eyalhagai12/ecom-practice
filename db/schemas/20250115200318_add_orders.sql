-- +goose Up
-- +goose StatementBegin
CREATE TYPE order_status AS ENUM (
  'pending',
  'processing',
  'shipping',
  'delivered',
  'cancelled'
);

CREATE TABLE IF NOT EXISTS "order" (
  id UUID PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status order_status NOT NULL DEFAULT 'pending',
  total_price FLOAT NOT NULL DEFAULT 0.0
);

CREATE TABLE IF NOT EXISTS order_item (
  id BIGSERIAL PRIMARY KEY,
  order_id UUID REFERENCES "order"(id) NOT NULL,
  product_id UUID REFERENCES product(id) NOT NULL,
  quantity INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_item;
DROP TABLE IF EXISTS "order";
DROP TYPE order_status;
-- +goose StatementEnd

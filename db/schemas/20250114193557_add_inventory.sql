-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS location (
  id BIGSERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(800) NOT NULL,
  address VARCHAR(800) NOT NULL
);

CREATE TABLE IF NOT EXISTS inventory (
  id UUID PRIMARY KEY NOT NULL,
  product_id UUID REFERENCES product(id) NOT NULL,
  quantity INT NOT NULL,
  location_id BIGSERIAL REFERENCES location(id) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS location;
-- +goose StatementEnd

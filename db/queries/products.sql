-- name: GetStoreProducts :many
SELECT * FROM product WHERE store_id = $1;

-- name: GetProductByID :one
SELECT * FROM product WHERE id = $1;

-- name: CreateProduct :one
INSERT INTO product (id, name, price, quantity, store_id) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateProduct :one
UPDATE product SET name = $2, price = $3, quantity = $4 WHERE id = $1 RETURNING *;

-- name: DeleteProduct :one
UPDATE product SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING *;
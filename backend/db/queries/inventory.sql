-- name: GetProductInventories :many
SELECT * FROM inventory WHERE product_id = $1;

-- name: GetInventoryByID :one
SELECT * FROM inventory WHERE id = $1;

-- name: CreateLocation :one
INSERT INTO location (name, address) VALUES ($1, $2) RETURNING *;

-- name: CreateInventory :one
INSERT INTO inventory (id, product_id, quantity, location_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateInventory :one
UPDATE inventory SET quantity = $2 WHERE id = $1 RETURNING *;

-- name: DeleteInventory :one
UPDATE inventory SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING *;
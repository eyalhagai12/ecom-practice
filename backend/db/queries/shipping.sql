-- name: GetShippingByOrderId :one
SELECT * FROM shipping WHERE order_id = $1;

-- name: GetShippingByUUID :one
SELECT * FROM shipping WHERE id = $1;

-- name: CreateShippingProcess :one
INSERT INTO shipping (id, order_id) VALUES ($1, $2) RETURNING *;

-- name: UpdateShippingStatus :one
UPDATE shipping SET status = $2 WHERE id = $1 RETURNING *;

-- name: DeleteShipping :one
UPDATE shipping SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING *;
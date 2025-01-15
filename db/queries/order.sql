-- name: GetOrderByUUID :many
SELECT sqlc.embed(o), sqlc.embed(oi) 
FROM "order" as o
    JOIN order_item oi ON oi.order_id = o.id
WHERE o.id = $1;

-- name: GetOrders :many
SELECT * FROM "order";

-- name: CreateOrder :one
INSERT INTO "order" (id, total_price) VALUES ($1, $2) RETURNING *;

-- name: UpdateOrderStatus :one
UPDATE "order" SET status = $2 WHERE id = $1 RETURNING *;

-- name: DeleteOrder :one
UPDATE "order" SET status = "cancelled" WHERE id = $1 RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_item (order_id, product_id, quantity) VALUES ($1, $2, $3) RETURNING *;
-- name: GetStoreByUUID :one
SELECT * FROM store
WHERE id = $1;

-- name: InsertNewStore :one
INSERT INTO store (
    id, name
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateStore :one
UPDATE store
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteStore :one
UPDATE store
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;
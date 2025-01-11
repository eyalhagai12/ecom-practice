-- name: GetStoreByUUID :one
SELECT * FROM store
WHERE id = $1;

-- name: InsertNewStore :one
INSERT INTO store (
    id, name
) VALUES (
    $1, $2
) RETURNING *;

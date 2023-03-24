-- name: GetChain :one
SELECT *
FROM chains
WHERE id = $1 LIMIT 1;

-- name: GetManyChain :many
SELECT *
FROM chains;
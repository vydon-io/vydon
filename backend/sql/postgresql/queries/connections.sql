-- name: GetConnectionById :one
SELECT * from vydon_api.connections WHERE id = $1;

-- name: GetConnectionsByIds :many
SELECT * from vydon_api.connections WHERE id = ANY($1::uuid[]);

-- name: GetConnectionByNameAndAccount :one
SELECT c.* from vydon_api.connections c
INNER JOIN vydon_api.accounts a ON a.id = c.account_id
WHERE a.id = sqlc.arg('accountId') AND c.name = sqlc.arg('connectionName');

-- name: GetConnectionsByAccount :many
SELECT c.* from vydon_api.connections c
INNER JOIN vydon_api.accounts a ON a.id = c.account_id
WHERE a.id = sqlc.arg('accountId')
ORDER BY c.created_at DESC;

-- name: RemoveConnectionById :exec
DELETE FROM vydon_api.connections WHERE id = $1;

-- name: RemoveConnectionByNameAndAccount :exec
DELETE FROM vydon_api.connections WHERE name = $1 and account_id = $2;

-- name: CreateConnection :one
INSERT INTO vydon_api.connections (
  name, account_id, connection_config, created_by_id, updated_by_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateConnection :one
UPDATE vydon_api.connections
SET name = $1, connection_config = $2,
updated_by_id = $3
WHERE id = $4
RETURNING *;

-- name: IsConnectionNameAvailable :one
SELECT count(c.id) from vydon_api.connections c
INNER JOIN vydon_api.accounts a ON a.id = c.account_id
WHERE a.id = sqlc.arg('accountId') and c.name = sqlc.arg('connectionName');

-- name: IsConnectionInAccount :one
SELECT count(c.id) from vydon_api.connections c
INNER JOIN vydon_api.accounts a ON a.id = c.account_id
WHERE a.id = sqlc.arg('accountId') and c.id = sqlc.arg('connectionId');


-- name: AreConnectionsInAccount :one
SELECT count(c.id) from vydon_api.connections c
INNER JOIN vydon_api.accounts a ON a.id = c.account_id
WHERE a.id = sqlc.arg('accountId') and c.id = ANY(sqlc.arg('connectionIds')::uuid[]);;

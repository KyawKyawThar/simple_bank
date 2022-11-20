-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currentcy)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccounts :one
SELECT *
FROM accounts
WHERE id = $1 LIMIT 1;


-- name: ListAccounts :many
SELECT *
FROM accounts
ORDER BY id LIMIT $1
OFFSET $2;

-- name: UpdateAccounts :one
UPDATE accounts
set balance = $2
WHERE id = $1 RETURNING *;

-- name: DeleteAccounts :exec
DELETE
FROM accounts
WHERE id = $1;

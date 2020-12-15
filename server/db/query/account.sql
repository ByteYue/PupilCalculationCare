-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  message,
  mistakes,
  password
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE owner = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY owner
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET message = $2,mistakes = $3,password = $4
WHERE owner = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE owner = $1;
-- name: AddCustomer :one
INSERT INTO customer (first_name, last_name, address, username, password)
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: RemoveCustomer :exec
DELETE FROM customer
WHERE id = $1;

-- name: GetCustomer :one
SELECT * FROM customer
WHERE id = $1;

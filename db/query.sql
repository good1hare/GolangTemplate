-- name: GetUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;

-- name: GetUserByPhone :one
SELECT *
FROM users
WHERE phone = ?
LIMIT 1;

-- name: ListAllUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateUser :execresult
INSERT INTO users (name, phone)
VALUES (?, ?);

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = ?;
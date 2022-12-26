-- name: CreateUser :one
INSERT INTO users (
	id, username, email, password, name 
)VALUES(
	$1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one 
SELECT * FROM users 
WHERE username = $1
OR email = $1 LIMIT 1;

-- name: Me :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

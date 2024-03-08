-- name: FindUserByUsername :one
SELECT uuid, username, password
FROM idm_users
WHERE username = $1;

-- name: CreateUser :exec
INSERT INTO idm_users (username, password, first_name, last_name, created_at, last_modified_at, deleted_at)
VALUES ($1, $2, $3, $4, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING uuid;

-- name: UpdateUser :exec
UPDATE idm_users
SET username = $2, password = $3, first_name = $4, last_name = $5, last_modified_at = NOW() AT TIME ZONE 'utc'
WHERE uuid = $1;

-- name: DeleteUserByUsername :exec
DELETE FROM idm_users
WHERE username = $1;
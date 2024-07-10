-- name: CreateUserFromGoogle :one
INSERT INTO users (authprovider, authprovider_id)
VALUES ('google', $1)
RETURNING id;

-- name: CreateUserFromEmail :one
INSERT INTO users (authprovider, authprovider_id)
VALUES ('email', $1)
RETURNING id;

-- name: GrantPermission :exec
INSERT INTO users_permissions (user_id, role)
VALUES ($1, $2);

-- name: RevokePermission :exec
DELETE FROM users_permissions WHERE user_id = $1 AND role = $2;



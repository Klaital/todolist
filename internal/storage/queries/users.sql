-- name: CreateUserFromGoogle :one
INSERT INTO users (authprovider, authprovider_id)
VALUES ('google', ?)
RETURNING id;

-- name: CreateUserFromEmail: one
INSERT INTO users (authprovider, authprovider_id)
VALUES ('email', ?)
RETURNING id;

-- name: GrantPermission :exec
INSERT INTO users_permissions (user_id, role)
VALUES (?, ?);

-- name: RevokePermission :exec
DELETE FROM users_permissions WHERE user_id = ? AND role = ?;



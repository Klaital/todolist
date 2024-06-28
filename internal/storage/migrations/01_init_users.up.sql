CREATE TABLE users (
    id INTEGER NOT NULL PRIMARY KEY,
    authprovider VARCHAR(32) NOT NULL,
    authprovider_id VARCHAR(128) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(authprovider, authprovider_id)
);
CREATE INDEX authprovider_users ON users (authprovider, authprovider_id);

CREATE TABLE users_permissions (
    id INTEGER NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    role VARCHAR(16) NOT NULL,
    UNIQUE(user_id, role)
);
CREATE INDEX permissions_for_user ON users_permissions (user_id);
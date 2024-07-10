CREATE TABLE todolists (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    name VARCHAR(128) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(owner_id, name)
);
CREATE INDEX list_owner ON todolists (owner_id);

CREATE TABLE list_items (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    todolist_id INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    status INTEGER NOT NULL DEFAULT 1,
    descr TEXT NOT NULL
);
CREATE INDEX item_owner ON list_items (todolist_id);

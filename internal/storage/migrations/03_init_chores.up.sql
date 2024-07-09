
CREATE TABLE chores (
    id INTEGER NOT NULL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    name VARCHAR(128) NOT NULL,
    last_completed TIMESTAMPTZ,
    next_deadline TIMESTAMPTZ,
    UNIQUE(owner_id, name)
);

CREATE INDEX chore_owner ON chores(owner_id);

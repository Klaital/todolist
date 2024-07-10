
CREATE TABLE chores (
    id BIGSERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    name VARCHAR(128) NOT NULL,
    last_completed TIMESTAMPTZ NOT NULL,
    next_deadline TIMESTAMPTZ NOT NULL,
    frequency_days INTEGER NOT NULL DEFAULT 7,
    UNIQUE(owner_id, name)
);

CREATE INDEX chore_owner ON chores(owner_id);

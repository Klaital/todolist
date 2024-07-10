-- name: CreateChore :one
INSERT INTO chores (owner_id, name, last_completed, next_deadline, frequency_days) VALUES ($1, $2, 0, 0, $3)
    RETURNING id;

-- name: UpdateChoreCompleted :exec
UPDATE chores SET last_completed=$2, next_deadline=$3 WHERE id=$1;

-- name: ListChores :many
SELECT * FROM chores WHERE owner_id=$1;

-- name: GetChore :one
SELECT * FROM chores WHERE id=$1;

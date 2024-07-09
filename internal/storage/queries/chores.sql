-- name: CreateChore :one
INSERT INTO chores (owner_id, name, last_completed, next_deadline)
VALUES (?, ?, 0, 0)
    RETURNING id;


-- name: UpdateChoreCompleted :exec
UPDATE chores SET last_completed=?, next_deadline=? WHERE id=?;

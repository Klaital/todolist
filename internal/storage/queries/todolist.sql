-- name: CreateList :one
INSERT INTO todolists (owner_id, name)
VALUES ($1, $2)
RETURNING id;

-- name: AddToList :one
INSERT INTO list_items (todolist_id, descr)
VALUES ($1, $2)
RETURNING id;

-- name: UpdateStatus :exec
UPDATE list_items SET status=$2 WHERE id=$1;

-- name: DeleteListItem :exec
DELETE FROM list_items WHERE id=$1;

-- name: DeleteList :exec
-- DELETE FROM list_items FROM list_items JOIN todolists
--                   ON list_items.todolist_id=todolists.id
--
--
-- DELETE FROM todolists WHERE id=?;
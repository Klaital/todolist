-- name: CreateList :one
INSERT INTO todolists (owner_id, name)
VALUES (?, ?)
RETURNING id;

-- name: AddToList :one
INSERT INTO list_items (todolist_id, descr)
VALUES (?, ?)
RETURNING id;

-- name: UpdateStatus :exec
UPDATE list_items SET status=? WHERE id=?;

-- name: DeleteListItem :exec
DELETE FROM list_items WHERE id=?;

-- name: DeleteList :exec
DELETE list_items FROM list_items JOIN todolists
                  ON list_items.todolist_id=todolists.id

DELETE FROM todolists WHERE id=?;
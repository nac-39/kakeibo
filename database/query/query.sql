-- name: GetBox :one
SELECT * from box
WHERE box_id = $1;

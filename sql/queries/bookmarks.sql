-- name: AddBookmark :one
INSERT INTO bookmarks(user_id, post_id)
VALUES (
  $1,
  $2
)
RETURNING *;

-- name: RemoveBookmark :exec
DELETE FROM bookmarks
WHERE user_id = $1 AND post_id = $2;

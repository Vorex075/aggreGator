-- name: AddBookmark :one
INSERT INTO bookmarks(user_id, post_id)
VALUES (
  $1,
  $2
)
RETURNING *;

-- name: UserHasBookmark :one
SELECT *
FROM bookmarks
WHERE user_id = $1 AND post_id = $2;

-- name: GetUserBookmarks :many
SELECT *
FROM bookmarks
WHERE user_id = $1 AND (created_at < $2 OR (created_at < $2 AND post_id < $3))
ORDER BY created_at DESC, post_id DESC
LIMIT $4;

-- name: RemoveBookmark :exec
DELETE FROM bookmarks
WHERE user_id = $1 AND post_id = $2;

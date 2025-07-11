-- name: CreatePost :one
INSERT INTO posts (created_at, updated_at, title, url, 
  description, published_at, feed_id)
VALUES(
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
)
RETURNING *;

-- name: GetPostForUser :many
SELECT *
FROM posts
WHERE feed_id IN (SELECT feed_follows.feed_id
  FROM feed_follows
  INNER JOIN users ON feed_follows.user_id = users.id
  WHERE users.id = $1)
LIMIT $2;

-- name: GetRecentPostsForUser :many
SELECT *
FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1 AND (published_at < $2)
ORDER BY published_at DESC
LIMIT $3;

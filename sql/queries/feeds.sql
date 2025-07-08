-- name: CreateFeed :one
INSERT INTO feeds (name, url, user_id, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: GetFeed :one
SELECT *
FROM feeds
WHERE url = $1; 

-- name: GetAllFeeds :many
SELECT feeds.name AS rss_name, feeds.url, users.name AS username
FROM users INNER JOIN feeds ON users.id = feeds.user_id;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $2
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at DESC NULLS FIRST
LIMIT 1;

-- name: GetFeedById :one
SELECT *
FROM feeds
WHERE id = $1;

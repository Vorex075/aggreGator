-- name: CreateFeedFollow :one
WITH inserted AS(
  INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
  VALUES (
    $1,
    $2,
    $3,
    $4
  )
  RETURNING *
)
SELECT inserted.*, users.name AS username, feeds.name AS feed_name
FROM inserted
INNER JOIN users ON inserted.user_id = users.id
INNER JOIN feeds ON inserted.feed_id = feeds.id;


-- name: GetFeedsFollowForUser :many
SELECT feed_follows.id,  users.name AS username, feeds.name AS feed_name, 
  feed_follows.created_at, feed_follows.updated_at, 
  feed_follows.user_id, feed_follows.feed_id
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE users.name = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE user_id = $1 AND feed_id = $2;


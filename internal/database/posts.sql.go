// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
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
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
	FeedID      int32
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostForUser = `-- name: GetPostForUser :many
SELECT id, created_at, updated_at, title, url, description, published_at, feed_id
FROM posts
WHERE feed_id IN (SELECT feed_follows.feed_id
  FROM feed_follows
  INNER JOIN users ON feed_follows.user_id = users.id
  WHERE users.id = $1)
LIMIT $2
`

type GetPostForUserParams struct {
	ID    uuid.UUID
	Limit int32
}

func (q *Queries) GetPostForUser(ctx context.Context, arg GetPostForUserParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostForUser, arg.ID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRecentPostsForUser = `-- name: GetRecentPostsForUser :many
SELECT posts.id, posts.created_at, posts.updated_at, title, url, description, published_at, posts.feed_id, feed_follows.id, feed_follows.created_at, feed_follows.updated_at, user_id, feed_follows.feed_id
FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1 AND (published_at < $2)
ORDER BY published_at DESC
LIMIT $3
`

type GetRecentPostsForUserParams struct {
	UserID      uuid.UUID
	PublishedAt time.Time
	Limit       int32
}

type GetRecentPostsForUserRow struct {
	ID          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt time.Time
	FeedID      int32
	ID_2        int32
	CreatedAt_2 time.Time
	UpdatedAt_2 time.Time
	UserID      uuid.UUID
	FeedID_2    int32
}

func (q *Queries) GetRecentPostsForUser(ctx context.Context, arg GetRecentPostsForUserParams) ([]GetRecentPostsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentPostsForUser, arg.UserID, arg.PublishedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentPostsForUserRow
	for rows.Next() {
		var i GetRecentPostsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.UserID,
			&i.FeedID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

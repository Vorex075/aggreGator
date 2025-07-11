// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: bookmarks.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addBookmark = `-- name: AddBookmark :one
INSERT INTO bookmarks(user_id, post_id)
VALUES (
  $1,
  $2
)
RETURNING user_id, post_id, created_at
`

type AddBookmarkParams struct {
	UserID uuid.UUID
	PostID int32
}

func (q *Queries) AddBookmark(ctx context.Context, arg AddBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, addBookmark, arg.UserID, arg.PostID)
	var i Bookmark
	err := row.Scan(&i.UserID, &i.PostID, &i.CreatedAt)
	return i, err
}

const getUserBookmarks = `-- name: GetUserBookmarks :many
SELECT user_id, post_id, created_at
FROM bookmarks
WHERE user_id = $1 AND (created_at < $2 OR (created_at < $2 AND post_id < $3))
ORDER BY created_at DESC, post_id DESC
LIMIT $4
`

type GetUserBookmarksParams struct {
	UserID    uuid.UUID
	CreatedAt time.Time
	PostID    int32
	Limit     int32
}

func (q *Queries) GetUserBookmarks(ctx context.Context, arg GetUserBookmarksParams) ([]Bookmark, error) {
	rows, err := q.db.QueryContext(ctx, getUserBookmarks,
		arg.UserID,
		arg.CreatedAt,
		arg.PostID,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bookmark
	for rows.Next() {
		var i Bookmark
		if err := rows.Scan(&i.UserID, &i.PostID, &i.CreatedAt); err != nil {
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

const removeBookmark = `-- name: RemoveBookmark :exec
DELETE FROM bookmarks
WHERE user_id = $1 AND post_id = $2
`

type RemoveBookmarkParams struct {
	UserID uuid.UUID
	PostID int32
}

func (q *Queries) RemoveBookmark(ctx context.Context, arg RemoveBookmarkParams) error {
	_, err := q.db.ExecContext(ctx, removeBookmark, arg.UserID, arg.PostID)
	return err
}

const userHasBookmark = `-- name: UserHasBookmark :one
SELECT user_id, post_id, created_at
FROM bookmarks
WHERE user_id = $1 AND post_id = $2
`

type UserHasBookmarkParams struct {
	UserID uuid.UUID
	PostID int32
}

func (q *Queries) UserHasBookmark(ctx context.Context, arg UserHasBookmarkParams) (Bookmark, error) {
	row := q.db.QueryRowContext(ctx, userHasBookmark, arg.UserID, arg.PostID)
	var i Bookmark
	err := row.Scan(&i.UserID, &i.PostID, &i.CreatedAt)
	return i, err
}

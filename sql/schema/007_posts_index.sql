-- +goose Up
CREATE INDEX idx_posts_pub_id
ON posts (published_at DESC);

-- +goose Down
DROP INDEX idx_posts_pub_id;

-- +goose Up
CREATE TABLE bookmarks(
  user_id uuid,
  post_id INTEGER,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  PRIMARY KEY (user_id, post_id),
  CONSTRAINT fk_user
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE,
  CONSTRAINT fk_post
  FOREIGN KEY (post_id)
  REFERENCES posts(id)
  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE bookmarks;

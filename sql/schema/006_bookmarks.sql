-- +goose Up
CREATE TABLE bookmarks(
  user_id uuid,
  post_id INTEGER,
  PRIMARY KEY (user_id, post_id),
  CONSTRAINT fk_user
  FOREIGN KEY (user_id)
  REFERENCES users(id),
  CONSTRAINT fk_post
  FOREIGN KEY (post_id)
  REFERENCES posts(id)
);

-- +goose Down

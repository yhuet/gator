-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id uuid NOT NULL REFERENCES users ON DELETE CASCADE,
    feed_id uuid NOT NULL REFERENCES feeds ON DELETE CASCADE,
    CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT FK_feed_id FOREIGN KEY (feed_id) REFERENCES feeds (id),
    CONSTRAINT U_user_feed UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feeds_follows;

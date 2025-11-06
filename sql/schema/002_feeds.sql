-- +goose Up
CREATE TABLE feeds (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id uuid NOT NULL REFERENCES users ON DELETE CASCADE,
    CONSTRAINT FK_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down
DROP TABLE feeds;

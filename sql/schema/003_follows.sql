-- +goose Up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id uuid NOT NULL
        REFERENCES users(id)
            ON DELETE CASCADE,
    feed_id uuid NOT NULL
        REFERENCES feeds(id)
            ON DELETE CASCADE,

    UNIQUE (user_id, feed_id)
);

ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
DROP TABLE feed_follows;
ALTER TABLE feeds DROP COLUMN last_fetched_at;
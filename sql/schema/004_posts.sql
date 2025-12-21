-- +goose Up
CREATE TABLE posts (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL,
    feed_id uuid NOT NULL
        REFERENCES feeds(id)
            ON DELETE CASCADE,

    UNIQUE (url)
);

-- +goose Down
DROP TABLE posts;

-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,

    UNIQUE (name)
);

CREATE TABLE IF NOT EXISTS feeds (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id uuid NOT NULL REFERENCES users(id)
    ON DELETE CASCADE,

    UNIQUE (url)
);

-- +goose Down
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS feeds;

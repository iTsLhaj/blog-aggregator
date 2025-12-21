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

CREATE TABLE IF NOT EXISTS feed_follows (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- "Add a unique constraint on user/feed pairs - we don't want duplicate follow records."
    -- it say's user/feed pairs so i guess i fucked-up
    user_id uuid NOT NULL REFERENCES users(id)
        ON DELETE CASCADE,
    feed_id uuid NOT NULL REFERENCES feeds(id)
        ON DELETE CASCADE,

    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE IF EXISTS feed_follows;
DROP TABLE IF EXISTS feeds;
DROP TABLE IF EXISTS users;
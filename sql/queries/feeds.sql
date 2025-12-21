-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: GetFeedByURL :one
SELECT *
FROM feeds
WHERE url = $1;

-- name: CreateFeedFollow :one
WITH inserted_follow AS (
    INSERT INTO feed_follows (user_id, feed_id)
        VALUES ($1, $2)
        RETURNING *
)
SELECT
    inserted_follow.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM inserted_follow
    JOIN users ON inserted_follow.user_id = users.id
    JOIN feeds ON inserted_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follows
    JOIN users ON feed_follows.user_id = users.id
    JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = now(),
    last_fetched_at = now()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;

-- name: DeleteFeedFollowForUser :exec
DELETE FROM feed_follows
WHERE user_id = $1
AND feed_id = $2;

-- name: DeleteFeeds :exec
DELETE FROM feeds;

-- name: DeleteFollows :exec
DELETE FROM feed_follows;

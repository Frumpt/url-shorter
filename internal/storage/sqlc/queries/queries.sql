
-- name: SaveURL :one
INSERT INTO url (id, alias, url)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetURL :one
SELECT url FROM url
WHERE alias = $1
    LIMIT 1;

-- name: DeleteURL :exec
WITH rows AS (
DELETE FROM url
WHERE url.alias = $1
    RETURNING *
)
SELECT count(*) FROM rows;
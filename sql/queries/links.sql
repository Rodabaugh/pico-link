-- name: CreateLink :one
INSERT INTO links (id, created_at, updated_at, link_name, link_url)
VALUES (
    gen_random_uuid(), NOW(), NOW(), $1, $2
)
RETURNING *;

-- name: GetAllLinks :many
SELECT * FROM links;

-- name: GetLinkByID :one
SELECT * FROM links WHERE id = $1;

-- name: GetLinkByName :one
SELECT * FROM links WHERE link_name = $1;

-- name: DeleteLinkByID :exec
DELETE FROM links WHERE id = $1;
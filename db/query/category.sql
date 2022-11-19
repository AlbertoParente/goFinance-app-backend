-- name: CreateCategory :one
INSERT INTO categories (
    user_id,
    title,
    type,
    description
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories WHERE user_id = $1 AND type = $2 AND title LIKE $3 description LIKE $4;

-- name: UpdateCategory :one
UPDATE categories SET title = $1, description = $2 WHERE ID = $1;

-- name: DelereCategory :exec
DELETE FROM categories WHERE id = $1;
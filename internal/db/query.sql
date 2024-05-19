-- Users

-- name: GetUser :one
SELECT *
FROM users
WHERE id = @id AND deleted_at IS NULL LIMIT 1;

-- name: SetUser :one
INSERT INTO users (id, name, bio)
  VALUES (@id, @name, @bio)
  ON CONFLICT(id) 
  DO UPDATE SET name=excluded.name, bio=excluded.bio, updated_at=date('now')
RETURNING *;

-- name: SearchUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL 
LIMIT @limit OFFSET @offset;

-- name: CountUsers :one
SELECT count(id) as countUsers
FROM users
WHERE deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at=date('now')
WHERE id = @id LIMIT 1;

-- Products

-- name: GetProduct :one
SELECT *
FROM products
WHERE id = @id AND deleted_at IS NULL LIMIT 1;

-- name: SetProduct :one
INSERT INTO products (id, name)
  VALUES (@id, @name)
  ON CONFLICT(id) 
  DO UPDATE SET name=excluded.name, updated_at=date('now')
RETURNING *;

-- name: SearchProducts :many
SELECT *
FROM products
WHERE deleted_at IS NULL 
LIMIT @limit OFFSET @offset;
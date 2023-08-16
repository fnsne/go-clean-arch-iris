
-- name: GetBook :one
SELECT * FROM books WHERE id = $1;

-- name: SaveBook :exec
INSERT INTO books (id, title, author, isbn) VALUES ($1, $2, $3, $4);
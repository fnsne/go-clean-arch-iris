// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package books

import (
	"context"
)

const getBook = `-- name: GetBook :one
SELECT id, title, author, isbn, created_at, updated_at FROM books WHERE id = $1
`

func (q *Queries) GetBook(ctx context.Context, id int32) (Book, error) {
	row := q.db.QueryRow(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Isbn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const saveBook = `-- name: SaveBook :exec
INSERT INTO books (id, title, author, isbn) VALUES ($1, $2, $3, $4)
`

type SaveBookParams struct {
	ID     int32
	Title  string
	Author string
	Isbn   string
}

func (q *Queries) SaveBook(ctx context.Context, arg SaveBookParams) error {
	_, err := q.db.Exec(ctx, saveBook,
		arg.ID,
		arg.Title,
		arg.Author,
		arg.Isbn,
	)
	return err
}

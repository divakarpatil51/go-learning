package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(ctx context.Context, post *Post) error
	}
	// User interface {
	// 	Create(db sql.DB, user interface{}) error
	// }
}

func NewStorage(db *sql.DB) Storage {

	return Storage{
		Posts: &PostStore{db},
	}
}

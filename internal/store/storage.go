package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ERRNOTFOUND = errors.New("Resource not found")
)

type Storage struct {
	PostsIF interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	UsersIF interface {
		Create(context.Context, *User) error
	}
	CommentsIF interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		PostsIF:    &PostStore{db},
		UsersIF:    &UsersStore{db},
		CommentsIF: &CommentsStore{db},
	}
}

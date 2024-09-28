package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound            = errors.New("resource not found")
	ErrConflict            = errors.New("same title already exists")
	ErrForeignKeyViolation = errors.New("invalid user_id, user does not exist")
	ErrDuplicateEmail      = errors.New("duplicate emailid")
	ErrDuplicateUsername   = errors.New("duplicate username")
	QueryTimeoutDuration   = time.Second * 5
)

type Storage struct {
	PostsIF interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetaData, error)
	}
	UsersIF interface {
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		CreateAndInvite(ctx context.Context, user *User, token string, exp time.Duration) error
		Activate(context.Context, string) error
		Delete(context.Context, int64) error
		GetByEmail(context.Context, string) (*User, error)
	}
	CommentsIF interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	FollowIF interface {
		Follow(context.Context, int64, int64) error
		UnFollow(context.Context, int64, int64) error
	}
	RoleIF interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		PostsIF:    &PostStore{db},
		UsersIF:    &UsersStore{db},
		CommentsIF: &CommentsStore{db},
		FollowIF:   &FollowerStore{db},
		RoleIF:     &RoleStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

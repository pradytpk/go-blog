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
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
		CreateAndInvite(ctx context.Context, user *User, token string) error
	}
	CommentsIF interface {
		Create(context.Context, *Comment) error
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
	FollowIF interface {
		Follow(context.Context, int64, int64) error
		UnFollow(context.Context, int64, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		PostsIF:    &PostStore{db},
		UsersIF:    &UsersStore{db},
		CommentsIF: &CommentsStore{db},
		FollowIF:   &followerStore{db},
	}
}

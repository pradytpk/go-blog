package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type followerStore struct {
	db *sql.DB
}

type Follower struct {
	UserID     int64  `json:"user_id"`
	FollowerID int64  `json:"follower_id"`
	CreatedAt  string `json:"created_at"`
}

func (s *followerStore) Follow(ctx context.Context, followerUserId int64, userId int64) error {
	query := `
	INSERT into followers (user_id, follower_id) values ($1,$2)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, userId, followerUserId)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return ErrConflict
			case "foreign_key_violation":
				return ErrForeignKeyViolation
			}
		}
		return err
	}
	return nil
}

func (s *followerStore) UnFollow(ctx context.Context, followerUserId int64, userId int64) error {
	query := `
	DELETE FROM followers WHERE user_id = $1 and follower_id= $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := s.db.ExecContext(ctx, query, userId, followerUserId)
	return err
}

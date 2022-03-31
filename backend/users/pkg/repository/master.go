package repository

import (
	"context"
	"time"

	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type Master interface {
	// User.
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, userID gocql.UUID) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID gocql.UUID) error
	LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error)
	LogOut(ctx context.Context, userID gocql.UUID) error
	LogInWithFB(ctx context.Context, userID gocql.UUID, token string) (*entity.User, error)
	// Phrase.
	CreatePhrase(ctx context.Context, phrase *entity.Phrase) error
	GetPhrases(ctx context.Context, userID gocql.UUID, before, after time.Time) ([]*entity.Phrase, error)
	DeletePhrase(ctx context.Context, userID gocql.UUID, curTime time.Time) error
}

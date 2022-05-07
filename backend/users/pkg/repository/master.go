package repository

import (
	"context"
	"time"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type Master interface {
	// User.
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID string) error
	LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error)
	LogOut(ctx context.Context, userID string) error
	LogInWithFB(ctx context.Context, userID, token string) (*entity.User, error)
	// Phrase.
	CreatePhrase(ctx context.Context, phrase *entity.Phrase) error
	GetPhrases(ctx context.Context, userID string, before, after time.Time) ([]*entity.Phrase, error)
	DeletePhrase(ctx context.Context, userID string, curTime time.Time) error
}

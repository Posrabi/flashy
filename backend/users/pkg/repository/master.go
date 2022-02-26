// nolint: dupl
package repository

import (
	"context"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type Master interface {
	// User.
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user_id, hash_password string) error
	LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error)
	LogOut(ctx context.Context, userID string) error
}

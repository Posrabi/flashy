package repository

import (
	"context"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type User interface {
	CreateUser(ctx context.Context, user *entity.User) (userID, authToken string, err error)
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User) error
	LogIn(ctx context.Context, username, hashPassword string) (userID, authToken string, err error)
	LogOut(ctx context.Context, userID string) error
	ValidateUser(ctx context.Context) error
}

package repository

import (
	"context"

	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type User interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, userID gocql.UUID) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID gocql.UUID) error
	LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error)
	LogOut(ctx context.Context, userID gocql.UUID) error
}

package api

import (
	"context"

	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	sc "github.com/Posrabi/flashy/backend/users/pkg/scylla"
)

type masterRepository struct {
	user repository.User
}

func NewMasterRepository(sess *gocql.Session) repository.Master {
	return &masterRepository{
		user: sc.NewUserRepository(sess),
	}
}

func (m *masterRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return m.user.CreateUser(ctx, user)
}

func (m *masterRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
	return m.user.GetUser(ctx, userID)
}

func (m *masterRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return m.user.UpdateUser(ctx, user)
}

func (m *masterRepository) DeleteUser(ctx context.Context, userID string) error {
	return m.user.DeleteUser(ctx, userID)
}

func (m *masterRepository) LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error) {
	return m.user.LogIn(ctx, username, hashPassword)
}

func (m *masterRepository) LogOut(ctx context.Context, userID string) error {
	return m.user.LogOut(ctx, userID)
}

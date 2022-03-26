package api

import (
	"context"
	"time"

	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
	sc "github.com/Posrabi/flashy/backend/users/pkg/scylla"
)

type masterRepository struct {
	user   repository.User
	phrase repository.Phrase
}

func NewMasterRepository(sess *gocql.Session) repository.Master {
	return &masterRepository{
		user:   sc.NewUserRepository(sess),
		phrase: sc.NewPhraseRepository(sess),
	}
}

func (m *masterRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return m.user.CreateUser(ctx, user)
}

func (m *masterRepository) GetUser(ctx context.Context, userID gocql.UUID) (*entity.User, error) {
	return m.user.GetUser(ctx, userID)
}

func (m *masterRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return m.user.UpdateUser(ctx, user)
}

func (m *masterRepository) DeleteUser(ctx context.Context, userID gocql.UUID) error {
	return m.user.DeleteUser(ctx, userID)
}

func (m *masterRepository) LogIn(ctx context.Context, username, hashPassword string) (*entity.User, error) {
	return m.user.LogIn(ctx, username, hashPassword)
}

func (m *masterRepository) LogOut(ctx context.Context, userID gocql.UUID) error {
	return m.user.LogOut(ctx, userID)
}

func (m *masterRepository) CreatePhrase(ctx context.Context, phrase *entity.Phrase) error {
	return m.phrase.CreatePhrase(ctx, phrase)
}

func (m *masterRepository) GetPhrases(ctx context.Context, userID gocql.UUID, start, end time.Time) ([]*entity.Phrase, error) {
	return m.phrase.GetPhrases(ctx, userID, start, end)
}

func (m *masterRepository) DeletePhrase(ctx context.Context, userID gocql.UUID, curTime time.Time) error {
	return m.phrase.DeletePhrase(ctx, userID, curTime)
}

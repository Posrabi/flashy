package repository

import (
	"context"
	"time"

	"github.com/gocql/gocql"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type Phrase interface {
	CreatePhrase(ctx context.Context, phrase *entity.Phrase) error
	GetPhrases(ctx context.Context, userID gocql.UUID, start, end time.Time) ([]*entity.Phrase, error)
	DeletePhrase(ctx context.Context, userID gocql.UUID, phraseTime time.Time) error
}

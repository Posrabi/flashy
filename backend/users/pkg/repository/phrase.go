package repository

import (
	"context"
	"time"

	"github.com/Posrabi/flashy/backend/users/pkg/entity"
)

type Phrase interface {
	CreatePhrase(ctx context.Context, phrase *entity.Phrase) error
	GetPhrases(ctx context.Context, userID string, start, end time.Time) ([]*entity.Phrase, error)
	DeletePhrase(ctx context.Context, userID string, phraseTime time.Time) error
}

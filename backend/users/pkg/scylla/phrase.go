package scylla

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"google.golang.org/grpc/codes"

	gerr "github.com/Posrabi/flashy/backend/common/pkg/error"
	"github.com/Posrabi/flashy/backend/users/pkg/entity"
	"github.com/Posrabi/flashy/backend/users/pkg/repository"
)

type phraseRepo struct {
	sess *gocql.Session
}

const (
	phraseColumns = "user_id, phrase_time, word, sentence, correct"
	phraseTable   = "users.phrases"
)

func NewPhraseRepository(sess *gocql.Session) repository.Phrase {
	return &phraseRepo{
		sess: sess,
	}
}

func (p *phraseRepo) CreatePhrase(ctx context.Context, phrase *entity.Phrase) error {
	q := `INSERT INTO %s (%s) VALUES (?, toTimestamp(now()), ?, ?, ?) IF NOT EXISTS`

	args := []interface{}{phrase.UserID, phrase.Word, phrase.Sentence, phrase.Correct}

	if err := p.sess.Query(fmt.Sprintf(q, phraseTable, phraseColumns), args...).Idempotent(true).WithContext(ctx).Exec(); err != nil {
		return gerr.NewScError(err, codes.Internal, fmt.Sprintf(q, info, phraseColumns), args)
	}

	return nil
}

func (p *phraseRepo) GetPhrases(ctx context.Context, userID gocql.UUID, start, end time.Time) ([]*entity.Phrase, error) {
	q := `SELECT %s FROM %s WHERE user_id = ? AND phrase_time > ? and phrase_time < ?`

	args := []interface{}{userID, start.UnixMilli(), end.UnixMilli()}

	var phrases []*entity.Phrase

	scanner := p.sess.Query(fmt.Sprintf(q, phraseColumns, phraseTable), args...).Idempotent(true).WithContext(ctx).Iter().Scanner()
	for scanner.Next() {
		var phrase entity.Phrase
		if err := scanner.Scan(&phrase.UserID, &phrase.Time, &phrase.Word, &phrase.Sentence, &phrase.Correct); err != nil {
			return nil, gerr.NewScError(err, codes.Internal, fmt.Sprintf(q, phraseColumns, phraseTable), args)
		}
		phrases = append(phrases, &phrase)
	}

	if err := scanner.Err(); err != nil {
		return nil, gerr.NewScError(err, codes.Internal, fmt.Sprintf(q, phraseColumns, phraseTable), args)
	}

	return phrases, nil
}

func (p *phraseRepo) DeletePhrase(ctx context.Context, userID gocql.UUID, phraseTime time.Time) error {
	q := `DELETE FROM %s WHERE user_id = ? AND phrase_time = ?`

	args := []interface{}{userID, phraseTime.UnixMilli()}

	if err := p.sess.Query(fmt.Sprintf(q, phraseTable), args...).Idempotent(true).Exec(); err != nil {
		return gerr.NewScError(err, codes.Internal, fmt.Sprintf(q, phraseTable), args)
	}

	return nil
}
